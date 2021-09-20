package presenters

import "fmt"

func PrintEntitlement(eulaAccepted, eligableToDownload bool) {
	fmt.Printf("\nEula Accepted:         %t\n", eulaAccepted)
	fmt.Printf("Eligable to Download:  %t\n\n", eligableToDownload)
}