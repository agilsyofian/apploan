package models

import (
	"fmt"
	"testing"

	"github.com/agilsyofian/golang/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomKontrak(t *testing.T) (*KontrakResponse, error) {
	konsumen, _, err := createRandomKonsumen(t)
	require.NoError(t, err)
	require.NotEmpty(t, konsumen)

	newUUID, _ := uuid.NewRandom()

	payload := Kontrak{
		No:         newUUID,
		KonsumenID: konsumen.ID,
		Otr:        float64(util.RandomInt(1000000, 5000000)),
		AdminFee:   float64(util.RandomInt(100000, 500000)),
		JmlCicilan: float64(util.RandomInt(100000, 500000)),
		JmlBunga:   float64(util.RandomInt(0, 1)),
		NamaAsset:  util.RandomString(10),
		Status:     "inpg",
	}
	result, err := testQueries.KontrakCreate(payload)
	return result, err
}

func TestCreateRandomKontrak(t *testing.T) {
	result, err := createRandomKontrak(t)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestGetKontrakKonsumen(t *testing.T) {
	kontrak, err := createRandomKontrak(t)
	require.NoError(t, err)
	require.NotEmpty(t, kontrak)

	result, err := testQueries.KontrakGetByKonsumen(kontrak.Kontrak.KonsumenID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestGetKontrakID(t *testing.T) {
	kontrak, err := createRandomKontrak(t)
	require.NoError(t, err)
	require.NotEmpty(t, kontrak)

	result, err := testQueries.KontrakGetByID(kontrak.Kontrak.No)
	fmt.Println(result.No)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, kontrak.Kontrak.No, result.No)
}

func TestUpdateKontrak(t *testing.T) {
	kontrak, err := createRandomKontrak(t)
	require.NoError(t, err)
	require.NotEmpty(t, kontrak)

	jmlCicilan := float64(util.RandomInt(100000, 4000000))

	updateData := &Kontrak{
		JmlCicilan: jmlCicilan,
		Status:     "cancel",
	}
	result, err := testQueries.KontrakUpdate(kontrak.Kontrak.No, updateData)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, jmlCicilan, result.JmlCicilan)
	require.Equal(t, "cancel", string(result.Status))
}
