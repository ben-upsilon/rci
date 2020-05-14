package main

import (
	"fmt"
	"log"
	"os/exec"
)

var TargetDir = "/Users/ben/testBuilderForRCI"

func CloneByTag() {
	cmd := exec.Command("git", "--version")
	cmd.Dir = TargetDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("wtf")
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

func BuildWorkspaceByBID() {

	cmd := exec.Command("mkdir", "workspace")
	cmd.Dir = TargetDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
