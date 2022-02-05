package abfactory

type adidas struct {
}

type adidasShirt struct {
	shirt
}

type adidasShoe struct {
	shoe
}

func (a *adidas) makeShoe() Shoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) makeShirt() Shirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
