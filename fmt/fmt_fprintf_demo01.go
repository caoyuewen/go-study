/**
 * Description: 
 * User: 1067
 * Date: 2018-09-14
 * Time: 17:59
 */

package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Fprintf(os.Stderr, "fp: an %s\n", "test")
	fmt.Printf("p: an %s\n", "test")

}
