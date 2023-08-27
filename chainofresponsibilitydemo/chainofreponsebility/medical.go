package chainofreponsebility

import "fmt"

/*Medical is struct*/
type Medical struct {
	next Department
}

/*Execute is function*/
func (m *Medical) Execute(p *Patient) {
	if p.isMedicineProvided {
		fmt.Println("Patient already provided medicine")
		m.next.Execute(p)
		return
	}
	fmt.Println("We are providing medicine to Patient")
	p.isMedicineProvided = true
	m.next.Execute(p)
}

/*SetNext is function*/
func (m *Medical) SetNext(next Department) {
	m.next = next
}
