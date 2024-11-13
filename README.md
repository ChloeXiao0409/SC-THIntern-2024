This is for submission of SafetyCulture Summer Internship 2024 Take-Home Challenge.

Copy with the instructions
## Instructions

- This technical assessment consists of 2 components:
- Component 1:

  - within `get_folder.go`.
    - We would like you to read through, and run, the code.
    - Implement `GetAllChildFolders` method in `get_folder.go` that returns all child folders of a given folder.
    - Write up some unit tests in `get_folder_test.go` for all methods in `get_folder.go`.

- Component 2:
  - within `move_folder.go`.
    - Implement `MoveFolder` method in `move_folder.go` that moves a folder from one parent to another. (more details under component 2 section)
    - Write up some unit tests in `move_folder_test.go` for the `MoveFolder` method.

## Path Structure

You are given a hierarchical tree where each node in the tree is represented by a path similar to `ltree` paths in PostgreSQL.

The tree structure is represented as a series of paths, where each path is folder name separated by dots (e.g., `"alpha.bravo.charlie"`). Each name in the path represents a node, and the full path represents that node’s position in the hierarchy.

we use `ltree` path for our site directory structure as well as our documents folder structure within the SC platform. This allow us to easily store and manipulate our folder structure using psql.

## Component 1

You will need to implement the following:

1. A method to get all child folders of a given folder.
2. The method should return a list of all child folders.
3. Implement any necessary error handling (e.g. invalid orgID, invalid paths, etc).

### Example Scenario

```go
folders := [
  {
    name: "alpha",
    path: "alpha",
    orgID: "org1",
  },
  {
    name: "bravo",
    path: "alpha.bravo",
    orgID: "org1",
  },
  {
    name: "charlie",
    path : "alpha.bravo.charlie",
    orgID: "org1",
  },
  {
    name: "delta",
    path: "alpha.delta",
    orgID: "org1",
  },
  {
    name: "echo",
    path: "echo",
    orgID: "org1",
  },
  {
    name: "foxtrot",
    path: "foxtrot",
    orgID: "org2",
  },
]

getAllChildFolders("org1", "alpha")
// Expected output
[
   {
    name: "bravo",
    path: "alpha.bravo",
    orgID: "org1",
  },
  {
    name: "charlie",
    path : "alpha.bravo.charlie",
    orgID: "org1",
  },
  {
    name: "delta",
    path: "alpha.delta",
    orgID: "org1",
  },
]

getAllChildFolders("org1", "bravo")
// Expected output
[
  {
    name: "charlie",
    path : "alpha.bravo.charlie",
    orgID: "org1",
  },
]

getAllChildFolders("org1", "charlie")
// Expected output
[]

getAllChildFolders("org1", "echo")
// Expected output
[]

getAllChildFolders("org1", "invalid_folder")
// Error: Folder does not exist

getAllChildFolders("org1", "foxtrot")
// Error: Folder does not exist in the specified organization
```

## Component 2

You will need to implement the following:

1. A method to move a subtree from one parent node to another, while maintaining the order of the children.
2. The method should return the new folder structure once the move has occurred.
3. Implement any necessary error handling (e.g. invalid paths, moving a node to a child of itself, moving folders to a different orgID, etc).
4. There is no need to persist state, we can assume each method call will be independent of the previous one.

### Example Scenario

```go

folders := [
  {
    name: "alpha",
    path: "alpha",
    orgID: "org1",
  },
  {
    name: "bravo",
    path: "alpha.bravo",
    orgID: "org1",
  },
  {
    name: "charlie",
    path: "alpha.bravo.charlie",
    orgID: "org1",
  },
  {
    name: "delta",
    path: "alpha.delta",
    orgID: "org1",
  },
  {
    name: "echo",
    path: "alpha.delta.echo",
    orgID: "org1",
  },
  {
    name: "foxtrot",
    path: "foxtrot",
    orgID: "org2",
  }
  {
    name: "golf",
    path: "golf",
    orgID: "org1",
  }
]

moveFolder("bravo", "delta")
// Expected output
[
  {
    name: "alpha",
    path: "alpha",
    orgID: "org1",
  },
  {
    name: "bravo",
    path: "alpha.delta.bravo",
    orgID: "org1",
  },
  {
    name: "charlie",
    path: "alpha.delta.bravo.charlie",
    orgID: "org1",
  },
  {
    name: "delta",
    path: "alpha.delta",
    orgID: "org1",
  },
  {
    name: "echo",
    path: "alpha.delta.echo",
    orgID: "org1",
  },
  {
    name: "foxtrot",
    path: "foxtrot",
    orgID: "org2",
  }
  {
    name: "golf",
    path: "golf",
    orgID: "org1",
  }
]

moveFolder("bravo", "golf")
// Expected output
[
  {
    name: "alpha",
    path: "alpha",
    orgID: "org1",
  },
  {
    name: "bravo",
    path: "golf.bravo",
    orgID: "org1",
  },
  {
    name: "charlie",
    path: "golf.bravo.charlie",
    orgID: "org1",
  },
  {
    name: "delta",
    path: "alpha.delta",
    orgID: "org1",
  },
  {
    name: "echo",
    path: "alpha.delta.echo",
    orgID: "org1",
  },
  {
    name: "foxtrot",
    pa th: "foxtrot",
    orgID: "org2",
  },
  {
    name: "golf",
    path: "golf",
    orgID: "org1",
  }
]

moveFolder("bravo", "charlie")
// Error: Cannot move a folder to a child of itself

moveFolder("bravo", "bravo")
// Error: Cannot move a folder to itself

moveFolder("bravo", "foxtrot")
// Error: Cannot move a folder to a different organization

moveFolder("invalid_folder", "delta")
// Error: Source folder does not exist

moveFolder("bravo", "invalid_folder")
// Error: Destination folder does not exist