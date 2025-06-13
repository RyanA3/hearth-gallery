package templates

import (
	"fmt"
	"html/template"
	"path/filepath"
)

var TemplateCache = make(map[string]*template.Template)

// InitTemplateCache initializes and caches all HTML templates.
func InitTemplateCache(dir string) error {
	// Get all page templates
	pages, err := filepath.Glob(filepath.Join(dir, "pages", "*.tmpl"))
	if err != nil {
		return err
	}

	fmt.Println("Loaded pages: ", pages);

	// Get all partial templates
	partials, err := filepath.Glob(filepath.Join(dir, "partials", "*.tmpl"))
	if err != nil {
		return err
	}

	fmt.Println("Loaded partials: ", partials);

	// Get the base template
	baseTemplatePath := filepath.Join(dir, "base.tmpl")

	for _, page := range pages {
		name := filepath.Base(page)

		// Create a new template set
		ts, err := template.ParseFiles(baseTemplatePath)
		if err != nil {
			return err
		}

		// Add partials to the template set
		ts, err = ts.ParseFiles(partials...)
		if err != nil {
			return err
		}

		// Add the specific page template to the set
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return err
		}

		TemplateCache[name] = ts
	}

	return nil
}
