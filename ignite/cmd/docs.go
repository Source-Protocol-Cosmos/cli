package ignitecmd

import (
	"github.com/spf13/cobra"

	"github.com/Source-Protocol-Cosmos/cli/docs"
	"github.com/Source-Protocol-Cosmos/cli/ignite/pkg/localfs"
	"github.com/Source-Protocol-Cosmos/cli/ignite/pkg/markdownviewer"
)

func NewDocs() *cobra.Command {
	c := &cobra.Command{
		Use:   "docs",
		Short: "Show Ignite CLI docs",
		Args:  cobra.NoArgs,
		RunE:  docsHandler,
	}
	return c
}

func docsHandler(cmd *cobra.Command, args []string) error {
	path, cleanup, err := localfs.SaveTemp(docs.Docs)
	if err != nil {
		return err
	}
	defer cleanup()

	return markdownviewer.View(path)
}
