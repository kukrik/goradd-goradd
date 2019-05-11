//** This file was code generated by got. DO NOT EDIT. ***

package browsertest

import (
	"bytes"
	"context"
)

func (form *JsUnitForm) AddHeadTags() {
	form.FormBase.AddHeadTags()
	if "Goradd JavaScript Unit Test" != "" {
		form.Page().SetTitle("Goradd JavaScript Unit Test")
	}

	// double up to deal with body attributes if they exist
	form.Page().BodyAttributes = ``
}

func (form *JsUnitForm) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
<h1>JavaScript Unit Tests</h1>
`)

	buf.WriteString(`
`)

	{
		err := form.RunButton.Draw(ctx, buf)
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
		err := form.Results.Draw(ctx, buf)
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
`)

	buf.WriteString(`
`)

	return
}
