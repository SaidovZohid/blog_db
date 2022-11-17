package utils

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := "1234567"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, password)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}