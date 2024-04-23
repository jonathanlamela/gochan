package utils

func SalutaChan(canale chan string) {
	canale <- "Ciao mondo"
}
