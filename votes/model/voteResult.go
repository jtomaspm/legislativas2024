package model

type VoteResult struct {
	LocationId                       string
	LocationName                     string
	BlankVotes                       int
	BlankVotesPercentage             float64
	NullVotes                        int
	NullVotesPercentage              float64
	NumberParishes                   int
	NumberVoters                     int
	PercentageVoters                 float64
	PS_Mandates                      int
	PS_Percentage                    float64
	PS_ValidVotesPercentage          float64
	PS_Votes                         int
	PSD_CDS_PPM_Mandates             int
	PSD_CDS_PPM_Percentage           float64
	PSD_CDS_PPM_ValidVotesPercentage float64
	PSD_CDS_PPM_Votes                int
	CH_Mandates                      int
	CH_Percentage                    float64
	CH_ValidVotesPercentage          float64
	CH_Votes                         int
	IL_Mandates                      int
	IL_Percentage                    float64
	IL_ValidVotesPercentage          float64
	IL_Votes                         int
	BE_Mandates                      int
	BE_Percentage                    float64
	BE_ValidVotesPercentage          float64
	BE_Votes                         int
	PCP_Mandates                     int
	PCP_Percentage                   float64
	PCP_ValidVotesPercentage         float64
	PCP_Votes                        int
	L_Mandates                       int
	L_Percentage                     float64
	L_ValidVotesPercentage           float64
	L_Votes                          int
	PAN_Mandates                     int
	PAN_Percentage                   float64
	PAN_ValidVotesPercentage         float64
	PAN_Votes                        int
	ADN_Mandates                     int
	ADN_Percentage                   float64
	ADN_ValidVotesPercentage         float64
	ADN_Votes                        int
	PSD_CDS_Mandates                 int
	PSD_CDS_Percentage               float64
	PSD_CDS_ValidVotesPercentage     float64
	PSD_CDS_Votes                    int
	RIR_Mandates                     int
	RIR_Percentage                   float64
	RIR_ValidVotesPercentage         float64
	RIR_Votes                        int
	JPP_Mandates                     int
	JPP_Percentage                   float64
	JPP_ValidVotesPercentage         float64
	JPP_Votes                        int
	ND_Mandates                      int
	ND_Percentage                    float64
	ND_ValidVotesPercentage          float64
	ND_Votes                         int
	PCTP_MRPP_Mandates               int
	PCTP_MRPP_Percentage             float64
	PCTP_MRPP_ValidVotesPercentage   float64
	PCTP_MRPP_Votes                  int
	VP_Mandates                      int
	VP_Percentage                    float64
	VP_ValidVotesPercentage          float64
	VP_Votes                         int
	E_Mandates                       int
	E_Percentage                     float64
	E_ValidVotesPercentage           float64
	E_Votes                          int
	MPT_A_Mandates                   int
	MPT_A_Percentage                 float64
	MPT_A_ValidVotesPercentage       float64
	MPT_A_Votes                      int
	PTP_Mandates                     int
	PTP_Percentage                   float64
	PTP_ValidVotesPercentage         float64
	PTP_Votes                        int
	NC_Mandates                      int
	NC_Percentage                    float64
	NC_ValidVotesPercentage          float64
	NC_Votes                         int
	PPM_Mandates                     int
	PPM_Percentage                   float64
	PPM_ValidVotesPercentage         float64
	PPM_Votes                        int
	TotalConsulates                  int
	TotalConsulatesApproved          int
	TotalParishes                    int
	TotalParishesApproved            int
}

func findMandates(name string, tr *TerritoryResults) int {
	for _, n := range tr.CurrentResults.ResultsParty {
		if n.Acronym == name {
			return n.Mandates
		}
	}
	return 0
}
func findPercentage(name string, tr *TerritoryResults) float64 {
	for _, n := range tr.CurrentResults.ResultsParty {
		if n.Acronym == name {
			return n.Percentage
		}
	}
	return 0
}
func findValidVotesPercentage(name string, tr *TerritoryResults) float64 {
	for _, n := range tr.CurrentResults.ResultsParty {
		if n.Acronym == name {
			return n.ValidVotesPercentage
		}
	}
	return 0
}
func findVotes(name string, tr *TerritoryResults) int {
	for _, n := range tr.CurrentResults.ResultsParty {
		if n.Acronym == name {
			return n.Votes
		}
	}
	return 0
}

