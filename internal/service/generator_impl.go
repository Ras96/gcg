package service

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/Ras96/gcg/internal/model"
	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
)

type generatorService struct {
	Tmpl *template.Template
	Opts *imports.Options
}

func NewGeneratorService() GeneratorService {
	return &generatorService{
		Tmpl: nil,
		Opts: &imports.Options{
			AllErrors: true,
			Comments:  true,
		},
	}
}

var fmap = template.FuncMap{
	"title": strings.Title,
}

func (r *generatorService) GenerateConstructors(file *model.File, output string, isPrivate bool) ([]byte, error) {
	r.Tmpl = template.New("constructor").Funcs(
		template.FuncMap{
			"title": strings.Title,
			"funcName": func(funcName string) string {
				if !isPrivate {
					return strings.Title(funcName)
				}

				return funcName
			},
		},
	)
	if _, err := r.Tmpl.Parse(model.GenTmpl); err != nil {
		return nil, errors.Wrap(err, "Could not parse templates")
	}

	w := &bytes.Buffer{}
	if err := r.writeConstructors(w, file); err != nil {
		return nil, errors.Wrap(err, "Could not write constructors")
	}

	out, err := r.format(w, output)
	if err != nil {
		return nil, errors.Wrap(err, "Could not format output")
	}

	return out, nil
}

func (r *generatorService) writeConstructors(w *bytes.Buffer, file *model.File) error {
	b := bufio.NewWriter(w)
	if err := r.Tmpl.Execute(b, file); err != nil {
		return errors.Wrap(err, "Could not execute template")
	}

	if err := b.Flush(); err != nil {
		return errors.Wrap(err, "Could not flush buffer")
	}

	return nil
}

func (r *generatorService) format(w *bytes.Buffer, filename string) ([]byte, error) {
	formatted, err := imports.Process(filename, w.Bytes(), r.Opts)
	if err != nil {
		if len(filename) == 0 {
			fmt.Fprintln(os.Stdout, w.String())
		} else {
			if err := ioutil.WriteFile(filename, w.Bytes(), fs.ModePerm); err != nil {
				return nil, errors.Wrap(err, "Could not write to file")
			}
		}

		fmt.Fprintln(os.Stderr, "Error occurred. Instead, gcg output the unformatted file")
		fmt.Fprintln(os.Stderr, "")

		return nil, errors.Wrap(err, "Could not format file")
	}

	return formatted, nil
}
