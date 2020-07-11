package zoia

import (
	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/caption"
)

// A Page holds details about page of a patch.
type Page struct {
	Raw  binary.Raw
	Name string
}

// ParsePage returns a new Header.
func ParsePage(page binary.Raw) *Page {
	p := &Page{
		Raw:  page,
		Name: caption.TrimToString(page[0:binary.SizeCaption]),
	}

	return p
}
