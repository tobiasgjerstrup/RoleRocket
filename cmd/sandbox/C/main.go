package main

/*
#cgo CFLAGS: -g -Wall
#include "mylib.c"
*/
import "C"
import "fmt"

func main() {
	C.say_hello()
	fmt.Println(C.return_number())
}
