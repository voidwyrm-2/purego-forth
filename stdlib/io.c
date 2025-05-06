#include "common.h"
#include <stdio.h>

int echo(push_t, pop_t, int slen) {
  if (slen == 0) {
    return 1;
  }

  int n = pop();

  printf("%d", n);

  return 0;
}
/*
int emit(push_t, pop_t, int slen) {
  if (slen == 0) {
    return 1;
  }

  int n = call_f(pop);

  printf("%c", n);

  return 0;
}
*/
