package service

import "demi-anchor/internal/models"

func (s *service) ValidateDonation(d *models.Donation) (bool, string) {
	if d.StreamerID == 0 {
		return false, "StreamerID can't be '0'"
	}
	if d.Author == "" {
		return false, "Author can't be empty"
	}
	if d.Money == 0 {
		return false, "Money can't be '0'"
	}
	if d.Time.IsZero() {
		return false, "Time is empty"
	}

	return true, ""
}

func (s *service) ValidatePeriod(d *models.Period) (bool, string) {
	if d.FirstDate.After(d.LastDate) {
		return false, "Wrong period. Last date < first date"
	}
	return true, ""
}
