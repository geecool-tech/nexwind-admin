package constant

const (
	PROD = "prod"
	DEV  = "dev"
)

// CONFIG_FILE config file map.
var CONFIG_FILE = map[string]string{
	PROD: "app.yaml",
	DEV:  "app.dev.yaml",
}
