package randomstring

import (
	"errors"
	"math/rand"
	"time"
	"unsafe"
)

// StdChars is a set of standard characters allowed in string.
var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var src = rand.NewSource(time.Now().UnixNano())

const (
	defaultLength = 4
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// New returns a new random string of the standard length, consisting of
// standard characters.
func New() string {
	return *randString(defaultLength)
}

// NewLen returns a new random string of the provided length, consisting of
// standard characters.
func NewLen(length int) string {
	return *randString(length)
}

// NewLenChars returns a new random string of the provided length, consisting
// of the provided string of allowed characters (maximum 63).
func NewLenChars(length int, chars string) (string, error) {
	if length == 0 {
		return "", nil
	}
	clen := len(chars)
	if clen < 2 || clen > 63 {
		return "", errors.New("Invalid length of character set, max 63 chars")
	}
	letterBytes = chars
	return *randString(length), nil
}

// source https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func randString(n int) *string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return (*string)(unsafe.Pointer(&b))
}
