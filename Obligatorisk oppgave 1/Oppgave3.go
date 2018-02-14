package main

import "fmt"
import "os"
import "os/signal"
import "syscall"



//Skriv et program som består av en evig løkke. Hvor mye minne og CPU bruker programmet når det kjører.
// Programmet skal skrive ut en avslutningsmeld- ing når programmet mottar et SIGINT signal.
// Generer ulike avslutningssig- naler til prosessen,
// og dokumenter hvilke avslutningskommandoer programmet håndterer og som trigger avslutningsmeldingen.



func main() {



	// Go signal notification fungerer med den gitte kanalen

	sig := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` registrer gitte kanalen.

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Denne rutinen blocker signalet til signalet er mottat. Jeg har gjort det sånn at når signalet
	// Blir sendt så starter "for" løkken.

	go func() {
		sig := <-sig
		fmt.Println()
		fmt.Println(sig)
		done <- true
		for
		{
			fmt.Println("Running")


		}

		}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

	// defer vil aldri bli kjørt hvis sysexit blir gitt
	defer fmt.Println("!")

	// Exit med status 3 SYSEXIT.
	os.Exit(3)


// Prossesen kjører så kort at det er ikke så synlig hvor mye prossesorkraft den bruker.
// Men den bruker et fast minne som den har blitt tildelt når prossesen starter.

}
