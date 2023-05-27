package utilitize

type Kontrak struct {
	Otr   float64
	Tenor int64
	Bunga float64
	Fee   float64
}

type KontrakCalculation struct {
	AdminFee   float64
	JmlCicilan float64
	JmlBunga   float64
}

func NewKontrak(otr, bunga, fee float64, tenor int64) *Kontrak {
	return &Kontrak{
		Otr:   otr,
		Tenor: tenor,
		Bunga: bunga,
		Fee:   fee,
	}
}

func (k *Kontrak) BuildKontrak() KontrakCalculation {
	var result KontrakCalculation

	result.AdminFee = k.Fee * k.Otr

	pinjamanPerbulan := k.Otr / float64(k.Tenor)
	bunga := k.Otr * k.Bunga
	bungaPerBulan := bunga / float64(k.Tenor)

	result.JmlBunga = bunga
	result.JmlCicilan = (pinjamanPerbulan + bungaPerBulan) * float64(k.Tenor)

	return result
}
