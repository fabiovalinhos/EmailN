package contract

// é o payload para request
type NewCampaign struct {
	Name      string
	Content   string
	Emails    []string
	CreatedBy string
}
