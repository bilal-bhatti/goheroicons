// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.646
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Radio() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\" data-slot=\"icon\"><path fill-rule=\"evenodd\" d=\"M17.45 3.473a.75.75 0 1 0-.4-1.446L5.313 5.265c-.84.096-1.671.217-2.495.362A2.212 2.212 0 0 0 1 7.816v7.934A2.25 2.25 0 0 0 3.25 18h13.5A2.25 2.25 0 0 0 19 15.75V7.816c0-1.06-.745-2-1.817-2.189a41.12 41.12 0 0 0-5.406-.59l5.673-1.564ZM16 9.5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0ZM14.5 16a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3Zm-9.26-5a.75.75 0 0 1 .75-.75H6a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75V11Zm2.75-.75a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75H8a.75.75 0 0 0 .75-.75V11a.75.75 0 0 0-.75-.75h-.01Zm-1.75-1.5A.75.75 0 0 1 6.99 8H7a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75v-.01Zm3.583.42a.75.75 0 0 0-1.06 0l-.007.007a.75.75 0 0 0 0 1.06l.007.007a.75.75 0 0 0 1.06 0l.007-.007a.75.75 0 0 0 0-1.06l-.007-.008Zm.427 2.08A.75.75 0 0 1 11 12v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75V12a.75.75 0 0 1 .75-.75h.01Zm-.42 3.583a.75.75 0 0 0 0-1.06l-.007-.007a.75.75 0 0 0-1.06 0l-.007.007a.75.75 0 0 0 0 1.06l.007.008a.75.75 0 0 0 1.06 0l.008-.008Zm-3.59.417a.75.75 0 0 1 .75-.75H7a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75v-.01Zm-1.013-1.484a.75.75 0 0 0-1.06 0l-.008.007a.75.75 0 0 0 0 1.06l.007.008a.75.75 0 0 0 1.061 0l.007-.008a.75.75 0 0 0 0-1.06l-.007-.007ZM3.75 11.25a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75V12a.75.75 0 0 1 .75-.75h.01Zm1.484-1.013a.75.75 0 0 0 0-1.06l-.007-.007a.75.75 0 0 0-1.06 0l-.007.007a.75.75 0 0 0 0 1.06l.007.007a.75.75 0 0 0 1.06 0l.007-.007ZM7.24 13a.75.75 0 0 1 .75-.75H8a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75V13Zm-1.25-.75a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75H6a.75.75 0 0 0 .75-.75V13a.75.75 0 0 0-.75-.75h-.01Z\" clip-rule=\"evenodd\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
