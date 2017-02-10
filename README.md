## Introduction
sumHash is a tiny application wich can calculate an input text's sum using those algorithms:
- SHA1
- SHA256
- SHA512
- SHA512_224
- SHA224
- SHA384
- SHA512_256
- MD5

## Build or run 

**Build an executable:**
```bash
~$ go build sumHash.go
```
Or you can

**Run directly:**
```bash
~$ go run sumHash.go
```
## Usage

Show the help:
```bash
~$ go build sumHash.go
~$ ./sumHash -h
```
Output:
```bash
Usage of ./sumHash:
  -s string
    	Hash name: 
		- SHA256
		- SHA224
		- SHA512
		- SHA384
		- SHA512_224
		- SHA512_256 
		- MD5
  -t string
    	Input text to be hashed
```

**Example of use:**

```bash
~$ ./sumHash -s sha256 -t "Hello world!"
```

Output:

```bash
~$ Text: Hello world!
SHA256: c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a
```