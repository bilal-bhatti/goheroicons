// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.646
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowDownOnSquareStack(styleClass string) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{styleClass}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\" data-slot=\"icon\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/s20/solid/arrow-down-on-square-stack.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><path fill-rule=\"evenodd\" d=\"M8 1a.75.75 0 0 1 .75.75V6h-1.5V1.75A.75.75 0 0 1 8 1Zm-.75 5v3.296l-.943-1.048a.75.75 0 1 0-1.114 1.004l2.25 2.5a.75.75 0 0 0 1.114 0l2.25-2.5a.75.75 0 0 0-1.114-1.004L8.75 9.296V6h2A2.25 2.25 0 0 1 13 8.25v4.5A2.25 2.25 0 0 1 10.75 15h-5.5A2.25 2.25 0 0 1 3 12.75v-4.5A2.25 2.25 0 0 1 5.25 6h2ZM7 16.75v-.25h3.75a3.75 3.75 0 0 0 3.75-3.75V10h.25A2.25 2.25 0 0 1 17 12.25v4.5A2.25 2.25 0 0 1 14.75 19h-5.5A2.25 2.25 0 0 1 7 16.75Z\" clip-rule=\"evenodd\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
