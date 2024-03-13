package model

type VoteResult struct {
	LocationId   string
	LocationName string
	NrVoters     int
}

func VoteResultFromTR(tr *TerritoryResults) *VoteResult {
	res := VoteResult{
		LocationId:   tr.TerritoryKey,
		LocationName: tr.TerritoryFullName,
		NrVoters:     tr.CurrentResults.NumberVoters,
	}
	return &res
}
