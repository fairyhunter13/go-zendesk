package zendesk

import (
	"fmt"
	"net/url"
)

// Option specifies the functional options for this package
type Option func(*Client) error

// WithHeader saves HTTP header in client. It will be included all API request
func WithHeader(key, value string) Option {
	return func(z *Client) error {
		z.headers[key] = value
		return nil
	}
}

// WithSubdomain saves subdomain in client. It will be used
// when call API
func WithSubdomain(subdomain string) Option {
	return func(z *Client) error {
		if !subdomainRegexp.MatchString(subdomain) {
			return fmt.Errorf("%s is invalid subdomain", subdomain)
		}

		baseURLString := fmt.Sprintf(baseURLFormat, subdomain)
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			return err
		}

		z.baseURL = baseURL
		return nil

	}
}

// WithEndpointURL replace full URL of endpoint without subdomain validation.
// This is mainly used for testing to point to mock API server.
func WithEndpointURL(newURL string) Option {
	return func(z *Client) error {
		baseURL, err := url.Parse(newURL)
		if err != nil {
			return err
		}

		z.baseURL = baseURL
		return nil
	}
}

// WithCredential saves credential in client. It will be set
// to request header when call API
func WithCredential(cred Credential) Option {
	return func(z *Client) error {
		z.credential = cred
		return nil
	}
}
