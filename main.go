package main

import "fmt"

//stor(er)

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongoDBNumberStore struct {
	// some value here
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(Number int) error {
	fmt.Println("store the number into the mongoDB storage")
	return nil
}

type Apiserver struct {
	numberStore NumberStorer
}

func main() {
	apiserver := Apiserver{
		numberStore: MongoDBNumberStore{},
	}
	numbers, err := apiserver.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(numbers)
}
