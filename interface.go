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

    //requests from qml
    _ func(number1, number2 string) string `slot:"calculator"`
}
//setup functions to communicate between front end and back end

//example of receiving data from frontend and returning a result
func (q *QmlBridge) ConfigureBridge(config Config) {
	//1. configure the hotloader
	q.business = BusinessInterface{}
	q.business.configureInterface()
	q.hotLoader = HotLoader{} //may not need it, defined in main.go

	//2. Configure signals
	//configure calculator
	q.ConnectCalculator(func(number1, number2 string) string {
	    return addingNumbers(number1, number2)
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

