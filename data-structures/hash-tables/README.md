Hashing transforms a key into an array index, thereby providing instantaneous access to the value stored in an array holding the key-value pairs in the map. This array is called a hash table.

A **hash function** transforms a key into a value in the range of indices of a hash table.

A **hash table** holds the key-value pairs whose locations are determined by a hash function applied to keys.

# Hashing

For a small set of key-value pairs, it is usually possible to allocate an array big enough to hold all items. But a small table holding keys with a large range of values will require that the function transform several keys to the same table location. When two or more keys are mapped to the same table location by a hash function is called a *collision*.

An implementation of hashing requires two things:

* A hash function to transform keys to hash table locations.
* A collision resolution scheme to deal with the collisions that will inevitably occur.

## Hash functions

A hash function must transform keys into integers in the range of `0` to `t-1`, where `t` is the size of the hash table. The best hash functions use the division method. For numeric values this is:

`hash(k) = k % t`

For this method to work at its best, `t` should be a prime number that is not close to a power of two.

For non-numeric keys, the key could be converted to a number before the division method is applied to it.

## Collision resolution schemes

An important value to consider with collision resolution schemes is the *load factor*. This is calculated as `n/t`, where `n` is the number of items in the table and `t` is the size of the table.

**Chaining** is the first approach to resolve a collision. It uses linked lists of key-value pairs that start in hash table locations. It is a very robust approach and has good performance, but requires extra memory because of the linked list nodes.

**Open-addressing** probes sequences to look through the table for an open spot to store a key-value pair in. It is space-efficient, but its performance degrades quickly as the load factor approaches one.

# Hash Tables

A hash table stores key-value pairs by the hash value of the key. Hash table keys must have a `hash()` function returning integer values and an equality comparison operation to test for collision.

Other than that, for a hash table to work, it must implement some methods to perform mutations on the table and its values, such as `set`, `get` and `delete`.

# References

* *Go Data Structures and Algorithms* by Christopher Fox (2018)
