package idiomatic

import "fmt"

// create a new type alias
type AMap map[int]string

/*

Importants !!
From docs: https://blog.golang.org/maps
Concurrency
Maps are not safe for concurrent use: it's not defined what happens when you read and write to them simultaneously.
If you need to read from and write to a map from concurrently executing goroutines, the accesses must be mediated by some kind of synchronization mechanism.
One common way to protect maps is with sync.RWMutex.
*/

func MapPlay() AMap {
	// define a map
	var map1 AMap

	// instanciate map
	map1 = make(AMap)
	fmt.Println("instantiated map", map1)
	// implicit map init
	map2 := make(AMap)
	fmt.Println("implicit map ", map2)

	// map literal
	map3 := AMap{
		1: "one",
		2: "two",
	}
	fmt.Println("literal map", map3)

	return map3
}

func (m AMap) addItem(key int, value string) {
	m[key] = value
}

func (m AMap) deleteItem(key int) {
	delete(m, key)
}

// reference here is needed due to copy by value and  map has header*.
// with no use of pointer new assignment only lives in method context.
func (m *AMap) purge() {
	*m = make(AMap)
}
