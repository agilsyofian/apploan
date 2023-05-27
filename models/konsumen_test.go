package models

import (
	"testing"

	"github.com/agilsyofian/golang/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomKonsumen(t *testing.T) (*Konsumen, string, error) {
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
	result, err := testQueries.CreateKonsumen(konsumen)
	return result, plainPass, err
}

func TestAuthKonsumen(t *testing.T) {
	konsumen, _, err := createRandomKonsumen(t)
	require.NoError(t, err)
	require.NotEmpty(t, konsumen)

	result, err := testQueries.AuthKonsumen(konsumen.Username)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, konsumen.Username, result.Username)
}

func TestCreateKonsumen(t *testing.T) {
	konsumen, plainPass, err := createRandomKonsumen(t)
	require.NoError(t, err)
	require.NotEmpty(t, konsumen)
	err = util.CheckPassword(plainPass, konsumen.Password)
	require.NoError(t, err)
}

func TestGetKonsumen(t *testing.T) {
	konsumen, _, err := createRandomKonsumen(t)
	require.NoError(t, err)

	result, err := testQueries.GetKonsumen(konsumen.ID)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, konsumen.Username, result.Username)
}

func TestGetListKonsumen(t *testing.T) {
	n := 10
	for i := 0; i < n; i++ {
		createRandomKonsumen(t)
	}
	result, err := testQueries.GetListKonsumen(1, 5)
	require.NoError(t, err)
	require.Equal(t, 5, len(result))
}

func TestUpdateKonsumen(t *testing.T) {
	konsumen, _, err := createRandomKonsumen(t)
	require.NoError(t, err)

	editedFullname := util.RandomString(10)
	var updateDataKonsumen Konsumen = Konsumen{
		FullName: editedFullname,
	}
	result, err := testQueries.UpdateKonsumen(konsumen.ID, &updateDataKonsumen)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, editedFullname, result.FullName)
}

func TestSoftDeleteKonsumen(t *testing.T) {
	konsumen, _, err := createRandomKonsumen(t)
	require.NoError(t, err)

	result, err := testQueries.SoftDeleteKonsumen(konsumen.ID)

	require.NoError(t, err)
	require.NotEmpty(t, result)
}
