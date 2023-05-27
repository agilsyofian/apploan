package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTagihanExist(t *testing.T) {
	createRandomKontrak, err := createRandomKontrak(t)
	require.NoError(t, err)
	require.NotEmpty(t, createRandomKontrak)
}

func TestTagihanList(t *testing.T) {
	createRandomKontrak, err := createRandomKontrak(t)
	require.NoError(t, err)
	require.NotEmpty(t, createRandomKontrak)

	result, err := testQueries.TagihanList(createRandomKontrak.Kontrak.No)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestTagihanUpdateBatch(t *testing.T) {
	createRandomKontrak, err := createRandomKontrak(t)
	require.NoError(t, err)
	require.NotEmpty(t, createRandomKontrak)

	var ids []string
	for _, tg := range createRandomKontrak.Tagihan {
		ids = append(ids, tg.ID.String())
	}

	tgl_paid := time.Now().Format("2006-01-02")
	payload := Tagihan{
		Status:  "paid",
		TglPaid: &tgl_paid,
	}
	err = testQueries.TagihanUpdateBatch(ids, payload)
	require.NoError(t, err)
}
