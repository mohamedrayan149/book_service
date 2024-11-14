package config

// Parameter Name Constants
const (
	IDParam         = "id"
	TitleParam      = "title"
	AuthorParam     = "author_name"
	PriceRangeParam = "price_range"
	UsernameParam   = "username"
)

// Error Message Constants
const (
	ErrParameterRequired       = "at least one of the following parameters is required: title, author_name, or price_range"
	ErrorUsernameRequired      = "username is required"
	ErrUserNameRequired        = "username is a required field."
	ErrIDRequired              = "id is a required field."
	ErrTitleRequired           = "title is a required field."
	ErrAuthorNameRequired      = "author name is a required field."
	ErrPriceRequired           = "price is a required field and must be a number."
	ErrFieldRequired           = "this field is required."
	ErrPublishDateRequired     = "publish date is a required field."
	ErrInvalidPriceRangeFormat = "invalid price range format, expected 'min-max'"
	ErrInvalidPriceRangeValues = "invalid price range values, expected numbers"
)

// Success Message Constants
const (
	SuccessBookDeleted = "book deleted successfully"
	SuccessBookUpdated = "book updated successfully"
)

// Response Field Constants
const (
	ActionsField         = "actions"
	BookCountField       = "book_count"
	DistinctAuthorsField = "distinct_authors"
	MessageField         = "message"
	ErrorField           = "error"
)

// Elasticsearch Constants
const (
	ElasticDevURL      = "http://es-search-7.fiverrdev.com:9200"
	ErrorInitClient    = "error creating Elasticsearch client: %s"
	IndexBooks         = "books_mohamed"
	FieldAuthorNameKey = "author_name.keyword"
	AggDistinctAuthors = "distinct_authors"
	NoResultFound      = "no result found"
)

var ElasticsearchErrorMap = map[int]string{
	400: "Bad Request: The request was invalid or improperly formatted.",
	401: "Unauthorized: Authentication required or failed.",
	403: "Forbidden: You do not have permission to perform this action.",
	404: "Not Found: The requested resource could not be found.",
	408: "Request Timeout: The request took too long to process.",
	409: "Conflict: A version conflict occurred, possibly due to concurrent updates.",
	429: "Too Many Requests: Elasticsearch is throttling due to high load.",
	500: "Internal Server Error: An error occurred within Elasticsearch.",
	503: "Service Unavailable: The cluster is unavailable, possibly due to maintenance or overload.",
	509: "Circuit Breaking Exception: Memory limit reached; Elasticsearch is protecting itself from OOM.",
}

// Redis Configuration Constants
const (
	RedisDevAddr          = "redis-search.fiverrdev.com:6382"
	RedisDB               = 0
	ErrorConnectingRedis  = "error connecting to Redis: %v"
	UserActionsKeyPattern = "user:%s:actions"
	ActionHistoryLimit    = -3
	EndOfList             = -1
)

// Route Constants
const (
	BooksBaseRoute   = "/books"
	StoreStatsRoute  = "/store"
	SearchBooksRoute = "/search"
	ActivityRoute    = "/activity"
)

// Miscellaneous Constants
const (
	Space               = " "
	Slash               = "/"
	PriceRangeDelimiter = "-"
	ParseBitSize        = 64
	EmptyString         = ""
	Zero                = 0
)

// Server Configuration
const (
	ServerPort = ":8080"
)

// Database Field Names
const (
	FieldID          = "id"
	FieldTitle       = "title"
	FieldAuthorName  = "author_name"
	FieldPrice       = "price"
	FieldPublishDate = "publish_date"
)
