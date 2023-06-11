package utils

// Wraps a string in ANSI that makes it red
func Red(s string) string {
	return "\033[31m" + s + "\033[0m"
}
