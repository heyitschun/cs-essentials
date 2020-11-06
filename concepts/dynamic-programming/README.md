**Dynamic programming** refers to simplifying a complicated problem by breaking it down into simple sub-problems in a recursive manner. Decisions that span several points in time often break apart recursively. If the sub-problems can be nested recursively inside larger problems, so that dynamic programming methods are applicable, then there is a relation between the value of the larger problem and the values of the sub-problems. This relationship is called the [Bellman equation](https://en.wikipedia.org/wiki/Bellman_equation).

# Optimal substructure

A problem can be solved with dynamic programming if it has two key attributes: *optimal substructure* and *overlapping sub-problems*. Problems that combine optimal solutions to non-overlapping sub-problems then it uses a *divide and conquer* approach (merge sort and quicksort).

*Optimal substructure* means that the solution to a given optimization problem can be obtained by the combination of optimal solutions to its sub-problems (recursion).

*Overlapping sub-problems* means that the space of sub-problems must be small. A recursive algorithm that solves the problem should solve the same sub-problem over and over, not generate new ones.
