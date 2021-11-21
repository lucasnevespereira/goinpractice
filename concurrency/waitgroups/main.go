package main

import (
	"compress/gzip"
	"io"
	"os"
)

const (
	gzFileExtension = ".gz"
)

func main() {
	// collects a list of files passed in the command line
	for _, file := range os.Args[1:] {
		compress(file)
	}
}

func compress(filename string) error {
	// opens the source file for reading
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	// opens a destination file with a .gz extension added to the source file name
	out, err := os.Create(filename + gzFileExtension)
	if err != nil {
		return err
	}
	defer out.Close()

	// gzip writer compresses data ans then writes it to the underlying file.
	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in) // does all the copying for us
	gzout.Close()

	return err
}
