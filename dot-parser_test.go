package dotparser

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type AppConfig struct {
	NetworkConfiguration struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"network"`
}

func TestYamlParser(t *testing.T) {
	t.Logf("Starting test YAMLParser")
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "test_files")
	config := DotFileConfiguration{
		Project:        "application",
		ConfigFileName: "config.yml",
		GetHomeDirFunc: func() (string, error) {
			return dir, nil
		},
	}

	data, err := os.ReadFile(ConstructPath(config))

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("✅ Successfully read file (%d bytes)", len(data))

	appConfig, err := YamlParser[AppConfig]{}.Parse(string(data))
	require.NoError(t, err)
	assert.Equal(t, appConfig.NetworkConfiguration.Host, "localhost")
	assert.Equal(t, appConfig.NetworkConfiguration.Port, 8080)

	t.Logf("Test for finished")
}
