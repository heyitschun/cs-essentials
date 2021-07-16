Quicksort is a divide-and-conquer algorithm. It picks a pivot element `x` from an array `a`, partitions `a` into a set of elements less than `x`, a set elements equal to `x` and a set of elements greater than `x`. Quicksort then recursively sorts the first and third sets in this partition.

![Quicksort](https://opendatastructures.org/ods-python/img4174.png)

```
quicksort(a, i, n)
    if n <= 1 {
        return
    }
    
    x <- a[i+randint(0, n)]
    (p, j, q) <- (i-1, i, i+n)
    while j < q {
        if a[j] < x {
            p <- p+1
            a[j], a[p] <- a[p], a[j]
            j <- j+1
        } else if a[j] > x {
            q <- q-1
            a[j], a[q] <- a[q], a[j]
        } else {
            j <- j+1
        }
        quicksort(a, i, p-i+1)
        quicksort(a, q, n-(q-i))
    }
```

The general approach to quicksort is to choose a random element `x` as the pivot, scan left to find an element greater than `x`, scan right to find an element less than `x`, then swap these elements so that they are on the correct side of `x`. So instead of making copies of subarrays that are sorted, quicksort sorts the subarray in-place.

For some sorting algorithms, the best case is when the array comes already sorted. Interestingly, quicksort's worst scenario is when the input array is sorted (either ascending or descending)! The best case for quicksort is if the pivot point ends up smack in the middle on each after each sub-partition, because the algorithm involves recursion this will lead to fewer sub-partitions; thus, fewer steps.

# Time complexity

The average and best case complexity of quicksort is `O(n log n)`. In the worst case, however, quicksort can actually become slow, performing `O(n^2)` comparisons.
