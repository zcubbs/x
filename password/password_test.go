package password

import (
	"github.com/stretchr/testify/require"
	"github.com/zcubbs/x/random"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	pwd := random.String(6)

	hashedPassword, err := Hash(pwd)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = Check(pwd, hashedPassword)
	require.NoError(t, err)

	wrongPass := random.String(6)
	err = Check(wrongPass, hashedPassword)
	require.ErrorIs(t, err, bcrypt.ErrMismatchedHashAndPassword)

	hashedPassword2, err := Hash(pwd)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword, hashedPassword2)
}
