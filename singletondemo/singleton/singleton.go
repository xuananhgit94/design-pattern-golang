package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func init() {
	instance = &singleton{100}
}

func GetInstance() Singleton {
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
