package contracts

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	manifest "github.com/estafette/estafette-ci-manifest"
	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

func TestUnmarshalBuilderConfigFromYaml(t *testing.T) {
	t.Run("ReturnsConfigWithoutErrors", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-api-test.yaml")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		// act
		err = yaml.Unmarshal(bytes, &config)

		assert.Nil(t, err)
	})

	t.Run("ReturnsCredentialsWithType", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-api-test.yaml")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		// act
		err = yaml.Unmarshal(bytes, &config)

		if !assert.Nil(t, err) {
			return
		}
		assert.Equal(t, 7, len(config.Credentials))
		assert.Equal(t, "container-registry", config.Credentials[0].Type)
		assert.Equal(t, "container-registry", config.Credentials[1].Type)
		assert.Equal(t, "kubernetes-engine", config.Credentials[2].Type)
		assert.Equal(t, "kubernetes-engine", config.Credentials[3].Type)
		assert.Equal(t, "bitbucket-api-token", config.Credentials[4].Type)
		assert.Equal(t, "github-api-token", config.Credentials[5].Type)
		assert.Equal(t, "slack-webhook", config.Credentials[6].Type)
	})

	t.Run("ReturnsCredentialsWithAdditionalProperties", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-api-test.yaml")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		// act
		err = yaml.Unmarshal(bytes, &config)

		if !assert.Nil(t, err) {
			return
		}
		assert.Equal(t, "extensions", config.Credentials[0].AdditionalProperties["repository"])
		assert.Equal(t, "username", config.Credentials[0].AdditionalProperties["username"])
		assert.Equal(t, "secret", config.Credentials[0].AdditionalProperties["password"])

		assert.Equal(t, "estafette-production", config.Credentials[2].AdditionalProperties["project"])
		assert.Equal(t, "europe-west2", config.Credentials[2].AdditionalProperties["region"])
		assert.Equal(t, "production-europe-west2", config.Credentials[2].AdditionalProperties["cluster"])
		assert.Equal(t, "estafette", config.Credentials[2].AdditionalProperties["defaultNamespace"])
		assert.Equal(t, "{}", config.Credentials[2].AdditionalProperties["serviceAccountKeyfile"])
	})

	t.Run("ReturnsTrustedImages", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-api-test.yaml")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		// act
		err = yaml.Unmarshal(bytes, &config)

		if !assert.Nil(t, err) {
			return
		}
		assert.Equal(t, 8, len(config.TrustedImages))
		assert.Equal(t, "extensions/docker", config.TrustedImages[0].ImagePath)
		assert.True(t, config.TrustedImages[0].RunDocker)
		assert.Equal(t, 1, len(config.TrustedImages[0].InjectedCredentialTypes))
		assert.Equal(t, "container-registry", config.TrustedImages[0].InjectedCredentialTypes[0])
		assert.Equal(t, "extensions/gke", config.TrustedImages[1].ImagePath)
		assert.False(t, config.TrustedImages[1].RunDocker)
		assert.Equal(t, 1, len(config.TrustedImages[1].InjectedCredentialTypes))
		assert.Equal(t, "kubernetes-engine", config.TrustedImages[1].InjectedCredentialTypes[0])
		assert.Equal(t, "extensions/bitbucket-status", config.TrustedImages[2].ImagePath)
		assert.False(t, config.TrustedImages[2].RunDocker)
		assert.Equal(t, 1, len(config.TrustedImages[2].InjectedCredentialTypes))
		assert.Equal(t, "bitbucket-api-token", config.TrustedImages[2].InjectedCredentialTypes[0])
		assert.Equal(t, "extensions/github-status", config.TrustedImages[3].ImagePath)
		assert.False(t, config.TrustedImages[3].RunDocker)
		assert.Equal(t, 1, len(config.TrustedImages[3].InjectedCredentialTypes))
		assert.Equal(t, "github-api-token", config.TrustedImages[3].InjectedCredentialTypes[0])
		assert.Equal(t, "extensions/slack-build-status", config.TrustedImages[4].ImagePath)
		assert.False(t, config.TrustedImages[4].RunDocker)
		assert.Equal(t, 1, len(config.TrustedImages[4].InjectedCredentialTypes))
		assert.Equal(t, "slack-webhook", config.TrustedImages[4].InjectedCredentialTypes[0])
	})

	t.Run("ReturnsManifestPreferences", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-api-test.yaml")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		// act
		err = yaml.Unmarshal(bytes, &config)

		if !assert.Nil(t, err) {
			return
		}

		if !assert.NotNil(t, config.ManifestPreferences) {
			return
		}

		assert.Equal(t, 1, len(config.ManifestPreferences.LabelRegexes))
		assert.Equal(t, "api|web|library|container", config.ManifestPreferences.LabelRegexes["type"])
		assert.Equal(t, 2, len(config.ManifestPreferences.BuilderOperatingSystems))
		assert.Equal(t, manifest.OperatingSystemLinux, config.ManifestPreferences.BuilderOperatingSystems[0])
		assert.Equal(t, manifest.OperatingSystemWindows, config.ManifestPreferences.BuilderOperatingSystems[1])
		assert.Equal(t, 2, len(config.ManifestPreferences.BuilderTracksPerOperatingSystem))
		assert.Equal(t, 3, len(config.ManifestPreferences.BuilderTracksPerOperatingSystem[manifest.OperatingSystemLinux]))
		assert.Equal(t, 3, len(config.ManifestPreferences.BuilderTracksPerOperatingSystem[manifest.OperatingSystemWindows]))
	})
}

