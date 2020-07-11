package mock

import (
	"github.com/fairyhunter13/go-zendesk/zendesk"
)

var _ zendesk.API = (*Client)(nil)
