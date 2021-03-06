package gistdb

import "github.com/google/go-github/github"

func isGistIdinGistList(searchID string, gistlist []*github.Gist) bool {
	for _, g := range gistlist {
		if *(g.ID) == searchID {
			return true
		}
	}
	return false
}
func isFilenameInGistFileList(searchID string, g map[github.GistFilename]github.GistFile) bool {
	gSearchID := github.GistFilename(searchID)
	for k := range g {
		if k == gSearchID {
			return true
		}
	}
	return false
}

func deepCopyGist(gist github.Gist) github.Gist {
	newGist := gist
	newGist.Files = make(map[github.GistFilename]github.GistFile)
	for k, v := range gist.Files {
		newGist.Files[k] = v
	}
	return newGist
}