func TestUnmarshalBuilderConfigFromJson(t *testing.T) {

	t.Run("ReturnsConfigWithoutErrors", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-builder-test.json")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		// act
		err = json.Unmarshal(bytes, &config)

		assert.Nil(t, err)
	})

	t.Run("ReturnsCredentialsWithType", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-builder-test.json")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		// act
		err = json.Unmarshal(bytes, &config)

		if !assert.Nil(t, err) {
			return
		}
		assert.Equal(t, 4, len(config.Credentials))
		assert.Equal(t, "container-registry", config.Credentials[0].Type)
		assert.Equal(t, "container-registry", config.Credentials[1].Type)
		assert.Equal(t, "kubernetes-engine", config.Credentials[2].Type)
		assert.Equal(t, "kubernetes-engine", config.Credentials[3].Type)
	})

	t.Run("ReturnsCredentialsWithAdditionalProperties", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-builder-test.json")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		// act
		err = json.Unmarshal(bytes, &config)

		if !assert.Nil(t, err) {
			return
		}
		assert.Equal(t, "extensions", config.Credentials[0].AdditionalProperties["repository"])
		assert.Equal(t, "username", config.Credentials[0].AdditionalProperties["username"])
		assert.Equal(t, "secret", config.Credentials[0].AdditionalProperties["password"])

		assert.Equal(t, "estafette-production", config.Credentials[2].AdditionalProperties["project"])
		assert.Equal(t, "europe-west2", config.Credentials[2].AdditionalProperties["region"])
		assert.Equal(t, "production-europe-west2", config.Credentials[2].AdditionalProperties["cluster"])
		assert.Equal(t, "estafette", config.Credentials[2].AdditionalProperties["defaultNamespace"])
		assert.Equal(t, "{}", config.Credentials[2].AdditionalProperties["serviceAccountKeyfile"])
	})

	t.Run("ReturnsTrustedImages", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-builder-test.json")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		// act
		err = json.Unmarshal(bytes, &config)

		if !assert.Nil(t, err) {
			return
		}
		assert.Equal(t, 5, len(config.TrustedImages))
		assert.Equal(t, "extensions/docker", config.TrustedImages[0].ImagePath)
		assert.True(t, config.TrustedImages[0].RunDocker)
		assert.Equal(t, 1, len(config.TrustedImages[0].InjectedCredentialTypes))
		assert.Equal(t, "container-registry", config.TrustedImages[0].InjectedCredentialTypes[0])
		assert.Equal(t, "extensions/gke", config.TrustedImages[1].ImagePath)
		assert.False(t, config.TrustedImages[1].RunDocker)
		assert.Equal(t, 1, len(config.TrustedImages[1].InjectedCredentialTypes))
		assert.Equal(t, "kubernetes-engine", config.TrustedImages[1].InjectedCredentialTypes[0])
		assert.Equal(t, "extensions/bitbucket-status", config.TrustedImages[2].ImagePath)
		assert.False(t, config.TrustedImages[2].RunDocker)
		assert.Equal(t, 1, len(config.TrustedImages[2].InjectedCredentialTypes))
		assert.Equal(t, "bitbucket-api-token", config.TrustedImages[2].InjectedCredentialTypes[0])
		assert.Equal(t, "extensions/github-status", config.TrustedImages[3].ImagePath)
		assert.False(t, config.TrustedImages[3].RunDocker)
		assert.Equal(t, 1, len(config.TrustedImages[3].InjectedCredentialTypes))
		assert.Equal(t, "github-api-token", config.TrustedImages[3].InjectedCredentialTypes[0])
		assert.Equal(t, "extensions/slack-build-status", config.TrustedImages[4].ImagePath)
		assert.False(t, config.TrustedImages[4].RunDocker)
		assert.Equal(t, 1, len(config.TrustedImages[4].InjectedCredentialTypes))
		assert.Equal(t, "slack-webhook", config.TrustedImages[4].InjectedCredentialTypes[0])
	})

	t.Run("ReturnsJobType", func(t *testing.T) {

		bytes, _ := ioutil.ReadFile("config-builder-in-builder-test.json")
		var config BuilderConfig

		// act
		_ = json.Unmarshal(bytes, &config)

		assert.Equal(t, JobTypeBuild, config.JobType)
	})

	t.Run("ReturnsTrack", func(t *testing.T) {

		bytes, _ := ioutil.ReadFile("config-builder-in-builder-test.json")
		var config BuilderConfig

		// act
		_ = json.Unmarshal(bytes, &config)

		assert.Equal(t, "dev", *config.Track)
	})

	t.Run("ReturnsGitConfig", func(t *testing.T) {

		bytes, _ := ioutil.ReadFile("config-builder-in-builder-test.json")
		var config BuilderConfig

		// act
		_ = json.Unmarshal(bytes, &config)

		assert.Equal(t, "github.com", config.Git.RepoSource)
		assert.Equal(t, "estafette", config.Git.RepoOwner)
		assert.Equal(t, "estafette-ci-contracts", config.Git.RepoName)
		assert.Equal(t, "master", config.Git.RepoBranch)
		assert.Equal(t, "3adf11c158811dbf0b94ca5bdbbdae79fffe7852", config.Git.RepoRevision)
	})

	t.Run("ReturnsVersionConfig", func(t *testing.T) {

		bytes, _ := ioutil.ReadFile("config-builder-in-builder-test.json")
		var config BuilderConfig

		// act
		_ = json.Unmarshal(bytes, &config)

		assert.Equal(t, "0.1.67-rc.1", config.Version.Version)
		assert.Equal(t, 0, *config.Version.Major)
		assert.Equal(t, 1, *config.Version.Minor)
		assert.Equal(t, "67-rc.1", *config.Version.Patch)
		assert.Equal(t, 67, *config.Version.AutoIncrement)
	})

	t.Run("ReturnsManifestPreferences", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-builder-test.json")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		// act
		err = json.Unmarshal(bytes, &config)

		if !assert.Nil(t, err) {
			return
		}

		if !assert.NotNil(t, config.ManifestPreferences) {
			return
		}

		assert.Equal(t, 1, len(config.ManifestPreferences.LabelRegexes))
		assert.Equal(t, "api|web|library|container", config.ManifestPreferences.LabelRegexes["type"])
		assert.Equal(t, 2, len(config.ManifestPreferences.BuilderOperatingSystems))
		assert.Equal(t, manifest.OperatingSystemLinux, config.ManifestPreferences.BuilderOperatingSystems[0])
		assert.Equal(t, manifest.OperatingSystemWindows, config.ManifestPreferences.BuilderOperatingSystems[1])
		assert.Equal(t, 2, len(config.ManifestPreferences.BuilderTracksPerOperatingSystem))
		assert.Equal(t, 3, len(config.ManifestPreferences.BuilderTracksPerOperatingSystem[manifest.OperatingSystemLinux]))
		assert.Equal(t, 3, len(config.ManifestPreferences.BuilderTracksPerOperatingSystem[manifest.OperatingSystemWindows]))
	})
}

