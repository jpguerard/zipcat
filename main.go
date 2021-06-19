package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

const zipcat = "zipcat"

func cat(file string, buffer []byte) error {
	z, err := zip.OpenReader(file)
	if err != nil {
		return fmt.Errorf("%s: error opening archive %s: %w", zipcat, file, err)
	}
	defer z.Close()
	for _, f := range z.File {
		c, err := f.Open()
		if err != nil {
			return fmt.Errorf("%s: error opening %s in archive %s: %w", zipcat, f.Name, file, err)
		}
		_, err = io.CopyBuffer(os.Stdout, c, buffer)
		if err != nil {
			c.Close()
			return fmt.Errorf("%s: error reading %s in archive %s: %w", zipcat, f.Name, file, err)
		}
		c.Close()
	}
	return nil
}

func main() {
	buffer := make([]byte, 4194304)
	for _, file := range os.Args[1:] {
		err := cat(file, buffer)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
