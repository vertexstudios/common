package entity

type Plugin struct {
	Id                string
	Owner             string
	Title             string
	Description       string
	License           string
	Repository        string
	Dependencies      []string
	Authors           []string
	SupportedVersions []string
	Price             float64
	UseSymmetry       bool
}

func NewPlugin(id, author, title string, price float64, description string, dependencies []string, authors []string, versions []string, license string, repository string, useSymmetry bool) *Plugin {
	return &Plugin{
		Id:                id,
		Owner:             author,
		Title:             title,
		Price:             price,
		Description:       description,
		Dependencies:      dependencies,
		Authors:           authors,
		SupportedVersions: versions,
		License:           license,
		Repository:        repository,
		UseSymmetry:       useSymmetry,
	}
}
