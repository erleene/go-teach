package main

import (
	"fmt"
	"log"

	"github.com/SeedJobs/devops-go-cron/services/google/iam"
)

func main() {
	iam, err := iam.NewIamClient
	if err != nil {
		log.Fatal(err)
	}

	list, err := &iam.ListServiceAccounts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list)
}
