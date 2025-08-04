#include <stdio.h>
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