package entity

type LicenseBundle struct {
	ID       string `bson:"_id"`
	PluginID string `bson:"plugin"`
	MaxNodes int    `bson:"max_nodes"`
	Title    string `bson:"title"`
}

func NewLicenseBundle(id, plugin, title string, maxNodes int) *LicenseBundle {
	return &LicenseBundle{
		PluginID: plugin,
		ID:       id,
		MaxNodes: maxNodes,
		Title:    title,
	}
}
