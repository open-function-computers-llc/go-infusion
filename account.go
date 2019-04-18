package infusion

import (
	"encoding/json"
	"io/ioutil"
)

// AccountProfile the details about the account for the authorized account
type AccountProfile struct {
	Address struct {
		CountryCode string `json:"country_code"`
		Field       string `json:"field"`
		Line1       string `json:"line1"`
		Line2       string `json:"line2"`
		Locality    string `json:"locality"`
		PostalCode  string `json:"postal_code"`
		Region      string `json:"region"`
		ZipCode     string `json:"zip_code"`
		ZipFour     string `json:"zip_four"`
	} `json:"address"`
	BusinessGoals          []string `json:"business_goals"`
	BusinessPrimaryColor   string   `json:"business_primary_color"`
	BusinessSecondaryColor string   `json:"business_secondary_color"`
	BusinessType           string   `json:"business_type"`
	CurrencyCode           string   `json:"currency_code"`
	Email                  string   `json:"email"`
	LanguageTag            string   `json:"language_tag"`
	LogoURL                string   `json:"logo_url"`
	Name                   string   `json:"name"`
	Phone                  string   `json:"phone"`
	PhoneExt               string   `json:"phone_ext"`
	TimeZone               string   `json:"time_zone"`
	Website                string   `json:"website"`
}

// GetAccountProfile Request the Account Details from Infusion
func GetAccountProfile() (AccountProfile, error) {
	ap := AccountProfile{}
	r, err := getRequest("/account/profile")
	if err != nil {
		return ap, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return ap, err
	}

	err = json.Unmarshal(body, &ap)
	if err != nil {
		return ap, err
	}

	log.Info("Account recieved for " + ap.Name)
	log.Info(string(body))
	return ap, err
}
