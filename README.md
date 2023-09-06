# cryptocalc
==========

A program to calculate md5/sha1/sha256/sha512 based on golang

Install:

```shell
$ go get -u github.com/shenyp09/cryptocalc
```

### Usage: cryptocal command file

Command: md5 sha1 sha256 sha512


e.g.

```shell
$ ./cryptocal md5 cryptocal.go

MD5: e49ee397aa3400341635b4f163badc1d
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