func VoteResultFromTR(tr *TerritoryResults) *VoteResult {
	res := VoteResult{
		LocationId:                       tr.TerritoryKey,
		LocationName:                     tr.TerritoryFullName,
		BlankVotes:                       tr.CurrentResults.BlankVotes,
		BlankVotesPercentage:             tr.CurrentResults.BlankVotesPercentage,
		NullVotes:                        tr.CurrentResults.NullVotes,
		NullVotesPercentage:              tr.CurrentResults.NullVotesPercentage,
		NumberParishes:                   tr.CurrentResults.NumberParishes,
		NumberVoters:                     tr.CurrentResults.NumberVoters,
		PercentageVoters:                 tr.CurrentResults.PercentageVoters,
		PS_Mandates:                      findMandates("PS", tr),
		PS_Percentage:                    findPercentage("PS", tr),
		PS_ValidVotesPercentage:          findValidVotesPercentage("PS", tr),
		PS_Votes:                         findVotes("PS", tr),
		PSD_CDS_PPM_Mandates:             findMandates("PPD/PSD.CDS-PP.PPM", tr),
		PSD_CDS_PPM_Percentage:           findPercentage("PPD/PSD.CDS-PP.PPM", tr),
		PSD_CDS_PPM_ValidVotesPercentage: findValidVotesPercentage("PPD/PSD.CDS-PP.PPM", tr),
		PSD_CDS_PPM_Votes:                findVotes("PPD/PSD.CDS-PP.PPM", tr),
		CH_Mandates:                      findMandates("CH", tr),
		CH_Percentage:                    findPercentage("CH", tr),
		CH_ValidVotesPercentage:          findValidVotesPercentage("CH", tr),
		CH_Votes:                         findVotes("CH", tr),
		IL_Mandates:                      findMandates("IL", tr),
		IL_Percentage:                    findPercentage("IL", tr),
		IL_ValidVotesPercentage:          findValidVotesPercentage("IL", tr),
		IL_Votes:                         findVotes("IL", tr),
		BE_Mandates:                      findMandates("B.E.", tr),
		BE_Percentage:                    findPercentage("B.E.", tr),
		BE_ValidVotesPercentage:          findValidVotesPercentage("B.E.", tr),
		BE_Votes:                         findVotes("B.E.", tr),
		PCP_Mandates:                     findMandates("PCP-PEV", tr),
		PCP_Percentage:                   findPercentage("PCP-PEV", tr),
		PCP_ValidVotesPercentage:         findValidVotesPercentage("PCP-PEV", tr),
		PCP_Votes:                        findVotes("PCP-PEV", tr),
		L_Mandates:                       findMandates("L", tr),
		L_Percentage:                     findPercentage("L", tr),
		L_ValidVotesPercentage:           findValidVotesPercentage("L", tr),
		L_Votes:                          findVotes("L", tr),
		PAN_Mandates:                     findMandates("PAN", tr),
		PAN_Percentage:                   findPercentage("PAN", tr),
		PAN_ValidVotesPercentage:         findValidVotesPercentage("PAN", tr),
		PAN_Votes:                        findVotes("PAN", tr),
		ADN_Mandates:                     findMandates("ADN", tr),
		ADN_Percentage:                   findPercentage("ADN", tr),
		ADN_ValidVotesPercentage:         findValidVotesPercentage("ADN", tr),
		ADN_Votes:                        findVotes("ADN", tr),
		PSD_CDS_Mandates:                 findMandates("PPD/PSD.CDS-PP", tr),
		PSD_CDS_Percentage:               findPercentage("PPD/PSD.CDS-PP", tr),
		PSD_CDS_ValidVotesPercentage:     findValidVotesPercentage("PPD/PSD.CDS-PP", tr),
		PSD_CDS_Votes:                    findVotes("PPD/PSD.CDS-PP", tr),
		RIR_Mandates:                     findMandates("R.I.R.", tr),
		RIR_Percentage:                   findPercentage("R.I.R.", tr),
		RIR_ValidVotesPercentage:         findValidVotesPercentage("R.I.R.", tr),
		RIR_Votes:                        findVotes("R.I.R.", tr),
		JPP_Mandates:                     findMandates("JPP", tr),
		JPP_Percentage:                   findPercentage("JPP", tr),
		JPP_ValidVotesPercentage:         findValidVotesPercentage("JPP", tr),
		JPP_Votes:                        findVotes("JPP", tr),
		ND_Mandates:                      findMandates("ND", tr),
		ND_Percentage:                    findPercentage("ND", tr),
		ND_ValidVotesPercentage:          findValidVotesPercentage("ND", tr),
		ND_Votes:                         findVotes("ND", tr),
		PCTP_MRPP_Mandates:               findMandates("PCTP/MRPP", tr),
		PCTP_MRPP_Percentage:             findPercentage("PCTP/MRPP", tr),
		PCTP_MRPP_ValidVotesPercentage:   findValidVotesPercentage("PCTP/MRPP", tr),
		PCTP_MRPP_Votes:                  findVotes("PCTP/MRPP", tr),
		VP_Mandates:                      findMandates("VP", tr),
		VP_Percentage:                    findPercentage("VP", tr),
		VP_ValidVotesPercentage:          findValidVotesPercentage("VP", tr),
		VP_Votes:                         findVotes("VP", tr),
		E_Mandates:                       findMandates("E", tr),
		E_Percentage:                     findPercentage("E", tr),
		E_ValidVotesPercentage:           findValidVotesPercentage("E", tr),
		E_Votes:                          findVotes("E", tr),
		MPT_A_Mandates:                   findMandates("MPT.A", tr),
		MPT_A_Percentage:                 findPercentage("MPT.A", tr),
		MPT_A_ValidVotesPercentage:       findValidVotesPercentage("MPT.A", tr),
		MPT_A_Votes:                      findVotes("MPT.A", tr),
		PTP_Mandates:                     findMandates("PTP", tr),
		PTP_Percentage:                   findPercentage("PTP", tr),
		PTP_ValidVotesPercentage:         findValidVotesPercentage("PTP", tr),
		PTP_Votes:                        findVotes("PTP", tr),
		NC_Mandates:                      findMandates("NC", tr),
		NC_Percentage:                    findPercentage("NC", tr),
		NC_ValidVotesPercentage:          findValidVotesPercentage("NC", tr),
		NC_Votes:                         findVotes("NC", tr),
		PPM_Mandates:                     findMandates("PPM", tr),
		PPM_Percentage:                   findPercentage("PPM", tr),
		PPM_ValidVotesPercentage:         findValidVotesPercentage("PPM", tr),
		PPM_Votes:                        findVotes("PPM", tr),
		TotalConsulates:                  tr.TotalConsulates,
		TotalConsulatesApproved:          tr.TotalConsulatesApproved,
		TotalParishes:                    tr.TotalParishes,
		TotalParishesApproved:            tr.TotalParishesApproved,
	}
	return &res
}
