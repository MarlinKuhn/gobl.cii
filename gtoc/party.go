package gtoc

import (
	"fmt"

	"github.com/invopop/gobl/org"
)

// NewParty creates the SellerTradeParty part of a EN 16931 compliant invoice
func NewParty(party *org.Party) *Party {
	if party == nil {
		return nil
	}
	p := &Party{
		Name:                      party.Name,
		Contact:                   newContact(party),
		PostalTradeAddress:        newPostalTradeAddress(party.Addresses),
		URIUniversalCommunication: newEmail(party.Emails),
	}

	if party.TaxID != nil {
		// Assumes VAT ID being used instead of possible tax number
		p.SpecifiedTaxRegistration = &SpecifiedTaxRegistration{
			ID:       party.TaxID.String(),
			SchemeID: "VA",
		}
	}

	return p
}

func newContact(p *org.Party) *Contact {
	if len(p.People) == 0 {
		return nil
	}

	c := new(Contact)
	if len(p.People) > 0 {
		c.PersonName = contactName(p.People[0].Name)
		if len(p.People[0].Emails) > 0 {
			c.Email = &Email{
				URIID: p.People[0].Emails[0].Address,
			}
		}
	}
	if len(p.Telephones) > 0 {
		c.Phone = &PhoneNumber{
			CompleteNumber: p.Telephones[0].Number,
		}
	}

	return c
}

func contactName(p *org.Name) string {
	g := p.Given
	s := p.Surname

	if g == "" && s == "" {
		return ""
	}
	if g == "" {
		return s
	}
	if s == "" {
		return g
	}
	return fmt.Sprintf("%s %s", g, s)
}
