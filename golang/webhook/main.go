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
}
