package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	arg1 := os.Args[1]
	command := "echo \"lenght=" + strconv.Itoa(len(arg1)) + "\" >> $GITHUB_OUTPUT"
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	fmt.Println("exec cmd " + command)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("End success" + arg1)
}
