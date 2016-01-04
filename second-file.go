// this second file is not a different package: it belongs to the same main package, so we repeat `package main`
package main

// as this file belongs to the main package, the hello function is accessible
func hello() string {
	return "hello"
}
