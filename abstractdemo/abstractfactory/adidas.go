package abstractfactory

/*Adidas is struct*/
type Adidas struct {
}

/*MakeShoe is function*/
func (a *Adidas) MakeShoe() IShoe {
	return &adidasShoe{
		shoe{
			"adidas", 14,
		},
	}
}

/*MakeShort is function*/
func (a *Adidas) MakeShort() IShort {
	return &adidasShort{
		short{
			"adidas", 14,
		},
	}
}
