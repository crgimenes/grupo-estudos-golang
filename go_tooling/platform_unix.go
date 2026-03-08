//go:build !windows

package main

func platformName() string {
	return "unix"
}
