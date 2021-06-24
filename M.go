/*
	M - The Merlin Command Wrapper
	Jonah G. Rongstad (c) 2021
*/

package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	var (
		spellbook string = "~/.merlin/spellbook.mn"
		notation string
		args = os.Args[1:]
	)

	if sb := os.Getenv("SPELLBOOK"); sb != "" {
		spellbook = sb
	}

	notation = spellbook + " ;spellbook "

	for i := len(args)-1; i >= 0; i-- {
		notation += args[i] + " ;summon "
	}

	merlin := exec.Command("merlin", notation)
	
	merlin.Stdout = os.Stdout
	merlin.Stdin = os.Stdin

	if err := merlin.Run(); err != nil {
		log.Fatal(err)
	}
}
