## Description
Here, In this template you will get to know about How to test the Terraform module using the Terratest in go language on Azure cloud.

---
### Prerequisite
* Azure Cloud account
* Go language install
* Terraform and Terratest install in IDE.
---
### Steps :-
1. Login into the AZ account using the `az login`.
2. Run the go file to test the terraform module using the commands:-
3. Go to `test` folder

`go mod init test` It will initialize your module with the test package name.  

`go mod tidy ` It will remove unwanted dependencies.

`go test -v` It will run your all test cases inside your test folder.

<b>Directory Structure</b>
```
├── envs
│   └── dev
│       └── variables.tfvars
├── README.md
├── terraform
│   ├── main.tf
│   ├── output.tf
│   ├── resource_group.tf
│   └── vars.tf
└── test
    └── simple_test.go
```

---
