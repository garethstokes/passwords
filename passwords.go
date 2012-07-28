package passwords

import (
	"math/rand"
	"time"
	"strings"
	"crypto/md5"
	"io"
	"encoding/base64"
	"bytes"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 abcdefghijklmnopqrstuvwxyz" +
	"~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"

type Password struct {
	Hash string
	Salt string
	Iterations int
}

func toBase64(input []byte) (string) {
	output := &bytes.Buffer{}

	encoder := base64.NewEncoder(base64.StdEncoding, output)
	encoder.Write(input)
	encoder.Close()

	return output.String()
}

func fromBase64(input string) ([]byte) {
	output := &bytes.Buffer{}

	decoder := base64.NewDecoder(base64.StdEncoding, output)
	decoder.Read([]byte(input))

	return output.Bytes()
}

func ComputeWithSalt(value string, iterations int, salt string) (Password) {
	h := md5.New()

	if iterations < 1 {
		iterations = 1
	}

	for i := 0; i < iterations; i++ {
		io.WriteString(h, value + salt)
		value = string(h.Sum(nil))
	}

	password := Password{
		Hash: toBase64([]byte(value)),
		Salt: salt,
		Iterations: iterations,
	}
	return password;
}

func Compute(value string, iterations int) (Password) {
	salt := CreateRandomSalt()
	return ComputeWithSalt(value, iterations, salt)
}

func CreateRandomSalt() (string) {
	rand.Seed(int64(time.Now().Nanosecond()))
	length := 8

	r := make([]string, length)
	ri := 0
	buf := make([]byte, length)
	known := map[string]bool{}

	for i := 0; i < length; i++ {
	retry:
		l := rand.Intn(length)
		for j := 0; j < l; j++ {
			buf[j] = chars[rand.Intn(len(chars))]
		}
		s := string(buf[0:l])
		if known[s] {
			goto retry
		}
		known[s] = true
		r[ri] = s
		ri++
	}
	return strings.Join(r, "")
}
