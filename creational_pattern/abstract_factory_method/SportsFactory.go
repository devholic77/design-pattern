package abfactory

import "fmt"

type SportsFactory interface {
	makeShoe() Shoe
	makeShirt() Shirt
}

func getSportsFactory(brand string) (SportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
