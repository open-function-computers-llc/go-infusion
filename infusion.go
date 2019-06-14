package infusion

import "github.com/Sirupsen/logrus"

var log *logrus.Logger
var tagMapper map[string]int

// Init Initialize the package with the required items from Infusionsoft and get it ready for use
func Init(c Config) error {
	err := c.validate()
	if err != nil {
		return err
	}

	authToken = c.AuthToken
	refreshToken = c.RefreshToken
	clientID = c.ClientID
	clientSecret = c.ClientSecret
	log = c.Logger
	tagMapper = c.TagMapper
	for key, val := range tagMapper {
		log.Info("New Infusion tag associated, key: ", key, " id: ", val)
	}

	log.Info("Infusion initialized with auth token " + authToken + " and refresh token " + refreshToken)
	log.Info("Checking to make sure token is valid")

	_, err = GetAccountProfile()
	if err != nil {
		if err.Error() == "Invalid status code from Infusion" {
			err = GetRefreshToken()
			if err != nil {
				return err
			}

			// try this again
			_, err = GetAccountProfile()
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
