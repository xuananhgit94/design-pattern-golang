package mediator

import "fmt"

/*GoodsTrain is struct*/
type GoodsTrain struct {
	Mediator Mediator
}

/*RequestArrival is function*/
func (g *GoodsTrain) RequestArrival() {
	if g.Mediator.CanLand(g) {
		fmt.Println("Good Train: Landing")
	} else {
		fmt.Println("Good Train: Wating")
	}
}

/*Departure is function*/
func (g *GoodsTrain) Departure() {
	fmt.Println("Good Train: Leaving")
	g.Mediator.NotifyFree()
}

/*PermitArrival is function*/
func (g *GoodsTrain) PermitArrival() {
	fmt.Println("Good Train: Arrival Permitted. Landing")
}
