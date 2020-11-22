Time complexity is the computational complexity that describes the amount of time it takes to run an algorithm. The exact amount of time it takes the computer to run an operation, of course, depends on the computer's power. In computer science, the runtime of an operation refers to the number of instructions performed during the operation. 

Runtimes are described with the big O notation and is not limited to time. Common runtimes are `O(n)`, `O(log n)`, `O(n log n)` and `O(n**2)`. But this list is not exhaustive and a runtime can have multiple variables.

Big O is about the expression of runtime *scaling*, more so than anything. Highly precise notation is actually not the point of using big O, which is why constants and non-dominant terms can be dropped in big O notation. It's all about scaling.

Big O notation can also be used to describe the amount of memory allocation the algorithm requires.

# Seven common functions

There are seven functions that are most commonly used in algorithms.

## Constant functions

The simplest function is a **constant function**.

```
f(n) -> c
```

In this constant function, it does not matter what the parameter `n` is; the function call will always be equal to the constant value of `c`. The most fundamental constant function is `f(n) = 1`.

## Logarithmic functions

The **logarithmic function** is defined as:

```
f(n) -> log_b_n
```

where `b` is the *base* of the log and greater than 1. In computer science, the most common log base is 2 as computers store integers in binary notation. Is is so common, that `log 2 n` is simply written as `log n`

## Linear functions

The **linear function** is another simple function:

```
f(n) -> n
```

Linear functions appear in algorithms any time there is a basic operation for each of `n` elements (think iterating over arrays). It represents the best running in any algorithm thata processes each of `n` objecsts that are not already in memory (since reading in `n` objects aready requires that many operations).

## N log N functions

The **n log n function** assigns to an input `n` the value of `n` times the logarithm base-two of `n`. 

```
f(n) -> n log n
```

The function grows more rapidly than a linear function but not as rapidly as a quadratic function. Sorting algorithms are often good examples of functions with a runtime of `n log n`.

## Quadratic functions

A function that grows rather rapidly is the **quadratic function**.

```
f(n) -> n^2
```

That is, given an input `n`, the function returns `n` squared. Quadratic functions appear in time complexity analyses when there are nested loops.

## Cubic functions and other polynomials

Another functions that is the power of the input is the **cubic function**.

```
f(n) -> n^3
```

While this functions will grow even faster than the quadratic function, it appears much less frequently.

## Exponential functions

An **exponential function** looks like this:

```
f(n) -> b^n
```

The *positive constant* `b` is called the base and parameter `n` is the *exponent*. Similar to the logarithmic function, the most common base for the exponential function is `b=2`.

## Comparison

| constant | logarithm | linear | n-log-n   | quadratic | cubic | exponential |
| `1`      | `log n`   | `n`    | `n log n` | `n^2`     | `n^3` | `a^n`       |

In an ideal world, data structure operations have a logarithm or constant runtime and algorithms have a linear or N log N runtime. Algorithms with quadratic or cubic runtimes are impractical and functions with exponential runtimes should be avoided for all but the smallest inputs.

![Time compexity comparison](https://upload.wikimedia.org/wikipedia/commons/7/7e/Comparison_computational_complexity.svg)

# Read more

- [Data Structures and Algorithms in Python](https://www.amazon.nl/Structures-Algorithms-Python-Michael-Goodrich/dp/1118290275/ref=sr_1_1?__mk_nl_NL=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=1D52V3F27RXYV&dchild=1&keywords=data+structures+and+algorithms+in+python&qid=1605781127&sprefix=data++struct%2Caps%2C151&sr=8-1)
- [Open Data Structures](https://opendatastructures.org/ods-python/1_3_Mathematical_Background.html#950)
- [Wikipedia](https://en.wikipedia.org/wiki/Time_complexity)
