package main

/*
#cgo CFLAGS: -g -Wall
#include "mylib.h"
#include "mylib.c"
*/
import "C"
import (
	"fmt"
	"math/rand"
	"unsafe"
)

func main() {
	C.say_hello()
	number := C.return_number()
	fmt.Println(number)

	// Proper capitalization and conversion
	fmt.Println(C.GoString(C.return_char()))

	number = C.return_another_number()
	fmt.Println(number)

	fmt.Println(C.GoString(C.return_char_again()))

	input := make([]C.int, 10000)
	for i := 0; i < 10000; i++ {
		input[i] = C.int(rand.Intn(100))
	}
	size := C.int(len(input))

	cInput := (*C.int)(C.malloc(C.size_t(size) * C.size_t(unsafe.Sizeof(C.int(0)))))
	if cInput == nil {
		panic("C.malloc failed")
	}

	for i := 0; i < int(size); i++ {
		*(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(cInput)) + uintptr(i)*unsafe.Sizeof(C.int(0)))) = input[i]
	}

	sortedPtr := C.bubbleSort(cInput, size)
	if sortedPtr == nil {
		panic("Sorting failed")
	}

	sorted := make([]int, int(size))
	for i := 0; i < int(size); i++ {
		value := *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(sortedPtr)) + uintptr(i)*unsafe.Sizeof(C.int(0))))
		sorted[i] = int(value)
	}

	fmt.Println("Sorted:", sorted)

	C.free(unsafe.Pointer(cInput))
	C.free(unsafe.Pointer(sortedPtr))
}
