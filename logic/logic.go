package logic

import (
	"fmt"
	"time"
)

//handles the interface between any go functions
//and the qml
func (l *LogicInterface) ConfigureLogic() {
	l.People = make(map[string]*Person)
}

func (l *LogicInterface) SearchForMatches(regex string, informant func(float64, bool)) {
	//if search takes a while then the informant can alert the front end
	var p Person
	p.FirstName = "Alex"
	p.LastName = "Walker"
	p.Email = "alex@walker.com"
	//just demoing using async if things are slow
	go func() {
		time.Sleep(1 * time.Second)
		informant(0.0, true) //we don't know how long this will take
		time.Sleep(1 * time.Second)
		informant(0.0, true) //we don't know how long this will take
		time.Sleep(1 * time.Second)
		informant(0.0, true) //we don't know how long this will take
		time.Sleep(1 * time.Second)
		//the informant knows whatever it was doing is ready whenit is sent a 1.0 value
		l.People[p.Email] = &p
		fmt.Println(l.People[p.Email])
		informant(1.0, true) //we don't know how long this will take
	}()
}

