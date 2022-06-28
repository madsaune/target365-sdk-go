package environments

type BaseURL string

const (
	BaseURLShared BaseURL = "https://shared.target365.io/api"
	BaseURLTest   BaseURL = "https://test.target365.io/api"
)
