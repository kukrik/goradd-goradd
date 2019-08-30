package panels

import (
	"context"
	. "github.com/goradd/goradd/pkg/bootstrap/control"
	"github.com/goradd/goradd/pkg/bootstrap/examples"
	"github.com/goradd/goradd/pkg/page"
	"github.com/goradd/goradd/pkg/page/action"
	"github.com/goradd/goradd/pkg/page/control"
	"github.com/goradd/goradd/pkg/url"
	"github.com/goradd/goradd/test/browsertest"
	"strings"
)

type SelectListPanel struct {
	control.Panel
}

func NewSelectListPanel(ctx context.Context, parent page.ControlI) {
	itemList := []control.ListValue{
		{"First", 1},
		{"Second", 2},
		{"Third", 3},
		{"Fourth", 4},
		{"Fifth", 5},
		{"Sixth", 6},
		{"Seventh", 7},
		{"Eighth", 8},
	}

	p := &SelectListPanel{}
	p.Panel.Init(p, parent, "selectListPanel")

	p.AddControls(ctx,
		FormGroupCreator{
			Label: "Standard SelectList",
			Child: SelectListCreator{
				ID: "singleSelectList",
				NilItem: "- Select One -",
				Items: itemList,
				ControlOptions:page.ControlOptions{
					IsRequired: true,
				},
			},
		},
		FormGroupCreator{
			Label: "SelectList With Size",
			Child: SelectListCreator{
				ID: "selectListWithSize",
				Items: itemList,
				Size: 4,
				ControlOptions:page.ControlOptions{
					IsRequired: true,
				},
			},
		},
		FormFieldsetCreator{
			Legend: "Radio List",
			Instructions: "A radio list",
			Child: RadioListCreator{
				ID: "radioList1",
				Items: itemList,
			},
		},
		FormFieldsetCreator{
			Legend: "Checkbox List",
			Child: CheckboxListCreator{
				ID: "checklist1",
				Items: itemList,
			},
		},
		ButtonCreator{
			ID:       "ajaxButton",
			Text:     "Submit Ajax",
			OnSubmit: action.Ajax("selectListPanel", ButtonSubmit),
		},
		ButtonCreator{
			ID:       "serverButton",
			Text:     "Submit Server",
			OnSubmit: action.Ajax("selectListPanel", ButtonSubmit),
		},

	)

}

func (p *SelectListPanel) Action(ctx context.Context, a page.ActionParams) {
	switch a.ID {
	case ButtonSubmit:
		GetFormFieldset(p, "checklist1-fs").SetInstructions(
			strings.Join(GetCheckboxList(p, "checklist1").SelectedIds(), ","))
	}
}


// testPlain exercises the plain text box
func testSelectListAjaxSubmit(t *browsertest.TestForm)  {
	var myUrl = url.NewBuilder(controlsFormPath).SetValue("control", "selectlist").String()
	f := t.LoadUrl(myUrl)

	testSelectListSubmit(t, f, f.Page().GetControl("ajaxButton"))

	t.Done("Complete")
}

func testSelectListServerSubmit(t *browsertest.TestForm)  {
	var myUrl = url.NewBuilder(controlsFormPath).SetValue("control", "selectlist").String()
	f := t.LoadUrl(myUrl)

	testSelectListSubmit(t, f, f.Page().GetControl("serverButton"))

	t.Done("Complete")
}

// testSelectListSubmit does a variety of submits using the given button. We use this to double check the various
// results we might get after a submission, as well as nsure that the ajax and server submits produce
// the same results.
func testSelectListSubmit(t *browsertest.TestForm, f page.FormI, btn page.ControlI) {

	// For testing purposes, we need to use the id of the list item, rather than the value of the list item,
	// since that is what is presented in the html.
	select1 := f.Page().GetControl("singleSelectList").(*SelectList)
	select2 := f.Page().GetControl("selectListWithSize").(*SelectList)
	radio1 := f.Page().GetControl("radioList1").(*RadioList)

	id,_ := select2.GetItemByValue(2)
	t.ChangeVal("selectListWithSize", id)

	t.Click(btn)

	t.AssertEqual(true, t.HasClass("singleSelectList-fg", "error"))


	t.AssertEqual(2, select2.IntValue())

	id,_ = select1.GetItemByValue(1)
	t.ChangeVal("singleSelectList", id)
	id,_ = select2.GetItemByValue(2)
	t.ChangeVal("selectListWithSize", id)
	id,_ = radio1.GetItemByValue(3)
	t.CheckGroup("radioList1", id)

	t.Click(btn)

	t.AssertEqual(1, select1.IntValue())
	t.AssertEqual(2, select2.IntValue())
	t.AssertEqual(3, radio1.IntValue())
	//t.AssertEqual(4, radio2.IntValue())
}

func init() {
	examples.RegisterPanel("lists", "Lists", NewSelectListPanel, 3)

	browsertest.RegisterTestFunction("Bootstrap Select List Ajax Submit", testSelectListAjaxSubmit)
	browsertest.RegisterTestFunction("Bootstrap Select List Server Submit", testSelectListServerSubmit)
}