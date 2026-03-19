// Code scaffolded by goctl. Safe to edit.
// goctl {{.version}}

package {{.PkgName}}

import (
	"net/http"

	{{.ImportPackages}}
	
	"github.com/5a6e/newzero/rest/httpx"
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			{{if .HasResp}}httpx.OkJsonCtx(r.Context(), w, &httpx.HttpResponse{
				Message: "ok",
				Data:    resp,
				Code:    0,
			}){{else}}httpx.Ok(w){{end}}
		}
	}
}