func TestMarshalBuilderConfigToJson(t *testing.T) {

	t.Run("ReturnsJsonForOriginalYamlConfig", func(t *testing.T) {

		bytes, err := ioutil.ReadFile("config-builder-in-api-test.yaml")
		if !assert.Nil(t, err) {
			return
		}
		var config BuilderConfig

		err = yaml.Unmarshal(bytes, &config)
		if !assert.Nil(t, err) {
			return
		}

		// act
		jsonBytes, err := json.Marshal(config)
		if !assert.Nil(t, err) {
			return
		}

		assert.Equal(t, "{\"manifestPreferences\":{\"labelRegexes\":{\"type\":\"api|web|library|container\"},\"builderOperatingSystems\":[\"linux\",\"windows\"],\"builderTracksPerOperatingSystem\":{\"linux\":[\"stable\",\"beta\",\"dev\"],\"windows\":[\"windowsservercore-1809\",\"windowsservercore-1909\",\"windowsservercore-ltsc2019\"]}},\"credentials\":[{\"name\":\"container-registry-extensions\",\"type\":\"container-registry\",\"additionalProperties\":{\"password\":\"secret\",\"repository\":\"extensions\",\"username\":\"username\"}},{\"name\":\"container-registry-estafette\",\"type\":\"container-registry\",\"additionalProperties\":{\"password\":\"secret\",\"repository\":\"estafette\",\"username\":\"username\"}},{\"name\":\"gke-estafette-production\",\"type\":\"kubernetes-engine\",\"additionalProperties\":{\"cluster\":\"production-europe-west2\",\"defaultNamespace\":\"estafette\",\"project\":\"estafette-production\",\"region\":\"europe-west2\",\"serviceAccountKeyfile\":\"{}\"}},{\"name\":\"gke-estafette-development\",\"type\":\"kubernetes-engine\",\"additionalProperties\":{\"cluster\":\"development-europe-west2\",\"defaultNamespace\":\"estafette\",\"project\":\"estafette-development\",\"region\":\"europe-west2\",\"serviceAccountKeyfile\":\"{}\"}},{\"name\":\"bitbucket-api-token\",\"type\":\"bitbucket-api-token\",\"additionalProperties\":{\"token\":\"sometoken\"}},{\"name\":\"github-api-token\",\"type\":\"github-api-token\",\"additionalProperties\":{\"token\":\"sometoken\"}},{\"name\":\"slack-webhook\",\"type\":\"slack-webhook\",\"additionalProperties\":{\"webhook\":\"somewebhookurl\"}}],\"trustedImages\":[{\"path\":\"extensions/docker\",\"runPrivileged\":false,\"runDocker\":true,\"allowCommands\":false,\"allowNotifications\":false,\"injectedCredentialTypes\":[\"container-registry\"]},{\"path\":\"extensions/gke\",\"runPrivileged\":false,\"runDocker\":false,\"allowCommands\":false,\"allowNotifications\":false,\"injectedCredentialTypes\":[\"kubernetes-engine\"]},{\"path\":\"extensions/bitbucket-status\",\"runPrivileged\":false,\"runDocker\":false,\"allowCommands\":false,\"allowNotifications\":false,\"injectedCredentialTypes\":[\"bitbucket-api-token\"]},{\"path\":\"extensions/github-status\",\"runPrivileged\":false,\"runDocker\":false,\"allowCommands\":false,\"allowNotifications\":false,\"injectedCredentialTypes\":[\"github-api-token\"]},{\"path\":\"extensions/slack-build-status\",\"runPrivileged\":false,\"runDocker\":false,\"allowCommands\":false,\"allowNotifications\":false,\"injectedCredentialTypes\":[\"slack-webhook\"]},{\"path\":\"docker\",\"runPrivileged\":false,\"runDocker\":true,\"allowCommands\":false,\"allowNotifications\":false},{\"path\":\"multiple-git-sources-test\",\"runPrivileged\":false,\"runDocker\":false,\"allowCommands\":true,\"allowNotifications\":false,\"injectedCredentialTypes\":[\"bitbucket-api-token\",\"github-api-token\"]},{\"path\":\"estafette/estafette-ci-builder\",\"runPrivileged\":true,\"runDocker\":false,\"allowCommands\":false,\"allowNotifications\":false}]}", string(jsonBytes))
	})
}

func TestGetTrustedImage(t *testing.T) {

	t.Run("ReturnsTrustedImageForContainerImageWithTag", func(t *testing.T) {

		containerImage := "extensions/gke:stable"
		bytes, _ := ioutil.ReadFile("config-builder-in-api-test.yaml")
		var config BuilderConfig
		yaml.Unmarshal(bytes, &config)

		// act
		trustedImage := config.GetTrustedImage(containerImage)

		if assert.NotNil(t, trustedImage) {
			assert.Equal(t, "extensions/gke", trustedImage.ImagePath)
			assert.Equal(t, false, trustedImage.RunDocker)
		}
	})

	t.Run("ReturnsNilForUntrustedContainerImage", func(t *testing.T) {

		containerImage := "golang:1.11.1-alpine"
		bytes, _ := ioutil.ReadFile("config-builder-in-api-test.yaml")
		var config BuilderConfig
		yaml.Unmarshal(bytes, &config)

		// act
		trustedImage := config.GetTrustedImage(containerImage)

		assert.Nil(t, trustedImage)
	})
}

