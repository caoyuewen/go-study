/**
 * Description: 
 * User: 1067
 * Date: 2018-09-14
 * Time: 16:50
 */

package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main() {
	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"

	scanner := bufio.NewScanner(strings.NewReader(input))

	// Set the split function for the scanning operation.
	//scanner.Split(bufio.ScanWords)//注释这行代码  默认会以\n拆分

	//count the words
	var count int
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr,"reading input:",err)
	}

	fmt.Printf("%d\n",count)
}
