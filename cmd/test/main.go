package main

import (
	"log"

	"fyne.io/fyne/v2/data/binding"
)

func main() {
	boundString := binding.NewString()
	s, _ := boundString.Get()
	log.Printf("Bound = '%s'", s)

	myInt := 5
	boundInt := binding.BindInt(&myInt)
	myInt = 10
	i, _ := boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)
}
