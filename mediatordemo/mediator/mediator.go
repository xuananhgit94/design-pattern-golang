package mediator

/*Mediator is interfacing*/
type Mediator interface {
	CanLand(train Train) bool
	NotifyFree()
}
