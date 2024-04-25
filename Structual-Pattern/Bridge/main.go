package main

func main() {
	hpPrinter := &HP{}
	//epsonPrinter := &Epson{}

	macComputer := &Mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()

}
