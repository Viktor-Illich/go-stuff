#! /bin/bash

gcc -c ./callclib/*.c
ar rs ./callC.a ./*.o
rm callC.o
go run ./call_clib.go
rm ./callC.a