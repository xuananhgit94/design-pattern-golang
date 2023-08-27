package main

import m "designpattern/mediatordemo/mediator"

func main() {
	stationManager := m.NewStationManager()

	passengerTrain := m.PassengerTrain{
		Mediator: stationManager,
	}

	goodsTrain := m.GoodsTrain{
		Mediator: stationManager,
	}

	passengerTrain.RequestArrival()
	goodsTrain.RequestArrival()
	passengerTrain.Departure()
}
