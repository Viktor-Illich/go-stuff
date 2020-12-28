#! /bin/bash

go build -o ./use_by_c.o -buildmode=c-shared use_by_c.go

gcc -o ./use_go ./use_go.c ./use_by_c.o

rm use_by_c.h
rm use_by_c.o

./use_go
rm use_go