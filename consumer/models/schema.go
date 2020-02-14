package models

type Action struct {
	ID          int    `json:"id"`
	EventName   string `json:"event_name"`
	ActionType  string `json:"action_type"`
	ActionSpec  string `json:"action_spec"`
	AccountName string `json:"created_by"`
}

type HTTPActionSpec struct {
	URL         string          `json:"url"`
	Method      string          `json:"method"`
	Body        string          `json:"request_body,omitempty"`
	ContentType HTTPContentType `json:"content_type,omitempty"`
	Headers     []Header        `json:"headers,omitempty"`
	Auth        *HTTPAuth       `json:"authentication,omitempty"`
	TLSVerify   bool            `json:"tls_verify,omitempty"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HTTPAuth struct {
	AuthType  AuthenticationType `json:"auth_type"`
	BasicAuth *HTTPBasicAuth     `json:"basic_auth"`
}

type AuthenticationType string

//define the types of authentication possible
const (
	BASIC = AuthenticationType("basic")
	NONE  = AuthenticationType("none")
)

type HTTPBasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//HTTPContentType specifies content type of the request
type HTTPContentType string

//these constants define the types of content possible for the request
const (
	JSON = HTTPContentType("application/json")
	XML  = HTTPContentType("application/xml")
	HTML = HTTPContentType("text/html")
)

type EmailActionSpec struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Password string `json:"password"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}

type CalmActionSpec struct {
	AppUUID    string `json:"app_uuid"`
	ActionUUID string `json:"action_uuid"`
}
