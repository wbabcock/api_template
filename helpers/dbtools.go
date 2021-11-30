package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"time"
)

// UniqueID : Generator to create a mathmatical hash id
func UniqueID(length int) string {
	rand.Seed(time.Now().Unix())
	characters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_")
	hash := make([]rune, length)
	for i := range hash {
		hash[i] = characters[rand.Intn(len(characters))]
	}
	return string(hash)
}

// UniqueCode : Same as above, but only with Alphanumeric Characters
func UniqueCode(length int) string {
	rand.Seed(time.Now().Unix())
	characters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hash := make([]rune, length)
	for i := range hash {
		hash[i] = characters[rand.Intn(len(characters))]
	}
	return string(hash)
}

// HashMD5 : Return a string hashed to MD5 string
func HashMD5(value string) string {
	hasher := md5.New()
	hasher.Write([]byte(value))
	return hex.EncodeToString(hasher.Sum(nil))
}

// NumericCleanser : Take a string value and return a string of numbers
func NumericCleanser(value string) string {
	re := regexp.MustCompile("[0-9]+")
	s := re.FindAllString(value, -1)
	output := ""
	for _, v := range s {
		output += v
	}
	return output
}
