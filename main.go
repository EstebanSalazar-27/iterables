package main

import (
	"fmt"

	"github.com/estebansalazar-27/iterables/users"
)

func courseExercise1() {
	//  Create an array with a list of goals to this course
	arr := [3]string{"Learn in dive go", "Get a new work in go", "Learn another programming language to expand my knowledge and get a wide vision of programming"}
	// Create a slice based on the array with the second and third elements
	slice := arr[1:3]

	// Print the arr
	fmt.Println(arr)
	//  Print the first goal
	fmt.Println(arr[0], "First goal")
	// Print the slice
	fmt.Println(slice, "slice")
}
func courseExercise2() {
	// Create an array with a list of goals to this course
	arr := []string{"Learn in dive go", "Get a new work in go", "Learn another programming language to expand my knowledge and get a wider vision of programming"}
	// Modify the second element in the array
	arr[1] = "Modified goal"
	//  Add to the array a new element
	arr = append(arr, "This is a new goal")
	//  Print the array
	fmt.Println(arr)

}
func main() {
	err := users.SaveUsersToFile()

	if err != nil {
		fmt.Println("Error al guardar usuarios:", err)
		return
	}

	usersCollection, err := users.GetUsers()
	//  A way to slice an array is with the notation [index where start to slice up : index where finish slice up]
	//  [0, 1, 2, 3] if i'd like to slice up the array omitting the rest of elements forward of an specific position it would be such as [0:3]
	// Here we're getting a slice of our original array from the position 0 to 3 therefore the len of sliceCollection is 3 but the capacity is 4 because go internally memorize all the other elements to the right of the selected slice
	//  instead if we set the first slice value as 1 that value is gonna be lost in our slice so the capacity will be 3
	slicedCollection := usersCollection[:3]
	slicedCollection[0] = users.User{
		Name: "pepe",
	}
	if err != nil {
		fmt.Println("Error al obtener usuarios:", err)
		return
	}
	fmt.Println(slicedCollection, "Sliced users collection")
	fmt.Println(usersCollection, "Getting users collection from json")
	fmt.Println(len(slicedCollection), cap(slicedCollection))

	slicedCollection = slicedCollection[:cap(slicedCollection)]
	fmt.Println(slicedCollection, "Sliced users collection")

	//  Exercises calls
	courseExercise1()
	courseExercise2()
}
