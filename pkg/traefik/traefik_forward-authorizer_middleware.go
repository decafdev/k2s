package traefik

type ForwardAuth struct {
	Address             string   `yaml:"address,omitempty"`
	TrustForwardHeader  bool     `yaml:"trustForwardHeader,omitempty"`
	AuthRequestHeaders  []string `yaml:"authRequestHeaders,omitempty"`
	AuthResponseHeaders []string `yaml:"authResponseHeaders,omitempty"`
}

// ForwardAuthorizer struct
type ForwardAuthorizer struct {
	ForwardAuth *ForwardAuth `yaml:"forwardAuth,omitempty"`
}

// NewForwardAuthorizerMiddleware function description
func NewForwardAuthorizerMiddleware(o *ForwardAuth) *ForwardAuthorizer {
	return &ForwardAuthorizer{ForwardAuth: o}
}
