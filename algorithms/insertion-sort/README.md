# Efficiency of Insertion Sort

The insertion sort algorithm involves four steps: removal, comparison, shift, insertion.

In the worst case (when the array is provided in descending order), comparisons happen `1 + 2 + 3 ... + (n-1)` times. For `n=5` this would be 10 comparisons. The same number of shifts are required if the array is in descending order.

Insertion and removal happen once every iteration so they each add `n-1` steps, bringing the algorithm to `O(n^2 + 2n - 2)`. This can be simplified to `O(n^2)`, since we only care about the highest order if `n`.
