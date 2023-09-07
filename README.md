# CRYPTOCALCIEC2018 ( General version )
==========

CHUHEI's here. 

A program to calculate NINE secure types of hash(s) , Based on ISO/IEC 10118-3:2018 and security officer's recommendation,  Powered by golang

The remaining eight are taken out from this branch due to the length of hash and details of cryptanalysis. 

Install:

```shell
$ go get -u github.com/chuhei1987/cryptocalc_GO
```

### Usage: cryptocal command file

Command: sm3 sha256 sha384 sha512 sha512-256 sha3-258 sha3-384 sha3-512 whirlpool


e.g.

```shell
$ ./cryptocal sha256 cryptocal.go

SHA2-256: 57acc54a01333cf860cbaace0a22bde817e5a7540a2678b34efb12b8c9006236

$ ./cryptocal sha512 cryptocal.go

SHA2-512: e49459573f9ec86104704e8f19d7fff1ec23fad77a4d775012e1f7be99ed3c7687ae2cb6f6a8be2f2f9857b9465f53aa3d430185834e5dc02189c24992d85294

$ ./cryptocal sm3 cryptocal.go

     SM3: b5055a161ceec9b267a583b0e0ae209bf9fe6850f1d7e599ff751156b8d6c818

$ ./cryptocal whirlpool cryptocal.go

WHIRLPOOL: 56d3579f67a689add42c69bcb07449ea8cb2ac63f10067da385a9403296bdac84df86872e40eb3f7f7c9d24e940c9ce14326fb2a832bc4142c1c377d81c80e6b

```

### For Compiling
You can try using go build the cryptocal.go if you install the GOlang-DEV or you point the standalone files of GOlang-dev by setting system environment.

### For Trying
Go run cryptocal.go may do.

Note: 

For the time being, I am still figuring how to get that box work, there is a penalty, perhaps 20% of overall performance penalty, so bloating but stable when using non C or non C++ when constructing what you need.

How to check if that thing is fine?
Try using an archive, around 200MiB if available (a jpeg picture may do),  to test your box built, a basic step is either using certutil (M$) , sha256sum (mainly for *nix) and 7-zip's SHA-256 hashing to check if result is exactly same as for hashing, The reading should be in binary mode to get right if you try searching what I found.

### Help about using certutil (M$)

certutil (M$) does not only just mean managing these SSL certs, that one also can check integrity of a file too.
By using such as certutil -hashfile (a file you like to check) ( sha256 or sha512 ), then shall be fine.

### Help about using sha256sum, sha384sum, sha512sum or something like that in GNU posix
Preferring one in one out, so most of authors follow the thing to save a day.
In most cases, such as sha256sum [text file] , sha256sum [archive file] output the same result.

### What next?
Try digging deeper, when new thing ready, may reveal some of the actual purposes.

The being is in public domain. 





