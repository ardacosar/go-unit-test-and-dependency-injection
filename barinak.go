package main

type hayvan struct {
	Ad     string
	Tur    string
	Saglik string
}

type Barinak interface {
	HayvaniGetir() (hayvan, error)
}

type barinak struct {
	Canli hayvan
}

var barinak1 barinak

func BarinakOlustur() {
	barinak1 = barinak{
		Canli: hayvan{
			Ad:     "Pamuk",
			Tur:    "Kedi",
			Saglik: "İyi",
		},
	}
}

func (b *barinak) HayvaniGetir() (hayvan, error) {
	//return hayvan{}, errors.New("herhangi bir sıkıntı olabilir")
	return b.Canli, nil
}

type sahteBarinak struct {
	saglik string
}

func (b *sahteBarinak) HayvaniGetir() (hayvan, error) {
	return hayvan{
		Ad:     "test",
		Tur:    "test",
		Saglik: b.saglik,
	}, nil
}
