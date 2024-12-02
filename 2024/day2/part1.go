package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
	"strconv"
)

func main() {

    filePath := os.Args[1]
    readFile, err := os.Open(filePath)

    if err != nil {
        fmt.Println(err)
    }

	var lines [][]int

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
	var text_line string

    for fileScanner.Scan() {
		text_line = fileScanner.Text()
		tmp := strings.Split(text_line, " ")
		values := make([]int, 0, len(tmp))
		for _, raw := range tmp {
			v, _ := strconv.Atoi(raw)
			values = append(values, v)
		}
		lines = append(lines, values)
    }

    readFile.Close()

	var safe int = 0
	
	for _, v := range lines {
		var is_safe bool = true
		var ascending int = 0
		var current int
		for j, v2 := range v {
			if j == 0 {
				current = v2
			} else {
				if ascending == 0 {
					if v2 > current {
						ascending = 1
					} else if v2 < current {
						ascending = -1
					}
				} else if ascending == 1 {
					if v2 < current {
						is_safe = false
					}
				} else {
					if v2 > current {
						is_safe = false
					}
				}
				if v2 == current {
					is_safe = false
				} else if ! (v2 - current <= 3 && v2 - current >= -3) {
					is_safe = false
				}
				if ! is_safe {
					break
				}
				current = v2
			}
		}
		if is_safe {
			safe = safe + 1
		}
	}
	
	fmt.Println(safe)
}