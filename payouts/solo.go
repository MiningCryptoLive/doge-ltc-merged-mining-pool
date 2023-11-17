package payouts

import (
	"fmt"
	"log"
	"time"

	"designs.capital/dogepool/persistence"
)

type SOLO struct{}

func (SOLO) UpdateMinerBalances(poolID string, remainingReward float32, confirmed persistence.Found) (time.Time, error) {
	log.Printf("Awarding %v %v SOLO reward to miner %v\n", remainingReward, confirmed.Chain, confirmed.Miner)

	usage := "SOLO REWARD FOR BLOCK %v"
	usage = fmt.Sprintf(usage, confirmed.BlockHeight)
	return confirmed.Created, persistence.Balances.AddAmount(poolID, confirmed.Chain, confirmed.Miner, usage, remainingReward)
}
