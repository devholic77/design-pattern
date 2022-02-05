package abfactory

type nike struct {
}

type nikeShirt struct {
	shirt
}

type nikeShoe struct {
	shoe
}

func (n *nike) makeShoe() Shoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) makeShirt() Shirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}
