package goperset

// Secure Endpoint
const (
	SecurityController   = "/api/v1/security/"
	RolesEndpoint        = SecurityController + "roles/"
	csrfEndpoint         = SecurityController + "csrf_token/"
	guestLoginEndpoint   = SecurityController + "guest_token/"
	loginEndpoint        = SecurityController + "login"
	refreshLoginEndpoint = SecurityController + "refresh"
)

// Async Events Endpoints
const (
	AsyncEventsController = "/api/v1/async_event/"
)

// Available Domains Endpoints
const (
	AvaliableDomainsController = "/api/v1/available_domains/"
)

// Cache Keys Endpoints
const (
	CacheKeysController = "/api/v1/cache_keys/"
)

// CSS Template Endpoints
const (
	CssTemplateController = "/api/v1/css_template/"
)

// Dashboard Endpoints
const (
	DashboardController      = "/api/v1/dashboard/"
	DashboardInfoEndpoint    = DashboardController + "_info"
	DashboardRelatedEndpoint = DashboardController + "related/"
	importDashboardEndpoint  = DashboardController + "import/"
)

// Database Endpoints
const (
	DatabaseController = "/api/v1/database/"
)

// Annotation Layer Endpoints
const (
	AnnotationLayerController         = "/api/v1/annotation_layer/"
	AnnotationLayerInfoEndpoint       = AnnotationLayerController + "_info"
	RelatedAnnotationLayerEndpoint    = AnnotationLayerController + "related/"
	AnnotationLayerAnnotationEndpoint = AnnotationLayerController + "annotation/"
)

// Chart Endpoints
const (
	ChartController             = "/api/v1/chart/"
	ChartInfoEndpoint           = ChartController + "_info"
	ChartDataEndpoint           = ChartController + "data"
	ExportChartEndpoint         = ChartController + "export/"
	CHartFavoriteStatusEndpoint = ChartController + "favorite_status/"
	ImportChartEndpoint         = ChartController + "import/"
	RelatedChartEndpoint        = ChartController + "related/"
)

// Advanced Data Type Endpoints
const (
	AdvancedDataTypeController      = "/api/v1/advanced_data_type/"
	ConvertAdvancedDataTypeEndpoint = AdvancedDataTypeController + "convert"
	AdvancedDataTypesEndpoint       = AdvancedDataTypeController + "types"
)

// Assets Endpoints
const (
	AssetsController     = "/api/v1/assets/"
	ImportAssetsEndpoint = AssetsController + "import/"
	ExportAssetsEndpoint = AssetsController + "export/"
)

// Dataset Endpoints
const (
	DatasetController = "/api/v1/dataset/"
)

// API Endpoint constants
const (
	InvalidateCacheKeyEndpoint = CacheKeysController + "invalidate"
	CssTemplateInfoEndpoint    = CssTemplateController + "_info"
	CssTemplateRelatedEndpoint = CssTemplateController + "related/"
)
