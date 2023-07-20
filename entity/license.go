package entity

import (
	"time"
)

type LicenseMode int64

const (
	Open LicenseMode = iota
	Closed
)

type License struct {
	Key          string      `bson:"key"`
	Holder       string      `bson:"holder"`
	PluginID     string      `bson:"plugin"`
	Bundle       string      `bson:"bundle"`
	AllowedNodes []string    `bson:"allowed_nodes"`
	CreatedOn    time.Time   `bson:"created_on"`
	ExpireOn     time.Time   `bson:"expire_on"`
	Mode         LicenseMode `bson:"mode"`
}

func NewLicense(holder, pluginID, key, bundle string, allowedNodes []string, createdOn, expireOn time.Time, mode LicenseMode) *License {
	return &License{
		Holder:       holder,
		AllowedNodes: allowedNodes,
		PluginID:     pluginID,
		Key:          key,
		Bundle:       bundle,
		CreatedOn:    createdOn,
		ExpireOn:     expireOn,
		Mode:         mode,
	}
}
