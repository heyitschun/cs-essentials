Merge sort is a divide-and-conquer algorithm that solves a large problem by dividing it into parts. The solutions of each part are then combined as form a solution to the original problem. The strategy of the merge sort is to recursively sort halves of an array and then merge the results into the final sorted array.

![Merge Sort](http://opendatastructures.org/versions/edition-0.1e/ods-java/img1364.png)

If the length of array `a` is 1 or 0, then no sorting is required. Otherwise, the array is split into two and recursively sorted. The sorted arrays are then merged.

```
merge_sort(a)
    if len(a) <= 1 {
        return a
    }
    
    m <- len(a)/2
    a0 <- merge_sort(a[m])
    a1 <- merge_sort(a[m])
    merge(a0, a1, a)
    return a
```

Elements are added to `a`, one at a time. If `a0` or `a1` are empty, then the next elements are added from the other array. Otherwise, the minimum of the next element in `a0` is taken and the next element in `a1`. These are then added to `a`.

```
merge(a0, a1, a)
    i0 <- i1 <- 0
    for i in range(len(a)-1) {
        if i0 = len(a0) {
            a[i] <- a1[i1]
            i1 <- i2
        }
        else if i1 = len(a1) {
            a[i] <- a0[i0]
            i0 <- i1
        }
        else if a0[i0] <= a1[i1] {
            a[i] <- a0[i0]
            i0 <- i1
        } else {
            a[i] <- a1[i1]
            i1 <- i2
        }
    }
```

# Time complexity

Merge sort turns the problem of sorting into two problems: sorting two arrays of length `n/2`. These two arrays turn into four. Four turn into eight until `n=2` at which point the next two problems each have a size of `n=1`. Each subproblem of size `n/2**i`, the time spent merging and copying data is `O(n/2**i)`. Since there are `2**i` subproblems of size `n/2**i`, the total time spent working on problems of size `2**i` is `2^i x O(n/2^i) = O(n)`. This this operation happens `log n` times, the time complexity of the merge sort algorithm is `O(n log n)`.
