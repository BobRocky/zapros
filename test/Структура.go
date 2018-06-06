// _range_ iterates over elements in a variety of data
// structures. Let's see how to use `range` with some
// of the data structures we've already learned.

package main

import "fmt"

func main() {

	// `range` on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "ghbdtn"}
	for k, v := range kvs {
		fmt.Printf("%s 5 %s\n", k, v)
		fmt.Println(v)
	}

	// `range` can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	for i, c := range "gog" {
		fmt.Println(i, c)
	}
}
