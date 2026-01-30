package main

import (
	"fmt"
	"math"
)

func dijkstra(graph [][]int, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	for i := 0; i < n-1; i++ {
		u := minDistance(dist, visited)
		visited[u] = true

		for v := 0; v < n; v++ {
			if !visited[v] && graph[u][v] != 0 &&
				dist[u] != math.MaxInt32 &&
				dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}
	return dist
}

func minDistance(dist []int, visited []bool) int {
	min := math.MaxInt32
	minIndex := -1

	for i := 0; i < len(dist); i++ {
		if !visited[i] && dist[i] <= min {
			min = dist[i]
			minIndex = i
		}
	}
	return minIndex
}

func main() {
	graph := [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}	

	start := 0
	distances := dijkstra(graph, start)

	fmt.Println("Jarak terpendek dari node", start)
	for i, d := range distances {
		fmt.Printf("Ke node %d = %d\n", i, d)
	}
}
