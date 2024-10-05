package folder

import (
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...
	folders := f.GetFoldersByOrgID(orgID)
	// Find parent folder
	var parentFolder Folder
	parentFound := false
	for _, folder := range folders {
		if folder.Name == name {
			parentFolder = folder
			parentFound = true
			break
		}
	}

	// If parent folder not found, return empty slice
	if !parentFound {
		return []Folder{}
	}

	// Find all children folders
	children := []Folder{}
	parentPathPrefix := parentFolder.Paths + "."
	for _, folder := range folders {
		if strings.HasPrefix(folder.Paths, parentPathPrefix) {
			// check if it's an immediate child
			restPath := strings.TrimPrefix(folder.Paths, parentPathPrefix)
			if !strings.Contains(restPath, ".") {
				children = append(children, folder)
			}
		}
	}

	// return []Folder{}
	return children
}
