package goperset

import (
	"context"
	"net/http"
)

type LoginPayload struct {
	Password string `json:"password"`
	Provider string `json:"provider"`
	Refresh  bool   `json:"refresh"`
	Username string `json:"username"`
}

type ClientToken struct {
	AccessToken string
	CsrfToken   string
}

type Goperset struct {
	BasePath string
	Client   *http.Client
	Context  context.Context
}

type DatabasePayload struct {
	Engine               string            `json:"engine"`
	ConfigurationMethod  string            `json:"configuration_method"`
	DatabaseName         string            `json:"database_name"`
	SQLAlchemyURI        string            `json:"sqlalchemy_uri"`
	AllowCTAS            bool              `json:"allow_ctas"`
	AllowCVAS            bool              `json:"allow_cvas"`
	AllowDML             bool              `json:"allow_dml"`
	AllowFileUpload      bool              `json:"allow_file_upload"`
	AllowRunAsync        bool              `json:"allow_run_async"`
	CacheTimeout         int               `json:"cache_timeout"`
	Driver               string            `json:"driver"`
	ExposeInSqlLab       bool              `json:"expose_in_sqllab"`
	ExternalURL          string            `json:"external_url"`
	Extra                string            `json:"extra"`
	ForceCTASSchema      string            `json:"force_ctas_schema"`
	ImpersonateUser      bool              `json:"impersonate_user"`
	IsManagedExternally  bool              `json:"is_managed_externally"`
	MaskedEncryptedExtra string            `json:"masked_encrypted_extra"`
	Parameters           map[string]string `json:"parameters"`
	ServerCert           string            `json:"server_cert"`
	SSHTunnel            SSHTunnelDetails  `json:"ssh_tunnel"`
	UUID                 string            `json:"uuid"`
}

type SSHTunnelDetails struct {
	ID                 int    `json:"id"`
	Password           string `json:"password"`
	PrivateKey         string `json:"private_key"`
	PrivateKeyPassword string `json:"private_key_password"`
	ServerAddress      string `json:"server_address"`
	ServerPort         int    `json:"server_port"`
	Username           string `json:"username"`
}

type EmbedPayload struct {
	AllowedDomains []string `json:"allowed_domains"`
}

type DashboardRolesPayload struct {
	Published      bool  `json:"published"`
	DashboardRoles []int `json:"roles"`
}

type RolePayload struct {
	RoleName string `json:"name"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type CsrfResponse struct {
	AccessToken string `json:"result"`
}

type RolesResponse struct {
	Result []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"result"`
}

type RoleResponse struct {
	ID     int `json:"id"`
	Result struct {
		Name string `json:"name"`
	} `json:"result"`
}

type EmbedResponse struct {
	Result struct {
		UUID        string `json:"uuid"`
		DashboardID string `json:"dashboard_id"`
	} `json:"result"`
}
