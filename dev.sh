#!/bin/sh
go run main.go &
oldcksum=$(find . -type f | xargs cksum | cksum)

while true; do
    newcksum=$(find . -type f | xargs cksum | cksum)
    if [ "$newcksum" != "$oldcksum" ]; then
        date=$(date '+%Y-%m-%d %H:%M:%S')
        echo "[${date}] go file update detected."
        oldcksum=$newcksum
        ps | grep main.go | awk '{print $1}' | xargs kill -9
        ps | grep exe/main | awk '{print $1}' | xargs kill -9
        go run main.go &
    fi
    sleep 1
done
