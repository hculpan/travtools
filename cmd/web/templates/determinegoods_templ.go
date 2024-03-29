// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func DetermineGoods() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container\"><h1>Step 2: Determine Goods Available</h1><form hx-get=\"/spec-trade-purchase\" class=\"mb-4\" hx-target=\"#content\" hx-swap=\"innerHTML\"><input type=\"text\" name=\"step\" id=\"step\" value=\"3\" hidden><div class=\"mb-4\"><label class=\"form-label\">System Population Code:</label> <input type=\"text\" class=\"form-control\" id=\"population\" name=\"population\" placeholder=\"9\" style=\"max-width: 250px;\"></div><div class=\"mb-4\"><label class=\"form-label\">System Trade Codes:</label> <input type=\"text\" class=\"form-control\" id=\"trade\" name=\"trade\" placeholder=\"\" style=\"max-width: 250px;\"><div class=\"form-text text-secondary\">Input two-letter codes only, separated by commas</div></div><div class=\"mb-4\"><label class=\"form-label\">Broker's Skill:</label> <input type=\"text\" class=\"form-control\" id=\"broker\" name=\"broker\" placeholder=\"\" style=\"max-width: 250px;\"><div class=\"form-text text-secondary\">Input the Broker skill level, not the result of a Broker roll</div></div><div class=\"mb-4\"><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"goodstype\" id=\"legal-goods\" value=\"legal-goods\" checked> <label class=\"form-check-label\" for=\"legal-goods\">Legal Goods</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"goodstype\" id=\"illegal-goods\" value=\"illegal-goods\"> <label class=\"form-check-label\" for=\"illegal-goods\">Illegal/Black Market Goods</label></div></div><button class=\"btn btn-primary\" type=\"submit\">Next</button></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
