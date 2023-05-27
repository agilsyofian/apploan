package models

import (
	"testing"
	"time"

	"github.com/agilsyofian/golang/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateRegister(t *testing.T) {
	newUUID, _ := uuid.NewRandom()
	plainPass := util.RandomString(6)
	password, _ := util.HashPassword(plainPass)
	var konsumen Konsumen = Konsumen{
		ID:          newUUID,
		Username:    util.RandomString(6),
		Password:    password,
		NIK:         util.RandomInt(332105010798001, 832105010798001),
		FullName:    util.RandomString(20),
		LegalName:   util.RandomString(20),
		TempatLahir: util.RandomString(9),
		TglLahir:    "1993-09-09",
		Gaji:        float64(util.RandomInt(10000000, 15000000)),
		FotoKTP:     util.RandomString(20),
		FotoSelfie:  util.RandomString(20),
	}

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

	result, err := testQueries.Register(konsumen, limits)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, konsumen.ID, result.Konsumen.ID)
	require.Equal(t, 4, len(result.Limit))
}
