package chainofreponsebility

import "fmt"

/*Reception is struct*/
type Reception struct {
	next Department
}

/*Execute is function*/
func (r *Reception) Execute(p *Patient) {
	if p.isRegistered {
		fmt.Println("Patient registration has already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.isRegistered = true
	r.next.Execute(p)
}

/*SetNext is function*/
func (r *Reception) SetNext(next Department) {
	r.next = next
}
