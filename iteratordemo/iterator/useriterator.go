package iterator

/*UserIterator is struct*/
type UserIterator struct {
	users []*User
	index int
}

/*HasNext is function*/
func (u *UserIterator) HasNext() bool {
	return u.index < len(u.users)
}

/*GetNext is function*/
func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
