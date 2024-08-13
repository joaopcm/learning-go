package main

import "fmt"

var moviesFromDb = []string{
	"The Shawshank Redemption",
	"The Godfather",
	"The Dark Knight",
	"The Lord of the Rings",
	"Pulp Fiction",
	"The Matrix",
	"The Silence of the Lambs",
	"The Departed",
	"The Lord of the Rings",
	"The Hobbit",
	"The Fellowship of the Ring",
	"The Two Towers",
	"The Return of the King",
}

func main() {
	// Creating a slice from an array
	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[1:4]
	// arr[2] = 15
	// slice[0] = 123
	// fmt.Println(arr, slice)

	// Creating a slice from scratch
	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println(slice)
	// fmt.Println(len(slice), cap(slice))

	// Nil slice
	// var slice []int
	// fmt.Println(slice, slice == nil)

	// Empty slice
	// slice := []int{}
	// fmt.Println(slice, slice == nil)

	// Loop and append a slice
	// resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// movies := make([]string, 0, len(resultsFromApi))
	// for _, id := range resultsFromApi {
	// 	movie := moviesFromDb[id]
	// 	fmt.Println(len(movies), cap(movies))
	// 	movies = append(movies, movie)
	// 	fmt.Println(len(movies), cap(movies))
	// }
	// fmt.Println(movies)

	// Full Slice Expression
	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[:2:2]
	// fmt.Println(slice, cap(slice))

	// Out of bounds
	// slice := []int{1, 2, 3, 4}
	// fmt.Println(slice[10])

	// Passed by value or reference?
	// slice := []int{1, 2, 3, 4}
	// foo(slice)
	// fmt.Println(slice)

	// Maps
	// m := make(map[string]string)
	// m["João"] = "Melo"
	// value, ok := m["João"]
	// fmt.Println(value, ok)
	// delete(m, "João")
	// value, ok = m["João"]
	// fmt.Println(value, ok)

	m := map[string]string{
		"Pedro": "Pessoa",
		"João":  "Melo",
	}
	// clear(m)
	// fmt.Println(m)

	// 	f := math.NaN()
	// 	f2 := math.NaN()
	// 	m := map[float64]string{
	// 		f:  "João",
	// 		f2: "Melo",
	// 	}
	// 	value, ok := m[f2]
	// 	fmt.Println(m, value, ok)
	// 	delete(m, f2)
	// 	value, ok = m[f2]
	// 	fmt.Println(m, value, ok)
	// 	clear(m)
	// 	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}

// func foo(slice []int) {
// 	slice[0] = 123
// }
