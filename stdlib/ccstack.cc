#include "common.h"

extern "C" int dup(push_t, pop_t, int slen) {
  expect_atleast(1);

  int n = pop();

  push(n);
  push(n);

  return 20;
}
