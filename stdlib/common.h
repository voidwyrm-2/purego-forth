#define push_t void (*push)(int n)
#define pop_t int (*pop)(void)

#define expect_atleast(count_)                                                 \
  if (slen < count_) {                                                         \
    return 1;                                                                  \
  }

#define stack_f(name_, body_) int name_(push_t, pop_t, int slen) body_

#define math_f(name_, op_)                                                     \
  int name_(push_t, pop_t, int slen) {                                         \
    expect_atleast(2);                                                         \
                                                                               \
    int y = pop();                                                             \
    int x = pop();                                                             \
                                                                               \
    push(x op_ y);                                                             \
                                                                               \
    return 0;                                                                  \
  }
