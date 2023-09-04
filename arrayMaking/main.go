package main

import "fmt"

func main() {
	fmt.Println("Array time!")

	// 1. var arrayName[length] datatype
	var players [11]string

	// 2. var arrayName = [length] datatype{ values }
	var fruits = [5]string{"Mango", "Orange", "Apple"}

	// 3. var arrayName = [...] datatype{ values }
	// U can even ignore the length of the array in the declaration & replace
	// it with ... and let the compiler find the length for you...

	var oddNumbers = [...]int{1, 5, 7, 77}
	var evenNumbers = [...]float64{40.0, 46.4, 80.8, 10.28}

	// Empty arrays
	var emptyArray [10]int

	fmt.Println(players)
	fmt.Println(fruits)
	fmt.Println(oddNumbers)
	fmt.Println(evenNumbers)
	fmt.Println(emptyArray)

	// Using shorthand notation to declare an array
	// (1): arrayName := [length] datatype {values}
	newOddNumbers := [5]int{3, 7, 11, 73, 55}

	// (2): arrayName := [length] datatype {values}
	newEvenNumbers := [...]float32{434, 32, 3290}
	fmt.Println(newOddNumbers)
	fmt.Println(newEvenNumbers)

	gspStudents := [...]string{"Munezero", "Sabato", "Joshua", "Peter"}
	fmt.Println(gspStudents)

	// Array initialization
	// (1) Uninitialized, empty array
	myArray1 := [6]float64{}

	// (2) Partially initialized, empty array
	myArray2 := [6]float64{43.0, 55.4, 60.5, 64.0}

	// (3) Fully initialized, empty array
	myArray3 := [6]float64{43.0, 55.4, 60.5, 64.0, 80.4, 110.3}
	fmt.Println(myArray1)
	fmt.Println(myArray2)
	fmt.Println(myArray3)

	// ASSIGNING ELEMENTS

	// Specific elements to some positions in array
	// Can be added like this...
	specializedArray := [7]int{1: 5, 3: 30, 5: 129, 6: 540}
	fmt.Println(specializedArray)

	// Index of array starts from 0 and ends at length - 1.
	var myBigArray [10]int // int64 array with length 10

	myBigArray[0] = 23456
	myBigArray[1] = 540
	myBigArray[2] = 800
	myBigArray[3] = 1450
	fmt.Println(myBigArray)
	// We get: [23456 540 800 1450 0 0 0 0 0 0]

	notNecessaryAssigned := [5]int{232, 343}
	fmt.Println(notNecessaryAssigned)
	// We get:[232 343 0 0 0] (Cuz of last 3 unassigned elems)

	// How about Array Length?
	// Length/size of an array denotes the number of elements it can contain
	// NB: The length of an array is always fixed and cannot be changed!
	oldOddNumbers := [6]int{3, 121, 33, 22}
	oldSchoolRappers := [...]string{"Big-L", "2Pac", "B.I.G", "IceCube", "GangStarr"}
	lengthOfTheList := len(oldSchoolRappers)
	fmt.Println("Old Odd Numbers are: ", oldOddNumbers)
	fmt.Println("The size of the array: ", len(oldOddNumbers))
	fmt.Println(oldSchoolRappers)
	fmt.Println(lengthOfTheList)

	// Accessing Array!
	fineNumbers := [8]int{30, 60, 44, 100, 142}
	fineTrainers := [...]string{"Ganza", "Clara", "Kim"}
	fmt.Println("The 3rd element ", fineNumbers[3])
	fmt.Println("The 1st element is ", fineTrainers[1])

	// Looping through array (Iterating arrays using range)!
	fiveLuckyNumbers := [5]int{10, 20, 30, 40, 50}
	for k := 0; k < len(fiveLuckyNumbers); k++ {
		fmt.Println(fiveLuckyNumbers[k])
	}

	var floatIntNumbers = [...]float64{40.4, 32, 60.4, 20.8, 50.4, 20.8}
	for m := 0; m < len(floatNumbers); m++ {
		fmt.Printf("%d th element of a is % .2f\n", m, floatIntNumbers[m])
	}

	/*
		Arrays are value types
		----------------------
		Arrays in Go are value types and not reference types. This means that
		when they are assigned to a new variable, a copy of the original array
		is assigned to the new variable.
		If changes are made to the new variable, it will not be reflected in
		the original array.
	*/

	dist1 := [...]string{"Muhanga", "Ruhango", "Kamonyi"}
	dist2 := dist1 // a copy of dist1 is assigned to dist1
	dist2[0] = "Huye"
	dist2[1] = "Nyaruguru"
	fmt.Println("Districts 1 is: ", dist1)
	fmt.Println("Districts 2 is: ", dist2)

}
