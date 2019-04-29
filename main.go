package main

/*
c.f.
- Using terraform with go
  - https://gist.github.com/gosuri/a1233ad6197e45d670b3

https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/

- Integration with terraform go-SDK
  https://stackoverflow.com/questions/46932679/integration-with-terraform-go-sdk
*/

import (
	"log"

	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	prev, _ := filepath.Abs(".")
	defer os.Chdir(prev)

	dir := "testdata"

	// change working directory
	os.Chdir(dir)

	// terraform install check
	versionOut, err := exec.Command("terraform", "-version").CombinedOutput()
	if err != nil {
		log.Println(string(versionOut))
		log.Fatal(err)
	}
	log.Println(string(versionOut))
	log.Print("install check finished")

	// terraform init
	// TODO skip option?
	initOut, err := exec.Command("terraform", "init").CombinedOutput()
	if err != nil {
		log.Println(string(initOut))
		log.Fatal(err)
	}
	log.Println(string(initOut))
	log.Print("init completed")

	// do validate before diff action
	validateOut, err := exec.Command("terraform", "validate").CombinedOutput()
	if err != nil {
		log.Println(string(validateOut))
		log.Fatal(err)
	}
	log.Println(string(validateOut))
	log.Print("validate finished")

	// terraform is valid format after validation
	planOut, err := exec.Command("terraform", "plan").CombinedOutput()
	if err != nil {
		log.Println(string(planOut))
		log.Fatal(err)
	}
	log.Println(string(planOut))

}
