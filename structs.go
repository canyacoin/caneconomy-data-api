package main

// SummaryResponse in the data structure returned in JSON format from the summaryHandler (/summary)
type SummaryResponse struct {
	TotalUsers            int64   `json:"totalUsers"`
	TotalJobsCompleted    int64   `json:"totalJobsCompleted"`
	TotalCanTransacted    int64   `json:"totalCanTransacted"`
	TotalCanInEscrow      float64 `json:"totalCanInEscrow"`
	TotalCanFees          float64 `json:"totalCanFees"`
	FeeCanRate            int64   `json:"feeCanRate"`
	EscrowContractAddress string  `json:"escrowContractAddress"`
}

// User structure from firestore
type User struct {
	Timezone   string   `json:"timezone"`
	State      string   `json:"state"`
	Email      string   `json:"email"`
	HourlyRate string   `json:"hourlyRate"`
	Colors     []string `json:"colors"`
	Type       string   `json:"type"`
	Avatar     struct {
		URI string `json:"uri"`
	} `json:"avatar"`
	EthereumLogin struct {
		Expiry int `json:"expiry"`
		Pin    int `json:"pin"`
	} `json:"ethereumLogin"`
	Phone            string        `json:"phone"`
	Work             string        `json:"work"`
	Address          string        `json:"address"`
	Context          string        `json:"@context"`
	EthAddress       string        `json:"ethAddress"`
	SkillTags        []string      `json:"skillTags"`
	Name             string        `json:"name"`
	Category         string        `json:"category"`
	Title            string        `json:"title"`
	EthAddressLookup string        `json:"ethAddressLookup"`
	UserType         string        `json:"@type"`
	Bio              string        `json:"bio"`
	Timestamp        string        `json:"timestamp"`
	Slug             string        `json:"slug"`
	Description      string        `json:"description"`
	WorkSkillTags    []interface{} `json:"workSkillTags"`
	URI              string        `json:"uri"`
}

type Job struct {
	ActionLog []struct {
		AmountCan           int    `json:"amountCan"`
		AmountUsd           int    `json:"amountUsd"`
		ExecutedBy          string `json:"executedBy"`
		Message             string `json:"message"`
		PaymentType         string `json:"paymentType"`
		Private             bool   `json:"private"`
		TimelineExpectation string `json:"timelineExpectation"`
		Timestamp           string `json:"timestamp"`
		Type                string `json:"type"`
		WeeklyCommitment    string `json:"weeklyCommitment"`
		WorkType            string `json:"workType"`
	} `json:"actionLog"`
	BoostVisibility bool   `json:"boostVisibility"`
	Budget          int    `json:"budget"`
	ClientID        string `json:"clientId"`
	Deadline        string `json:"deadline"`
	Draft           bool   `json:"draft"`
	Slug            string `json:"slug"`
	HexID           string `json:"hexId"`
	ID              string `json:"id"`
	Information     struct {
		Attachments         []interface{} `json:"attachments"`
		Description         string        `json:"description"`
		InitialStage        string        `json:"initialStage"`
		ProviderType        string        `json:"providerType"`
		Skills              []string      `json:"skills"`
		TimelineExpectation string        `json:"timelineExpectation"`
		Title               string        `json:"title"`
		WeeklyCommitment    int           `json:"weeklyCommitment"`
		WorkType            string        `json:"workType"`
	} `json:"information"`
	PaymentLog  []interface{} `json:"paymentLog"`
	PaymentType string        `json:"paymentType"`
	State       string        `json:"state"`
}
