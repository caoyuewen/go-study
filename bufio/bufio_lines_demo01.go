/**
 * Description: 
 * User: 1067
 * Date: 2018-09-14
 * Time: 16:41
 */

package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	var inputArr []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "exit" {
			break
		}
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		inputArr = append(inputArr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}


	fmt.Println(inputArr)

}
