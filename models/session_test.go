package models

import (
	"testing"
	"time"

	"github.com/agilsyofian/golang/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomSession(t *testing.T) (*Session, error) {
	konsumen, _, err := createRandomKonsumen(t)
	require.NoError(t, err)
	require.NotEmpty(t, konsumen)

	newUUID, _ := uuid.NewRandom()
	var sesData Session = Session{
		ID:           newUUID,
		KonsumenID:   konsumen.ID,
		RefreshToken: util.RandomString(60),
		UserAgent:    util.RandomString(20),
		ClientIP:     util.RandomString(10),
		ExpiredAt:    time.Now(),
		IsBlocked:    false,
	}
	result, err := testQueries.SessionCreate(sesData)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, konsumen.ID, result.KonsumenID)

	return result, err
}

func TestCreateRandomSession(t *testing.T) {
	result, err := createRandomSession(t)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestGetSession(t *testing.T) {
	session, _ := createRandomSession(t)

	result, err := testQueries.SessionGet(session.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, session.ID, result.ID)
	require.Equal(t, session.KonsumenID, result.KonsumenID)
}
