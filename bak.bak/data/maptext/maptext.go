package maptext

import (
	"crypto"
	"crypto/md5"
	"fmt"

	"github.com/skeptycal/goutil/repo/errorlogger"
)

type direction = int

const (
	none direction = iota
	left
	right
	both
)

const (
	b1 = 0b0001
	b2 = 0b0010
	b4 = 0b0100
	b8 = 0b1000
)

var log = errorlogger.Log

func init() {
	if ok := crypto.MD5.Available(); !ok {
		log.Fatal("crypto.MD5 hashing is not available: %v", ok)
	}

	// h := md5.New()s

}

var BlankData = []byte{}

type (
	Any = interface{}

	id struct {
		n int
		s string
	}

	textMap struct {
		m map[int]string
	}

	edge struct {
		n1  *node
		n2  *node
		dir direction
	}
)

func (i *id) String() string {
	if i.s == "" {
		i.s = fmt.Sprintf("%09d", i.n)
	}
	return i.s
}

func (e *edge) CanLeft() bool { return e.dir&b2 != 0 }

func (e *edge) CanRight() bool { return e.dir&b1 != 0 }

func CheckSum(data []byte) [16]byte { return md5.Sum(data) }
