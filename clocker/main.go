package clocker

import (
	"log"
)

type Clocker struct {
	Name string
}

func NewClocker(newSchema Clocker) {
	log.Println(newSchema)
	log.Println("new clocker created")
}
