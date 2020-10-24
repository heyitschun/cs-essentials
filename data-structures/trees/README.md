Trees are the basis for several important data types and structures. Several sorting algorithms are based on trees. So it is important to understand the basics of them.

A **tree** is a special type of graph. It is a graph with a distinguished node root `r`, such that there is exactly one simple path between each node in the tree and `r`.

![Tree](https://computersciencewiki.org/images/5/5f/Binary_tree.svg.png)

In the tree above, the root `r` is labeled as `2`. It has two child nodes labeled `7` and `5` respectively. Each of these have their own children and so on.

# Terminology

- *Children*--Immediate descendants of a node.
- *Siblings*--Nodes that have the same parent node.
- *Descendants*--All the nodes of which a node is an ancestor.
- *Terminal node*--A node without children. Also known as a terminal leaf.
- *Internal node*--A node with children. Also known as a non-terminal leaf or node.
- *Sub-tree*--A graph consisting of a node in a tree, all its descendants and the edges connecting them.
- *Forest*--A graph consisting of several trees.
- *Level*--Represents the number of edges in the path from the node to the root.
- *Height*--The maximum level in the tree.

Tree can be ordered. This means that the order of the children of each node is specific. These trees are not usually drawn in a distinct way. Other mechanisms must be used to specify whether a tree is ordered.

# Binary trees

When it comes to designing and creating data structures, binary trees are especially important.

A **binary tree** is an *ordered* tree whose nodes have at most two children. Binary trees are called **complete** when every level has two child nodes--except the last node, this node must have a *left child* and *no right child*.

Some important properties of trees:

- A tree with `n` nodes has `n-1` edges.
- A complete tree with `n` *internal* nodes has either `n` or `n+1` leaves.
- The height of a binary tree with `n` nodes is `log n`.

These properties are handy to have memorized when it comes to determining the runtime of a function that uses this data structure.