func TestGetCredentialsByType(t *testing.T) {

	t.Run("ReturnsListOfCredentialsMatchingType", func(t *testing.T) {

		bytes, _ := ioutil.ReadFile("config-builder-in-api-test.yaml")
		var config BuilderConfig
		yaml.Unmarshal(bytes, &config)

		// act
		credentials := config.GetCredentialsByType("container-registry")

		if assert.Equal(t, 2, len(credentials)) {
			assert.Equal(t, "container-registry-extensions", credentials[0].Name)
			assert.Equal(t, "container-registry", credentials[0].Type)
			assert.Equal(t, "container-registry-estafette", credentials[1].Name)
			assert.Equal(t, "container-registry", credentials[1].Type)
		}
	})

	t.Run("ReturnsEmptyListIfNoCredentialsMatchType", func(t *testing.T) {

		bytes, _ := ioutil.ReadFile("config-builder-in-api-test.yaml")
		var config BuilderConfig
		yaml.Unmarshal(bytes, &config)

		// act
		credentials := config.GetCredentialsByType("aws-token")

		assert.Equal(t, 0, len(credentials))
	})
}

func TestGetCredentialsForTrustedImage(t *testing.T) {

	t.Run("ReturnsListOfCredentialsMatchingTypesOfTrustedImage", func(t *testing.T) {

		bytes, _ := ioutil.ReadFile("config-builder-in-api-test.yaml")
		var config BuilderConfig
		yaml.Unmarshal(bytes, &config)
		trustedImage := config.GetTrustedImage("extensions/docker")

		// act
		credentialMap := config.GetCredentialsForTrustedImage(*trustedImage)

		if assert.Equal(t, 1, len(credentialMap)) {
			if assert.Equal(t, 2, len(credentialMap["container-registry"])) {
				assert.Equal(t, "container-registry-extensions", credentialMap["container-registry"][0].Name)
				assert.Equal(t, "container-registry", credentialMap["container-registry"][0].Type)
				assert.Equal(t, "container-registry-estafette", credentialMap["container-registry"][1].Name)
				assert.Equal(t, "container-registry", credentialMap["container-registry"][1].Type)
			}
		}
	})

	t.Run("ReturnsListOfCredentialsMatchingTypesOfTrustedImageWithMultipleAssociatedCredentialTypes", func(t *testing.T) {

		bytes, _ := ioutil.ReadFile("config-builder-in-api-test.yaml")
		var config BuilderConfig
		yaml.Unmarshal(bytes, &config)
		trustedImage := config.GetTrustedImage("multiple-git-sources-test")

		// act
		credentialMap := config.GetCredentialsForTrustedImage(*trustedImage)

		if assert.Equal(t, 2, len(credentialMap)) {
			if assert.Equal(t, 1, len(credentialMap["bitbucket-api-token"])) {
				assert.Equal(t, "bitbucket-api-token", credentialMap["bitbucket-api-token"][0].Name)
				assert.Equal(t, "bitbucket-api-token", credentialMap["bitbucket-api-token"][0].Type)
			}
			if assert.Equal(t, 1, len(credentialMap["github-api-token"])) {
				assert.Equal(t, "github-api-token", credentialMap["github-api-token"][0].Name)
				assert.Equal(t, "github-api-token", credentialMap["github-api-token"][0].Type)
			}
		}
	})

	t.Run("ReturnsListOfCredentialsMatchingTypesOfTrustedImageWithMultipleAssociatedCredentialTypesAndNonExistingTypes", func(t *testing.T) {

		bytes, _ := ioutil.ReadFile("config-builder-in-api-test.yaml")
		var config BuilderConfig
		yaml.Unmarshal(bytes, &config)
		trustedImage := &TrustedImageConfig{
			InjectedCredentialTypes: []string{
				"bitbucket-api-token",
				"github-api-token",
				"gitlab-api-token",
			},
		}

		// act
		credentialMap := config.GetCredentialsForTrustedImage(*trustedImage)

		if assert.Equal(t, 2, len(credentialMap)) {
			if assert.Equal(t, 1, len(credentialMap["bitbucket-api-token"])) {
				assert.Equal(t, "bitbucket-api-token", credentialMap["bitbucket-api-token"][0].Name)
				assert.Equal(t, "bitbucket-api-token", credentialMap["bitbucket-api-token"][0].Type)
			}
			if assert.Equal(t, 1, len(credentialMap["github-api-token"])) {
				assert.Equal(t, "github-api-token", credentialMap["github-api-token"][0].Name)
				assert.Equal(t, "github-api-token", credentialMap["github-api-token"][0].Type)
			}
		}
	})

	t.Run("ReturnsEmptyListIfNoCredentialsMatchTypesOfTrustedImage", func(t *testing.T) {

		bytes, _ := ioutil.ReadFile("config-builder-in-api-test.yaml")
		var config BuilderConfig
		yaml.Unmarshal(bytes, &config)
		trustedImage := config.GetTrustedImage("docker")

		// act
		credentialMap := config.GetCredentialsForTrustedImage(*trustedImage)

		assert.Equal(t, 0, len(credentialMap))
	})
}

