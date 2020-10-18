A **heap** is a complete binary tree whose every node has the *heap-order* property, which means that the value stored at the node is greater than or equal to the values stored at its descendants.

# Binary heap

Heaps can be implemented in a variety of ways, but as they are complete binary trees, storing them in contiguous memory locations becomes very efficient. *Eytzinger's method* allows for the representation of a complete binary tree as an array by laying out the nodes of the tree in breadth-first order. This way, the root's left child is at index 1, the right child at index 2, the left child of the left child of the root at index 3 and so on.

By applying this method, the left child of the node at index `i` is at index `left(i) = 2i+1`. The right child of the node at index `i` will be at `right(i) = 2i+2`. The parent of the node at index `i` will be at `parent(i) = (i-1)/2`.

![Binary heap](https://opendatastructures.org/ods-python/img3864.png)

A *BinaryHeap* uses this technique to implicitly represent a complete binary tree in which the elements are heap-ordered.

## Adding to a binary heap

```
add(x)
    if len(array) < n + 1 {
        resize(array)
    }
    array[n] <- x
    n <- n + 1
    bubble_up(n-1)
    return true
    
bubble_up(i)
    p <- parent(i)
    while i > 0 and array[i] < array[p] {
        array[i], array[p] <- array[p], array[i]
        i <- p
        p <- parent(i)
    }
```

The `bubble_up` function is called after the value `x` is added to the array, to ensure that the array maintains the heap-ordered property. 
