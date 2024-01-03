package builder

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/forbole/juno/v4/node"
	nodeconfig "github.com/forbole/juno/v4/node/config"
	"github.com/forbole/juno/v4/node/remote"
)

func BuildNode(cfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (node.Node, error) {
	switch cfg.Type {
	case nodeconfig.TypeRemote:
		return remote.NewNode(cfg.Details.(*remote.Details), encodingConfig.Codec)
	case nodeconfig.TypeLocal:
		return nil, fmt.Errorf("local mode is currently not supported")
	case nodeconfig.TypeNone:
		return nil, nil

	default:
		return nil, fmt.Errorf("invalid node type: %s", cfg.Type)
	}
}
