package data_service

import "fmt"

type AnalyticalDataService interface {
	SendXMLDocument()
}

type XMLDocument struct {
}

func (doc *XMLDocument) SendXMLDocument() {
	fmt.Println("Sending XML...")
}
