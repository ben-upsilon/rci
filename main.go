package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	testCheckJava()
}
func testCheckJava() {
	cmd := exec.Command("java", "--version")
	cmd.Dir = ""
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("wtf")
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
