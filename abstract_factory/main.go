package main

import "fmt"

// Animal is the type for our abstract factory
type Animal interface {
	Says()
	LikesWater() bool
}

// Dog is the concrete factory for dogs
type Dog struct{}

func (d *Dog) Says() {
	fmt.Println("Woof")
}

func (d *Dog) LikesWater() bool {
	return true
}

// Cat is the concrete factory for cats
type Cat struct{}

func (c *Cat) Says() {
	fmt.Println("Meow")
}

func (d *Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct {
}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct {
}

func (df *CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	fmt.Println("abstract factory")

	// create one of each dog and cat factory
	DogFactory := DogFactory{}
	CatFactory := CatFactory{}

	//call the new method to create dog and cat
	dog := DogFactory.New()
	cat := CatFactory.New()
	dog.Says()
	cat.Says()

	fmt.Println("A dog likes water:", dog.LikesWater())
	fmt.Println("A cat likes water:", cat.LikesWater())
}
