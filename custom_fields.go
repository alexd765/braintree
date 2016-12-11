package braintree

import (
	"encoding/xml"
	"errors"
	"fmt"
)

// CustomFields contains braintree custom fields as a map
type CustomFields map[string]string

// MarshalXML encodes CustomFields into xml
func (cf CustomFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tokens := []xml.Token{start}

	for key, value := range cf {
		start := xml.StartElement{Name: xml.Name{Local: key}}
		data := xml.CharData(value)
		end := xml.EndElement{Name: xml.Name{Local: key}}
		tokens = append(tokens, start, data, end)
	}

	tokens = append(tokens, xml.EndElement{Name: start.Name})

	for _, token := range tokens {
		if err := e.EncodeToken(token); err != nil {
			return err
		}
	}

	return e.Flush()
}

// UnmarshalXML decodes XML into CustomFields
func (cf *CustomFields) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	result := map[string]string{}

	key := ""
	value := ""

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch token := token.(type) {
		case xml.StartElement:
			key = token.Name.Local
		case xml.CharData:
			value = fmt.Sprintf("%s", token)
		case xml.EndElement:
			switch token.Name.Local {
			case key:
				result[key] = value
			case start.Name.Local:
				*cf = result
				return nil
			default:
				return errors.New("unexpected end element")
			}
		}
	}
}
