package algorithms

import (
	"fmt"
	"strconv"
)

// Run Length Encoding
// ===================
// Runs through the file and finds repeating characters.
// The repeating characters are then reduced to the count
// of them followed by the character itself.
func Run_length_encoding(input []byte) (output []byte, err error) {
	var encoded_file []byte
	if len(input) == 0 {
		return input, fmt.Errorf("File is empty")
	}

	currentChar := input[0]
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == currentChar {
			count++
		} else {
			if count > 1 {
				encoded_file = append(encoded_file, []byte(strconv.Itoa(count))...)
				encoded_file = append(encoded_file, '|')
			}
			encoded_file = append(encoded_file, currentChar)
			currentChar = input[i]
			count = 1
		}
	}

	// Handle the last run
	if count > 1 {
		encoded_file = append(encoded_file, []byte(strconv.Itoa(count))...)
		encoded_file = append(encoded_file, '|')
	}
	encoded_file = append(encoded_file, currentChar)

	return encoded_file, nil
}

func Run_length_decoding(input []byte) (output []byte, err error) {
	var decoded_file []byte
	if len(input) == 0 {
		return input, fmt.Errorf("File is empty")
	}

	count_string := ""
	is_count := false

	for i := 0; i < len(input); i++ {
		if input[i] == '|' {
			is_count = false
			count, err := strconv.Atoi(count_string)
			if err != nil {
				return input, fmt.Errorf("Couldn't convert to int: %s", count_string)
			}
			for l := 0; l < count-1; l++ {
				decoded_file = append(decoded_file, input[i+1])
			}
			count_string = ""
		} else if is_count || (input[i] >= '0' && input[i] <= '9' && (i == 0 || input[i-1] == '|')) {
			count_string += string(input[i])
			is_count = true
		} else {
			decoded_file = append(decoded_file, input[i])
		}
	}

	return decoded_file, nil
}
