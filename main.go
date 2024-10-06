package main

import (
	"bytes"
	"flag"
	"log"
	"os"

	"github.com/cakePhone/compressing-algorithm/algorithms"
)

func main() {
	decode := flag.Bool("d", false, "Decode the file")
	input_file_name := flag.String("i", "", "Input file path")
	output_file_name := flag.String("o", "", "Output file path")

	flag.Parse()

	content, err := os.ReadFile(*input_file_name)
	if err != nil {
		log.Fatal("No such file.")
	}

	output_file, err := os.Create(*output_file_name)
	if err != nil {
		log.Fatal("For some reason, the OS refused to create a file to output.")
	}
	defer output_file.Close()

	output_string := content

	content = bytes.TrimSpace(content)
	if *decode {
		output_string, err = algorithms.Run_length_decoding(content)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		output_string, err = algorithms.Run_length_encoding(content)
		if err != nil {
			log.Fatal(err)
		}

		output_string, _ = algorithms.Huffman_encoding(output_string)
	}

	output_file.Write(output_string)
}

func byteIsDigit(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	} else {
		return false
	}
}
