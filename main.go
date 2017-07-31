package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var file *os.File
	var err error

	if len(os.Args) > 1 {
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file = os.Stdin
	}
	defer file.Close()

	hash, err := txt2json(file)
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(hash, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func txt2json(file io.Reader) (map[string]interface{}, error) {

	reSectionKey := regexp.MustCompile("^# (.*) #+$")
	reKeyValue := regexp.MustCompile("^\\s*(.*)\\s*\\|\\s*(.*)\\s*$")

	var key string
	hash := make(map[string]interface{})

	scanner := bufio.NewScanner(file)

	tmpSlice := []string{}
	tmpHash := make(map[string]interface{})

	for scanner.Scan() {
		m := reSectionKey.FindStringSubmatch(scanner.Text())
		if len(m) > 0 {
			if key != "" {
				if len(tmpSlice) > 0 {
					hash[key] = tmpSlice
				} else if len(tmpHash) > 0 {
					hash[key] = tmpHash
				}
			}
			key = m[1]

			tmpSlice = []string{}
			tmpHash = make(map[string]interface{})
			continue
		}

		m = reKeyValue.FindStringSubmatch(scanner.Text())
		if len(m) > 0 && key != "" {
			itemKey := m[1]
			itemVal := m[2]
			if itemKey == "" || strings.HasPrefix(itemKey, "-") {
				continue
			}
			tmpHash[itemKey] = itemVal
			continue
		}

		tmpSlice = append(tmpSlice, scanner.Text())
	}
	if key != "" {
		if len(tmpSlice) > 0 {
			hash[key] = tmpSlice
		} else if len(tmpHash) > 0 {
			hash[key] = tmpHash
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hash, nil
}
