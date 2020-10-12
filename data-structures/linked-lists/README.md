Linked lists are collections of data elements where one element points to the next element.

![Singly linked list](https://upload.wikimedia.org/wikipedia/commons/6/6d/Singly-linked-list.svg)

The structure of a linked list consists of nodes, which collectively represent a sequence. A linked list can be ordered or unordered and contain unique values or duplicates. There are no restrictions in that regard, as long as the nodes contain a value and a reference (a.k.a. link) to the next node in the sequence. This structure allows for efficient insertion and removal of elements from any position *during iteration*. Access time to element `n`, has a time complexity of `O(n)`, which is a big drawback if speed is required to access any arbitrary element without being in the node before it.

Nodes in a singly linked list (SLL) store a data value and a reference to the next node in the sequence:

```python
class Node:
    val
    nxt
```

The SLL itself could be represented in a class object or interface that implement functions on the elements.

# Read more

* [Open Data Structures](https://opendatastructures.org/ods-java/3_1_SLList_Singly_Linked_Li.html)
* [Wikipedia](https://en.wikipedia.org/wiki/Linked_list)
