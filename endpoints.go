package goperset

// Secure Endpoint
const (
	SecurityController   = "/api/v1/security/"
	csrfEndpoint         = SecurityController + "csrf_token/"
	guestLoginEndpoint   = SecurityController + "guest_token/"
	loginEndpoint        = SecurityController + "login"
	refreshLoginEndpoint = SecurityController + "refresh"
)

const (
	DatabaseController         = "/api/v1/database/"
	DashboardController        = "/api/v1/dashboard/"
	AdvancedDataTypeController = "/api/v1/advanced_data_type/"
	AnnotationLayerController  = "/api/v1/annotation_layer/"
	AssetsController           = "/api/v1/assets/"
	AsyncEventsController      = "/api/v1/async_event/"
	AvaliableDomainsController = "/api/v1/available_domains/"
	CacheKeysController        = "/api/v1/cache_keys/"
	ChartController            = "/api/v1/chart/"
	CssTemplateController      = "/api/v1/css_template/"
)

// API Endpoint constants
const (
	importDashboardEndpoint           = DashboardController + "import/"
	RolesEndpoint                     = SecurityController + "roles/"
	ConvertAdvancedDataTypeEndpoint   = AdvancedDataTypeController + "convert"
	AdvancedDataTypesEndpoint         = AdvancedDataTypeController + "types"
	AnnotationLayerInfoEndpoint       = AnnotationLayerController + "_info"
	RelatedAnnotationLayerEndpoint    = AnnotationLayerController + "related/"
	AnnotationLayerAnnotationEndpoint = AnnotationLayerController + "annotation/"
	ImportAssetsEndpoint              = AssetsController + "import/"
	ExportAssetsEndpoint              = AssetsController + "export/"
	InvalidateCacheKeyEndpoint        = CacheKeysController + "invalidate"
	ChartInfoEndpoint                 = ChartController + "_info"
	ChartDataEndpoint                 = ChartController + "data"
	ExportChartEndpoint               = ChartController + "export/"
	CHartFavoriteStatusEndpoint       = ChartController + "favorite_status/"
	ImportChartEndpoint               = ChartController + "import/"
	RelatedChartEndpoint              = ChartController + "related/"
	CssTemplateInfoEndpoint           = CssTemplateController + "_info"
	CssTemplateRelatedEndpoint        = CssTemplateController + "related/"
)
