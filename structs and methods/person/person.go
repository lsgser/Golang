package person

type Person struct{
	firstName string
	lastName string
}
/*Getters*/
func (p *Person) FirstName() string{
	return p.firstName
}

func (p *Person) LastName() string{
	return p.lastName
}

/*Setters*/
func (p *Person) SetFirstName(name string){
	p.firstName = name
}	

func (p *Person) SetLastName(surname string){
	p.lastName = surname
}
