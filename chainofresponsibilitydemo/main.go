package main

import c "designpattern/chainofresponsibilitydemo/chainofreponsebility"

func main() {
	cashier := &c.Cashier{}

	medical := &c.Medical{}
	medical.SetNext(cashier)

	doctor := &c.Doctor{}
	doctor.SetNext(medical)

	reception := c.Reception{}
	reception.SetNext(doctor)

	patient := &c.Patient{Name: "Xuan Anh"}

	reception.Execute(patient)
}
