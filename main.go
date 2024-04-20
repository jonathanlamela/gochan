package main

import (
	"fmt"
	"time"
)

func SalutaChan(canale chan string) {
	canale <- "Ciao mondo"
}

func main() {
	currentTime := time.Now()

	fmt.Printf("-- Task avviato %s --\n", currentTime.Format("02-01-2006 15:04:05"))

	//Crea un canale di tipo string
	channel1 := make(chan string)

	//Crea un canale di tipo int
	channel2 := make(chan int)

	//Chiude il canale alla fine della funzione main
	defer close(channel1)

	defer close(channel2)

	//Lancia una routine che esegue la funzione SalutaChan
	go SalutaChan(channel1)

	//Ricevo la risposta da chan
	risposta := <-channel1

	//Stampo la variabile risposta
	fmt.Println(risposta)

	//Creo una routine che contiene una funzione che sta in ascolto del channel
	go func() {
		for m := range channel2 {
			fmt.Println(m)
		}
	}()

	//Invio dei numeri al canale 2
	for i := 0; i <= 100; i++ {
		channel2 <- i * 2
	}

	fmt.Printf("-- Task completato %s --\n", currentTime.Format("02-01-2006 15:04:05"))

}
