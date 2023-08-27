package iterator

/*UserCollection is struct*/
type UserCollection struct {
	Users []*User
}

/*CreateIterator is function*/
func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		users: u.Users,
	}
}
