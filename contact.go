package infusion

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

// ContactsQueryResponse this is the data structre that infusion returns when querying for contacts
type ContactsQueryResponse struct {
	Contacts []Contact `json:"contacts"`
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
}

// Contact this is the data structure that infusion returns when showing the information for a single contact
type Contact struct {
	Addresses []struct {
		CountryCode string `json:"country_code"`
		Field       string `json:"field"`
		Line1       string `json:"line1"`
		Line2       string `json:"line2"`
		Locality    string `json:"locality"`
		PostalCode  string `json:"postal_code"`
		Region      string `json:"region"`
		ZipCode     string `json:"zip_code"`
		ZipFour     string `json:"zip_four"`
	} `json:"addresses"`
	Anniversary string `json:"anniversary"`
	Birthday    string `json:"birthday"`
	Company     struct {
		CompanyName string `json:"company_name"`
		ID          int    `json:"id"`
	} `json:"company"`
	ContactType  string `json:"contact_type"`
	CustomFields []struct {
		Content struct {
		} `json:"content"`
		ID int `json:"id"`
	} `json:"custom_fields"`
	DateCreated    string `json:"date_created"`
	EmailAddresses []struct {
		Email string `json:"email"`
		Field string `json:"field"`
	} `json:"email_addresses"`
	EmailOptedIn bool   `json:"email_opted_in"`
	EmailStatus  string `json:"email_status"`
	FamilyName   string `json:"family_name"`
	FaxNumbers   []struct {
		Field  string `json:"field"`
		Number string `json:"number"`
		Type   string `json:"type"`
	} `json:"fax_numbers"`
	GivenName    string `json:"given_name"`
	ID           int    `json:"id"`
	JobTitle     string `json:"job_title"`
	LastUpdated  string `json:"last_updated"`
	LeadSourceID int    `json:"lead_source_id"`
	MiddleName   string `json:"middle_name"`
	OwnerID      int    `json:"owner_id"`
	PhoneNumbers []struct {
		Extension string `json:"extension"`
		Field     string `json:"field"`
		Number    string `json:"number"`
		Type      string `json:"type"`
	} `json:"phone_numbers"`
	PreferredLocale string `json:"preferred_locale"`
	PreferredName   string `json:"preferred_name"`
	Prefix          string `json:"prefix"`
	SocialAccounts  []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"social_accounts"`
	SourceType string `json:"source_type"`
	SpouseName string `json:"spouse_name"`
	Suffix     string `json:"suffix"`
	TimeZone   string `json:"time_zone"`
	Website    string `json:"website"`
}

// LookupContact reach out to Infusion and query contacts based on name and email
func LookupContact(firstName, lastName, email string) (ContactsQueryResponse, error) {
	response := ContactsQueryResponse{}

	r, err := getRequest("/contacts?email=" + url.QueryEscape(email) + "&family_name=" + url.QueryEscape(lastName) + "&given_name=" + url.QueryEscape(firstName))
	if err != nil {
		return response, err
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(string(body))
		return response, err
	}

	log.Info("Client Lookup complete. Result: " + string(body))
	return response, nil
}

// LookupContact reach out to Infusion and query contacts based on email address only
func LookupContactByEmail(email string) (ContactsQueryResponse, error) {
	response := ContactsQueryResponse{}

	r, err := getRequest("/contacts?email=" + url.QueryEscape(email))
	if err != nil {
		return response, err
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(string(body))
		return response, err
	}

	log.Info("Client Lookup complete. Result: " + string(body))
	return response, nil
}

// CreateContact Send a POST request to infusion to create a contact
func CreateContact(firstName, lastName, email string) (Contact, error) {
	response := Contact{}

	type createContactBodyEmail struct {
		Email string `json:"email"`
		Field string `json:"field"`
	}
	type createContactBody struct {
		EmailAddresses []createContactBodyEmail `json:"email_addresses"`
		FamilyName     string                   `json:"family_name"`
		GivenName      string                   `json:"given_name"`
	}
	emailObject := createContactBodyEmail{
		Email: email,
		Field: "EMAIL1",
	}
	emails := make([]createContactBodyEmail, 1)
	emails[0] = emailObject
	body := createContactBody{
		FamilyName:     lastName,
		GivenName:      firstName,
		EmailAddresses: emails,
	}
	data, _ := json.Marshal(body)

	r, err := postRequest("/contacts", data)
	if err != nil {
		return response, err
	}

	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(rBody, &response)
	if err != nil {
		return response, err
	}
	if r.StatusCode != 201 {
		return response, errors.New("Error from Infusion: " + string(rBody))
	}

	log.Info("Contact creation complete. Result: " + string(rBody))
	return response, nil
}

// TagContact tag a contact in Infusion by contact and tag ID
func TagContact(contactID, tagID int) error {
	type taggerResponse []struct {
		Key string `json:"key"`
	}
	response := taggerResponse{}

	r, err := postRequest("/contacts/"+strconv.Itoa(contactID)+"/tags", []byte("{\"tagIds\": ["+strconv.Itoa(tagID)+"]}"))
	if err != nil {
		return err
	}

	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rBody, &response)
	if err != nil {
		return err
	}
	if r.StatusCode != 201 {
		return errors.New("Error from Infusion: " + string(rBody) + " CODE" + r.Status)
	}

	log.Info("Contact was tagged successfully. Result: " + string(rBody))
	return nil
}

// UnTagContact untag a contact in Infusion by contact and tag ID
func UnTagContact(contactID, tagID int) error {
	type taggerResponse []struct {
		Key string `json:"key"`
	}

	r, err := deleteRequest("/contacts/" + strconv.Itoa(contactID) + "/tags/" + strconv.Itoa(tagID))
	if err != nil {
		return err
	}

	if r.StatusCode != 204 {
		return errors.New("Got a different status code than 204. Status: " + r.Status)
	}
	log.Info("Contact untagged successfully.")
	return nil
}
