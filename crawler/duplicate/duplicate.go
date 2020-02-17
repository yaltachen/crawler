package duplicate

var visited map[string]bool

func init() {
	visited = make(map[string]bool)
}

// IsDuplicate check duplicate
func IsDuplicate(key string) bool {
	if _, exists := visited[key]; exists {
		return true
	}
	visited[key] = true
	return false
}
