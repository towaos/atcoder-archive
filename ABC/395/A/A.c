#include <stdio.h>

int main() {
  int N;
  int variables;
  int A[1001];

  variables = scanf("%d", &N);

  for (int i = 0; i < N; i++) {
    variables = scanf("%d", &A[i]);
  }

  int num = A[0];
  int is_ascending = 1;


  for (int i = 1; i < N; i++) {
    if (num >= A[i]) {
      is_ascending = 0;
      break;
    }
    num = A[i];
  }

  if (is_ascending) {
    puts("Yes");
  } else {
    puts("No");
  }

  return 0;
}