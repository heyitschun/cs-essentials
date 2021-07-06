# The (in)efficiency of Bubble Sort

Bubble sort compares two numbers, at consecutive indices, to determine which one is larger. If they are out of order, their indices will be swapped. 

It does this over multiple iterations over the array. So if the array has a size of `n=5`, then the first iteration will have four comparisons, the second will have three comparisons, and so on. In total, the algorithm makes `(n-1) + (n-2) + (n-3) ... + 1` comparisons. In the case of `n=5` that would be 10 comparisons.

# Time Complexity

The worst case scenario is if the array is given in descending order, because this means that the algorithm needs to perform the maximum number of swaps. For `n=5` that would mean 10 swaps, bringing the total number of steps to 20 (10 comparisons + 10 swaps). For `n=10` the total number of steps would be 90. For `n=20` the steps would be 380. So as the number of elements increases, the number of steps grows *exponentially*. The growth rate of the number of steps is approximately `N^2`.

The time complexity of bubble sort can be expressed as `O(n^2)`; thus, it is a relatively inefficient algorithm.
