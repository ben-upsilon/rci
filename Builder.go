package main

import (
	"fmt"
	"log"
	"os/exec"
)

func CloneByTag() {
	cmd := exec.Command("git", "--version")
	cmd.Dir = ""
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("wtf")
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
