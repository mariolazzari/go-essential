package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

func main() {
	// Creating a budget with all fields filled
	// The Expires field is set to a specific time in the future
	// Here we set it to 7 days from now
	b1 := Budget{
		CampaignID: "12345",
		Balance:    1000.00,
		Expires:    time.Now().Add(7 * 24 * time.Hour),
	}
	fmt.Println("Budget 1:", b1)
	fmt.Printf("%v\n", b1)

	// Creating a second budget without the Expires field
	// This will use the zero value for time.Time, which is the zero time
	// (January 1, year 1, 00:00:00 UTC)
	b2 := Budget{
		CampaignID: "67890",
		Balance:    500.00,
	}
	fmt.Println("Budget 2:", b2)
	fmt.Printf("%v\n", b2)

	// Creating a third budget without any fields set
	// This will use the zero values for all fields
	b3 := Budget{}
	fmt.Println("Budget 3:", b3)
	fmt.Printf("%v\n", b3)
}
