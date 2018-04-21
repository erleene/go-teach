package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

var branch string

func main() {
	branch := flag.String("branch", "", "the type of branch to list")

	flag.Parse()

	if *branch == "" {
		fmt.Println("Please provide the type of branch to list")
	}

	//check to see if we are currently inside the repo
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory is: ", dir)
	//Walk dir and do something to verify it is a repo

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == ".git" {
				fmt.Println("This is a github repository")

				/**only list branches that are not checked out**/

				// We instance a new repository targeting the given path (the .git folder)
				repo, err := git.PlainOpen(dir)
				if err != nil {
					return err
				}

				// // Get the working directory for the repository
				// w, err := repo.Worktree()
				//
				// //print the status of the current Repository
				// w.Status()

				//retrieve all references to branches
				localBr, _ := repo.Branches()

				/**from localBr, check if master, if not then delete **/

				//for each localBr reference print their name and check if it's a branch
				localBr.ForEach(func(ref *plumbing.Reference) error {
					refName := ref.Name() //returns a ReferenceName string type
					fmt.Println("Reference name:", refName)

					// //check if this is a branch
					// if refName.IsBranch() {
					// 	fmt.Println("It's a branch")
					// }
					if refName != "refs/heads/master" {
						fmt.Println("This is not the master branch")
						//DELETE
						//create new Hash from reference
						hashForRefHash := ref.Hash()
						NewHashReference(refName, hashForRefHash)

						//convert ref

						err := repo.DeleteBranch(refName)
						//
					} else {
						fmt.Println("THIS IS THE MASTER BRANCH")
					}

					return nil
				})

			}
		}
		return nil
	})
}
