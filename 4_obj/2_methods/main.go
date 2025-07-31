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

func (b Budget) IsExpired() bool {
	return time.Now().After(b.Expires)
}

func (b Budget) TimeLeft() time.Duration {
	if b.IsExpired() {
		return 0
	}
	return b.Expires.Sub(time.Now().UTC())
}

// Update method to add a sum to the Balance
// This method modifies the Budget's Balance field
// It does not return a new Budget, but updates the existing one
// This allows for modifying the budget without creating a new instance
func (b *Budget) Update(sum float64) {
	b.Balance += sum
}

func main() {
	b := Budget{
		CampaignID: "12345",
		Balance:    1000.00,
		Expires:    time.Now().Add(7 * 24 * time.Hour),
	}
	fmt.Println("Time left:", b.TimeLeft())

	b.Update(500.00)
	fmt.Println("Updated Balance:", b.Balance)
}
