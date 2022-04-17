package main

import (
	"fmt"
)

const (
	_ = iota
	 a = iota
	 b
	 c
)

const (
	 a2 = iota
)

const (
	isAdmin = 1 << iota
	isHeadquaters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	//const myConst int = 42
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%v\n",roles);
	fmt.Printf("isAdmin ? %v\n", isAdmin & roles == isAdmin);
	fmt.Printf("%v\n",c);
	fmt.Printf("%v\n",a2);

}