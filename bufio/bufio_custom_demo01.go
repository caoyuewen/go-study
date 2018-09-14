/**
 * Description: 
 * User: 1067
 * Date: 2018-09-14
 * Time: 16:30
 */

package main

import (
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	// An artificial input source.人工输入源
	const input = "1234 5678 1234567901234567890"

	scanner := bufio.NewScanner(strings.NewReader(input))

	//Create a custom split function by wrapping the existing ScanWords function.
	//通过包装现有的ScanWords函数来创建一个自定义拆分函数。
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)

		if err != nil && token != nil {
			strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	scanner.Split(split)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
