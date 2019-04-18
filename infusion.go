package infusion

import "github.com/Sirupsen/logrus"

var log *logrus.Logger

// Init Initialize the package with the required items from Infusionsoft and get it ready for use
func Init(c Config) error {
	err := c.validate()
	if err != nil {
		return err
	}

	authToken = c.AuthToken
	refreshToken = c.RefreshToken
	log = c.Logger

	log.Info("Infusion initialized with auth token " + authToken)

	return nil
}
