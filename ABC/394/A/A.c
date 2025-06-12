#include <stdio.h>
#include <string.h>

int main() {
  char S[101];
  char result[101];
  int result_index = 0;
 
  scanf("%s", S);
 
  for (int i = 0; i < strlen(S); i++) {
    if (S[i] == '2') {
      result[result_index] = '2';
      result_index++;
    }
  }
 
  result[result_index] = '\0';
 
  printf("%s\n", result);
 
  return 0;
}