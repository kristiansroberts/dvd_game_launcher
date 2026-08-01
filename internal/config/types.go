package config

type Game struct {
	Name    string
	Dir     string
	ExePath string
}

type Extras struct {
	Name string
	Dir  string
}

// Art, OST, Manual, and Video structs are used to represent additional resources included with the game
type Art struct {
	Name string
	Dir  string
}

type OST struct {
	Name string
	Dir  string
}

type Manual struct {
	Name string
	Dir  string
}

type Video struct {
	Name string
	Dir  string
}

type ExtraType int

const (
	ExtraTypeUnknown ExtraType = iota
	ExtraTypeArt
	ExtraTypeOST
	ExtraTypeManual
	ExtraTypeVideo
)

// item level struct to represent an instance of an extra resource
type ExtraItem struct {
	Name      string
	ExtraType ExtraType
	FilePath  string
}
