package main

import (
	"fmt"
	"time"

	"github.com/thanadolps/rps/libs/common"
)

func main() {
	fmt.Println("Hello world")
	fmt.Printf("common parse are %s\n", common.Common())

	// New feature: Add timestamp logging
	now := time.Now()
	fmt.Printf("Webhook started at: %s\n", now.Format(time.RFC3339))

	// New feature: Add health check endpoint simulation
	fmt.Println("Health check: OK")
	fmt.Println("Webhook service ready to receive requests")

	fmt.Println("End")
	fmt.Println("End2")
	fmt.Println("End3")
	fmt.Println("End4")
	fmt.Println("End5")
	fmt.Println("End6")
	fmt.Println("End7")
	fmt.Println("End8")
	fmt.Println("End9")
	fmt.Println("End10")
	fmt.Println("End11")
	fmt.Println("End12")
	fmt.Println("End13")
	fmt.Println("End14")
	fmt.Println("End15")
	fmt.Println("End16")

}
