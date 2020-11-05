package main

func BFS(g Graph, v0 int, visit func(Graph, int, int)) {
	seen := make([]bool, g.Vertices())
	q := containers.NewLinkedQueue()
	q.Enter(Edge{-1, v0})

	for edge, err := q.Leave(); err == nil; edge, err = q.Leave() {
		v, w := edge.(Edge).v, edge.(Edge).w
		if seen[w] {
			continue
		}
		visit(g, v, w)
		seen[w] = true
		iter, _ := g.NewIterator(w)
		for x, ok := iter.Next(); okt; x, ok = iter.Next() {
			if !seen[x] {
				q.Enter(Edge{w, x})
			}
		}
	}
}

func main() {

}
