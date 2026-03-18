package upgrade

import "github.com/5a6e/newzero/tools/goctl/internal/cobrax"

// Cmd describes an upgrade command.
var Cmd = cobrax.NewCommand("upgrade", cobrax.WithRunE(upgrade))
