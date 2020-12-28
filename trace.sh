#! /bin/bash

strace -c go run ./unsafe/unsafe.go 1>/dev/null
dtruss -c go run ./unsafe/unsafe.go 2>&1