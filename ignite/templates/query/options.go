package query

import (
	"github.com/Source-Protocol-Cosmos/cli/ignite/pkg/multiformatname"
	"github.com/Source-Protocol-Cosmos/cli/ignite/templates/field"
)

// Options ...
type Options struct {
	AppName     string
	AppPath     string
	ModuleName  string
	ModulePath  string
	QueryName   multiformatname.Name
	Description string
	ResFields   field.Fields
	ReqFields   field.Fields
	Paginated   bool
}
