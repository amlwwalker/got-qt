package logic

import ()

//this handles interfacing with any business logic occuring elsewhere
type LogicInterface struct {
	People map[string]*Person
}

type Person struct {
	FirstName string
	LastName  string
	Email     string
}
