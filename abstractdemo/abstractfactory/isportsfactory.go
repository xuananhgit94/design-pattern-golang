package abstractfactory

/*ISportsFactory is interfacing*/
type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShort() IShort
}

/*GetSportsFactory is function*/
func GetSportsFactory(brand string) ISportsFactory {
	switch brand {
	case "adidas":
		return &Adidas{}
	case "nike":
		return &Nike{}
	}
	return nil
}
