/*
Package greeting provides utilities
for printing greeting strings
*/
package greeting

const testVersion = 3

// Returns a greeting string for the supplied name.
// If no name is supplied, greets the world.
func HelloWorld(name string) string {
	if len(name) > 0 {
		return "Hello, " + name + "!"
	}

	return "Hello, World!"
}
