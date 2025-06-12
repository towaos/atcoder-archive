#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int length_sort(const void *a, const void *b) {
  char *str1 = *(char**)a;
  char *str2 = *(char**)b;

  int len1 = strlen(str1);
  int len2 = strlen(str2);

  return len1 - len2;
}

int main() {
  int N;
  int variables;
  char list[50][51];
  char *pointers[50];
  char result[1276] = "";

  variables = scanf("%d", &N);
  
  for (int i = 0; i < N; i++) {
    variables = scanf("%s", list[i]);
    pointers[i] = list[i];
  }

  qsort(pointers, N, sizeof(char*), length_sort);

  for (int i = 0; i < N; i++) {
    strcat(result, pointers[i]);
  }

  printf("%s\n", result);

  return 0;
}