func TestFilterTrustedImages(t *testing.T) {

	t.Run("ReturnsEmptyListIfStagesIsEmpty", func(t *testing.T) {

		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath: "extensions/gke",
			},
			&TrustedImageConfig{
				ImagePath: "extensions/docker",
			},
		}
		stages := []*manifest.EstafetteStage{}

		// act
		filteredTrustedImages := FilterTrustedImages(trustedImages, stages, "github.com/estafette/estafette-ci-contracts")

		assert.Equal(t, 0, len(filteredTrustedImages))
	})

	t.Run("ReturnsListWithTrustedImagesUsedInStages", func(t *testing.T) {

		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath: "extensions/gke",
			},
			&TrustedImageConfig{
				ImagePath: "extensions/docker",
			},
		}
		stages := []*manifest.EstafetteStage{
			&manifest.EstafetteStage{
				ContainerImage: "extensions/gke:stable",
			},
		}

		// act
		filteredTrustedImages := FilterTrustedImages(trustedImages, stages, "github.com/estafette/estafette-ci-contracts")

		if assert.Equal(t, 1, len(filteredTrustedImages)) {
			assert.Equal(t, "extensions/gke", filteredTrustedImages[0].ImagePath)
		}
	})

	t.Run("ReturnsListWithTrustedImagesUsedInNestedStages", func(t *testing.T) {

		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath: "extensions/gke",
			},
			&TrustedImageConfig{
				ImagePath: "extensions/docker",
			},
		}
		stages := []*manifest.EstafetteStage{
			&manifest.EstafetteStage{
				ParallelStages: []*manifest.EstafetteStage{
					&manifest.EstafetteStage{
						ContainerImage: "extensions/gke:stable",
					},
				},
			},
		}

		// act
		filteredTrustedImages := FilterTrustedImages(trustedImages, stages, "github.com/estafette/estafette-ci-contracts")

		if assert.Equal(t, 1, len(filteredTrustedImages)) {
			assert.Equal(t, "extensions/gke", filteredTrustedImages[0].ImagePath)
		}
	})

	t.Run("ReturnsListWithTrustedImagesUsedInServices", func(t *testing.T) {

		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath: "extensions/gke",
			},
			&TrustedImageConfig{
				ImagePath: "extensions/docker",
			},
			&TrustedImageConfig{
				ImagePath: "bsycorp/kind",
			},
		}
		stages := []*manifest.EstafetteStage{
			&manifest.EstafetteStage{
				Services: []*manifest.EstafetteService{
					&manifest.EstafetteService{
						ContainerImage: "bsycorp/kind:latest-1.15",
					},
				},
			},
		}

		// act
		filteredTrustedImages := FilterTrustedImages(trustedImages, stages, "github.com/estafette/estafette-ci-contracts")

		if assert.Equal(t, 1, len(filteredTrustedImages)) {
			assert.Equal(t, "bsycorp/kind", filteredTrustedImages[0].ImagePath)
		}
	})

	t.Run("ReturnsListWithTrustedImagesUsedInStagesDeduplicated", func(t *testing.T) {

		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath: "extensions/gke",
			},
			&TrustedImageConfig{
				ImagePath: "extensions/docker",
			},
		}
		stages := []*manifest.EstafetteStage{
			&manifest.EstafetteStage{
				ContainerImage: "extensions/gke:stable",
			},
			&manifest.EstafetteStage{
				ContainerImage: "extensions/gke:stable",
			},
		}

		// act
		filteredTrustedImages := FilterTrustedImages(trustedImages, stages, "github.com/estafette/estafette-ci-contracts")

		if assert.Equal(t, 1, len(filteredTrustedImages)) {
			assert.Equal(t, "extensions/gke", filteredTrustedImages[0].ImagePath)
		}
	})

	t.Run("ReturnsListWithTrustedImagesUsedInStagesAndNestedStagesDeduplicated", func(t *testing.T) {

		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath: "extensions/gke",
			},
			&TrustedImageConfig{
				ImagePath: "extensions/docker",
			},
		}
		stages := []*manifest.EstafetteStage{
			&manifest.EstafetteStage{
				ContainerImage: "extensions/gke:stable",
			},
			&manifest.EstafetteStage{
				ContainerImage: "extensions/gke:stable",
			},
			&manifest.EstafetteStage{
				ParallelStages: []*manifest.EstafetteStage{
					&manifest.EstafetteStage{
						ContainerImage: "extensions/gke:stable",
					},
				},
			},
		}

		// act
		filteredTrustedImages := FilterTrustedImages(trustedImages, stages, "github.com/estafette/estafette-ci-contracts")

		if assert.Equal(t, 1, len(filteredTrustedImages)) {
			assert.Equal(t, "extensions/gke", filteredTrustedImages[0].ImagePath)
		}
	})

	t.Run("ReturnsListWithTrustedImagesAllowedForThisPipeline", func(t *testing.T) {

		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath:        "extensions/gke",
				AllowedPipelines: "github.com/estafette/estafette-ci-contracts",
			},
			&TrustedImageConfig{
				ImagePath:        "extensions/docker",
				AllowedPipelines: "github.com/estafette/estafette-ci-api",
			},
		}
		stages := []*manifest.EstafetteStage{
			&manifest.EstafetteStage{
				ContainerImage: "extensions/gke:stable",
			},
			&manifest.EstafetteStage{
				ContainerImage: "extensions/docker:stable",
			},
		}

		// act
		filteredTrustedImages := FilterTrustedImages(trustedImages, stages, "github.com/estafette/estafette-ci-contracts")

		if assert.Equal(t, 1, len(filteredTrustedImages)) {
			assert.Equal(t, "extensions/gke", filteredTrustedImages[0].ImagePath)
		}
	})
}

