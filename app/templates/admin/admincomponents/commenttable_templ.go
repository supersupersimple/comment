// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package admincomponents

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/supersupersimple/comment/app/model"

func CommentsTable(comments []*model.AdminComment) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"overflow-x-auto\"><table class=\"table\"><!-- head --><thead><tr><th class=\"w-1/6\">Operation</th><th>PublishTime</th><th>Status</th><th class=\"w-1/4 xl:w-1/3\">Content</th><!-- added w-1/3 xl:w-1/2 classes to set width --><th>Username</th><th>Email</th><th>PageTitle</th><th>PageSlug</th><th>ID</th><th>Depth</th><th>ReplyTo</th></tr></thead> <tbody>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, comment := range comments {
			templ_7745c5c3_Err = CommentRow(comment).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
