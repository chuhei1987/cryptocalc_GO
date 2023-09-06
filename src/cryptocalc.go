/*

cryptocalc.go

init by  shenyp09 of github -dot- com , updated in 2016.

Edited by chuhei1987 of github -dot- com aka chuhei , updated in 2023

Whatever you like, caveat emptor. This being is following ISO/IEC 10118-3:2018, there are 17 hashing functions in side this box.

For this version, As there are some insecqure function so this version is designed to take them out. 

If you like some view, you may read https://webstore.iec.ch/preview/info_isoiec10118-3%7Bed4.0%7Den.pdf for the start.

朱熹第一個整合之哈希程式，主要仿傚ＣＥＲＴＵＴＩＬ作法，提供ISO/IEC 10118-3:2018　所指的拾柒方法。
不管如何，用者自慎。

*/

package main

import (
	
	"crypto/sha256" //You may that being like sha256.h and sha256.cpp , and vice versa.
	"crypto/sha512"
	"golang.org/x/crypto/sha3"
	
	"github.com/tjfoc/gmsm/sm3"
	"github.com/jzelinskie/whirlpool"
	
	
	
	"fmt"
	
	"io"
	"os"
)

var Usage = func() {
	fmt.Println("cryptocalcIEC2018 tuned by chuhei1987 of github -dot- com aka chuhei ")
	fmt.Println("Caveat Emptor even the functions shown are claimed fine")
	fmt.Println("Usage: cryptocalc command filepath/filename")
	fmt.Println("The command are: ")
	
	fmt.Println("sm3")
	fmt.Println("whirlpool")
	fmt.Println("sha3-256 sha3-384 sha3-512")
	fmt.Println("sha256 sha384 sha512 sha512_256")
	
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
		
		
			
		
		
		//SHA2,FIPS-180-4
		
		case "sha256":
			hash := sha256.New()
			io.Copy(hash, infile)
			fmt.Printf("SHA2-256: %x\n", hash.Sum([]byte("")))
		case "sha384":
			hash := sha512.New384()
			io.Copy(hash, infile)
			fmt.Printf("SHA2-384: %x\n", hash.Sum([]byte("")))
		case "sha512":
			hash := sha512.New()
			io.Copy(hash, infile)
			fmt.Printf("SHA2-512: %x\n", hash.Sum([]byte("")))
		
		case "sha512_256":
			hash := sha512.New512_256()
			io.Copy(hash, infile)
			fmt.Printf("SHA2-512/256: %x\n", hash.Sum([]byte("")))	
		
		
		//SHA3,FIPS-202	
		
		case "sha3-256":
			hash := sha3.New256()
			io.Copy(hash, infile)
			fmt.Printf("SHA3-256: %x\n", hash.Sum([]byte("")))
		case "sha3-384":
			hash := sha3.New384()
			io.Copy(hash, infile)
			fmt.Printf("SHA3-384: %x\n", hash.Sum([]byte("")))
		case "sha3-512":
			hash := sha3.New512()
			io.Copy(hash, infile)
			fmt.Printf("SHA3-512: %x\n", hash.Sum([]byte("")))
		
		//Shang Mi 3 (SM3,商密3) (NCA PRC)	
		case "sm3":
			hash := sm3.New() //REFER https://pkg.go.dev/github.com/tjfoc/gmsm/sm3
			io.Copy(hash, infile)
			fmt.Printf("SM3: %x\n", hash.Sum([]byte("")))
		
		//NESSIE WHIRLPOOL
		case "whirlpool":
			hash := whirlpool.New() //REFER https://pkg.go.dev/github.com/jzelinskie/whirlpool#section-readme
			io.Copy(hash, infile)
			fmt.Printf("WHIRLPOOL: %x\n", hash.Sum([]byte("")))
			
		
			
		
			
		default:
			Usage()
		}
	} //Implies only run once.
	return
}
