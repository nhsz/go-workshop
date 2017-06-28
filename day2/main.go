package main

import (
	"./binarytree"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputList := os.Args[1:]
	list := make([]int, len(inputList))
	for i, v := range inputList {
		list[i], _ = strconv.Atoi(v)
	}

	aBinaryTree := binarytree.GenerateBinaryTreeFromList(list)
	sortedList := binarytree.GenerateListFromBinaryTree(aBinaryTree)

	fmt.Println("Original input: ", list)
	fmt.Println("Sorted input: ", sortedList)
}
