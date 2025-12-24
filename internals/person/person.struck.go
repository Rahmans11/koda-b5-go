package person

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

func NewPerson(Name string, Address string, Phone string) *Person {
	return &Person{
		Name:    Name,
		Address: Address,
		Phone:   Phone,
	}
}

func (p *Person) PrintPerson() string {
	return fmt.Sprintf("Nama: %s\nAlamat: %s\nNomor Telepon: %s", p.Name, p.Address, p.Phone)
}

func (p *Person) GreetWithNamePerson() string {
	return fmt.Sprintf("Apa kabar %s", p.Name)
}

func (p *Person) UpdateNamePerson(NewName string) {
	p.Name = NewName
}
