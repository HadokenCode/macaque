package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("varA", "valueA")
	assert.Equal(t, "valueA", GetEnv(("varA"), "valueB"))
	assert.Equal(t, "valueB", GetEnv(("varB"), "valueB"))
}
