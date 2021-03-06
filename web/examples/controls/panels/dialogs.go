package panels

import (
	"context"
	"github.com/goradd/goradd/pkg/page"
	"github.com/goradd/goradd/pkg/page/action"
	. "github.com/goradd/goradd/pkg/page/control"
)

type DialogsPanel struct {
	Panel
}

const (
	ButtonAlert = iota + 11010
	ButtonMessage
	MessageAction
)

func NewDialogsPanel(ctx context.Context, parent page.ControlI) {
	p := &DialogsPanel{}
	p.Self = p
	p.Init(ctx, parent, "checkboxPanel")
}

func (p *DialogsPanel) Init(ctx context.Context, parent page.ControlI, id string) {
	p.Panel.Init(parent, "dialogsPanel")
	p.AddControls(ctx,
		PanelCreator{
			ID:             "result",
		},
		ButtonCreator{
			ID:       "alertButton",
			Text:     "Alert",
			OnSubmit: action.Ajax("dialogsPanel", ButtonAlert),
		},
		ButtonCreator{
			ID:       "messageButton",
			Text:     "Message",
			OnSubmit: action.Server("dialogsPanel", ButtonMessage),
		},

	)

	// This is really specific to this demo because we are switching back and forth between this and bootstrap dialogs.
	// You do not normally need to do this.
	RestoreNewDialogFunction()
}


func (p *DialogsPanel) Action(ctx context.Context, a page.ActionParams) {
	switch a.ID {
	case ButtonAlert:
		Alert(p, "Look out!", nil)
	case ButtonMessage:
		// GetDialogPanel here will create a new one if one is not already in the form, or will retrieve an old
		// one if its just hidden
		dp,isNew := GetDialogPanel(p, "msg-dlg")

		if isNew {
			// Set up the dialog since it was just created.
			dp.SetTitle("What do you want for Christmas?")
			tb := NewTextbox(dp, "msg-txt")
			tb.SetIsRequired(true)
			dp.AddButton("For Me", "forme", &DialogButtonOptions{
				IsPrimary:true,
				Validates:true,
				//OnClick:action.Ajax(p.ID(), ForMeAction), // You can handle button actions this way
			})
			dp.AddButton("For You", "foryou", &DialogButtonOptions{
				Validates:true,
				//OnClick:action.Ajax(p.ID(), ForYouAction),
			})
			dp.AddCloseButton("Cancel", "cancel")
			dp.OnButton(action.Ajax(p.ID(), MessageAction)) // or handle button actions this way
		} else {
			GetTextboxI(p, "msg-txt").SetText("") // reset the text in case it was just hidden
		}
		dp.Show()
	case MessageAction:
		// A dialog button was pressed
		btnID := a.EventValueString()
		switch btnID {
		case "forme":
			GetPanel(p,"result").SetText("You want to get a " + GetTextboxI(p, "msg-txt").Text())
		case "foryou":
			GetPanel(p,"result").SetText("You want to give a " + GetTextboxI(p, "msg-txt").Text())
		}
		dp,_ := GetDialogPanel(p, "msg-dlg")
		dp.Hide()
	}
}


func init() {
	page.RegisterControl(&DialogsPanel{})
}

