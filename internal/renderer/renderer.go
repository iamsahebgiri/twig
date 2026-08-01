package renderer

import (
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/iamsahebgiri/twig/internal/assets"
)

type Renderer struct {
	templates *template.Template
}

func New() *Renderer {
	t := template.Must(template.New("").
		Option("missingkey=error").
		Funcs(template.FuncMap{
			"upper": funcMap()["upper"],
		}).
		ParseFS(assets.TemplatesFS, "templates/*.tmpl", "templates/**/*.tmpl"),
	)
	return &Renderer{
		templates: t,
	}
}

func (r *Renderer) Render(w io.Writer, name string, data any) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

// r.RenderFile("dist/index.html", "home", homeData)
func (r *Renderer) RenderFile(path, template string, data any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return r.Render(f, template, data)
}
