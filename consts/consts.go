package consts

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/consts"
)

const (
	// Human-readable part for addresses
	HRP = "myhyperchain"

	// Name and symbol of the chain and token
	Name   = "MyCustomChain"
	Symbol = "MCT"

	// Additional token-specific constants
	MaxAssetAmount = 1000000
	TokenDecimals  = 18
)

var ID ids.ID

func init() {
	// Generate the VM ID based on the chosen chain name
	b := make([]byte, consts.IDLen)
	copy(b, []byte(Name))
	vmID, err := ids.ToID(b)
	if err != nil {
		panic(err)
	}
	ID = vmID
}

// Instantiate registry here so it can be imported by any package.
var (
	ActionRegistry *codec.TypeParser[chain.Action, *warp.Message, bool]
	AuthRegistry   *codec.TypeParser[chain.Auth, *warp.Message, bool]
)
