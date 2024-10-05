package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...
	folders := f.folders

	sourceIndex := -1
	destIndex := -1
	for i, folder := range folders {
		if folder.Name == name {
			sourceIndex = i
		}
		if folder.Name == dst {
			destIndex = i
		}
	}

	if sourceIndex == -1 {
		return nil, errors.New("source folder does not exist")
	}
	if destIndex == -1 && dst != "" {
		return nil, errors.New("destination folder does not exist")
	}

	sourceFolder := folders[sourceIndex]
	var destFolder Folder
	if dst != "" {
		destFolder = folders[destIndex]
	}

	// check if moving to itself
	if name == dst {
		return nil, errors.New("cannot move a folder to itself")
	}
	// check if moving to a child of itself
	if strings.HasPrefix(destFolder.Paths, sourceFolder.Paths + ".") {
		return nil, errors.New("cannot move a folder to a child of itself")
	}
	// check if moving to a different organization
	if dst != "" && sourceFolder.OrgId != destFolder.OrgId {
		return nil, errors.New("cannot move a folder to a different organization")
	}

	newPath := destFolder.Paths
	if newPath != "" {
		newPath += "."
	}
	newPath += sourceFolder.Name

	// Update path
	result := make([]Folder, len(folders))
	copy(result, folders)
	for i, folder := range result {
		if strings.HasPrefix(folder.Paths, sourceFolder.Paths) {
			result[i].Paths = strings.Replace(folder.Paths, sourceFolder.Paths, newPath, 1)
		}
	}

	// return []Folder{}, nil
	return result, nil
}
