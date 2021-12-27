package main

import (
	"testing"
)

func TestSagligiKontrolEt(t *testing.T) {

	t.Run("Eğer hayvan sağlıklı ise", func(t *testing.T) {
		sahteBarinak1 := sahteBarinak{"İyi"}
		Beklenen := "Sağlıklı"

		Gercek, _ := SagligiKontrolEt(&sahteBarinak1)

		if Gercek != Beklenen {
			t.Error("Sağlık yanlış ölçüldü")
		}
	})

	t.Run("Eğer hayvan sağlıklız ise", func(t *testing.T) {
		sahteBarinak1 := sahteBarinak{"Kötü"}
		Beklenen := "Sağlıksız"

		Gercek, _ := SagligiKontrolEt(&sahteBarinak1)

		if Gercek != Beklenen {
			t.Error("Sağlık yanlış ölçüldü")
		}
	})

}
