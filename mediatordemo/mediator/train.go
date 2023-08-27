package mediator

/*Train is interfacing*/
type Train interface {
	RequestArrival()
	Departure()
	PermitArrival()
}
