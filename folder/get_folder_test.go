package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// // feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		name:  "Matching folders",
			orgID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folders: []folder.Folder{
				{Name: "Folder1", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "Folder1"},
				{Name: "Folder2", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "Folder2"},
				{Name: "Folder3", OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")), Paths: "Folder3"},
			},
			want: []folder.Folder{
				{Name: "Folder1", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "Folder1"},
				{Name: "Folder2", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "Folder2"},
			},
		},
		{
			name:  "No matching folders",
			orgID: uuid.Must(uuid.FromString("33333333-3333-3333-3333-333333333333")),
			folders: []folder.Folder{
				{Name: "Folder1", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")), Paths: "Folder1"},
				{Name: "Folder2", OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")), Paths: "Folder2"},
			},
			want: []folder.Folder{},
		},
		{
			name:    "Empty folder list",
			orgID:   uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folders: []folder.Folder{},
			want:    []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// f := folder.NewDriver(tt.folders)
			// get := f.GetFoldersByOrgID(tt.orgID)
			f := folder.NewDriver(tt.folders)
			got := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, got)

		})
	}
}

