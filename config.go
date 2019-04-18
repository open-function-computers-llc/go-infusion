package infusion

import (
	"github.com/Sirupsen/logrus"
)

// Config This is where we will add anything we need to make this library work
// the way we want it to. Infusionsoft uses OAUTH2 to authenticate and authorize users, so we need those fields to be mapped here
type Config struct {
	Logger       *logrus.Logger
	AuthToken    string
	RefreshToken string
}

func (c Config) validate() error {
	return nil
}
