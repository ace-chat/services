package pkg

import (
	"fmt"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	code := GenerateCode()
	fmt.Println(code)
}
