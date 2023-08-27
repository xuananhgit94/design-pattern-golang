package command

/*Button is struct*/
type Button struct {
	Command Command
}

/*Press is function*/
func (b *Button) Press() {
	b.Command.execute()
}
