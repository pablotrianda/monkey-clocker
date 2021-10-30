package clocker

import (
	"encoding/json"
	"errors"
)

type Clocker struct {
	Name       string
	WorkHours  []int      `json:"work_hours"`
	WorkAgenda []BusyHour `json:"work_agenda"`
}

type BusyHour struct {
	Day       string
	BusyHours []int `json:"busy_hours"`
}

func NewClocker(agendaReq []byte) (Clocker, error) {
	var newClocker Clocker

	err := json.Unmarshal(agendaReq, &newClocker)
	if err != nil {
		return Clocker{}, errors.New("New clocker can't be created")
	}

	return newClocker, nil
}
