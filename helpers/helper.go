package helpers

import "math/rand"

type CompanyDetails struct {
	CompanyName string
	CompanyType string
}

func PrintCompanyDetails() string {
	return "shit"
}

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}
