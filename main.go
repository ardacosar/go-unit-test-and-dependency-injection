package main

import "fmt"

func main() {
	BarinakOlustur()

	Sonuc, err := SagligiKontrolEt(&barinak1)
	if err != nil {
		panic(err)
	}

	fmt.Println(Sonuc)
}
