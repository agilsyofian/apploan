package utilitize

import (
	"time"

	"github.com/agilsyofian/kreditplus/models"
)

type LimitFactory struct {
	Konsumen models.Konsumen
}

func NewFactoryLimit(k models.Konsumen) *LimitFactory {
	return &LimitFactory{
		Konsumen: k,
	}
}

func (l *LimitFactory) BuildLimit() []models.Limit {
	var result []models.Limit

	rpc := 0.30 * l.Konsumen.Gaji
	currentTime := time.Now()

	for i := 1; i <= 4; i++ {
		result = append(result, models.Limit{
			KonsumenID: l.Konsumen.ID,
			Tenor:      int64(i),
			Limit:      float64(i * int(rpc)),
			CreatedAt:  currentTime,
			UpdatedAt:  currentTime,
		})
	}

	return result
}
