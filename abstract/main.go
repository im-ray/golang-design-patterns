package main

import "fmt"

// Animal is the type for our abstract factory
type Animal interface {
	Says()
	LikesWater() bool
}

// Dog is the concrete factory for dogs.
type Dog struct{}

func (d *Dog) Says() {
	fmt.Println("Woof")
}

func (d *Dog) LikesWater() bool {
	return true
}

// implement the abstract factory for dogs

// create a concrete factory for cats

// implement the abstract factory for cats
type Cat struct{}

func (d *Cat) Says() {
	fmt.Println("Meow")
}

func (d *Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct{}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}




func main() {

	// Create one each of a DogFactory and a Catfactory
	dogFactory := DogFactory{}
	catFactory := CatFactory{}

	//Call the New method to create a dog and a cat.
	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	cat.Says()

	fmt.Println("A dog likes water:", dog.LikesWater())
	fmt.Println("A cat likes water:", cat.LikesWater())

}