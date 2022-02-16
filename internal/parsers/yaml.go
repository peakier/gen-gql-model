package parsers

type Config struct {
	GraqhqlConfigs []GraqhqlConfig `yaml:"graqhql_configs,flow"`
	TypeConfigs    []TypeConfig    `yaml:"type_configs,flow"`
}

type GraqhqlConfig struct {
	Path        string `yaml:"path"`
	OutputDir   string `yaml:"output_dir"`
	PackageName string `yaml:"package_name"`
}

type TypeConfig struct {
	GqlType    string `yaml:"gql_type"`
	GoType     string `yaml:"go_type"`
	ImportType string `yaml:"import_type"`
}

type ReaderMode string

const (
	ReaderModeJson ReaderMode = "json"
	ReaderModeUrl  ReaderMode = "url"
)
