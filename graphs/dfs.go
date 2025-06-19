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


func dfsIterative(graph map[string][]string, src, dst string) bool {
	stack := []string{src}
	visited := map[string]bool{
		src: true,
	}

	for len(stack) > 0 {
		current := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if current == dst {
			return true
		}
		
		for _, neighbor := range graph[current] {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = true
				stack = append(stack, neighbor)
			}
		}
		
	}

	return false
}