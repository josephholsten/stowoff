package app

import (
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
)

// Application describes the config files for each application
type Application struct {
	Name  string
	Files []string
}

func Load(name string) (*Application, error) {
	cfgPath := filepath.Join("mackup", "mackup", "applications", name+".cfg")

	return LoadSource(name, cfgPath)
}

// LoadSource application from data source, which can be file name with string type, or raw data in []byte
func LoadSource(name string, source interface{}) (*Application, error) {
	cfg, err := ini.LoadSources(
		ini.LoadOptions{UnparseableSections: []string{"configuration_files"}},
		source,
	)

	rawFiles := cfg.Section("configuration_files").Body()
	files := strings.Split(rawFiles, "\n")
	for i := range files {
		files[i] = strings.TrimSpace(files[i])
	}

	app := Application{
		Name:  cfg.Section("application").Key("name").Value(),
		Files: files,
	}

	return &app, err
}
