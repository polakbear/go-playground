package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	dir := "./sample"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	count := 0
	var toRename []string
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("dir: ", file.Name())
		} else {
			_, err := match(file.Name(), 4)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}

	for _, origFilename := range toRename {
		origPath := filepath.Join(dir, origFilename)
		newFilename, err := match(origFilename, count)
		if err != nil {
			panic(err)
		}
		newPath := filepath.Join(dir, newFilename)
		err = os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("mv %s => %s\n", origPath, newPath)
	}
}

// match returns the new filename
func match(filename string, total int) (string, error) {
	// "birthday", "001", "txt"
	pieces := strings.Split(filename, ".")
	extension := pieces[len(pieces)-1]
	temp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(temp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern", filename)
	}

	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, extension), nil
}
