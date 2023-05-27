package utilitize

type Limit struct {
	Tenor int64   `json:"tenor"`
	Limit float64 `json:"limit"`
}

type LimitFactory struct {
	PersenGaji float64
	Gaji       float64
	Cicilan    float64
}

func NewFactoryLimit(persenGaji, gaji, cicilan float64) *LimitFactory {
	return &LimitFactory{
		PersenGaji: persenGaji,
		Gaji:       gaji,
		Cicilan:    cicilan,
	}
}

func (l *LimitFactory) BuildLimit() []Limit {
	var result []Limit

	rpc := l.BuildRPC()
	for i := 1; i <= 4; i++ {
		result = append(result, Limit{
			Tenor: int64(i),
			Limit: float64(i * int(rpc)),
		})
	}

	return result
}

func (l *LimitFactory) BuildRPC() float64 {
	return l.PersenGaji*l.Gaji - l.Cicilan
}