func TestFilterCredentials(t *testing.T) {

	t.Run("ReturnsEmptyListIfTrustedImagesIsEmpty", func(t *testing.T) {

		credentials := []*CredentialConfig{
			&CredentialConfig{
				Name: "gke-a",
				Type: "kubernetes-engine",
			},
			&CredentialConfig{
				Name: "gke-b",
				Type: "kubernetes-engine",
			},
		}
		trustedImages := []*TrustedImageConfig{}

		// act
		filteredCredentials := FilterCredentials(credentials, trustedImages, "github.com/estafette/estafette-ci-contracts", "master")

		assert.Equal(t, 0, len(filteredCredentials))
	})

	t.Run("ReturnsEmptyListIfTrustedImagesSpecifyNoCredentials", func(t *testing.T) {

		credentials := []*CredentialConfig{
			&CredentialConfig{
				Name: "gke-a",
				Type: "kubernetes-engine",
			},
			&CredentialConfig{
				Name: "gke-b",
				Type: "kubernetes-engine",
			},
		}
		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath:               "extensions/gke",
				InjectedCredentialTypes: []string{},
			},
			&TrustedImageConfig{
				ImagePath:               "extensions/docker",
				InjectedCredentialTypes: []string{},
			},
		}

		// act
		filteredCredentials := FilterCredentials(credentials, trustedImages, "github.com/estafette/estafette-ci-contracts", "master")

		assert.Equal(t, 0, len(filteredCredentials))
	})

	t.Run("ReturnsListOfCredentialsSpecifiedForTrustedImages", func(t *testing.T) {

		credentials := []*CredentialConfig{
			&CredentialConfig{
				Name: "gke-a",
				Type: "kubernetes-engine",
			},
			&CredentialConfig{
				Name: "gke-b",
				Type: "kubernetes-engine",
			},
			&CredentialConfig{
				Name: "docker-hub",
				Type: "docker-registry",
			},
			&CredentialConfig{
				Name: "gcr-io",
				Type: "docker-registry",
			},
		}
		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath: "extensions/gke",
				InjectedCredentialTypes: []string{
					"kubernetes-engine",
				},
			},
			&TrustedImageConfig{
				ImagePath:               "extensions/docker",
				InjectedCredentialTypes: []string{},
			},
		}

		// act
		filteredCredentials := FilterCredentials(credentials, trustedImages, "github.com/estafette/estafette-ci-contracts", "master")

		assert.Equal(t, 2, len(filteredCredentials))
		assert.Equal(t, "gke-a", filteredCredentials[0].Name)
		assert.Equal(t, "kubernetes-engine", filteredCredentials[0].Type)
		assert.Equal(t, "gke-b", filteredCredentials[1].Name)
		assert.Equal(t, "kubernetes-engine", filteredCredentials[1].Type)
	})

	t.Run("ReturnsListOfCredentialsSpecifiedForTrustedImagesDeduplicated", func(t *testing.T) {

		credentials := []*CredentialConfig{
			&CredentialConfig{
				Name: "gke-a",
				Type: "kubernetes-engine",
			},
			&CredentialConfig{
				Name: "gke-b",
				Type: "kubernetes-engine",
			},
			&CredentialConfig{
				Name: "docker-hub",
				Type: "docker-registry",
			},
			&CredentialConfig{
				Name: "gcr-io",
				Type: "docker-registry",
			},
		}
		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath: "extensions/gke",
				InjectedCredentialTypes: []string{
					"kubernetes-engine",
				},
			},
			&TrustedImageConfig{
				ImagePath: "extensions/docker",
				InjectedCredentialTypes: []string{
					"kubernetes-engine",
				},
			},
		}

		// act
		filteredCredentials := FilterCredentials(credentials, trustedImages, "github.com/estafette/estafette-ci-contracts", "master")

		assert.Equal(t, 2, len(filteredCredentials))
		assert.Equal(t, "gke-a", filteredCredentials[0].Name)
		assert.Equal(t, "kubernetes-engine", filteredCredentials[0].Type)
		assert.Equal(t, "gke-b", filteredCredentials[1].Name)
		assert.Equal(t, "kubernetes-engine", filteredCredentials[1].Type)
	})

	t.Run("ReturnsListOfCredentialsSpecifiedForTrustedImagesAllowedForTrustedImages", func(t *testing.T) {

		credentials := []*CredentialConfig{
			&CredentialConfig{
				Name:                 "gke-a",
				Type:                 "kubernetes-engine",
				AllowedTrustedImages: "extensions/gke",
			},
			&CredentialConfig{
				Name:                 "gke-b",
				Type:                 "kubernetes-engine",
				AllowedTrustedImages: "extensions/port-forward",
			},
			&CredentialConfig{
				Name: "docker-hub",
				Type: "docker-registry",
			},
			&CredentialConfig{
				Name: "gcr-io",
				Type: "docker-registry",
			},
		}
		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath: "extensions/gke",
				InjectedCredentialTypes: []string{
					"kubernetes-engine",
				},
			},
			&TrustedImageConfig{
				ImagePath:               "extensions/docker",
				InjectedCredentialTypes: []string{},
			},
		}

		// act
		filteredCredentials := FilterCredentials(credentials, trustedImages, "github.com/estafette/estafette-ci-contracts", "master")

		assert.Equal(t, 1, len(filteredCredentials))
		assert.Equal(t, "gke-a", filteredCredentials[0].Name)
		assert.Equal(t, "kubernetes-engine", filteredCredentials[0].Type)
	})

	t.Run("ReturnsListOfCredentialsSpecifiedForTrustedImagesAllowedForPipeline", func(t *testing.T) {

		credentials := []*CredentialConfig{
			&CredentialConfig{
				Name:             "gke-a",
				Type:             "kubernetes-engine",
				AllowedPipelines: "github.com/estafette/estafette-ci-api",
			},
			&CredentialConfig{
				Name:             "gke-b",
				Type:             "kubernetes-engine",
				AllowedPipelines: "github.com/estafette/estafette-ci-contracts",
			},
			&CredentialConfig{
				Name: "docker-hub",
				Type: "docker-registry",
			},
			&CredentialConfig{
				Name: "gcr-io",
				Type: "docker-registry",
			},
		}
		trustedImages := []*TrustedImageConfig{
			&TrustedImageConfig{
				ImagePath: "extensions/gke",
				InjectedCredentialTypes: []string{
					"kubernetes-engine",
				},
			},
			&TrustedImageConfig{
				ImagePath:               "extensions/docker",
				InjectedCredentialTypes: []string{},
			},
		}

		// act
		filteredCredentials := FilterCredentials(credentials, trustedImages, "github.com/estafette/estafette-ci-contracts", "master")

		assert.Equal(t, 1, len(filteredCredentials))
		assert.Equal(t, "gke-b", filteredCredentials[0].Name)
		assert.Equal(t, "kubernetes-engine", filteredCredentials[0].Type)
	})
}

