A **queue** is an abstract data type (ADT) that holds an ordered sequence of elements. Access to the queue is restricted to the front and back of the queue. *Entering* the queue is done at the back and *leaving* the queue is done from the front; therefore, a queue is an ADT that follows the first-in-first-out protocol.

![Queue](https://upload.wikimedia.org/wikipedia/commons/5/52/Data_Queue.svg)

# Common methods

- `enqueue(x)`--Add element `x` to the queue.
- `dequeue()`--Remove the element at the front of the queue, provided that the queue is not empty.
- `empty()`--Check if the queue contains any elements.
- `read()`--Return the front element of the queue without removing it, provided that the queue is not empty.
