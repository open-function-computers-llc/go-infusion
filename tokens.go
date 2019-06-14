package infusion

// GetCurrentTokens return the two different tokens we are currently using to talk with infusionsoft
func GetCurrentTokens() (string, string) {
	return authToken, refreshToken
}
