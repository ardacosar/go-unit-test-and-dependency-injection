package main

import (
	"errors"
)

func SagligiKontrolEt(enjecteEdilenBarinak Barinak) (string, error) {
	getirilenHayvan, err := enjecteEdilenBarinak.HayvaniGetir()
	if err != nil {
		return "", errors.New("araba kaza yaptı")
	}

	if getirilenHayvan.Saglik == "İyi" {
		return "Sağlıklı", nil
	} else {
		return "Sağlıksız", nil
	}
}
