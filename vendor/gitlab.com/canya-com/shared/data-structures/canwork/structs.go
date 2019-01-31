package canwork

// UserDocument represents the firestote document object structure for canwork users
type UserDocument struct {
	Timezone    string   `json:"timezone"`
	State       string   `json:"state"`
	Email       string   `json:"email"`
	HourlyRate  string   `json:"hourlyRate"`
	Colors      []string `json:"colors"`
	Type        string   `json:"type"`
	WhiteListed string   `json:"whitelisted"`
	Avatar      struct {
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

// JobDocument represents the firestote document object structure for canwork jobs
type JobDocument struct {
	Information struct {
		WeeklyCommitment    int           `json:"weeklyCommitment"`
		Attachments         []interface{} `json:"attachments"`
		InitialStage        string        `json:"initialStage"`
		Title               string        `json:"title"`
		TimelineExpectation string        `json:"timelineExpectation"`
		WorkType            string        `json:"workType"`
		Description         string        `json:"description"`
		Skills              []string      `json:"skills"`
	} `json:"information"`
	BoostVisibility bool   `json:"boostVisibility"`
	ClientID        string `json:"clientId"`
	ID              string `json:"id"`
	ProviderID      string `json:"providerId"`
	BudgetCan       int    `json:"budgetCan"`
	OtherParty      struct {
		Avatar struct {
			URI string `json:"uri"`
		} `json:"avatar"`
		Name string `json:"name"`
	} `json:"otherParty"`
	Budget      int    `json:"budget"`
	PaymentType string `json:"paymentType"`
	PaymentLog  []struct {
		Timestamp string `json:"timestamp"`
		AmountCan int    `json:"amountCan"`
		TxID      string `json:"txId"`
	} `json:"paymentLog"`
	HexID     string `json:"hexId"`
	ActionLog []struct {
		Timestamp           string `json:"timestamp"`
		TimelineExpectation string `json:"timelineExpectation,omitempty"`
		Private             bool   `json:"private"`
		WorkType            string `json:"workType,omitempty"`
		ExecutedBy          string `json:"executedBy"`
		CAN                 int    `json:"CAN,omitempty"`
		WeeklyCommitment    int    `json:"weeklyCommitment,omitempty"`
		PaymentType         string `json:"paymentType,omitempty"`
		USD                 int    `json:"USD,omitempty"`
		Type                string `json:"type"`
		AmountCan           int    `json:"amountCan,omitempty"`
		TxID                string `json:"txId,omitempty"`
		Message             string `json:"message,omitempty"`
	} `json:"actionLog"`
	CanInEscrow int    `json:"canInEscrow"`
	State       string `json:"state"`
	Slug        string `json:"slug"`
}
