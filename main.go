package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed faces.txt
var facesfile string

func makeBorder(s string) string {
	splitted := strings.Split(s, "\n")
	for i := range splitted {
		splitted[i] = "│ " + splitted[i] + " │"
	}
	s = strings.Join(splitted, "\n")
	return fmt.Sprintf("╭─────╮\n%s\n╰─────╯", s)
}

func main() {
	facesfile = strings.ReplaceAll(facesfile, "-", " ")
	facesfile = strings.ReplaceAll(facesfile, "x", "•")
	faces := strings.Split(facesfile, "\n\n")

	rand.Seed(time.Now().UnixNano())
	// random number from 1 to 6
	r := rand.Intn(6)
	fmt.Println(makeBorder(faces[r]))
}
