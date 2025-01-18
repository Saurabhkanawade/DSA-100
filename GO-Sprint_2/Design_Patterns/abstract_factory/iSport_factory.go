package main

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

type NikeFactory struct{}

func (n *NikeFactory) makeShoe() IShoe {
	return &NikeShoe{
		logo: "",
		size: 0,
	}
}

func (n *NikeFactory) makeShirt() IShirt {
	return &NikeShirt{logo: "Nike", size: 38}
}

type AdidasFactory struct{}

func (a *AdidasFactory) makeShoe() IShoe {
	return &AdidasShoe{logo: "Adidas", size: 40}
}

func (a *AdidasFactory) makeShirt() IShirt {
	return &AdidasShirt{logo: "Adidas", size: 36}
}

func GetSportFactory(brand string) (ISportsFactory, error) {
	if brand == "nike" {
		return &NikeFactory{}, nil
	}
	if brand == "adidas" {
		return &AdidasFactory{}, nil
	}
	return nil, fmt.Errorf("unknown brand")
}
