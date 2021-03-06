package service

import (
	"testing"

	"github.com/mesg-foundation/core/service"
	"github.com/stvp/assert"
)

func TestDefaultPath(t *testing.T) {
	assert.Equal(t, defaultPath([]string{}), "./")
	assert.Equal(t, defaultPath([]string{"foo"}), "foo")
	assert.Equal(t, defaultPath([]string{"foo", "bar"}), "foo")
}

func TestBuildDockerImagePathDoNotExist(t *testing.T) {
	_, err := buildDockerImage("/doNotExist")
	assert.NotNil(t, err)
}

func TestGitCloneRepositoryDoNotExist(t *testing.T) {
	path, _ := createTempFolder()
	defer removeTempFolder(path, true)
	err := gitClone("/doNotExist", path)
	assert.NotNil(t, err)
}

func TestDownloadServiceIfNeededAbsolutePath(t *testing.T) {
	path := "/users/paul/service-js-function"
	newPath, didDownload, err := downloadServiceIfNeeded(path)
	assert.Nil(t, err)
	assert.Equal(t, path, newPath)
	assert.Equal(t, false, didDownload)
}

func TestDownloadServiceIfNeededRelativePath(t *testing.T) {
	path := "./service-js-function"
	newPath, didDownload, err := downloadServiceIfNeeded(path)
	assert.Nil(t, err)
	assert.Equal(t, path, newPath)
	assert.Equal(t, false, didDownload)
}

func TestDownloadServiceIfNeededUrl(t *testing.T) {
	path := "https://github.com/mesg-foundation/awesome.git"
	newPath, didDownload, err := downloadServiceIfNeeded(path)
	defer removeTempFolder(newPath, true)
	assert.Nil(t, err)
	assert.NotEqual(t, path, newPath)
	assert.Equal(t, true, didDownload)
}

func TestCreateTempFolder(t *testing.T) {
	path, err := createTempFolder()
	defer removeTempFolder(path, true)
	assert.Nil(t, err)
	assert.NotEqual(t, "", path)
}

func TestRemoveTempFolder(t *testing.T) {
	path, _ := createTempFolder()
	err := removeTempFolder(path, true)
	assert.Nil(t, err)
}

func TestInjectConfigurationInDependencies(t *testing.T) {
	s := &service.Service{}
	injectConfigurationInDependencies(s, "TestInjectConfigurationInDependencies")
	assert.Equal(t, s.Dependencies["service"].Image, "TestInjectConfigurationInDependencies")
}

func TestInjectConfigurationInDependenciesWithConfig(t *testing.T) {
	s := &service.Service{
		Configuration: &service.Dependency{
			Command: "xxx",
			Image:   "yyy",
		},
	}
	injectConfigurationInDependencies(s, "TestInjectConfigurationInDependenciesWithConfig")
	assert.Equal(t, s.Dependencies["service"].Image, "TestInjectConfigurationInDependenciesWithConfig")
	assert.Equal(t, s.Dependencies["service"].Command, "xxx")
}

func TestInjectConfigurationInDependenciesWithDependency(t *testing.T) {
	s := &service.Service{
		Dependencies: map[string]*service.Dependency{
			"test": &service.Dependency{
				Image: "xxx",
			},
		},
	}
	injectConfigurationInDependencies(s, "TestInjectConfigurationInDependenciesWithDependency")
	assert.Equal(t, s.Dependencies["service"].Image, "TestInjectConfigurationInDependenciesWithDependency")
	assert.Equal(t, s.Dependencies["test"].Image, "xxx")
}

func TestInjectConfigurationInDependenciesWithDependencyOverride(t *testing.T) {
	s := &service.Service{
		Dependencies: map[string]*service.Dependency{
			"service": &service.Dependency{
				Image: "xxx",
			},
		},
	}
	injectConfigurationInDependencies(s, "TestInjectConfigurationInDependenciesWithDependencyOverride")
	assert.Equal(t, s.Dependencies["service"].Image, "TestInjectConfigurationInDependenciesWithDependencyOverride")
}
