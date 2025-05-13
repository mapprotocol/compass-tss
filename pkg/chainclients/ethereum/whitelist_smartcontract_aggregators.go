package ethereum

import (
	"github.com/mapprotocol/compass-tss/common"
	"gitlab.com/thorchain/thornode/v3/x/thorchain/aggregators"
)

func LatestAggregatorContracts() []common.Address {
	addrs := []common.Address{}
	for _, agg := range aggregators.DexAggregators() {
		if agg.Chain.Equals(common.ETHChain) {
			addrs = append(addrs, common.Address(agg.Address))
		}
	}
	return addrs
}
