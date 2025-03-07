#include <stdio.h>

int memcmp(const void *str1, const void *str2, size_t n) {
    printf("%s\n", str1);
    printf("%s\n", str2);
    return 0;
}