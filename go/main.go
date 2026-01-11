
package main

import (
	"fmt"
	"os"
	"github.com/fsnotify/fsnotify"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("There was an error creating the watcher: %v\n", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add("./")
	if err != nil {
		fmt.Printf("There was an error adding the directory to the watcher: %v\n", err)
		return
	}

	fmt.Println("Watching for new files...")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				fmt.Printf("New file created: %s\n", event.Name)
				// Call the handleNewFile function here
				handleNewFile(event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Printf("Error watching: %v\n", err)
		}
	}
}

func handleNewFile(filePath string) {
	// Extract the file extension from the file name using `path.Ext()`
	ext := filepath.Ext(filePath)
	ext = strings.TrimPrefix(ext, ".")

	// Check if the directory for the file extension exists
	dir := filepath.Dir(filePath)
	if !filepath.HasPrefix(dir, ext) {
		// If the directory does not exist, create it using `os.MkdirAll()`
		err := os.MkdirAll(ext, 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
	}

	// Move the file to the directory using `os.Rename()`
	err := os.Rename(filePath, filepath.Join(ext, filepath.Base(filePath)))
	if err != nil {
		fmt.Printf("Error moving file: %v\n", err)
		return
	}

	// Call the sorter function from the sorter.py script
	sortFile(filepath.Join(ext, filepath.Base(filePath)))
}

func sortFile(filePath string) {
	// Call the sorting logic from the sorter.py script
	// You may need to import the sorter.py script and use its functions directly
	// For example:
	// import "sorter"
	// sorter.SortFile(filePath)
	// Alternatively, you can execute the sorter.py script as a subprocess
	cmd := exec.Command("python", "sorter.py", filePath)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing sorter.py: %v\n", err)
		return
	}
}