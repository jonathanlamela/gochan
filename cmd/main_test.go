package main_test

import (
	"gochan/pkgs/utils"
	"testing"
)

func TestSalutaChan(t *testing.T) {
	channel1 := make(chan string)
	defer close(channel1)

	go utils.SalutaChan(channel1)

	//Ricevo la risposta da chan
	risposta := <-channel1

	if risposta != "Ciao mondo" {
		t.Fail()
	}

}
