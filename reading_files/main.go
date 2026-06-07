package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// check is a small helper function.
//
// Instead of repeatedly writing:
//
//	if err != nil {
//	    panic(err)
//	}
//
// we centralize error handling here.
//
// In production code, we usually return errors gracefully.
// In learning examples, panic keeps the code shorter.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	//
	// Getwd() returns the directory from which
	// the program is currently running.
	//
	// Example:
	// C:\Users\91969\vslf\go
	//
	// Useful when building relative file paths.
	cwd, err := os.Getwd()
	check(err)

	fmt.Println("=== Current Working Directory ===")
	fmt.Println(cwd)
	fmt.Println()

	//
	// TempDir() returns the OS temporary directory.
	//
	// Windows:
	// C:\Users\<user>\AppData\Local\Temp
	//
	// Linux:
	// /tmp
	//
	// Useful for:
	// - temporary files
	// - caches
	// - intermediate processing
	tempDir := os.TempDir()

	fmt.Println("=== Temporary Directory ===")
	fmt.Println(tempDir)
	fmt.Println()

	//
	// Never manually concatenate paths:
	//
	// "folder/file.txt"
	//
	// because separators differ across OSes.
	//
	// filepath.Join automatically chooses:
	//
	// Windows -> \
	// Linux   -> /
	samplePath := filepath.Join(
		"reading_files",
		"sample.txt",
	)

	fmt.Println("=== Sample File Path ===")
	fmt.Println(samplePath)
	fmt.Println()

	//
	// ReadFile:
	// - opens file
	// - reads entire content
	// - closes file
	//
	// Returns:
	// []byte
	//
	// Best for small files.
	data, err := os.ReadFile(samplePath)
	check(err)

	fmt.Println("=== Read Entire File ===")
	fmt.Println(string(data))
	fmt.Println()

	//
	// Stat returns information about a file:
	//
	// Name
	// Size
	// Last Modified Time
	// Whether it's a directory
	info, err := os.Stat(samplePath)
	check(err)

	fmt.Println("=== File Information ===")
	fmt.Println("Name:", info.Name())
	fmt.Println("Size:", info.Size(), "bytes")
	fmt.Println("Is Directory:", info.IsDir())
	fmt.Println("Modified:", info.ModTime())
	fmt.Println()

	//
	// ReadDir lists everything inside a directory.
	//
	// Similar to:
	//
	// ls      (Linux)
	// dir     (Windows)
	entries, err := os.ReadDir("reading_files")
	check(err)

	fmt.Println("=== Directory Contents ===")

	for _, entry := range entries {

		if entry.IsDir() {
			fmt.Printf("[DIR ] %s\n", entry.Name())
		} else {
			fmt.Printf("[FILE] %s\n", entry.Name())
		}
	}

	fmt.Println()

	//
	// Here we create a file inside the
	// operating system's temp directory.
	tempFile := filepath.Join(
		tempDir,
		"basalt-temp.txt",
	)

	err = os.WriteFile(
		tempFile,
		[]byte("Hello from Basalt!"),
		0644,
	)
	check(err)

	fmt.Println("=== Temporary File Created ===")
	fmt.Println(tempFile)
	fmt.Println()

	//
	// Verify the contents we just wrote.
	tempData, err := os.ReadFile(tempFile)
	check(err)

	fmt.Println("=== Temporary File Content ===")
	fmt.Println(string(tempData))
	fmt.Println()

	//
	// Remove deletes a file.
	//
	// Good practice for temporary files.
	err = os.Remove(tempFile)
	check(err)

	fmt.Println("=== Temporary File Deleted ===")
	fmt.Println(tempFile)
}