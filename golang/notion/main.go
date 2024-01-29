package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	TitleType       string = "title"
	MultiSelectType string = "multi_select"
	CheckboxType    string = "checkbox"
	TextType        string = "text"
	ParagraphType   string = "paragraph"
	Heading2Type    string = "heading_2"
)

type Text struct {
	Content string `json:"content,omitempty"`
}

type Title struct {
	Type string `json:"type"`
	Text Text   `json:"text,omitempty"`
}

type Color string

var (
	Red    Color = "red"
	Blue   Color = "blue"
	Green  Color = "green"
	Purple Color = "purple"
	Pink   Color = "pink"
)

type MultiSelect struct {
	Name  string `json:"name,omitempty"`
	Color Color  `json:"color,omitempty"`
}

type Property struct {
	Type        string        `json:"type"`
	Title       []Title       `json:"title,omitempty"`
	MultiSelect []MultiSelect `json:"multi_select,omitempty"`
	Checkbox    bool          `json:"checkbox,omitempty"`
}

type RichText struct {
	Type string `json:"type,omitempty"`
	Text Text   `json:"text"`
}

type Paragraph struct {
	RichText []RichText `json:"rich_text,omitempty"`
}

type Heading2 struct {
	RichText []RichText `json:"rich_text,omitempty"`
}

type Block struct {
	Object    string     `json:"object"`
	Type      string     `json:"type"`
	Paragraph *Paragraph `json:"paragraph,omitempty"`
	Heading2  *Heading2  `json:"heading_2,omitempty"`
}

type Parent struct {
	Type       string `json:"type,omitempty"`
	DatabaseId string `json:"database_id,omitempty"`
}

type Page struct {
	Parent     Parent              `json:"parent"`
	Properties map[string]Property `json:"properties,omitempty"`
	Children   []Block             `json:"children,omitempty"`
}

func main() {
	body := Page{
		Parent: Parent{
			Type:       "database_id",
			DatabaseId: "42bc30c24c184db0a890018c009b69fd",
		},
		Properties: map[string]Property{
			"Name": {
				Type: TitleType,
				Title: []Title{
					{
						Type: TextType,
						Text: Text{
							Content: "test",
						},
					},
				},
			},
			"Tags": {
				Type: MultiSelectType,
				MultiSelect: []MultiSelect{
					{
						Name:  "Test",
						Color: Pink,
					},
				},
			},
		},
		Children: []Block{
			{
				Object: "block",
				Type:   Heading2Type,
				Heading2: &Heading2{
					RichText: []RichText{
						{
							Type: TextType,
							Text: Text{
								Content: "123",
							},
						},
					},
				},
			},
		},
	}

	res, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(res))

}
