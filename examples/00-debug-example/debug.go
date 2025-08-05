//go:build debug
// +build debug

// only compiled with -tags=debug
package main

func logDebug(msg string) {
	println("[DEBUG]", msg)
}
