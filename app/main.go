package main

import "awesomeProject/app/decisionstree"

func main(){

	context := map[string]string{
		"color": "white",
	}

	exampleTree := decisionstree.NewTree(context)

	result := exampleTree.Run()
	print(result)

}
