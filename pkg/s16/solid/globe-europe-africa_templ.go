// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.646
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func GlobeEuropeAfrica() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" fill=\"currentColor\" aria-hidden=\"true\" data-slot=\"icon\"><path fill-rule=\"evenodd\" d=\"M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1ZM5.657 3.023a5.5 5.5 0 1 0 7.584 3.304l-.947-.63a.431.431 0 0 0-.544.053.431.431 0 0 1-.544.054l-.467-.312a.475.475 0 0 0-.689.608l.226.453a2.119 2.119 0 0 1 0 1.894L10.1 8.8a.947.947 0 0 0-.1.424v.11a2 2 0 0 1-.4 1.2L8.8 11.6A1 1 0 0 1 7 11v-.382a1 1 0 0 0-.553-.894l-.422-.212A1.854 1.854 0 0 1 6.855 6h.707a.438.438 0 1 0-.107-.864l-.835.209a1.129 1.129 0 0 1-1.305-1.553l.342-.77Z\" clip-rule=\"evenodd\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}