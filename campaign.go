package infusion

import (
	"strconv"
)

//Campaign - Infusionsoft Campaign
type Campaign struct {
	ActiveContactCount    int    `json:"active_contact_count"`
	CompletedContactCount int    `json:"completed_contact_count"`
	CreatedBy             int    `json:"created_by_globaal_id"`
	DateCreated           string `json:"date_created"`
	ErrorMessage          string `json:""`
	Goals                 []struct {
		HistoricalContactCounts struct {
			TwentyFourHours int `json:"24_hours"`
			ThirtyDays      int `json:"30_days"`
		} `json:"historical_contact_counts"`
		ID                  int    `json:"id"`
		Name                string `json:"name"`
		NextSquenceIds      []int  `json:"next_sequence_ids"`
		PreviousSequenceIds []int  `json:"previous_sequence_ids"`
		Type                string `json:"type"`
	} `json:"goals"`
	ID                int    `json:"id"`
	Locked            bool   `json:"locked"`
	PublishedDate     string `json:"published_date"`
	PublishedStatus   bool   `json:"published_status"`
	PublishedTimezone string `json:"published_time_zone"`
	Sequences         []Sequence
	Timezone          string `json:"time_zone"`
}

//Sequence - This is sequence in a campaign
type Sequence struct {
	ActiveContactCount      int `json:"active_contact_count"`
	CompletedContactCount   int `json:"completed_contact_count"`
	HistoricalContactCounts struct {
		TwentyFourHours int `json:"24_hours"`
		ThirtyDays      int `json:"30_days"`
	} `json:"historical_contact_counts"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Paths []SequencePath
}

//SequencePath - Sequence Path
type SequencePath struct {
	ID    int `json:"id"`
	Items []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		NextItemID     int    `json:"next_item_id"`
		PreviousItemID int    `json:"previous_item_id"`
		Type           string `json:"type"`
	}
}

// AddEmailToCampaign - Adds a new email to a campaign. This Infusion method returns 204(no content) when successful, and we emulated that response.
func AddEmailToCampaign(campaignID, sequenceID, contactID int) error {

	cID := strconv.Itoa(campaignID)
	sID := strconv.Itoa(sequenceID)
	ccID := strconv.Itoa(contactID)

	emptyBody := []byte{}
	_, err := postRequest("/campaigns/"+cID+"/sequences/"+sID+"/contacts/"+ccID, emptyBody)
	if err != nil {
		return err
	}

	return nil
}
