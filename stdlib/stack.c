#include "common.h"

math_f(add, +);
math_f(sub, -);
math_f(mul, *);
math_f(div, /);
math_f(rem, %);

stack_f(dup, {
  check_underflow();

  int n = pop();

  push(n);
  push(n);

  return 0;
})
