package campaign

import (
	"emailn/internal/internalerrors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending  string = "Pending"
	Canceled        = "Canceled"
	Deleted         = "Deleted"
	Started         = "Started"
	Done            = "Done"
)

type Contacts struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string     `validate:"required" gorm:"size:50"`
	Name      string     `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time  `validate:"required"`
	Content   string     `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contacts `validate:"min=1,dive"`
	Status    string     `gorm:"size:20"`
	CreatedBy string     `validate:"email" gorm:"size:50"`
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Delete() {
	c.Status = Deleted
}

func NewCampaign(name string, content string, emails []string, createdBy string) (*Campaign, error) {

	// Na unha
	contacts := make([]Contacts, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].ID = xid.New().String()

	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
		CreatedBy: createdBy,
	}

	err := internalerrors.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}
