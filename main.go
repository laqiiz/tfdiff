package main

/*
c.f.
- Using terraform with go
  - https://gist.github.com/gosuri/a1233ad6197e45d670b3

https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/

- Integration with terraform go-SDK
  https://stackoverflow.com/questions/46932679/integration-with-terraform-go-sdk

- Add JSON Output Format to terraform plan
  > We do intend to make machine-readable plan output eventually, but we are intentionally postponing this for now since we don't feel ready to commit to a stable JSON format for plans.
  https://github.com/hashicorp/terraform/issues/11883#issuecomment-303474598

*/

import (
	"log"

	"os"
	"os/exec"
	"path/filepath"

	"github.com/laqiiz/tfdiff/tfplan"
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
	planOut, err := exec.Command("terraform", "plan", "-no-color").CombinedOutput()
	if err != nil {
		log.Println(string(planOut))
		log.Fatal(err)
	}
	//log.Println(string(planOut))
	log.Print("terraform plan finished")

	diffs, err := tfplan.NewParser().Do(planOut)
	if err != nil {
		log.Fatal(err)
	}

	for _, diff := range diffs {
		log.Printf("%+v\n", diff)
	}

}
