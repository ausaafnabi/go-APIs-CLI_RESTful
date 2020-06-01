package main

import(
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
 	Left *Tree
	Value int
	Right *Tree
} 

func Traverse(t *Tree) {
	if t==nil {
		return
	}
	Traverse(t.Left)
	fmt.Print(t.Value, " ")
	Traverse(t.Right)
}

func Create(n int) *Tree{
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i:=0;i<2*n;i++ {
		temp:=rand.Intn(n)
		t = insert(t,temp)
	}
	return t
}

func insert(t *Tree,v int) *Tree{
	if t == nil {
		return &Tree{nil,v,nil}
	}

	if v==t.Value {
		return t
	}

	if v<t.Value {
		t.Left = insert(t.Left,v)
		return t
	}

	t.Right = insert(t.Right,v)
	return t
}

func main() {
	tree := Create(30)
	Traverse(tree)
	fmt.Println()
	fmt.Println("The Value of Root Tree is ", tree.Value)
}
