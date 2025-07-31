package main

import "fmt"

func main() {
	worker()
}

func worker() {
	r1, err := acquire("resource1")
	if err != nil {
		fmt.Println("Error acquiring resource1:", err)
		return
	}
	defer release(r1)

	r2, err := acquire("resource2")
	if err != nil {
		fmt.Println("Error acquiring resource2:", err)
		return
	}
	defer release(r2)

	fmt.Println("Working with", r1)
	fmt.Println("Working with", r2)
	fmt.Println("Done working")
}

func acquire(name string) (string, error) {
	return name, nil
}

func release(name string) {
	fmt.Println("Releasing", name)
}
