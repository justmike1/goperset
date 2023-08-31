package goperset

import "net/http"

type LoginPayload struct {
	Password string `json:"password"`
	Provider string `json:"provider"`
	Refresh  bool   `json:"refresh"`
	Username string `json:"username"`
}

type Goperset struct {
	BasePath string
	Client   *http.Client
}

type DatabasePayload struct {
	Engine              string `json:"engine"`
	ConfigurationMethod string `json:"configuration_method"`
	DatabaseName        string `json:"database_name"`
	SQLAlchemyURI       string `json:"sqlalchemy_uri"`
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
