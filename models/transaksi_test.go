package models

import (
	"testing"
	"time"

	"github.com/agilsyofian/golang/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomTransaksi(t *testing.T) (*Transaksi, error) {
	kontrak, err := createRandomKontrak(t)
	require.NoError(t, err)
	require.NotEmpty(t, kontrak)

	newUUID, _ := uuid.NewRandom()

	payload := Transaksi{
		ID:        newUUID,
		KontrakNo: kontrak.No,
		Jenis:     "debit",
		Jml:       float64(util.RandomInt(100000, 500000)),
		CreatedAt: time.Now(),
	}
	result, err := testQueries.TransaksiCreate(payload)
	return result, err
}

func TestCreateRandomTransaksi(t *testing.T) {
	result, err := createRandomTransaksi(t)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestGetTransaksiByKontrak(t *testing.T) {
	transaksi, err := createRandomTransaksi(t)
	require.NoError(t, err)
	require.NotEmpty(t, transaksi)

	result, err := testQueries.TransaksiGetByKontrak(transaksi.KontrakNo)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestGetTransaksiID(t *testing.T) {
	transaksi, err := createRandomTransaksi(t)
	require.NoError(t, err)
	require.NotEmpty(t, transaksi)

	result, err := testQueries.TransaksiGetByID(transaksi.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, transaksi.KontrakNo, result.KontrakNo)
}
