package entity

import "time"

type LicenseLog struct {
	LicenseKey string      `bson:"license"`
	Type       int         `bson:"type"`
	Details    interface{} `bson:"details"`
	Date       time.Time   `bson:"date"`
}

func NewLicenseLog(editType int, details interface{}, date time.Time, licenseKey string) *LicenseLog {
	return &LicenseLog{
		LicenseKey: licenseKey,
		Type:       editType,
		Details:    details,
		Date:       date,
	}
}
