package script

import (
	"embed"
	"path/filepath"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/plush/v4"
	"github.com/ignite/cli/v28/ignite/pkg/placeholder"
	"github.com/ignite/cli/v28/ignite/pkg/xgenny"
	"github.com/ignite/cli/v28/ignite/templates/field/plushhelpers"
)

//go:embed files/* files/**/*
var fsRunScript embed.FS

// NewRunScript returns the generator to scaffold a chain run script.
func NewRunScript(path, log, home, binaryPath, account, denoms, output string) (string, error) {
	var (
		g         = genny.New()
		runScript = xgenny.NewEmbedWalker(
			fsRunScript,
			"files/",
			output,
		)
	)
	if err := g.Box(runScript); err != nil {
		return "", err
	}

	ctx := plush.NewContext()
	ctx.Set("path", path)
	ctx.Set("log", log)
	ctx.Set("home", home)
	ctx.Set("binaryPath", binaryPath)
	ctx.Set("binary", filepath.Base(binaryPath))
	ctx.Set("account-name", account)
	ctx.Set("denoms", denoms)

	plushhelpers.ExtendPlushContext(ctx)
	g.Transformer(xgenny.Transformer(ctx))

	_, err := xgenny.RunWithValidation(placeholder.New(), g)
	return filepath.Join(output, "run.sh"), err
}
