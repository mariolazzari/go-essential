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

func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {

	if campaignID == "" {
		return nil, fmt.Errorf("campaignID cannot be empty")
	}

	if balance < 0 {
		return nil, fmt.Errorf("balance cannot be negative")
	}

	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("expires time must be in the future")
	}

	b := Budget{
		CampaignID: campaignID,
		Balance:    balance,
		Expires:    expires,
	}

	return &b, nil
}

func main() {
	expires := time.Now().Add(7 * 24 * time.Hour)
	b1, err := NewBudget("12345", 1000.00, expires)
	if err != nil {
		fmt.Println("Error creating budget:", err)
		return
	}
	fmt.Println("Budget created with CampaignID:", b1.CampaignID)
	fmt.Println("Initial Balance:", b1.Balance)
	fmt.Println("Expires at:", b1.Expires)

	b2, err := NewBudget("", 500.00, expires)
	if err != nil {
		fmt.Println("Error creating budget:", err)
	} else {
		fmt.Printf("Budget created with CampaignID: %s\n", b2.CampaignID)
	}
}
