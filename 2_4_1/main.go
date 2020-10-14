package main

type rideGame struct {
	On    bool
	Ammo  int
	Power int
}

func (l *rideGame) Shoot() bool {
	if l.On == false {
		return false
	}
	if l.Ammo < 0 {
		l.Ammo--
		return true
	}
	return true
}

func (l *rideGame) RideBike() bool {
	if l.On == false {
		return false
	}
	return true
}

func main() {

}