func TestValidate(t *testing.T) {
	t.Run("ReturnsNoErrorWhenBuildIsSetForJobTypeBuild", func(t *testing.T) {

		config := getBuilderConfig()
		config.JobType = JobTypeBuild
		config.Build = &Build{}

		// act
		err := config.Validate()

		assert.Nil(t, err)
	})

	t.Run("ReturnsErrorWhenBuildIsNotSetForJobTypeBuild", func(t *testing.T) {

		config := getBuilderConfig()
		config.JobType = JobTypeBuild
		config.Build = nil

		// act
		err := config.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "build needs to be set for jobType build", err.Error())
	})

	t.Run("ReturnsNoErrorWhenReleaseIsSetForJobTypeRelease", func(t *testing.T) {

		config := getBuilderConfig()
		config.JobType = JobTypeRelease
		config.Release = &Release{}

		// act
		err := config.Validate()

		assert.Nil(t, err)
	})

	t.Run("ReturnsErrorWhenReleaseIsNotSetForJobTypeRelease", func(t *testing.T) {

		config := getBuilderConfig()
		config.JobType = JobTypeRelease
		config.Release = nil

		// act
		err := config.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "release needs to be set for jobType release", err.Error())
	})

	t.Run("ReturnsNoErrorWhenBotIsSetForJobTypeBot", func(t *testing.T) {

		config := getBuilderConfig()
		config.JobType = JobTypeBot
		config.Bot = &Bot{}

		// act
		err := config.Validate()

		assert.Nil(t, err)
	})

	t.Run("ReturnsErrorWhenBotIsNotSetForJobTypeBot", func(t *testing.T) {

		config := getBuilderConfig()
		config.JobType = JobTypeBot
		config.Bot = nil

		// act
		err := config.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "bot needs to be set for jobType bot", err.Error())
	})

	t.Run("ReturnsNoErrorWhenGitIsSet", func(t *testing.T) {

		config := getBuilderConfig()
		config.Git = &GitConfig{}

		// act
		err := config.Validate()

		assert.Nil(t, err)
	})

	t.Run("ReturnsErrorWhenGitIsNotSet", func(t *testing.T) {

		config := getBuilderConfig()
		config.Git = nil

		// act
		err := config.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "git needs to be set", err.Error())
	})

	t.Run("ReturnsNoErrorWhenVersionIsSet", func(t *testing.T) {

		config := getBuilderConfig()
		config.Version = &VersionConfig{}

		// act
		err := config.Validate()

		assert.Nil(t, err)
	})

	t.Run("ReturnsErrorWhenVersionIsNotSet", func(t *testing.T) {

		config := getBuilderConfig()
		config.Version = nil

		// act
		err := config.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "version needs to be set", err.Error())
	})

	t.Run("ReturnsNoErrorWhenManifestIsSet", func(t *testing.T) {

		config := getBuilderConfig()
		config.Manifest = &manifest.EstafetteManifest{}

		// act
		err := config.Validate()

		assert.Nil(t, err)
	})

	t.Run("ReturnsErrorWhenManifestIsNotSet", func(t *testing.T) {

		config := getBuilderConfig()
		config.Manifest = nil

		// act
		err := config.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "manifest needs to be set", err.Error())
	})
}

