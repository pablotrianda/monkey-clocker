package clocker

import (
	"errors"
)

type Clocker struct {
	Name      string `json:"name"`
	BusyHours []int  `json:"busy_hours"`
}

func NewClocker(name string, busyHours []int) (Clocker, error) {
	if name == "" || len(busyHours) == 0 {
		return Clocker{}, errors.New("New clocker can't be created")
	}

	var newClocker Clocker
	newClocker.Name = name
	newClocker.BusyHours = busyHours

	return newClocker, nil
}
