package gtoc

import (
	"encoding/xml"
	"fmt"

	"github.com/invopop/gobl"
	"github.com/invopop/gobl/bill"
)

// Date defines date in the UDT structure
type Date struct {
	Date   string `xml:",chardata"`
	Format string `xml:"format,attr,omitempty"`
}

type Conversor struct {
	cii *Document
}

func NewConversor() *Conversor {
	c := new(Conversor)
	c.cii = new(Document)
	return c
}

func (c *Conversor) GetDocument() *Document {
	return c.cii
}

// ConvertToCII converts a GOBL envelope into a CIIdocument
func (c *Conversor) ConvertToCII(env *gobl.Envelope) (*Document, error) {
	inv, ok := env.Extract().(*bill.Invoice)
	if !ok {
		return nil, fmt.Errorf("invalid type %T", env.Document)
	}

	transaction, err := NewTransaction(inv)
	if err != nil {
		return nil, err
	}

	ciiDoc := Document{
		RSMNamespace:           RSM,
		RAMNamespace:           RAM,
		QDTNamespace:           QDT,
		UDTNamespace:           UDT,
		BusinessProcessContext: BusinessProcess,
		GuidelineContext:       GuidelineContext,
		ExchangedDocument:      NewHeader(inv),
		Transaction:            transaction,
	}

	c.cii = &ciiDoc
	return c.cii, nil
}

// Bytes returns the XML representation of the document in bytes
func (d *Document) Bytes() ([]byte, error) {
	bytes, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return nil, err
	}
	return append([]byte(xml.Header), bytes...), nil
}
