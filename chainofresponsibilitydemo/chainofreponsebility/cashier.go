package chainofreponsebility

import "fmt"

/*Cashier is struct*/
type Cashier struct {
	next Department
}

/*Execute is function*/
func (c *Cashier) Execute(p *Patient) {
	if p.isPaid {
		fmt.Println("Patient already paid their build")
		c.next.Execute(p)
		return
	}
	fmt.Println("Patient is paying the bill")
	p.isPaid = true
}

/*SetNext is function*/
func (c *Cashier) SetNext(next Department) {
	c.next = next
}
