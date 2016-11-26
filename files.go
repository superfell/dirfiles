package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// check each directory in the tree rooted at dir, and return one per line
// the directories that have at least one file matching the need regex
// and no files matching the missing regex.
func main() {
	dir := flag.String("d", ".", "directory to start processing from")
	need := flag.String("n", "cover\\.png", "regex of file to look for")
	missing := flag.String("m", "cover\\.jpg", "regex of file to be missing")

	flag.Parse()
	ds, err := os.Stat(*dir)
	if err != nil {
		log.Fatalf("Error verifiying supplied root directory: %v\n", err)
	}
	if !ds.IsDir() {
		log.Fatalf("Supplied root directory isn't a directory: %v\n", *dir)
	}
	needRegex, err := regexp.Compile(*need)
	if err != nil {
		log.Fatalf("Unable to compile regex for needed files: %v\n", err)
	}
	missingRegex, err := regexp.Compile(*missing)
	if err != nil {
		log.Fatalf("Unable to compile regex for missing files: %v\n", err)
	}
	filepath.Walk(*dir, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			files, err := ioutil.ReadDir(p)
			if err != nil {
				return err
			}
			hasNeed := false
			hasMissing := false
			for _, f := range files {
				hasNeed = hasNeed || needRegex.MatchString(f.Name())
				hasMissing = hasMissing || missingRegex.MatchString(f.Name())
				if hasMissing {
					break
				}
			}
			if hasNeed && !hasMissing {
				fmt.Printf("%s\n", p)
			}
		}
		return err
	})
}
