package entity

import "time"

type PluginRelease struct {
	ID         string
	Plugin     string
	Version    string
	IssuedDate time.Time
}

func NewPluginRelease(plugin, id, version string, issued time.Time) *PluginRelease {
	return &PluginRelease{
		Plugin:     plugin,
		Version:    version,
		IssuedDate: issued,
		ID:         id,
	}
}
