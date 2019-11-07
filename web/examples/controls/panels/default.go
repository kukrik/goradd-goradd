package panels

import (
	"context"
	"github.com/goradd/goradd/pkg/page"
	. "github.com/goradd/goradd/pkg/page/control"
)

// shared
const controlsFormPath = "/goradd/examples/controls.g"

const (
	AjaxSubmit int = iota + 1
	ServerSubmit
	ButtonSubmit
	ResetStateSubmit
	ProxyClick
)


type DefaultPanel struct {
	Panel
}

func NewDefaultPanel(ctx context.Context, parent page.ControlI) {
	p := &DefaultPanel{}
	p.Self = p
	p.Panel.Init(parent, "defaultPanel")
}

func init() {
	page.RegisterControl(&DefaultPanel{})
}