package utils

import (
	"scsystem/pkg/jwt"
	"testing"
)

func TestHash(t *testing.T) {
	pass, _ := jwt.GenPassword("12345")
	pass_str := string(pass)
	print(pass_str)
	print(jwt.ComparePassword("$12$GLl9WEfPrgJvDUasNuJc/eKeN013tmojMzQcgOB1NuaxtG0ACU1rC", "12345"))
}
