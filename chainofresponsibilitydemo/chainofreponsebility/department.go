package chainofreponsebility

/*Department is interfacing*/
type Department interface {
	Execute(patient *Patient)
	SetNext(department Department)
}
