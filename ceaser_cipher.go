package crypto_tasks

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func DecryptCeaser(message string, offset int) string {
	result := make([]byte, len(message))

	for i, c := range message {
		r := int(c) - 'a'
		r = (r+offset)%26 + 'a'
		result[i] = byte(r)
	}
	return string(result)
}

func DecryptVigenere(message, keyword string) string {
	result := make([]byte, len(message))

	for i, c := range message {
		k := int32(keyword[i%len(keyword)])
		d := (c-k+26)%26 + 'a'
		result[i] = byte(d)
	}
	return string(result)
}

func DecryptPolybius(message string) (string, error) {
	table := [][]rune{
		{'a', 'b', 'c', 'd', 'e'},
		{'f', 'g', 'h', 'i', 'j'},
		{'k', 'l', 'm', 'n', 'o'},
		{'p', 'q', 'r', 's', 't'},
		{'u', 'v', 'w', 'x', 'y'},
		{'z', ' ', ' ', ' ', ' '},
	}

	if len(message)%2 != 0 {
		return "", errors.New("len of the message is not even")
	}
	result := make([]byte, len(message)/2)

	for i := 0; i < len(message); i += 2 {
		row := message[i] - '0' - 1
		col := message[i+1] - '0' - 1

		result[i/2] = byte(table[col][row])
	}

	return string(result), nil
}

func FrequencyDictionary(message []string) map[string]int {
	result := make(map[string]int)
	for _, letter := range message {
		result[letter] += 1
	}
	return result
}

func Normalize(text string) (string, error) {
	text = strings.ToLower(text)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	text = reg.ReplaceAllString(text, "")
	return text, nil
}

/* -- Extra --------------------------------------------------------------------------------------------------------- */

func VowelConsonantsBinarization(message string) string {
	vowel := map[rune]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'y': {}}
	result := make([]rune, len(message))

	for i, c := range message {
		if _, ok := vowel[c]; ok {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}

	return string(result)
}

func DecryptPolybiusSpiral(message string) (string, error) {
	table := [][]rune{
		/*0*/ /*1*/ /*2*/ /*3*/ /*4*/ /*5*/
		/*0*/ {'f', 'g', 'h', 'i', 'j', 'k'},
		/*1*/ {'e', 'x', 'y', 'z', '0', 'l'},
		/*2*/ {'d', 'w', '7', '8', '1', 'm'},
		/*3*/ {'c', 'v', '6', '9', '2', 'n'},
		/*4*/ {'b', 'u', '5', '4', '3', 'o'},
		/*5*/ {'a', 't', 's', 'r', 'q', 'p'},
	}

	if len(message)%2 != 0 {
		return "", errors.New("len of the message is not even")
	}
	result := make([]byte, len(message)/2)

	for i := 0; i < len(message); i += 2 {
		row := message[i] - '0'
		col := message[i+1] - '0'
		result[i/2] = byte(table[row][col])
	}

	return string(result), nil
}

func EncryptPolybius(charTable [][]rune, message string, zeroIdx bool) string {
	flattened := make(map[rune]string)

	for i, row := range charTable {
		for j, char := range row {
			if zeroIdx {
				flattened[char] = strconv.Itoa(i) + strconv.Itoa(j)
			} else {
				flattened[char] = strconv.Itoa(i+1) + strconv.Itoa(j+1)
			}
		}
	}
	result := ""
	for _, char := range message {
		result += flattened[char]
	}
	return result
}

func StrXOR(first, second string) string {
	var min int
	if len(first) < len(second) {
		min = len(first)
	} else {
		min = len(second)
	}

	result := make([]rune, min)
	for i := 0; i < min; i++ {
		if first[i] == second[i] {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return string(result)
}
