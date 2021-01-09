package configs

type Config struct {
	DB      Database `yaml:"database"`
	Details `yaml:"details"`
}

type Database struct {
	Server string `yaml:"server"`
	Port   string `yaml:"port"`
}

type Details struct {
	Name       string `yaml:"database_name"`
	Collection string `yaml:"collection_name"`
}
