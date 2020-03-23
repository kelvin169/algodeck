package main 

import (
 "fmt"
)

var toBeSorted [10] int  = [10] int {2,1,4,3,5,7,8,9,4,5}

func bubbleSort(input [10] int) {
	// n being the number of items in our list 
	n := 10 
	// set swaped to true
	swapped := true 
	//loop 
	for swapped {
//set wapped to false 
 swapped = false 
//iterate for all the elements in the list 
for i := 1; i < n ; i ++{

	// if the current element is greater than the next elemnt  then swap them 
	if input[i-1] > input [i] {

		//log that you are swapping the elements for prosterity 
		fmt.Println("swapping")
		//swap values using go's tuple assignment 
		input[i], input[i-1]= input[i-1],input[i]

		//set swapped to true 
		//if the loop ends and the swp is still equal to false 
		// our algorithm will assume the list is fully sorted 
		swapped = true 
	}
}

	}
	fmt.Println (input)
}

func main (){
	fmt.Println("awesome Stuff!")
	bubbleSort(toBeSorted)
}
fu