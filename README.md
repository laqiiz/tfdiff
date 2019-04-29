# tfdiff
A tool to help terraform developers know a diff between terraform and actual environment by terraform plan


## Inside design

* Hashicorp terraform-sdk cannot be used for that it does not work if the version is different
* At the moment `terraform plan` does not have option that output with json format
* Tools must be parse for console output of `terraform plan`.

