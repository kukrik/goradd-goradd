package control_base

import (
	"github.com/microcosm-cc/bluemonday"
	bootstrapBase "github.com/spekary/goradd/bootstrap/control/control_base"
	"github.com/spekary/goradd/html"
	"github.com/spekary/goradd/page"
	gr_control_base "github.com/spekary/goradd/page/control/control_base"
)

var sanitizer *bluemonday.Policy

type TextboxI interface {
	gr_control_base.TextboxI
}

// The local Textbox override. All textboxes will descend from this one. You can make changes here that wil impact
// all the text fields in the system.
type Textbox struct {
	gr_control_base.Textbox
}

func NewTextbox(parent page.ControlI, id string) *Textbox {
	t := &Textbox{}
	t.Init(t, parent, id)
	return t
}

func (t *Textbox) Init(self gr_control_base.TextboxI, parent page.ControlI, id string) {
	t.Textbox.Init(self, parent, id)
	t.Textbox.SetSanitizer(sanitizer)
}

func (t *Textbox) DrawingAttributes() *html.Attributes {
	a := t.Textbox.DrawingAttributes()

	if t.HasWrapper() {
		bootstrapBase.FormControlAddAttributes(t, a) // extend all textboxes to be form controls
	}

	return a
}

func init() {
	sanitizer = bluemonday.UGCPolicy() // Create a standard sanitizer.
}