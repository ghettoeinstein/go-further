package main

import (
	"fmt"
)

type messagePrinterInterface interface {
	Print()
}


func PrintSomething(mp messagePrinterInterface) {
    mp.Print()
}
type MessagePrinterExample struct {}

func (m MessagePrinterExample) Print() {
    fmt.Println("Hello Caleb from MessagePrinterExample")
}

type ErrorMessagePrinterExample struct {}

func (m ErrorMessagePrinterExample) Print() {
    fmt.Println("Hello Caleb  from ErrorMessagePrinterExample ")
}




func main() {
	mpe := MessagePrinterExample{}
    PrintSomething(mpe)

	empe := ErrorMessagePrinterExample{}
	PrintSomething(empe)

	
}