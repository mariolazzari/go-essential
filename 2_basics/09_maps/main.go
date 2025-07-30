package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"GOOGL": 2800.00,
		"AAPL":  150.00,
		"AMZN":  3400.00,
		"MSFT":  299.00,
	}

	// map length
	fmt.Println(len(stocks))

	// print value
	fmt.Println(stocks["AAPL"])
	// zero value for non-existing key
	fmt.Println(stocks["TSLA"])
	// check if key exists
	price, ok := stocks["TSLA"]
	if ok {
		fmt.Println("TSLA exists with price:", price)
	} else {
		fmt.Println("TSLA does not exist in the map.")
	}

	// set value
	stocks["TSLA"] = 800.00
	fmt.Println("After adding TSLA:", stocks)

	// delete value
	delete(stocks, "AMZN")
	fmt.Println("After deleting AMZN:", stocks)

	// print all keys
	for key := range stocks {
		fmt.Println(key)
	}

	// print all keys and values
	for key, value := range stocks {
		fmt.Println(key, ":", value)
	}

}
