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

// String method for ExtraType to provide an actual string representation of the enum value
func (et ExtraType) String() string {
	switch et {
	case ExtraTypeArt:
		return "Art"
	case ExtraTypeOST:
		return "OST"
	case ExtraTypeManual:
		return "Manual"
	case ExtraTypeVideo:
		return "Video"
	default:
		return "Unknown"
	}
}

// item level struct to represent an instance of an extra resource
type ExtraItem struct {
	Name      string
	ExtraType ExtraType
	FilePath  string
}
