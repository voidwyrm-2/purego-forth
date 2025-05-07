#include "common.h"
#include <stdio.h>
#include <unistd.h>

// NOTICE: -fPIC does not fix the issue stdout/in when on darwin/arm64
// do not change this until a solution is figured out

stack_f(echo, {
  expect_atleast(1);

  int n = pop();

  char buf[64];
  int sz = sprintf(buf, "%d\n", n);
  write(1, buf, sz);

  return 0;
});

stack_f(emit, {
  expect_atleast(1);

  int n = pop();

  char buf[3];
  int sz = sprintf(buf, "%c\n", n);
  write(1, buf, sz);

  return 0;
});

// experimental, doesn't work that well
/*
stack_f(inputc, {
  char buf[1];
  read(0, buf, 1);

  push((int)buf[0]);

  return 0;
});
*/
