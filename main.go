package main

<<<<<<< HEAD
//what we want is to get a list of service accounts of a project
//we need to make a request to the API to do this by creating a connection
//create a connection to the api
//we need a Context
//we need an ServiceAccountIterator

func main() {
=======
import (
	"fmt"
	"log"

	"github.com/erleene/go-teach/iam"
)

func main() {
	iam, err := iam.NewIamContext("beamery-preview")
	if err != nil {
		log.Fatal(err)
	}

	list := iam.ListAllServiceAccounts()
	if err != nil {
		log.Fatal(err)
	}
>>>>>>> more problems

	// iam, err := iam.NewIamClient
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// list, err := &iam.ListServiceAccounts()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println(list)
}
