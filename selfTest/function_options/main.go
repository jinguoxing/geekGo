package main

import "fmt"

type House struct {

	Material string
	HashFireplace bool
	Floors  int
}


type HouseOption func(*House)

func WithConcrete()  HouseOption{

	return func(house *House) {

		house.Material = "concrete"
	}
}


func WithFloors (floors int) HouseOption {

	return func(house *House) {
		house.Floors = floors
	}
}


func NewHouse(opts ...HouseOption) *House {

	const (
		defaultFloors       = 2
		defaultHasFireplace = true
		defaultMaterial     = "wood"
	)

	h := &House{
		Material: defaultMaterial,
		HashFireplace: defaultHasFireplace,
		Floors: defaultFloors,
	}

	for _,opt := range opts {
		opt(h)
	}

	return h
}


func main(){

	h := NewHouse(WithConcrete(),WithFloors(5))

	fmt.Printf("%+v",h)

}