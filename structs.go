package main

// SummaryResponse in the data structure returned in JSON format from the summaryHandler (/summary)
type SummaryResponse struct {
	TotalUsers            int     `json:"totalUsers"`
	TotalJobsCompleted    int     `json:"totalJobsCompleted"`
	TotalCanTransacted    int     `json:"totalCanTransacted"`
	TotalCanInEscrow      float64 `json:"totalCanInEscrow"`
	TotalCanFees          float64 `json:"totalCanFees"`
	FeeCanRate            int     `json:"feeCanRate"`
	EscrowContractAddress string  `json:"escrowContractAddress"`
}
