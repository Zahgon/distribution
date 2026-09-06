package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/distribution/distribution/v3/registry/api/errcode"
	v2 "github.com/distribution/distribution/v3/registry/api/v2"
)

var spaceRegex = regexp.MustCompile(`\n\s*`)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("please specify a template to execute.")
	}

	path := os.Args[1]
	filename := filepath.Base(path)

	funcMap := template.FuncMap{
		"removenewlines": func(s string) string {
			return spaceRegex.ReplaceAllString(s, " ")
		},
		"statustext":    http.StatusText,
		"prettygorilla": prettyGorillaMuxPath,
	}

	tmpl := template.Must(template.New(filename).Funcs(funcMap).ParseFiles(path))

	data := struct {
		RouteDescriptors []v2.RouteDescriptor
		ErrorDescriptors []errcode.ErrorDescriptor
	}{
		RouteDescriptors: v2.APIDescriptor.RouteDescriptors,
		ErrorDescriptors: append(errcode.GetErrorCodeGroup("registry.api.v2"),

			errcode.ErrorCodeUnauthorized.Descriptor(),
			errcode.ErrorCodeDenied.Descriptor(),
			errcode.ErrorCodeUnsupported.Descriptor()),
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		log.Fatalln(err)
	}
}

func prettyGorillaMuxPath(s string) string { _ = "STUB: not implemented"; return "" }
