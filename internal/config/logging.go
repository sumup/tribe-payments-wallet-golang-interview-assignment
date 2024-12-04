package config

type Log struct {
	// Level with possible values: `DEBUG|ERROR|FATAL|INFO|PANIC|WARN` E.g INFO.
	Level string `default:"INFO" envconfig:"LOG_LEVEL"`

	// StdoutEnabled is a toggle whether to log in stdout or not.
	StdoutEnabled bool `default:"true" envconfig:"STDOUT_LOG_ENABLED"`
}
