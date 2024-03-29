package person

import (
	"fmt"
	"sync"
)

// Person struct represents a person
// Name = person's name
// GettingReadyTime = time needed for the person to get ready
// Ready = this channel will close when person is ready to go
type Person struct {
	Name             string
	GettingReadyTime int
	Ready            chan struct{}
}

// NewPerson is a constructor
func NewPerson(name string) *Person {
	return &Person{name, 0, make(chan struct{})}
}

// Person behaviours

// GrabGlasses simulates a person grabbing their glasses
func (p *Person) GrabGlasses() {
	dur := waitRandomTime(5, 10)
	p.GettingReadyTime += dur

	fmt.Printf("%s grabs the glasses \n It took %s %d seconds\n\n", p.Name, p.Name, dur)
}

// TideShoes simulates a person tiding their shoes
// returns true when done
func (p *Person) TideShoes(wg *sync.WaitGroup) {

	dur := waitRandomTime(20, 35)
	p.GettingReadyTime += dur

	fmt.Printf("%s ties the shoes \n It took %s %d seconds\n\n", p.Name, p.Name, dur)

	wg.Done()
}

// CloseWindow simulates a person closing the windows
func (p Person) CloseWindow(c chan bool) {

	if !<-c {

		dur := waitRandomTime(2, 5)
		p.GettingReadyTime += dur

		fmt.Printf("%s closes the window \n It took %s %d seconds\n\n", p.Name, p.Name, dur)

		// "mark" window as closed
		c <- true
	}
}

// TurnOffTheFan simulates a person turning off the ceiling fan
func (p *Person) TurnOffTheFan(c chan bool) {

	// check if fan is turned off already
	if !<-c {

		dur := waitRandomTime(3, 6)
		p.GettingReadyTime += dur

		fmt.Printf("%s turns off the fan \n It took %s %d seconds\n\n", p.Name, p.Name, dur)

		// "mark" fan as turned off
		c <- true
	}
}

// PocketBelongings simulates a person pocketing their belongings
func (p *Person) PocketBelongings() {

	dur := waitRandomTime(5, 40)
	p.GettingReadyTime += dur

	fmt.Printf("%s pockets the belongings \n It took %s %d seconds\n\n", p.Name, p.Name, dur)

}

// TightenBelt simulates a person tightenning their belt
func (p *Person) TightenBelt() {

	dur := waitRandomTime(5, 12)
	p.GettingReadyTime += dur

	fmt.Printf("%s tightens the belt \n It took %s %d seconds\n\n", p.Name, p.Name, dur)

}
