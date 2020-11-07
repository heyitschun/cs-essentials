A **stack** is an abstract data type (ADT) that involves dynamic or static implementation. Its essential feature is that it is ordered and that access to it is restricted to one end. Stacks follow the last-in-first-out (LIFO) protocol. An implementation of a stack must include at least two operations: *pushing* to the stack and *popping* from the stack. It must also include a pointer to the element on the top of the stack.

**Pushing** refers to adding a new element to the top of the stack, while **popping** refers to removing the element that is at the top of the stack.

![Stack](https://en.wikipedia.org/wiki/File:Lifo_stack.png)

A stack implementation might also include other methods such as `empty()` to check if the stack has anything on it. Or `top()` to retrieve the top element of the stack without removing it.
