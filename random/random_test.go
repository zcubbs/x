package random

import (
	"strings"
	"testing"
)

func TestRandomInt(t *testing.T) {
	minInt := int64(5)
	maxInt := int64(10)
	for i := 0; i < 1000; i++ {
		result := Int(minInt, maxInt)
		if result < minInt || result > maxInt {
			t.Errorf("RandomInt(%d, %d) = %d; want value in range [%d, %d]", minInt, maxInt, result, minInt, maxInt)
		}
	}
}

func TestRandomString(t *testing.T) {
	length := 10
	result := String(length)
	if len(result) != length {
		t.Errorf("RandomString(%d) = %s; want length of %d", length, result, length)
	}
	for _, char := range result {
		if !strings.ContainsRune(alphabet, char) {
			t.Errorf("RandomString(%d) contains unexpected character '%c'", length, char)
		}
	}
}

func TestRandomDomainName(t *testing.T) {
	domain := DomainName()
	if !strings.HasSuffix(domain, ".com") || len(domain) <= 4 {
		t.Errorf("RandomDomainName() = %s; want domain ending with .com and length greater than 4", domain)
	}
}

func TestRandomEmail(t *testing.T) {
	email := Email()
	if !strings.HasSuffix(email, "@email.com") || len(email) <= 10 {
		t.Errorf("RandomEmail() = %s; want email ending with @email.com and length greater than 10", email)
	}
}
