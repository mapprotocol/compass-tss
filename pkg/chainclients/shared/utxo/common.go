package utxo

import (
	"fmt"

	"github.com/mapprotocol/compass-tss/common"
	"github.com/mapprotocol/compass-tss/mapclient"
)

func GetAsgardAddress(chain common.Chain, bridge mapclient.ThorchainBridge) ([]common.Address, error) {
	vaults, err := bridge.GetAsgardPubKeys()
	if err != nil {
		return nil, fmt.Errorf("fail to get asgards : %w", err)
	}

	newAddresses := make([]common.Address, 0)
	for _, v := range vaults {
		var addr common.Address
		addr, err = v.PubKey.GetAddress(chain)
		if err != nil {
			continue
		}
		newAddresses = append(newAddresses, addr)
	}
	return newAddresses, nil
}
