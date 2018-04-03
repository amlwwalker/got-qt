package main

import (
	"fmt"
	"time"
	"github.com/therecipe/qt/core"
)

type QmlBridge struct {
    core.QObject
    hotLoader HotLoader
    business BusinessInterface
    //messages to qml
    _ func(p string)        `signal:"updateLoader"`
    _ func(author, mode, date, host, version, port string, hotload bool)        `signal:"updateSettings"`
    _ func(data string) 	`signal:"sendTime"`
    _ func(c float64, indeterminate bool) 	`signal:"updateProcessStatus"`

    //requests from qml
    _ func(number1, number2 string) string `slot:"calculator"`
    _ func() `slot:"startAsynchronousProcess"`
    _ func(regex string) `slot:"searchFor"`

}
//setup functions to communicate between front end and back end

//example of receiving data from frontend and returning a result
func (q *QmlBridge) ConfigureBridge(config Config) {
	//1. configure the hotloader
	q.business = BusinessInterface{}
	q.business.configureInterface()
	q.hotLoader = HotLoader{} //may not need it, specified in main.go

	//2. Configure signals
	//configure calculator
	q.ConnectCalculator(func(number1, number2 string) string {
	    return addingNumbers(number1, number2)
	})
	q.ConnectStartAsynchronousProcess(func() {
		//inform process has started
		//so this needs to be signaled to start a long process.
		//the frontend will assume a value of 1.0 is process complete
		q.UpdateProcessStatus(0.0, false)
		q.business.startAsynchronousRoutine(q.UpdateProcessStatus)
	})
	q.ConnectSearchFor(func(regex string) {
		//in here we are going to add matches to the search model
		//that way the front end will be updated live
		//inform front end work has begun
		q.UpdateProcessStatus(0.0, true)
	    q.business.searchForMatches(regex, q.UpdateProcessStatus)
	})
		//example signalling the frontend with settings
	go func() {
		//send the settings to the front end after a period of time
        time.Sleep(5 * time.Second)
        fmt.Println("updating settings with ", config)
        q.UpdateSettings(config.Author, config.Mode, config.Date, config.Host, config.Version, config.Port, config.Hotload)
    }()
    	//example of external function signalling the front end
    q.sendCurrentTime()
	q.business.demo()
}

//example of sending data to the frontend via a signal
func (q *QmlBridge) sendCurrentTime() {
	go func() {
		for t := range time.NewTicker(time.Second * 1).C {
			q.SendTime(t.Format(time.ANSIC))
		}
	}()
}
//example of handling a receive from frontend via slot
func addingNumbers(number1, number2 string) string {
	fmt.Println("addingNumbers")
	return number1 + number1 + number2 + number2
}

//this handles interfacing with any business logic occuring elsewhere
type BusinessInterface struct {
	pModel *PersonModel
	sModel *PersonModel
	fModel *PersonModel
}
//handles the interface between any go functions
//and the qml
func (b *BusinessInterface) configureInterface() {
	b.pModel = NewPersonModel(nil)
	b.sModel = NewPersonModel(nil)
	b.fModel = NewPersonModel(nil)
}

func (b *BusinessInterface) searchForMatches(regex string, informant func(float64, bool)) {
	//if search takes a while then the informant can alert the front end
	var p = NewPerson(nil)
	p.SetFirstName(regex)
	p.SetLastName("Search")
	p.SetEmail(regex + "@hotmail.com")
	//just demoing using async if things are slow
	go func() {
		time.Sleep(1 * time.Second)
		informant(0.3, true) //we don't know how long this will take
		time.Sleep(1 * time.Second)
		informant(0.5, true) //we don't know how long this will take
		time.Sleep(1 * time.Second)
		informant(7.0, true) //we don't know how long this will take
		time.Sleep(1 * time.Second)
		informant(1.0, true) //we don't know how long this will take
		fmt.Println("found match...")
		b.sModel.AddPerson(p)
	}()
}
//the interface needs to know how to inform the front end on progress
//so takes a function that takes a value that the front end will use
func (b *BusinessInterface) startAsynchronousRoutine(informant func(float64, bool)) {
	//on a go routine, count up to 10
	//each tick, inform the front end of your percentage progress
	//when it reaches ten, inform the front end it is complete

	go func() {
		var c float64
		c = 0.0
		for (c < 1.0) {
			fmt.Printf("sending %.2f\r\n", c)
			informant(c, false) //we know how long this process will take
			time.Sleep(1 * time.Second)
			c = c + 0.1
		}
		return
	}()
}
func (b *BusinessInterface) demo() {
	var p = NewPerson(nil)
	p.SetFirstName("john")
	p.SetLastName("doe")
	p.SetEmail("john@doe.com")
	//add the person to the PersonModel
	b.pModel.SetPeople([]*Person{p})

	//make changes on a routine
	//to demo updates
	go func() {
		fmt.Println("3 seconds to adding new people")
		time.Sleep(3 * time.Second)

		//add person
		for i := 0; i < 3; i++ {
			var p = NewPerson(nil)
			p.SetFirstName("hello")
			p.SetLastName("world")
			p.SetEmail("hello@world.com")
			b.pModel.AddPerson(p)
		}
		fmt.Println("3 seconds to editing people")
		time.Sleep(3 * time.Second)
		//edit person (demo not changing a field by leaving it blank (see the model code))
		b.pModel.EditPerson(1, "bob", "", "bob@gmail.com")
		b.pModel.EditPerson(3, "", "john", "john@hotmail.com")
		fmt.Println("3 seconds to remove person")
		time.Sleep(3 * time.Second)
		//remove person
		b.pModel.RemovePerson(2)
	}()
}

