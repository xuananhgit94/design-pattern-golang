package mediator

import "fmt"

/*PassengerTrain is struct*/
type PassengerTrain struct {
	Mediator Mediator
}

/*RequestArrival is function*/
func (p *PassengerTrain) RequestArrival() {
	if p.Mediator.CanLand(p) {
		fmt.Println("Passenger Train: Lending")
	} else {
		fmt.Println("Passenger Train: Waiting")
	}
}

/*PermitArrival is function*/
func (p *PassengerTrain) PermitArrival() {
	fmt.Println("Passenger Train: Arrival Permitted. Landing")
}

/*Departure is function*/
func (p *PassengerTrain) Departure() {
	fmt.Println("Passenger Train: Leaving")
	p.Mediator.NotifyFree()
}
