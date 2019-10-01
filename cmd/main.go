package main

import (
	"dailywalk/person"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Concurrency Excersise 1")
	fmt.Println("")

	// Bob person setup
	bob := person.NewPerson("Bob")

	// Alice person setup
	alice := person.NewPerson("Alice")

	// Channels to control events only occuring once
	// (if the window is already closed it does not to be closed again)
	windowClosedC := make(chan bool)
	fanOffC := make(chan bool)

	// Set windowClosedC and fanOffC to false by default
	go func() {
		windowClosedC <- false
		fanOffC <- false

		time.Sleep(time.Millisecond * 100)
	}()

	// Goroutine for Bob
	go morningRoutine(bob, windowClosedC, fanOffC)

	// Goroutine for Alice
	go morningRoutine(alice, windowClosedC, fanOffC)

	// "Wait" untill both persons are ready
	<-bob.Ready
	<-alice.Ready
}

// morning routine describes a morning routine for each person
func morningRoutine(p person.Person, w, f chan bool) {
	fmt.Printf("%s starts getting ready\n\n", p.Name)

	p.GrabGlasses()
	p.TightenBelt()

	p.CloseWindow(w)
	p.TurnOffTheFan(f)

	p.PocketBelongings()

	p.Ready <- struct{}{}

	fmt.Printf("%s is ready to go!\n\n", p.Name)
}
