package main

import "fmt"

// CountComponents returns how many connected components the undirected
// graph has: every DFS launched from a yet-unvisited node discovers exactly
// one component.
func CountComponents(n int, edges [][2]int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	var walk func(v int)
	walk = func(v int) {
		visited[v] = true
		for _, w := range adj[v] {
			if !visited[w] {
				walk(w)
			}
		}
	}
	components := 0
	for v := 0; v < n; v++ {
		if !visited[v] {
			components++
			walk(v)
		}
	}
	return components
}

func main() {
	fmt.Println(CountComponents(5, [][2]int{{0, 1}, {1, 2}, {3, 4}}))
}
