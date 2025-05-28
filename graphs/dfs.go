package graphs

func dfs(graph map[string][]string, src, dst string, visited map[string]bool) bool {
	if src == dst {
		return true
	}

	if seen, ok := visited[src]; ok && seen {
		return false
	}

	visited[src] = true
	for _, neighbor := range graph[src] {
		if dfs(graph, neighbor, dst, visited) {
			return true
		}
	}

	return false
}

func HasPath(graph map[string][]string, src, dst string) bool {
	return dfs(graph, src, dst, make(map[string]bool))
}
