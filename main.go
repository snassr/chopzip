package main

import (
	"archive/zip"
	"flag"
	"io"
	"log"
	"os"
)

var (
	// cmd flag vars
	in     string
	out    string
	format string
)

func init() {
	flag.StringVar(&in, "in", "", "input file")
	flag.StringVar(&out, "out", "", "output file")
}

func main() {
	// parse flags
	flag.Parse()

	// check flag values // TODO: use unicode isSpace?
	if in == "" || out == "" {
		log.Fatal("in or out file was not specified")
	}

	// print flag values
	log.Printf("in: %v\n", in)
	log.Printf("out: %v\n", out)

	// open input file
	infile, err := os.Open(in)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer infile.Close()

	// create output file
	outFile, err := os.Create(out)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer outFile.Close()

	// Create a new zip archive
	w := zip.NewWriter(outFile)

	// create file in archive
	f, err := w.Create("zippedfile")
	if err != nil {
		log.Fatal(err)
	}

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := infile.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := f.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	defer w.Close()
}
