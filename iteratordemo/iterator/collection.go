package iterator

/*Collection is interfacing*/
type Collection interface {
	CreateIterator() Iterator
}
