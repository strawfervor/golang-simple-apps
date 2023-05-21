/*
This program is go implementation of ls
*/

package main

import (
	"fmt"
	"os"
	"time"
)

func CurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	} else {
		return dir
	}
}

func ListDir() {
	files, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, e := range files {
			fmt.Println(e.Type(), " ", e.Name())
		}
	}
}

func ListDirSingleFile() {
	files, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, e := range files {
			info, err := os.Stat(e.Name())
			if err != nil {
				fmt.Println("Error occured")
			} else {
				fmt.Println(info.Mode().Type(), "\t", info.Size(), "\t", (info.ModTime().Format(time.RFC822)), "\t", info.Name())
			}
		}
	}
}

func main() {
	fmt.Println(CurrentDir())
	ListDirSingleFile()
}
