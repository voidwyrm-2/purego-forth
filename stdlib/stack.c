#include "common.h"

// math operations

math_f(add, +);
math_f(sub, -);
math_f(mul, *);
math_f(div, /);
math_f(rem, %);

// stack operations

/*
stack_f(dup, {
  expect_atleast(1);

  int n = pop();

  push(n);
  push(n);

  return 0;
});
*/

stack_f(swap, {
  expect_atleast(2);

  int a = pop();
  int b = pop();

  push(b);
  push(a);

  return 0;
});