func TestUnmarshalBuilderConfig(t *testing.T) {
	t.Run("UnmarshalBuilderConfig", func(t *testing.T) {

		jsonData := `
		{
			"jobType":"build",
			"track":"stable",
			"dockerConfig":{
				 "runType":"dind",
				 "mtu":1460,
				 "bip":"192.168.1.1/24",
				 "networks":[
						{
							 "name":"estafette",
							 "driver":"default",
							 "subnet":"192.168.2.1/24",
							 "gateway":"192.168.2.1",
							 "durable":false
						}
				 ],
				 "registryMirror":"https://mirror.gcr.io"
			},
			"manifest":{
				 "Builder":{
						"Track":"stable",
						"OperatingSystem":"linux",
						"StorageMedium":"",
						"BuilderType":"docker"
				 },
				 "Stages":[
						{
							 "Name":"injected-before",
							 "ContainerImage":"",
							 "Shell":"",
							 "WorkingDirectory":"",
							 "Commands":null,
							 "RunCommandsInForeground":false,
							 "When":"status == 'succeeded'",
							 "EnvVars":null,
							 "AutoInjected":true,
							 "Retries":0,
							 "ParallelStages":[
									{
										 "Name":"git-clone",
										 "ContainerImage":"extensions/git-clone:stable",
										 "Shell":"/bin/sh",
										 "WorkingDirectory":"/estafette-work",
										 "Commands":null,
										 "RunCommandsInForeground":false,
										 "When":"status == 'succeeded'",
										 "EnvVars":null,
										 "AutoInjected":true,
										 "Retries":0,
										 "ParallelStages":null,
										 "Services":null,
										 "CustomProperties":null
									},
									{
										 "Name":"set-pending-build-status",
										 "ContainerImage":"extensions/bitbucket-status:stable",
										 "Shell":"/bin/sh",
										 "WorkingDirectory":"/estafette-work",
										 "Commands":null,
										 "RunCommandsInForeground":false,
										 "When":"status == 'succeeded'",
										 "EnvVars":null,
										 "AutoInjected":true,
										 "Retries":0,
										 "ParallelStages":null,
										 "Services":null,
										 "CustomProperties":{
												"status":"pending"
										 }
									},
									{
										 "Name":"envvars",
										 "ContainerImage":"extensions/envvars:stable",
										 "Shell":"/bin/sh",
										 "WorkingDirectory":"/estafette-work",
										 "Commands":null,
										 "RunCommandsInForeground":false,
										 "When":"status == 'succeeded'",
										 "EnvVars":null,
										 "AutoInjected":true,
										 "Retries":0,
										 "ParallelStages":null,
										 "Services":null,
										 "CustomProperties":{
												
										 }
									}
							 ],
							 "Services":null,
							 "CustomProperties":null
						},
						{
							 "Name":"injected-after",
							 "ContainerImage":"",
							 "Shell":"",
							 "WorkingDirectory":"",
							 "Commands":null,
							 "RunCommandsInForeground":false,
							 "When":"status == 'succeeded' || status == 'failed'",
							 "EnvVars":null,
							 "AutoInjected":true,
							 "Retries":0,
							 "ParallelStages":[
									{
										 "Name":"set-build-status",
										 "ContainerImage":"extensions/bitbucket-status:stable",
										 "Shell":"/bin/sh",
										 "WorkingDirectory":"/estafette-work",
										 "Commands":null,
										 "RunCommandsInForeground":false,
										 "When":"status == 'succeeded' || status == 'failed'",
										 "EnvVars":null,
										 "AutoInjected":true,
										 "Retries":0,
										 "ParallelStages":null,
										 "Services":null,
										 "CustomProperties":null
									}
							 ],
							 "Services":null,
							 "CustomProperties":null
						}
				 ]
			},
			"stages":[
				{
					 "Name":"injected-before",
					 "ContainerImage":"",
					 "Shell":"",
					 "WorkingDirectory":"",
					 "Commands":null,
					 "RunCommandsInForeground":false,
					 "When":"status == 'succeeded'",
					 "EnvVars":null,
					 "AutoInjected":true,
					 "Retries":0,
					 "ParallelStages":[
							{
								 "Name":"git-clone",
								 "ContainerImage":"extensions/git-clone:stable",
								 "Shell":"/bin/sh",
								 "WorkingDirectory":"/estafette-work",
								 "Commands":null,
								 "RunCommandsInForeground":false,
								 "When":"status == 'succeeded'",
								 "EnvVars":null,
								 "AutoInjected":true,
								 "Retries":0,
								 "ParallelStages":null,
								 "Services":null,
								 "CustomProperties":null
							},
							{
								 "Name":"set-pending-build-status",
								 "ContainerImage":"extensions/bitbucket-status:stable",
								 "Shell":"/bin/sh",
								 "WorkingDirectory":"/estafette-work",
								 "Commands":null,
								 "RunCommandsInForeground":false,
								 "When":"status == 'succeeded'",
								 "EnvVars":null,
								 "AutoInjected":true,
								 "Retries":0,
								 "ParallelStages":null,
								 "Services":null,
								 "CustomProperties":{
										"status":"pending"
								 }
							},
							{
								 "Name":"envvars",
								 "ContainerImage":"extensions/envvars:stable",
								 "Shell":"/bin/sh",
								 "WorkingDirectory":"/estafette-work",
								 "Commands":null,
								 "RunCommandsInForeground":false,
								 "When":"status == 'succeeded'",
								 "EnvVars":null,
								 "AutoInjected":true,
								 "Retries":0,
								 "ParallelStages":null,
								 "Services":null,
								 "CustomProperties":{
										
								 }
							}
					 ],
					 "Services":null,
					 "CustomProperties":null
				},
				{
					 "Name":"injected-after",
					 "ContainerImage":"",
					 "Shell":"",
					 "WorkingDirectory":"",
					 "Commands":null,
					 "RunCommandsInForeground":false,
					 "When":"status == 'succeeded' || status == 'failed'",
					 "EnvVars":null,
					 "AutoInjected":true,
					 "Retries":0,
					 "ParallelStages":[
							{
								 "Name":"set-build-status",
								 "ContainerImage":"extensions/bitbucket-status:stable",
								 "Shell":"/bin/sh",
								 "WorkingDirectory":"/estafette-work",
								 "Commands":null,
								 "RunCommandsInForeground":false,
								 "When":"status == 'succeeded' || status == 'failed'",
								 "EnvVars":null,
								 "AutoInjected":true,
								 "Retries":0,
								 "ParallelStages":null,
								 "Services":null,
								 "CustomProperties":null
							}
					 ],
					 "Services":null,
					 "CustomProperties":null
				}
		 	],
			"jobName":"build-something-663703472164241425",
			"triggerEvents":[
				 {
						"fired":true,
						"git":{
							 "event":"push",
							 "repository":"bitbucket.org/org/something",
							 "branch":"master"
						}
				 }
			],
			"ciServer":{
				 "baseUrl":"https://ci.estafette.io/",
				 "builderEventsUrl":"https://ci.estafette.io/api/commands",
				 "postLogsUrl":"https://ci.estafette.io/api/pipelines/bitbucket.org/xivart/edge/builds/663703472164241425/logs",
				 "jwt":"***"
			},
			"build":{
				 "ID":"663703472164241425"
			},
			"git":{
				 "repoSource":"bitbucket.org",
				 "repoOwner":"xivart",
				 "repoName":"edge",
				 "repoBranch":"master",
				 "repoRevision":"886de85bb620a8217a884cb18a34ad15423a26e9"
			},
			"version":{
				 "version":"1.0.7456",
				 "major":1,
				 "minor":0,
				 "patch":"7456",
				 "label":"master",
				 "autoincrement":7456,
				 "currentCounter":7456,
				 "maxCounter":7456,
				 "maxCounterCurrentBranch":7456
			},
			"credentials":[
				 {
						"name":"bitbucket-api-token",
						"type":"bitbucket-api-token",
						"additionalProperties":{
							 "token":"estafette.secret(***)"
						}
				 }
			],
			"trustedImages":[
				 {
						"path":"extensions/git-clone",
						"runPrivileged":false,
						"runDocker":false,
						"allowCommands":false,
						"injectedCredentialTypes":[
							 "bitbucket-api-token",
							 "github-api-token",
							 "cloudsource-api-token"
						]
				 },
				 {
						"path":"extensions/bitbucket-status",
						"runPrivileged":false,
						"runDocker":false,
						"allowCommands":false,
						"injectedCredentialTypes":[
							 "bitbucket-api-token"
						]
				 }
			]
	 }`

		//act
		var config BuilderConfig
		err := json.Unmarshal([]byte(jsonData), &config)

		assert.Nil(t, err)
		assert.Equal(t, "1.0.7456", config.Version.Version)
		assert.Equal(t, "663703472164241425", config.Build.ID)
		assert.Equal(t, JobTypeBuild, config.JobType)
	})

}

func getBuilderConfig() BuilderConfig {
	return BuilderConfig{
		JobType:  JobTypeBot,
		Git:      &GitConfig{},
		Version:  &VersionConfig{},
		Manifest: &manifest.EstafetteManifest{},
		Build:    &Build{},
		Release:  &Release{},
		Bot:      &Bot{},
	}
}
