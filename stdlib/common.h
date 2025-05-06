#define push_t void (*push)(int n)
#define pop_t int (*pop)(void)

#define check_underflow()                                                      \
  if (slen == 0) {                                                             \
    return 1;                                                                  \
  }

#define stack_f(name_, body_) int name_(push_t, pop_t, int slen) body_

#define math_f(name_, op_)                                                     \
  int name_(push_t, pop_t, int slen) {                                         \
    if (slen < 2) {                                                            \
      return 1;                                                                \
    }                                                                          \
                                                                               \
    int y = pop();                                                             \
    int x = pop();                                                             \
                                                                               \
    push(x op_ y);                                                             \
                                                                               \
    return 0;                                                                  \
  }
