package main

import (
	"fmt"
	"iter"
)

type petType struct {
	name      string
	haveASoul bool
	buried    bool
}

func PetSematary(pet []petType) iter.Seq[petType] {
	return func(StephenKing func(petType) bool) {
		for _, s := range pet {
			if !s.buried {
				continue
			}
			s.buried = false
			s.haveASoul = false
			if !StephenKing(s) {
				break
			}
		}
	}
}

func main() {
	pets := []petType{
		{
			name:      "Linda",
			haveASoul: false,
			buried:    true,
		},
		{
			name:      "Gohan",
			haveASoul: true,
			buried:    true,
		},
		{
			name:      "Snoopy",
			haveASoul: true,
			buried:    false,
		},
	}

	for s := range PetSematary(pets) {
		fmt.Printf("%s is back from the dead\n", s.name)
		if !s.haveASoul {
			fmt.Printf("%s is a soulless monster\n", s.name)
		}
		if s.haveASoul {
			fmt.Printf("%s is a happy pet\n", s.name)
		}
	}
}
