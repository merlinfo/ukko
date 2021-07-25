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

	// get path env var, or use default

	if sb := os.Getenv("SPELLBOOK"); sb != "" {
		spellbook = sb
	}

	notation = spellbook + " ;spellbook "

	// string of summons to open all of the file in args

	for i := len(args)-1; i >= 0; i-- {
		notation += args[i] + " ;summon "
	}

	// run it!

	merlin := exec.Command("merlin", notation)
	
	// with std...

	merlin.Stdout = os.Stdout
	merlin.Stdin = os.Stdin
	merlin.Stderr = os.Stderr

	if err := merlin.Run(); err != nil {
		log.Fatal(err)
	}
}
