package main

import (
	"html/template"
	"net/url"
	"path/filepath"
	"time"

	"github.com/AhmedARmohamed/snippetbox/pkg/models"
)

type templateData struct {
	CurrentYear int
	FormData url.Values
	FormErrors map[string]string
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// humanDate returns a nicely formatted string representation of time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.gohtml"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.gohtml"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.gohtml"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}

	return cache, nil
}
