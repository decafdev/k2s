package traefik

// StripPathPrefixRegex struct
type StripPathPrefixRegex struct {
	StripPrefixRegex StripPrefixRegex `yaml:"stripPrefixRegex,omitempty"`
}

// StripPrefixRegex struct
type StripPrefixRegex struct {
	Regex []string `yaml:"regex,omitempty"`
}

// NewStripPathPrefixRegexMiddleware function description
func NewStripPathPrefixRegexMiddleware(regex []string) *StripPathPrefixRegex {
	return &StripPathPrefixRegex{
		StripPrefixRegex: StripPrefixRegex{
			Regex: regex,
		},
	}
}
