package dotparser

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type YAMLParser struct{}

func (YAMLParser) Parse(content string) (AppConfig, error) {
	var cfg AppConfig
	err := yaml.Unmarshal([]byte(content), &cfg)

	return cfg, err
}

func TestYamlParser(t *testing.T) {
	t.Logf("Starting test YAMLParser")
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "test_files", "config.yml")
	data, err := os.ReadFile(path)

	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}
	t.Logf("✅ Successfully read file (%d bytes)", len(data))

	appConfig, err := YAMLParser{}.Parse(string(data))
	require.NoError(t, err)
	assert.Equal(t, appConfig.Host, "localhost")
	assert.Equal(t, appConfig.Port, 8080)

	t.Logf("Test for finished")
}
