package mylib // lib name to export 

// Exported function (starts with uppercase)
func Add(a, b int) int {
	return a + b
}

// Unexported function (starts with lowercase - cannot be used outside)
func subtract(a, b int) int {
	return a - b
}


