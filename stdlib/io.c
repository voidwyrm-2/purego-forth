#include "common.h"
#include <stdio.h>
#include <unistd.h>

int echo(push_t, pop_t, int slen) {
  if (slen == 0) {
    return 1;
  }

  int n = pop();
  char buf[64];
  int sz = sprintf(buf, "%d\n", n);
  write(1, buf, sz);

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
