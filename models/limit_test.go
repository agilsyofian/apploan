package models

import (
	"testing"
	"time"

	"github.com/agilsyofian/golang/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomLimit(t *testing.T) ([]Limit, uuid.UUID, error) {
	konsumen, _, err := createRandomKonsumen(t)
	require.NoError(t, err)
	require.NotEmpty(t, konsumen)

	var limits []Limit
	for i := 1; i <= 4; i++ {
		limits = append(limits, Limit{
			KonsumenID: konsumen.ID,
			Tenor:      int64(i),
			Limit:      float64(util.RandomInt(100000, 4000000)),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		})
	}

	result, err := testQueries.LimitCreate(limits)
	return result, konsumen.ID, err
}

func TestCreateRandomLimit(t *testing.T) {
	result, _, err := createRandomLimit(t)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestGetLimitKonsumen(t *testing.T) {
	limit, konsumenID, err := createRandomLimit(t)
	require.NoError(t, err)
	require.NotEmpty(t, limit)

	result, err := testQueries.LimitGetByKonsumen(konsumenID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestGetLimitID(t *testing.T) {
	limit, _, err := createRandomLimit(t)
	require.NoError(t, err)
	require.NotEmpty(t, limit)

	result, err := testQueries.LimitGetByID(limit[0].ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, limit[0].ID, result.ID)
}

func TestUpdateLimit(t *testing.T) {
	limit, _, err := createRandomLimit(t)
	require.NoError(t, err)
	require.NotEmpty(t, limit)

	limitUpdate := float64(util.RandomInt(100000, 4000000))

	updateData := &Limit{
		Tenor: util.RandomInt(1, 4),
		Limit: limitUpdate,
	}
	result, err := testQueries.LimitUpdate(limit[0].ID, updateData)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, limitUpdate, result.Limit)
}
