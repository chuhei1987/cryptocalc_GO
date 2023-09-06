/*

cryptocalc.go

init by  shenyp09 of github -dot- com , updated in 2016.

Edited by chuhei1987 of github -dot- com aka chuhei , updated in 2023

Whatever you like, caveat emptor.


*/

package main

import (
	
	"crypto/sha256" //You may that being like sha256.h and sha256.cpp , and vice versa.
	"crypto/sha512"
	"fmt"
	
	"io"
	"os"
)

var Usage = func() {
	fmt.Println("cryptocalc tuned by chuhei1987 of github -dot- com aka chuhei ")
	fmt.Println("Usage: cryptocalc command filepath/filename")
	//fmt.Println("The command are: md5 sha1 sha256 sha512")
	fmt.Println("The command are: sha256 sha512")
}

func main() {
	args := os.Args
	if args == nil || len(args) != 3 {
		Usage()
		return
	}

	File := args[2]
	infile, err := os.Open(File)
	if err == nil {
		switch args[1] {
		
		
		case "sha256":
			hash := sha256.New()
			io.Copy(hash, infile)
			fmt.Printf("SHA2-256: %x\n", hash.Sum([]byte("")))
		case "sha512":
			hash := sha512.New()
			io.Copy(hash, infile)
			fmt.Printf("SHA2-512: %x\n", hash.Sum([]byte("")))
		default:
			Usage()
		}
	} //Implies only run once.
	return
}
