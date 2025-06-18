package cmd

type person struct {
	id      string
	name    string
	deposit int
}

type transaction struct {
	sum         int
	date        string
	description string
}

func newPerson(id string, name string, deposit int) *person {
	return &person{
		id:      id,
		name:    name,
		deposit: deposit,
	}
}

func newTransaction(id string, description string) *transaction {
	// default date : today*
	return &transaction{
		sum:         0,
		description: description,
	}
}
