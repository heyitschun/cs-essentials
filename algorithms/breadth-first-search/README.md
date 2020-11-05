The **breadth-first** search algorithm starts at a node `n` and looks at neighbor `n+1`, then at neighbor `n+2`, and so on.

This algorithms uses a queue `q`, that initially contains only node `n`. It then repeatedly extracts an element from `q` and adds its neighbors to `q` (assuming these neighbors have not been in `q` before). Remember that a tree is a graph where each node `n` has one simple path leading back to the root node `r`, so in the breadth-first search algorithm there is no need to add the same node `n` to `q` more than once. Tracking seen nodes, is done by using an auxiliary array of boolean values.

```
breadth_first_search(g, r)
    seen <- NewArray(n)
    q <- SLL()
    q.add(r)
    seen[r] <- true
    while q.size() > 0 {
        i <- q.remove()
        for j in g.outer_edges(i) {
            if seen[j] == false {
                q.add(j)
                seen[j] <- true
            }
        }
    }
```
