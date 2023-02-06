package lib

import (
	"github.com/chigopher/pathlib"
	"github.com/deeprave/go-testutils/test"
	"github.com/spf13/afero"
	"testing"
)

var testData = []byte("test:\n\t- first item\n\t- second item\n")
var testHome = "/home/testing"
var testFileName = "config.yml"

func setUpHome(home string, fs afero.Fs, configFileName string) *pathlib.Path {
	// make a "home" path
	homePath := pathlib.NewPathAfero(home, fs)
	_ = homePath.MkdirAll()
	_, _ = homePath.IsDir()

	// create a file in the home path with some data
	configFile := pathlib.NewPathAfero(configFileName, fs)
	configPath := homePath.JoinPath(configFile)
	_ = configPath.WriteFile(testData)
	return homePath
}

func TestFindPath(t *testing.T) {

	fs := afero.NewMemMapFs()
	homePath := setUpHome(testHome, fs, testFileName)
	configFile := pathlib.NewPathAfero(testFileName, fs)

	// how try to locate it
	foundPath, err := FindPath(configFile, ".", homePath)
	test.ShouldBeNoError(t, err, "%v", err)
	size, err := foundPath.Size()
	test.ShouldBeNoError(t, err, "%v (size): %v", configFile, err)
	test.ShouldBeEqual(t, len(testData), int(size))
	t.Logf("successfully found (path): %v, %d bytes", foundPath, size)
}

func TestFindFile(t *testing.T) {

	fs := afero.NewMemMapFs()
	homePath := setUpHome(testHome, fs, testFileName)

	// how try to locate it
	foundPath, err := FindFile(testFileName, fs, ".", homePath)
	test.ShouldBeNoError(t, err, "%v", err)
	if err == nil {
		t.Logf("successfully found (file): %v", foundPath)
	}
}
