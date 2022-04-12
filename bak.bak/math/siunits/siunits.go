package siunits

import (
	"strconv"
)

type Any interface{}

// Maximum File Size by Operating System
/*
File system	Maximum size
	APFS	 8 EB
	exFAT	16 EB
	FAT12	16 MB (4 KB clusters) or 32 MB (8 KB clusters)
	FAT16B	 2 GB (without LFS) or 4 GB (with LFS)
	FAT32	 4 GB
	HFS		 2 GB
	HFS+	 8 EB
	HPFS	 2 GB
	NTFS	16 EB

	Prefix meanings have been notoriously inconsistent throughout
	the history of the computer industry: https://en.wikipedia.org/wiki/Binary_prefix

	For clarity (and consistency with my science background) I use the SI system:
	https://en.wikipedia.org/wiki/International_System_of_Units

*/
const (
	maxPrefixLength = 18
	maxPrefix       = "Gb"
)

// var FS FSPrefix = &fileSize{}

type (
	FSPrefix interface {
		// Name(key int) string
		// Symbol(key int) string
		Size() int
	}
	fileSize struct{ int64 }
)

func (f fileSize) Len() int { return len(f.ToString()) }

func (f fileSize) ToString() string {
	return strconv.FormatInt(f.int64, 10)
}

func (f fileSize) Size() int64 {
	return f.int64
}

func (f fileSize) String() string {

	/// ---> f.int64 file size in bytes
	//  example 1743295

	v := f.ToString()

	// w is length of number of bytes
	//example 1743295 is 7 length
	w := len(v)

	// find the largest SI prefix that works with the file size
	// example 1743295 : start with 7 length
	// first loop = no value for key
	// decriment n
	// second loop = 6 is found
	n := w

	for {
		if v := SI.Symbol(n); v != "" {
			// n is the number to use for the prefix
			break
		}
		n -= 1
		if n == 0 {
			// stop at zero (file size cannot be negative exponents ... the number is int64)
			break
		}
	}

	return v + "   "
}
