package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// Int generates a random integer between min and max
func Int(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// String generates a random string of length n
func String(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// DomainName generates a random domain name
func DomainName() string {
	return String(10) + ".com"
}

// Email generates a random email
func Email() string {
	return fmt.Sprintf("%s@email.com", String(6))
}
