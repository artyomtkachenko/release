package config

type Config struct {
	Port    string
	LogFile string
	DataDir string
}

func New(port string, logFile string, dataDir string) *Config {
	return &Config{
		Port:    port,
		LogFile: logFile,
		DataDir: dataDir,
	}
}
