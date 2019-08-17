//** This file was code generated by got. DO NOT EDIT. ***

package browsertest

import (
	"bytes"
	"context"
)

func (ctrl *JsUnitForm) AddHeadTags() {
	ctrl.FormBase.AddHeadTags()
	if "Goradd JavaScript Unit Test" != "" {
		ctrl.Page().SetTitle("Goradd JavaScript Unit Test")
	}

	// double up to deal with body attributes if they exist
	ctrl.Page().BodyAttributes = ``
}

func (ctrl *JsUnitForm) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
<h1>JavaScript Unit Tests</h1>
`)

	buf.WriteString(`
`)

	{
		err := ctrl.Page().GetControl("form.RunButton").Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
<h3>Results</h3>
`)

	buf.WriteString(`
`)

	{
		err := ctrl.Page().GetControl("form.Results").Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
<h3>Test Space</h3>
<p id = "intro">
The area below is used by the unit test code to exercise the code.
</p>
<div id="testspace">
<p id="testP">
I am here
</p>
<div id="testD" data-animal-type="bird" spellcheck>
a div
</div>
</div>
<ul id="listener"><li id="outer"><span id="inner">a span</span></li></ul>

`)

	buf.WriteString(`
`)

	return
}
