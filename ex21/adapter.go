package ex21

import "fmt"

type Car interface {
	Drive()
}

type Sedan struct {
}

func (s *Sedan) Drive() {
	fmt.Println("i'm driving")
}

type Train interface {
	RideTheRailroad()
}

type CarToRailAdapter struct {
	WV Sedan
}

func (a *CarToRailAdapter) RideTheRailroad() {
	a.WV.Drive()
	fmt.Println("i'm on the railroad")
}
