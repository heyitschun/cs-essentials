Recursion in computer science is a method where the solution to a problem depends on the solution to smaller instances of the same problem. Put more simply, a recursive function is a function that calls itself until a condition is met.

```
factorial(n)
    if n == 1 {
        return 1
    }
    else {
        return n * factorial(n-1)
    }
```

# Performance

## Speed

Recursion is a competitor to iteration; whatever can be done recursively, can be done iteratively and vice versa. One could be given the preference over the other, but this depends on the problem and the programming language. 

In languages that favor iterative looping structures (e.g. C, Java), there is significant time and space cost associated with recursion. This is because there is an overhead to managing the stack and function calls are relatively slow. In functional languages (e.g. Haskell), function calls are fast, so the difference between a recursive approach and an iterative approach is less noticeable. 

## Space

In some languages, the maximum size of the call stack is much less that the space available on the heap. Recursive solutions tend to ask for more space than iterative solutions. So in these languages (e.g. Python), there is sometimes a limit to the depth of recursion to avoid stack overflows.

## Time complexity

A recursive function that calls itself once has a runtime of `O(n)`. But this rapidly becomes worse when the functions starts making multiple calls. 

```
multiple_calls(n)
    if n <= 1 {
        return 1
    }
    return multiple_calls(n-1) + multiple_calls(n-1)
```

If each call of `multiple_calls` is a node in a function tree, then each node has two children. So in total, there will be `2^0 + 2^1 + 2^2 + ... + 2^n` nodes of `multiple_calls`, where `2` is the number of branches each recursive call has and `n` is the depth of the function tree. This is a common runtime pattern for recursive functions with multiple calls: `O(branches^depth)`.

# Read more

* https://en.wikipedia.org/wiki/Recursion_(computer_science)
