package intermediate

import (
	"fmt"
	"path/filepath"
)

func FilePathDemo() {
	
	// Joining
	joinedPath := filepath.Join("downloads", "file.txt")
	fmt.Println("Joined Path:", joinedPath)

	// Cleaning up redundant path
	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath)

	// Seperating the file and the directory
	dir, file := filepath.Split("home/user/projects/file.txt")
	fmt.Println("File:", file)
	fmt.Println("Path:", dir)

	// Print the extension of the file
	fmt.Println(filepath.Ext(file))

	// Find relative path from one file to another
	rel, err := filepath.Rel("a/b/c", "a/b/d/e/file")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(rel)
}