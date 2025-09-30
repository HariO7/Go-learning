package main

import (
	"fmt"
)

func main() {
	hobbies := [3]string{"Football", "Chess", "video Gaming"}

	//1
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	fmt.Println()

	//2
	hobbySlice := hobbies[:2]
	fmt.Println(hobbySlice)
	fmt.Println()

	//3
	hobbySlice = hobbySlice[1:3]
	fmt.Println(hobbySlice)
	fmt.Println()

	dynamicHobbies := []string{"learnGo", "complete project"}
	fmt.Println(dynamicHobbies)
	fmt.Println()

	dynamicHobbies[1] = "create own project"
	dynamicHobbies = append(dynamicHobbies, "advanceGo")
	fmt.Println(dynamicHobbies)
	fmt.Println()

	type Product struct {
		title string
		id    string
		price int64
	}

	speckerProduct := []Product{
		{
			title: "JBLClip4",
			id:    "po112",
			price: 200,
		},
		{
			title: "soundCore",
			id:    "pos12",
			price: 123,
		},
	}
	fmt.Println(speckerProduct)
	fmt.Println()

	speckerProduct = append(speckerProduct, Product{
		title: "skullcandy",
		id:    "por234",
		price: 89,
	})
	fmt.Println(speckerProduct)
	fmt.Println()

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
