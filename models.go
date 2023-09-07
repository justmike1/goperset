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

type DashboardPayload struct {
	CertificationDetails *string `json:"certification_details"`
	CertifiedBy          *string `json:"certified_by"`
	Css                  *string `json:"css"`
	DashboardTitle       *string `json:"dashboard_title"`
	ExternalUrl          *string `json:"external_url"`
	IsManagedExternally  *bool   `json:"is_managed_externally"`
	JsonMetadata         *string `json:"json_metadata"`
	Owners               []*int  `json:"owners"`
	PositionJson         *string `json:"position_json"`
	Published            *bool   `json:"published"`
	Roles                []*int  `json:"roles"`
	Slug                 *string `json:"slug"`
}

type GetDashboardParams struct {
	Columns        []*string `json:"columns"`
	Filters        []Filter  `json:"filters"`
	Keys           []*string `json:"keys"`
	OrderColumn    *string   `json:"order_column"`
	OrderDirection *string   `json:"order_direction"`
	Page           *int      `json:"page"`
	PageSize       *int      `json:"page_size"`
}

type Filter struct {
	Col   *string `json:"col"`
	Opr   *string `json:"opr"`
	Value *int    `json:"value"`
}

type PageInfo struct {
	Page     *int `json:"page"`
	PageSize *int `json:"page_size"`
}

type DashboardInfoParams struct {
	AddColumns  map[string]PageInfo `json:"add_columns"`
	EditColumns map[string]PageInfo `json:"edit_columns"`
	Keys        []*string           `json:"keys"`
}

type DatabasePayload struct {
	Engine               *string            `json:"engine"`
	ConfigurationMethod  *string            `json:"configuration_method"`
	DatabaseName         *string            `json:"database_name"`
	SQLAlchemyURI        *string            `json:"sqlalchemy_uri"`
	AllowCTAS            *bool              `json:"allow_ctas"`
	AllowCVAS            *bool              `json:"allow_cvas"`
	AllowDML             *bool              `json:"allow_dml"`
	AllowFileUpload      *bool              `json:"allow_file_upload"`
	AllowRunAsync        *bool              `json:"allow_run_async"`
	CacheTimeout         *int               `json:"cache_timeout"`
	Driver               *string            `json:"driver"`
	ExposeInSqlLab       *bool              `json:"expose_in_sqllab"`
	ExternalURL          *string            `json:"external_url"`
	Extra                *string            `json:"extra"`
	ForceCTASSchema      *string            `json:"force_ctas_schema"`
	ImpersonateUser      *bool              `json:"impersonate_user"`
	IsManagedExternally  *bool              `json:"is_managed_externally"`
	MaskedEncryptedExtra *string            `json:"masked_encrypted_extra"`
	Parameters           *map[string]string `json:"parameters"`
	ServerCert           *string            `json:"server_cert"`
	SSHTunnel            *SSHTunnelDetails  `json:"ssh_tunnel"`
	UUID                 *string            `json:"uuid"`
}

type SSHTunnelDetails struct {
	ID                 *int    `json:"id"`
	Password           *string `json:"password"`
	PrivateKey         *string `json:"private_key"`
	PrivateKeyPassword *string `json:"private_key_password"`
	ServerAddress      *string `json:"server_address"`
	ServerPort         *int    `json:"server_port"`
	Username           *string `json:"username"`
}

type DatasetPayload struct {
	Database            *int    `json:"database"`
	ExternalURL         *string `json:"external_url"`
	IsManagedExternally *bool   `json:"is_managed_externally"`
	Owners              *[]int  `json:"owners"`
	Schema              *string `json:"schema"`
	SQL                 *string `json:"sql"`
	TableName           *string `json:"table_name"`
}

type ChartPayload struct {
	CacheTimeout           *int      `json:"cache_timeout"`
	CertificationDetails   *string   `json:"certification_details"`
	CertifiedBy            *string   `json:"certified_by"`
	Dashboards             *[]int    `json:"dashboards"`
	DatasourceID           *int      `json:"datasource_id"`
	DatasourceName         *string   `json:"datasource_name"`
	DatasourceType         *string   `json:"datasource_type"`
	Description            *string   `json:"description"`
	ExternalURL            *string   `json:"external_url"`
	IsManagedExternally    *bool     `json:"is_managed_externally"`
	Owners                 *[]int    `json:"owners"`
	Params                 *string   `json:"params"`
	QueryContext           *string   `json:"query_context"`
	QueryContextGeneration *bool     `json:"query_context_generation"`
	SliceName              *string   `json:"slice_name"`
	VizType                *[]string `json:"viz_type"`
}

type EmbedPayload struct {
	AllowedDomains *[]string `json:"allowed_domains"`
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
