/*

cryptocalc.go

init by  shenyp09 of github -dot- com , updated in 2016.

Edited by chuhei1987 of github -dot- com aka chuhei , updated in 2023

Whatever you like, caveat emptor. This being is following ISO/IEC 10118-3:2018, there are 17 hashing functions in side this box.

If you like some view, you may read https://webstore.iec.ch/preview/info_isoiec10118-3%7Bed4.0%7Den.pdf for the start.

朱熹第一個整合之哈希程式，主要仿傚ＣＥＲＴＵＴＩＬ作法，提供ISO/IEC 10118-3:2018　所指的拾柒方法。
不管如何，用者自慎。

*/

package main

import (
	"crypto/sha1" 
	"crypto/sha256" //You may that being like sha256.h and sha256.cpp , and vice versa.
	"crypto/sha512"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/ripemd160"
	"github.com/zhimoe/ripemd128"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/jzelinskie/whirlpool"
	"github.com/pkositsyn/streebog-hash"
	
	
	"fmt"
	
	"io"
	"os"
)

var Usage = func() {
	fmt.Println("cryptocalcIEC2018 tuned by chuhei1987 of github -dot- com aka chuhei ")
	fmt.Println("Caveat Emptor especially when you notice result with an exclamation mark")
	fmt.Println("Usage: cryptoIEC command filepath/filename")
	fmt.Println("The command are: ")
	fmt.Println("sha")
	fmt.Println("sm3")
	fmt.Println("whirlpool")
	fmt.Println("ripemd160 ripemd128")
	fmt.Println("streebog256 streebog512")
	fmt.Println("sha3-224 sha3-256 sha3-384 sha3-512")
	fmt.Println("sha224 sha256 sha384 sha512 sha512_224 sha512_256")
	
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
		//RIPEMD
		case "ripemd160":
			hash := ripemd160.New() //refer https://pkg.go.dev/golang.org/x/crypto/ripemd160
			io.Copy(hash, infile)
			fmt.Printf("!RIPEMD-160: %x\n", hash.Sum([]byte("")))
		case "ripemd128":
			hash := ripemd128.New() //refer https://pkg.go.dev/github.com/zhimoe/ripemd128
			io.Copy(hash, infile)
			fmt.Printf("!RIPEMD-128: %x\n", hash.Sum([]byte("")))
		
			
		
		//SHA1
		case "sha":
			hash := sha1.New() //refer https://pkg.go.dev/crypto/sha256@go1.21.0
			io.Copy(hash, infile)
			fmt.Printf("!SHA1: %x\n", hash.Sum([]byte("")))
		//SHA2,FIPS-180-4
		case "sha224":
			hash := sha256.New224() 
			io.Copy(hash, infile)
			fmt.Printf("!SHA2-224: %x\n", hash.Sum([]byte("")))
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
		case "sha512_224":
			hash := sha512.New512_224()
			io.Copy(hash, infile)
			fmt.Printf("!SHA2-512/224: %x\n", hash.Sum([]byte("")))
		case "sha512_256":
			hash := sha512.New512_256()
			io.Copy(hash, infile)
			fmt.Printf("SHA2-512/256: %x\n", hash.Sum([]byte("")))	
		
		
		//SHA3,FIPS-202	
		case "sha3-224":
			hash := sha3.New224() //refer https://pkg.go.dev/golang.org/x/crypto/sha3
			io.Copy(hash, infile)
			fmt.Printf("!SHA3-224: %x\n", hash.Sum([]byte("")))
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
			
		//ГОСТ-P 34.11-2012 "Стрибог" 
		case "streebog256":
			hash := streebog.New256() //REFER https://pkg.go.devgithub.com/pkositsyn/streebog-hash
			io.Copy(hash, infile)
			fmt.Printf("!Стрибог-256: %x\n", hash.Sum([]byte("")))	
		case "streebog512":
			hash := streebog.New512() //REFER https://pkg.go.devgithub.com/pkositsyn/streebog-hash
			io.Copy(hash, infile)
			fmt.Printf("!Стрибог-512: %x\n", hash.Sum([]byte("")))		
			
		
			
		default:
			Usage()
		}
	} //Implies only run once.
	return
}
