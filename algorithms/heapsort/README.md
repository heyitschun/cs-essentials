The *heapsort* algorithm is an in-place sorting algorithm. It uses [binary heaps](https://github.com/heyitschun/cs-essentials/tree/master/data-structures/heaps), which represent a heap as an array. Heapsort stores `n` elements in an array `a` with the smallest value at the root `a[0]`. After `a` is transformed into a binary heap, heapsort repeatedly swaps `a[0]` and `a[n-1]`, decrements `n` and trickles down `0` so that `a[0], ...a[n-2]` are heap-ordered. After this process is finished, the elements of `a` are stored in decreasing order. 

```
heapsort(a)
    h <- BinaryHeap()
    h.a <- a
    h.n <- len(a)
    m <- h.n // 2
    for i in range(m) {
        h.trickle_down(i)
    }
    while h.n > 1 {
        h.n <- h.n-1
        h.a[h.n], h.a[0] <- h.a[0], h.a[h.n]
        h.trickle_down(0)
    }
    a.reverse()
```

# Time complexity

The number of comparisons done by heapsort are in `O(n log n)`.
