A doubly linked list is like a [linked list](https://github.com/heyitschun/cs-essentials/tree/master/data-structures/linked-lists), except that each node has two links. One link points to the next node, the other to the previous node. It also tracks the memory location of the first *and* last node.

![Doubly linked list](https://upload.wikimedia.org/wikipedia/commons/5/5e/Doubly-linked-list.svg)

Since doubly linked lists can access the front and end points so easily, they are ideal to build a [queue](https://github.com/heyitschun/cs-essentials/tree/master/data-structures/queues) with. Queues only give access to the first and last elements of the queue and doubly linked lists can reach both of these points in `O(1)`.

Nodes in a double linked list store a data value and a reference to the next node in the sequence:

```text
class Node:
    val
    nxt
    pre
```

# Read more

* [Open Data Structures](https://opendatastructures.org/ods-cpp/3_2_Doubly_Linked_List.html)
* [Wikipedia](https://en.wikipedia.org/wiki/Doubly_linked_list)
