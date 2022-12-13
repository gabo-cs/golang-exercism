// Package twofer implements a simple two-fer message
package twofer

import "fmt"

// ShareWith prints the two-fer message
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
