package main 

func main () {
	hpPrinter := &HpPrinter{}
	hpAdapter := &HpAdapter{printer: hpPrinter}



	canonPrinter := &CanonPrinter{}
	canonAdapter := &CanonAdapter{
		printer: canonPrinter,
	}

	printer := Printer{}
	printer.print(hpAdapter)
	printer.print(canonAdapter)

	
}