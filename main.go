package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"strconv"
)

func main() {
	decode := flag.Bool("d", false, "Decode the file")
	input_file := flag.String("i", "", "Input file path")
	output_file := flag.String("o", "", "Output file path")

	flag.Parse()

	content, err := os.ReadFile(*input_file)
	if err != nil {
		log.Fatal("No such file.")
	}

	content = bytes.TrimSpace(content)
	if *decode {
		run_length_decoding(content, *output_file)
	} else {
		run_length_encoding(content, *output_file)
	}
}

// Run Length Encoding
// ===================
// Runs through the file and finds repeating characters.
// The repeating characters are then reduced to the count
// of them followed by the character itself.
func run_length_encoding(file []byte, output string) {
	var encoded_file []byte
	if len(file) == 0 {
		return
	}

	currentChar := file[0]
	count := 1

	for i := 1; i < len(file); i++ {
		if file[i] == currentChar {
			count++
		} else {
			if count > 1 {
				encoded_file = append(encoded_file, []byte(strconv.Itoa(count))...)
				encoded_file = append(encoded_file, '|')
			}
			encoded_file = append(encoded_file, currentChar)
			currentChar = file[i]
			count = 1
		}
	}

	// Handle the last run
	if count > 1 {
		encoded_file = append(encoded_file, []byte(strconv.Itoa(count))...)
		encoded_file = append(encoded_file, '|')
	}
	encoded_file = append(encoded_file, currentChar)

	output_file, err := os.Create(output)
	if err != nil {
		log.Fatal("Couldn't create encoded file: ", err)
	}
	defer output_file.Close()

	_, err = output_file.Write(encoded_file)
	if err != nil {
		log.Fatal("Couldn't write encoded text to file: ", err)
	}
}

func run_length_decoding(file []byte, output string) {
	var decoded_file []byte
	count_string := ""
	is_count := false

	for i := 0; i < len(file); i++ {
		if file[i] == '|' {
			is_count = false
			count, err := strconv.Atoi(count_string)
			if err != nil {
				log.Fatal("Couldn't convert to int: ", count_string)
			}
			for l := 0; l < count-1; l++ {
				decoded_file = append(decoded_file, file[i+1])
			}
			count_string = ""
		} else if is_count || (file[i] >= '0' && file[i] <= '9' && (i == 0 || file[i-1] == '|')) {
			count_string += string(file[i])
			is_count = true
		} else {
			decoded_file = append(decoded_file, file[i])
		}
	}

	output_file, err := os.Create(output)
	if err != nil {
		log.Fatal("Couldn't create decoded file: ", err)
	}
	defer output_file.Close()

	_, err = output_file.Write(decoded_file)
	if err != nil {
		log.Fatal("Couldn't write decoded text to file: ", err)
	}
}

func byteIsDigit(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	} else {
		return false
	}
}
