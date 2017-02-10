/*
	sumHash is a tiny application wich can calculate an input text's sum using those algorithms:
		- SHA1
		- SHA256
		- SHA224
		- SHA512
		- SHA512_224
		- SHA512_256
		- SHA384
		- MD5

	Author: Chiheb Nexus
	License: GPLv3
	Year: 2017
	Go version: go1.6.2 linux/amd64
*/
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var (
	hashname = flag.String("s", "", `Hash name: 
		- SHA256
		- SHA224
		- SHA512
		- SHA384
		- SHA512_224
		- SHA512_256 
		- MD5`)
	text = flag.String("t", "", "Input text to be hashed")
)

// Return SHA256 sum
func ReturnSHA256Hash(text string) [32]byte {
	return sha256.Sum256([]byte(text))
}

// Return SHA224 sum
func ReturnSHA224HASH(text string) [28]byte {
	return sha256.Sum224([]byte(text))
}

// Return SHA512 sum
func ReturnSHA512Hash(text string) [64]byte {
	return sha512.Sum512([]byte(text))
}

// Return SHA384 sum
func ReturnSHA384Hash(text string) [48]byte {
	return sha512.Sum384([]byte(text))
}

// Return SHA512_224 sum
func ReturnSHA512_224Hash(text string) [28]byte {
	return sha512.Sum512_224([]byte(text))
}

// Return SHA512_256 sum
func ReturnSHA512_256Hash(text string) [32]byte {
	return sha512.Sum512_256([]byte(text))
}

// Return MD5 sum
func ReturnMD5Hash(text string) [16]byte {
	return md5.Sum([]byte(text))
}

func ReturnSHA1Hash(text string) [20]byte {
	return sha1.Sum([]byte(text))
}

func main() {
	flag.Parse()

	sample := "Text: %s\n%s: %x\n"
	if *hashname == "sha256" {
		fmt.Printf(sample, *text, "SHA256", ReturnSHA256Hash(*text))
	}
	if *hashname == "sha224" {
		fmt.Printf(sample, *text, "SHA224", ReturnSHA224HASH(*text))
	}
	if *hashname == "sha512" {
		fmt.Printf(sample, *text, "SHA512", ReturnSHA512Hash(*text))
	}
	if *hashname == "sha384" {
		fmt.Printf(sample, *text, "SHA348", ReturnSHA384Hash(*text))
	}
	if *hashname == "sha512_224" {
		fmt.Printf(sample, *text, "SHA512_224", ReturnSHA512_224Hash(*text))
	}
	if *hashname == "sha512_256" {
		fmt.Printf(sample, *text, "SHA512_256", ReturnSHA512_256Hash(*text))
	}
	if *hashname == "md5" {
		fmt.Printf(sample, *text, "MD5", ReturnMD5Hash(*text))
	}
	if *hashname == "sha1" {
		fmt.Printf(sample, *text, "SHA1", ReturnSHA1Hash(*text))
	}
}
