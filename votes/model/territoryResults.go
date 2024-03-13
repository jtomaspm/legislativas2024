package model

type TerritoryResults struct {
	CurrentResults struct {
		AvailableMandates    int         `json:"availableMandates"`
		BlankVotes           int         `json:"blankVotes"`
		BlankVotesPercentage float64     `json:"blankVotesPercentage"`
		Compensation         bool        `json:"compensation"`
		DisplayMessage       interface{} `json:"displayMessage"`
		HasNoVoting          bool        `json:"hasNoVoting"`
		NullVotes            int         `json:"nullVotes"`
		NullVotesPercentage  float64     `json:"nullVotesPercentage"`
		NumberParishes       int         `json:"numberParishes"`
		NumberVoters         int         `json:"numberVoters"`
		PercentageVoters     float64     `json:"percentageVoters"`
		ResultsParty         []struct {
			AbsoluteMajority     interface{} `json:"absoluteMajority"`
			Acronym              string      `json:"acronym"`
			ConstituenctyCounter int         `json:"constituenctyCounter"`
			ImageKey             string      `json:"imageKey"`
			Mandates             int         `json:"mandates"`
			Percentage           float64     `json:"percentage"`
			Presidents           interface{} `json:"presidents"`
			ValidVotesPercentage float64     `json:"validVotesPercentage"`
			Votes                int         `json:"votes"`
		} `json:"resultsParty"`
		SubscribedVoters      int         `json:"subscribedVoters"`
		Tie                   bool        `json:"tie"`
		TieMessage            interface{} `json:"tieMessage"`
		TotalBoycotts         int         `json:"totalBoycotts"`
		TotalForeignBoycotts  int         `json:"totalForeignBoycotts"`
		TotalLocalBoycotts    int         `json:"totalLocalBoycotts"`
		TotalMandates         int         `json:"totalMandates"`
		TotalParishesApproved int         `json:"totalParishesApproved"`
		TotalVoters           int         `json:"totalVoters"`
	} `json:"currentResults"`
	DisplayMessage         interface{} `json:"displayMessage"`
	ForeignCounterMessage  interface{} `json:"foreignCounterMessage"`
	NationalCounterMessage interface{} `json:"nationalCounterMessage"`
	PreviousResults        struct {
		BlankVotes           int         `json:"blankVotes"`
		BlankVotesPercentage float64     `json:"blankVotesPercentage"`
		Compensation         bool        `json:"compensation"`
		DisplayMessage       interface{} `json:"displayMessage"`
		NullVotes            int         `json:"nullVotes"`
		NullVotesPercentage  float64     `json:"nullVotesPercentage"`
		PercentageVoters     float64     `json:"percentageVoters"`
		ResultsParty         []struct {
			AbsoluteMajority     interface{} `json:"absoluteMajority"`
			Acronym              string      `json:"acronym"`
			ConstituenctyCounter interface{} `json:"constituenctyCounter"`
			ImageKey             string      `json:"imageKey"`
			Mandates             int         `json:"mandates"`
			Percentage           float64     `json:"percentage"`
			Presidents           interface{} `json:"presidents"`
			ValidVotesPercentage float64     `json:"validVotesPercentage"`
			Votes                int         `json:"votes"`
		} `json:"resultsParty"`
		SubscribedVoters int `json:"subscribedVoters"`
		TotalMandates    int `json:"totalMandates"`
		TotalVoters      int `json:"totalVoters"`
	} `json:"previousResults"`
	RefreshTimeout          int         `json:"refreshTimeout"`
	TerritoryFullName       string      `json:"territoryFullName"`
	TerritoryKey            string      `json:"territoryKey"`
	TerritoryMessage        interface{} `json:"territoryMessage"`
	TerritoryName           string      `json:"territoryName"`
	Time                    string      `json:"time"`
	TimeOnServer            int64       `json:"timeOnServer"`
	TotalConsulates         int         `json:"totalConsulates"`
	TotalConsulatesApproved int         `json:"totalConsulatesApproved"`
	TotalParishes           int         `json:"totalParishes"`
	TotalParishesApproved   int         `json:"totalParishesApproved"`
	WarningMessage          interface{} `json:"warningMessage"`
}
