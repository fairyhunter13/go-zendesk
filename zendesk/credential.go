package zendesk

// Type for authentication method in http specification
const (
	TypeBasicAuth = `Basic`
	TypeBearer    = `Bearer`
)

// Credential is interface of API credential
type Credential interface {
	Email() string
	Secret() string
	Type() string
}

// BasicAuthCredential is type of credential for Basic authentication
type BasicAuthCredential struct {
	email    string
	password string
}

// NewBasicAuthCredential creates BasicAuthCredential and returns its pointer
func NewBasicAuthCredential(email string, password string) *BasicAuthCredential {
	return &BasicAuthCredential{
		email:    email,
		password: password,
	}
}

// Email is accessor which returns email address
func (c BasicAuthCredential) Email() string {
	return c.email
}

// Secret is accessor which returns password
func (c BasicAuthCredential) Secret() string {
	return c.password
}

// Type is the type of credential being used.
func (c BasicAuthCredential) Type() string {
	return TypeBasicAuth
}

// APITokenCredential is type of credential for API token authentication
type APITokenCredential struct {
	email    string
	apiToken string
}

// NewAPITokenCredential creates APITokenCredential and returns its pointer
func NewAPITokenCredential(email string, apiToken string) *APITokenCredential {
	return &APITokenCredential{
		email:    email,
		apiToken: apiToken,
	}
}

// Email is accessor which returns email address
func (c APITokenCredential) Email() string {
	return c.email + "/token"
}

// Secret is accessor which returns API token
func (c APITokenCredential) Secret() string {
	return c.apiToken
}

// Type is the type of credential being used.
func (c APITokenCredential) Type() string {
	return TypeBasicAuth
}

// OAuthCredential is a type of credential for OAuth.
type OAuthCredential struct {
	oAuthToken string
}

// NewOAuthCredential creates OAuthCredential and returns its pointer
func NewOAuthCredential(token string) *OAuthCredential {
	return &OAuthCredential{
		oAuthToken: token,
	}
}

// Email is accessor which returns email address
func (c OAuthCredential) Email() string {
	return ""
}

// Secret is accessor which returns API token
func (c OAuthCredential) Secret() string {
	return c.oAuthToken
}

// Type is the type of credential being used.
func (c OAuthCredential) Type() string {
	return TypeBearer
}
