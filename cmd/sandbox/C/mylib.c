#include <stdio.h>
#include <stdlib.h>
#include "mylib.h"

void say_hello() {
    printf("Hello from external C!\n");
}

int return_number() {
    return 512;
}

const char* return_char() {
    return "Hello from C!";
}

int return_another_number() {
    return 256;
}

const char* return_char_again() {
    return "Hello from C!... AGAIN!";
}

int* bubbleSort(int* input, int size) {
    int* sorted = malloc(size * sizeof(int));
    if (!sorted) return NULL;

    // Copy input array
    for (int i = 0; i < size; i++)
        sorted[i] = input[i];

    // Sort the new array
    for (int i = 0; i < size - 1; i++) {
        for (int j = 0; j < size - i - 1; j++) {
            if (sorted[j] > sorted[j + 1]) {
                int temp = sorted[j];
                sorted[j] = sorted[j + 1];
                sorted[j + 1] = temp;
            }
        }
    }

    return sorted;
}