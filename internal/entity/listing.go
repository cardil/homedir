package entity

// Listing holds a listing f a given path.
type Listing struct {
	Path  string `json:"path"`
	Files []File `json:"files"`
}

// File holds information about given file.
type File struct {
	Name string `json:"name"`
	Mode string `json:"mode"`
	Size int    `json:"size"`
}
