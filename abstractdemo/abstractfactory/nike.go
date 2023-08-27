package abstractfactory

/*Nike is struct*/
type Nike struct {
}

/*MakeShoe is function*/
func (n *Nike) MakeShoe() IShoe {
	return &nikeShoe{
		shoe{
			"nike", 12,
		},
	}
}

/*MakeShort is function*/
func (n *Nike) MakeShort() IShort {
	return &nikeShort{
		short{
			"nike", 12,
		},
	}
}
