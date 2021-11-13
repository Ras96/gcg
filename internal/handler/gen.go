package handler

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/Ras96/gcg/internal/repository"
	"github.com/Ras96/gcg/internal/util/errors"
	"github.com/spf13/cobra"
)

type genHandler struct {
	repo *repository.Repositories
	opts *genOpts
}

type genOpts struct {
	output *string
}

func (h *genHandler) Run(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cobra.CheckErr(errors.New("Please provide an argument"))
	}

	file, err := h.repo.Analyzer.AnalyzeFile(args[0])
	errors.CheckErr(err, "Could not analyze file")

	res, err := h.repo.Generator.GenerateConstructors(file, *h.opts.output)
	errors.CheckErr(err, "Could not generate constructors")

	if len(*h.opts.output) == 0 {
		fmt.Fprintln(os.Stdout, string(res))
	} else {
		_ = ioutil.WriteFile(*h.opts.output, res, fs.ModePerm)
	}
}
