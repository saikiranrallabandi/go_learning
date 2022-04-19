package main

import (
	"fmt"
)

type Doctor struct {
	    number int
		actorName string
		companions []string
}

type Animal struct {
       Name string
	   Origin string
}

type Bird struct {
        Animal
		SpeedKPH    float32
		CanFly      bool
}


func main()  {
	     aDoctor := Doctor {
			      number: 3,
				  actorName: "John Pertwee",
				  companions: []string {
					  "Liz Shaw",
					  "Jo Grant",
					  "Sarah Jane Smith",
				},
		}
		bDoctor := struct{name string}{name: "John Pertwee"}
		fmt.Println(aDoctor.companions[1], bDoctor)
		b := Bird{}
		b.Name = "Emu"
		b.Origin = "Australia"
		b.SpeedKPH = 48
		b.CanFly = false
		fmt.Println(b.Name)
	}