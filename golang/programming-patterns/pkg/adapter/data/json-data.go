package data

type JSONDocument struct {
}

func (doc *JSONDocument) ConvertToXML() string {
	return "<xml></xml>"
}

type JSONDocumentAdapter struct {
	JSONDocument
}

func (adapter JSONDocumentAdapter) SendXMLDocument() {
	adapter.ConvertToXML()
	print("Sending XML...")
}
