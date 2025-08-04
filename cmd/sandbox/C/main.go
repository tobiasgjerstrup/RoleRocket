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
	"sort"
	"time"
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

	input := generateInput(100000)
	cInput := toCIntSlice(input)

	// Obviously GO cheats in this example lol
	benchmarkBubbleSort(cInput)
	benchmarkQuickSort(cInput)
	benchmarkGoSort(input)
}

func benchmarkBubbleSort(input []C.int) {
	size := C.int(len(input))
	cInput := copyToC(input)

	start := time.Now()
	sortedPtr := C.bubbleSort(cInput, size)
	duration := time.Since(start)

	if sortedPtr == nil {
		panic("Bubble Sort failed")
	}

	C.free(unsafe.Pointer(cInput))
	C.free(unsafe.Pointer(sortedPtr))

	fmt.Printf("Bubble Sort took %v\n", duration)
}

func benchmarkQuickSort(input []C.int) {
	size := C.int(len(input))
	cInput := copyToC(input)

	start := time.Now()
	sortedPtr := C.quickSortWrapper(cInput, size)
	duration := time.Since(start)

	if sortedPtr == nil {
		panic("Quick Sort failed")
	}

	C.free(unsafe.Pointer(cInput))
	C.free(unsafe.Pointer(sortedPtr))

	fmt.Printf("Quick Sort took %v\n", duration)
}

func copyToC(input []C.int) *C.int {
	size := len(input)
	cInput := (*C.int)(C.malloc(C.size_t(size) * C.size_t(unsafe.Sizeof(C.int(0)))))
	for i := 0; i < size; i++ {
		*(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(cInput)) + uintptr(i)*unsafe.Sizeof(C.int(0)))) = input[i]
	}
	return cInput
}

func generateInput(size int) []int {
	input := make([]int, size)
	for i := 0; i < size; i++ {
		input[i] = rand.Intn(100)
	}
	return input
}

func benchmarkGoSort(input []int) {
	copyInput := make([]int, len(input))
	copy(copyInput, input)

	start := time.Now()
	sort.Ints(copyInput)
	duration := time.Since(start)

	fmt.Printf("Go Default Sort took %v\n", duration)
}

func toCIntSlice(goSlice []int) []C.int {
	cSlice := make([]C.int, len(goSlice))
	for i, v := range goSlice {
		cSlice[i] = C.int(v)
	}
	return cSlice
}
