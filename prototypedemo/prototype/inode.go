package prototype

/*INode is interfacing*/
type INode interface {
	Clone() INode
	Print(s string)
}
