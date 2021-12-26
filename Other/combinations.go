package main

import (
	"fmt"
)

func GenCombosRecur(arr []string) [][]string {
	if len(arr) == 0 {
		return [][]string{{}}
	}
	fmt.Println("end of genCombos func")
	firstElm := arr[0]
	restElms := arr[1:]

	combosWithoutFirst := GenCombosRecur(restElms)
	combosWithFirst := [][]string{}

	for _, item := range combosWithoutFirst { 
		comboWithFirst := append(item, firstElm)
		combosWithFirst = append(combosWithFirst, comboWithFirst)
	}
	
	return append(combosWithFirst, combosWithoutFirst...)
}

// this solution is heavily influenced by the recursive
// solution above but I wanted to implement something on my own.

// inputArr = ["a", "b", "c"]
// finalCombosArr = [ [], ["a"], ["b"], ["a", "b"], ["c"], ["a", "c"], ["b", "c"], ["a", "b", "c"]  ]
// arrItem = "c"
// currComboArr = []
func GenCombosLoop(inputArr []string) [][]string {
	finalCombosArr := [][]string{{}}
	for _, arrItem := range inputArr {
		fmt.Println("current arrItem:", arrItem)
		currComboArr := [][]string{}
		for _, combo := range finalCombosArr {
			newCombo := append(combo, arrItem)
			currComboArr = append(currComboArr, newCombo)
		}
		finalCombosArr = append(finalCombosArr, currComboArr...)
	}
	return finalCombosArr
}