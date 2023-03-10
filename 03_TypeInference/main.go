package main

import (
	"fmt"

	"github.comLinkedInLearning/generics-in-go-2492215/business"
)

// main is our simple "playground" for the course.
// Note, that in production code, it is a good practice to keep the main function short.
func main() {
	// Create three different energy offers of kineteco
	solar2k := business.Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := business.Solar{Name: "Solar 3000", Netto: 4.000}
	windwest := business.Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Printf(solar2k.Print())
	fmt.Printf(solar3k.Print())
	fmt.Printf(windwest.Print())
	fmt.Printf("our first generic function: %s\n", business.PrintGeneric(solar2k))
	fmt.Printf("our first generic function with wind: %s\n", business.PrintGeneric[business.Wind](windwest))

	// Print functions for 01_04
	ss := []business.Solar{solar2k, solar3k}
	business.PrintSlice[business.Solar](ss)
	business.PrintSlice[business.Wind]([]business.Wind{windwest, windwest})

	business.Cost(10, solar2k.Netto)
	//business.Cost(0.45, 10)
	business.Cost(0.45, float64(10))
	business.Cost[float64](0.45, float64(10))

	ss2 := business.SolarSlice{solar2k, solar3k}
	business.PrintSlice(ss2)
	business.PrintSlice2(ss2)
}