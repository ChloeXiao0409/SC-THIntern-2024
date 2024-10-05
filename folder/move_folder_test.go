package folder_test

import (
	"testing"

	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	orgID1 := uuid.FromStringOrNil(folder.DefaultOrgID)
	orgID2 := uuid.Must(uuid.NewV4())

	initialFolders := []folder.Folder{
		{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
		{Name: "bravo", OrgId: orgID1, Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: orgID1, Paths: "alpha.bravo.charlie"},
		{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
		{Name: "echo", OrgId: orgID1, Paths: "alpha.delta.echo"},
		{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
		{Name: "golf", OrgId: orgID1, Paths: "golf"},
	}

	tests := []struct {
		name        string
		sourceName  string
		destName    string
		expected    []folder.Folder
		expectError bool
		errorMessage string
	}{
		{
			name:       "Move bravo to delta",
			sourceName: "bravo",
			destName:   "delta",
			expected: []folder.Folder{
				{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
				{Name: "bravo", OrgId: orgID1, Paths: "alpha.delta.bravo"},
				{Name: "charlie", OrgId: orgID1, Paths: "alpha.delta.bravo.charlie"},
				{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: orgID1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
				{Name: "golf", OrgId: orgID1, Paths: "golf"},
			},
			expectError: false,
		},
		{
			name:       "Move bravo to golf",
			sourceName: "bravo",
			destName:   "golf",
			expected: []folder.Folder{
				{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
				{Name: "bravo", OrgId: orgID1, Paths: "golf.bravo"},
				{Name: "charlie", OrgId: orgID1, Paths: "golf.bravo.charlie"},
				{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: orgID1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
				{Name: "golf", OrgId: orgID1, Paths: "golf"},
			},
			expectError: false,
		},
		{
			name:         "Move to child of itself",
			sourceName:   "bravo",
			destName:     "charlie",
			expectError:  true,
			errorMessage: "Cannot move a folder to a child of itself",
		},
		{
			name:         "Move to itself",
			sourceName:   "bravo",
			destName:     "bravo",
			expectError:  true,
			errorMessage: "Cannot move a folder to itself",
		},
		{
			name:         "Move to different organization",
			sourceName:   "bravo",
			destName:     "foxtrot",
			expectError:  true,
			errorMessage: "Cannot move a folder to a different organization",
		},
		{
			name:         "Invalid source folder",
			sourceName:   "invalid_folder",
			destName:     "delta",
			expectError:  true,
			errorMessage: "Source folder does not exist",
		},
		{
			name:         "Invalid destination folder",
			sourceName:   "bravo",
			destName:     "invalid_folder",
			expectError:  true,
			errorMessage: "Destination folder does not exist",
		},
		{
			name:       "Move to root",
			sourceName: "bravo",
			destName:   "",
			expected: []folder.Folder{
				{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
				{Name: "bravo", OrgId: orgID1, Paths: "bravo"},
				{Name: "charlie", OrgId: orgID1, Paths: "bravo.charlie"},
				{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: orgID1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
				{Name: "golf", OrgId: orgID1, Paths: "golf"},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := folder.NewDriver()
			d.SetFolders(initialFolders)

			result, err := d.MoveFolder(tt.sourceName, tt.destName)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				} else if err.Error() != tt.errorMessage {
					t.Errorf("Expected error message '%s', but got '%s'", tt.errorMessage, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !folder.EqualFolderSlices(result, tt.expected) {
					t.Errorf("Expected %v, but got %v", tt.expected, result)
				}
			}
		})
	}
}


// package folder_test

// import (
// 	"testing"
// )

// func Test_folder_MoveFolder(t *testing.T) {
// 	// TODO: your tests here
// }