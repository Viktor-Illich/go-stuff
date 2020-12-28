#! /bin/bash

gcc -c ./callclib/*.c
ar rs ./callC.a ./*.o
rm callC.o
go build ./call_clib.go
rm ./callC.a
./call_clib
rm ./call_clib