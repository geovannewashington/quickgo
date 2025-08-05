//go:build !debug
// +build !debug

// compiled when not using -tags=debug
package main

func logDebug(msg string) {
	// No-op in release build
}
