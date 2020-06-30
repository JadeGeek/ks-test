package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	 * stackoverflow.com › questions › how-to-join-a-slice-of-strings-
	 * into-a-...How to join a slice of strings into a single string?
	 *
	 * Use a slice, not an arrray. Just create it using
	 */

	reg := []string{"a", "b", "c"}

	/*
	 * An alternative would have been to convert your array to a slice when
	 * joining :
	 */

	fmt.Println(strings.Join(reg[:], ","))

	/*
	 * Read the Go blog about the differences between slices and arrays
	 * (http:blog.golang.org/go-slices-usage-and-internals).
	 *
	 * [Denys S&#233;guret] [so/q/28799110] [cc by-sa 3.0]
	 */

}
