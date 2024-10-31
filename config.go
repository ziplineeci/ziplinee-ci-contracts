package contracts

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	manifest "github.com/estafette/estafette-ci-manifest"
)

// ContainerRepositoryCredentialConfig is used to authenticate for (private) container repositories (will be replaced by CredentialConfig eventually)
type ContainerRepositoryCredentialConfig struct {
	Repository string `yaml:"repository"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}

type JobType string

const (
	JobTypeUnknown JobType = ""
	JobTypeBuild   JobType = "build"
	JobTypeRelease JobType = "release"
	JobTypeBot     JobType = "bot"
)

// BuilderConfig parameterizes a build/release job
type BuilderConfig struct {
	JobType JobType        `yaml:"jobType,omitempty" json:"jobType,omitempty"`
	Build   *Build         `yaml:"build,omitempty" json:"build,omitempty"`
	Release *Release       `yaml:"release,omitempty" json:"release,omitempty"`
	Bot     *Bot           `yaml:"bot,omitempty" json:"bot,omitempty"`
	Git     *GitConfig     `yaml:"git,omitempty" json:"git,omitempty"`
	Version *VersionConfig `yaml:"version,omitempty" json:"version,omitempty"`

	Track               *string                                `yaml:"track,omitempty" json:"track,omitempty"`
	DockerConfig        *DockerConfig                          `yaml:"dockerConfig,omitempty" json:"dockerConfig,omitempty"`
	Manifest            *manifest.EstafetteManifest            `yaml:"manifest,omitempty" json:"manifest,omitempty"`
	ManifestPreferences *manifest.EstafetteManifestPreferences `yaml:"manifestPreferences,omitempty" json:"manifestPreferences,omitempty"`
	JobName             *string                                `yaml:"jobName,omitempty" json:"jobName,omitempty"`
	Events              []manifest.EstafetteEvent              `yaml:"triggerEvents,omitempty" json:"triggerEvents,omitempty"`
	CIServer            *CIServerConfig                        `yaml:"ciServer,omitempty" json:"ciServer,omitempty"`
	Stages              []*manifest.EstafetteStage             `yaml:"stages,omitempty" json:"stages,omitempty"`
	Credentials         []*CredentialConfig                    `yaml:"credentials,omitempty" json:"credentials,omitempty"`
	TrustedImages       []*TrustedImageConfig                  `yaml:"trustedImages,omitempty" json:"trustedImages,omitempty"`
}

func (bc *BuilderConfig) Validate() (err error) {

	if bc.Git == nil {
		return errors.New("git needs to be set")
	}
	if bc.Version == nil {
		return errors.New("version needs to be set")
	}
	if bc.Manifest == nil {
		return errors.New("manifest needs to be set")
	}

	switch bc.JobType {
	case JobTypeBuild:
		if bc.Build == nil {
			return errors.New("build needs to be set for jobType build")
		}
	case JobTypeRelease:
		if bc.Release == nil {
			return errors.New("release needs to be set for jobType release")
		}
	case JobTypeBot:
		if bc.Bot == nil {
			return errors.New("bot needs to be set for jobType bot")
		}
	}

	return nil
}

// CredentialConfig is used to store credentials for every type of authenticated service you can use from docker registries, to kubernetes engine to, github apis, bitbucket;
// in combination with trusted images access to these centrally stored credentials can be limited
type CredentialConfig struct {
	Name                 string                 `yaml:"name" json:"name"`
	Type                 string                 `yaml:"type" json:"type"`
	AllowedPipelines     string                 `yaml:"allowedPipelines,omitempty" json:"allowedPipelines,omitempty"`
	AllowedTrustedImages string                 `yaml:"allowedTrustedImages,omitempty" json:"allowedTrustedImages,omitempty"`
	AllowedBranches      string                 `yaml:"allowedBranches,omitempty" json:"allowedBranches,omitempty"`
	AdditionalProperties map[string]interface{} `yaml:",inline" json:"additionalProperties,omitempty"`
}

// UnmarshalYAML customizes unmarshalling an EstafetteStage
func (cc *CredentialConfig) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {

	var aux struct {
		Name                 string                 `yaml:"name" json:"name"`
		Type                 string                 `yaml:"type" json:"type"`
		AllowedPipelines     string                 `yaml:"allowedPipelines,omitempty" json:"allowedPipelines,omitempty"`
		AllowedTrustedImages string                 `yaml:"allowedTrustedImages,omitempty" json:"allowedTrustedImages,omitempty"`
		AllowedBranches      string                 `yaml:"allowedBranches,omitempty" json:"allowedBranches,omitempty"`
		AdditionalProperties map[string]interface{} `yaml:",inline" json:"additionalProperties,omitempty"`
	}

	// unmarshal to auxiliary type
	if err := unmarshal(&aux); err != nil {
		return err
	}

	// map auxiliary properties
	cc.Name = aux.Name
	cc.Type = aux.Type
	cc.AllowedPipelines = aux.AllowedPipelines
	cc.AllowedTrustedImages = aux.AllowedTrustedImages
	cc.AllowedBranches = aux.AllowedBranches

	// fix for map[interface{}]interface breaking json.marshal - see https://github.com/go-yaml/yaml/issues/139
	cc.AdditionalProperties = cleanUpStringMap(aux.AdditionalProperties)

	return nil
}

// TrustedImageConfig allows trusted images to run docker commands or receive specific credentials
type TrustedImageConfig struct {
	ImagePath               string   `yaml:"path" json:"path"`
	RunPrivileged           bool     `yaml:"runPrivileged" json:"runPrivileged"`
	RunDocker               bool     `yaml:"runDocker" json:"runDocker"`
	AllowCommands           bool     `yaml:"allowCommands" json:"allowCommands"`
	AllowNotifications      bool     `yaml:"allowNotifications" json:"allowNotifications"`
	InjectedCredentialTypes []string `yaml:"injectedCredentialTypes,omitempty" json:"injectedCredentialTypes,omitempty"`
	AllowedPipelines        string   `yaml:"allowedPipelines,omitempty" json:"allowedPipelines,omitempty"`
}

// GitConfig contains all information for cloning the git repository for building/releasing a specific version
type GitConfig struct {
	RepoSource   string `json:"repoSource"`
	RepoOwner    string `json:"repoOwner"`
	RepoName     string `json:"repoName"`
	RepoBranch   string `json:"repoBranch"`
	RepoRevision string `json:"repoRevision"`
}

// VersionConfig contains all information regarding the version number to build or release
type VersionConfig struct {
	Version                 string  `json:"version"`
	Major                   *int    `json:"major,omitempty"`
	Minor                   *int    `json:"minor,omitempty"`
	Patch                   *string `json:"patch,omitempty"`
	Label                   *string `json:"label,omitempty"`
	AutoIncrement           *int    `json:"autoincrement,omitempty"`
	CurrentCounter          int     `json:"currentCounter,omitempty"`
	MaxCounter              int     `json:"maxCounter,omitempty"`
	MaxCounterCurrentBranch int     `json:"maxCounterCurrentBranch,omitempty"`
}

// CIServerConfig has a number of config items related to communication or linking to the CI server
type CIServerConfig struct {
	BaseURL          string    `json:"baseUrl"`
	BuilderEventsURL string    `json:"builderEventsUrl"`
	PostLogsURL      string    `json:"postLogsUrl"`
	CancelJobURL     string    `json:"cancelJobUrl"`
	JWT              string    `json:"jwt"`
	JWTExpiry        time.Time `json:"jwtExpiry"`
}

// DockerNetworkConfig has settings for creating a user defined docker network to make service containers accessible by name from other containers
type DockerNetworkConfig struct {
	Name    string `json:"name"`
	Driver  string `json:"driver"`
	Subnet  string `json:"subnet"`
	Gateway string `json:"gateway"`
	Durable bool   `json:"durable"`
}

type DockerRunType string

const (
	// DockerRunTypeUnknown indicates the value couldn't be mapped
	DockerRunTypeUnknown DockerRunType = ""
	// DockerRunTypeDinD represents docker-inside-docker
	DockerRunTypeDinD DockerRunType = "dind"
	// DockerRunTypeDoD represents docker-outside-docker
	DockerRunTypeDoD DockerRunType = "dod"
)

// DockerConfig has configuration to configure docker in estafette-ci-builder
type DockerConfig struct {
	RunType        DockerRunType         `yaml:"runType,omitempty" json:"runType,omitempty"`
	MTU            int                   `yaml:"mtu,omitempty" json:"mtu,omitempty"`
	BIP            string                `yaml:"bip,omitempty" json:"bip,omitempty"`
	Networks       []DockerNetworkConfig `yaml:"networks,omitempty" json:"networks,omitempty"`
	RegistryMirror string                `yaml:"registryMirror,omitempty" json:"registryMirror,omitempty"`
}

// BuildParamsConfig has config specific to builds
type BuildParamsConfig struct {
	BuildID int `json:"buildID"`
}

// ReleaseParamsConfig has config specific to releases
type ReleaseParamsConfig struct {
	ReleaseName   string `json:"releaseName"`
	ReleaseID     int    `json:"releaseID"`
	ReleaseAction string `json:"releaseAction,omitempty"`
	TriggeredBy   string `json:"triggeredBy,omitempty"`
}

// BotParamsConfig has config specific to releases
type BotParamsConfig struct {
	BotName string `json:"botName"`
	BotID   int    `json:"botID"`
}

// GetCredentialsByType returns all credentials of a certain type
func (c *BuilderConfig) GetCredentialsByType(filterType string) []*CredentialConfig {

	filteredCredentials := []*CredentialConfig{}

	for _, cred := range c.Credentials {
		if cred.Type == filterType {
			filteredCredentials = append(filteredCredentials, cred)
		}
	}

	return filteredCredentials
}

// GetCredentialsByType returns all credentials of a certain type
func GetCredentialsByType(credentials []*CredentialConfig, filterType string) []*CredentialConfig {

	filteredCredentials := []*CredentialConfig{}

	for _, cred := range credentials {
		if cred.Type == filterType {
			filteredCredentials = append(filteredCredentials, cred)
		}
	}

	return filteredCredentials
}

// FilterCredentialsByTrustedImagesAllowList returns the list of credentials filtered by the AllowedTrustedImages property on the credentials
func FilterCredentialsByTrustedImagesAllowList(credentials []*CredentialConfig, trustedImage TrustedImageConfig) (filteredCredentials []*CredentialConfig) {

	filteredCredentials = make([]*CredentialConfig, 0)
	for _, c := range credentials {
		if IsAllowedTrustedImageForCredential(*c, trustedImage) {
			filteredCredentials = append(filteredCredentials, c)
		}
	}

	return
}

// IsAllowedTrustedImageForCredential returns true if AllowedTrustedImages is empty or matches the trusted image Path property
func IsAllowedTrustedImageForCredential(credential CredentialConfig, trustedImage TrustedImageConfig) bool {

	if credential.AllowedTrustedImages == "" {
		return true
	}

	pattern := fmt.Sprintf("^(%v)$", strings.TrimSpace(credential.AllowedTrustedImages))
	isMatch, _ := regexp.Match(pattern, []byte(trustedImage.ImagePath))

	return isMatch
}

// FilterCredentialsByPipelinesAllowList returns the list of credentials filtered by the AllowedPipelines property on the credentials
func FilterCredentialsByPipelinesAllowList(credentials []*CredentialConfig, fullRepositoryPath string) (filteredCredentials []*CredentialConfig) {

	filteredCredentials = make([]*CredentialConfig, 0)
	for _, c := range credentials {
		if IsAllowedPipelineForCredential(*c, fullRepositoryPath) {
			filteredCredentials = append(filteredCredentials, c)
		}
	}

	return
}

// FilterCredentialsByBranchesAllowList returns the list of credentials filtered by the AllowedBranches property on the credentials
func FilterCredentialsByBranchesAllowList(credentials []*CredentialConfig, branch string) (filteredCredentials []*CredentialConfig) {

	filteredCredentials = make([]*CredentialConfig, 0)
	for _, c := range credentials {
		if IsAllowedBranchForCredential(*c, branch) {
			filteredCredentials = append(filteredCredentials, c)
		}
	}

	return
}

// IsAllowedPipelineForCredential returns true if AllowedPipelines is empty or matches the pipelines full path
func IsAllowedPipelineForCredential(credential CredentialConfig, fullRepositoryPath string) bool {

	if credential.AllowedPipelines == "" {
		return true
	}

	pattern := fmt.Sprintf("^(%v)$", strings.TrimSpace(credential.AllowedPipelines))
	isMatch, _ := regexp.Match(pattern, []byte(fullRepositoryPath))

	return isMatch
}

// IsAllowedBranchForCredential returns true if AllowedBranches is empty or matches the build/release job branch
func IsAllowedBranchForCredential(credential CredentialConfig, branch string) bool {

	if credential.AllowedBranches == "" {
		return true
	}

	pattern := fmt.Sprintf("^(%v)$", strings.TrimSpace(credential.AllowedBranches))
	isMatch, _ := regexp.Match(pattern, []byte(branch))

	return isMatch
}

// FilterTrustedImagesByPipelinesAllowList returns the list of trusted images filtered by the AllowedTrustedPipelines property on the trusted images
func FilterTrustedImagesByPipelinesAllowList(trustedImages []*TrustedImageConfig, fullRepositoryPath string) (filteredTrustedImages []*TrustedImageConfig) {

	filteredTrustedImages = make([]*TrustedImageConfig, 0)
	for _, ti := range trustedImages {
		if IsAllowedPipelineForTrustedImage(*ti, fullRepositoryPath) {
			filteredTrustedImages = append(filteredTrustedImages, ti)
		}
	}

	return
}

// IsAllowedPipelineForTrustedImage returns true if AllowedPipelines is empty or matches the pipelines full path
func IsAllowedPipelineForTrustedImage(trustedImage TrustedImageConfig, fullRepositoryPath string) bool {

	if trustedImage.AllowedPipelines == "" {
		return true
	}

	pattern := fmt.Sprintf("^(%v)$", strings.TrimSpace(trustedImage.AllowedPipelines))
	isMatch, _ := regexp.Match(pattern, []byte(fullRepositoryPath))

	return isMatch
}

// GetCredentialsForTrustedImage returns all credentials of a certain type
func (c *BuilderConfig) GetCredentialsForTrustedImage(trustedImage TrustedImageConfig) map[string][]*CredentialConfig {
	return GetCredentialsForTrustedImage(c.Credentials, trustedImage)
}

// GetCredentialsForTrustedImage returns all credentials of a certain type
func GetCredentialsForTrustedImage(credentials []*CredentialConfig, trustedImage TrustedImageConfig) map[string][]*CredentialConfig {

	credentialMap := map[string][]*CredentialConfig{}

	for _, filterType := range trustedImage.InjectedCredentialTypes {
		credsByType := GetCredentialsByType(credentials, filterType)
		// filter by allow list
		credsByType = FilterCredentialsByTrustedImagesAllowList(credsByType, trustedImage)
		if len(credsByType) > 0 {
			credentialMap[filterType] = credsByType
		}
	}

	return credentialMap
}

// GetTrustedImage returns a trusted image if the path without tag matches any of the trustedImages
func (c *BuilderConfig) GetTrustedImage(imagePath string) *TrustedImageConfig {
	return GetTrustedImage(c.TrustedImages, imagePath)
}

// GetTrustedImage returns a trusted image if the path without tag matches any of the trustedImages
func GetTrustedImage(trustedImages []*TrustedImageConfig, imagePath string) *TrustedImageConfig {

	imagePathSlice := strings.Split(imagePath, ":")
	imagePathWithoutTag := imagePathSlice[0]

	for _, trustedImage := range trustedImages {
		if trustedImage.ImagePath == imagePathWithoutTag {
			return trustedImage
		}
	}

	return nil
}

// FilterTrustedImages returns only trusted images used in the stages
func FilterTrustedImages(trustedImages []*TrustedImageConfig, stages []*manifest.EstafetteStage, fullRepositoryPath string) []*TrustedImageConfig {

	filteredImages := []*TrustedImageConfig{}

	for _, s := range stages {
		ti := GetTrustedImage(trustedImages, s.ContainerImage)
		if ti != nil {
			alreadyAdded := false
			for _, fi := range filteredImages {
				if fi.ImagePath == ti.ImagePath {
					alreadyAdded = true
					break
				}
			}

			if !alreadyAdded {
				filteredImages = append(filteredImages, ti)
			}
		}

		if len(s.ParallelStages) > 0 {
			for _, ps := range s.ParallelStages {
				ti := GetTrustedImage(trustedImages, ps.ContainerImage)
				if ti != nil {
					alreadyAdded := false
					for _, fi := range filteredImages {
						if fi.ImagePath == ti.ImagePath {
							alreadyAdded = true
							break
						}
					}

					if !alreadyAdded {
						filteredImages = append(filteredImages, ti)
					}
				}
			}
		}

		if len(s.Services) > 0 {
			for _, svc := range s.Services {
				ti := GetTrustedImage(trustedImages, svc.ContainerImage)
				if ti != nil {
					alreadyAdded := false
					for _, fi := range filteredImages {
						if fi.ImagePath == ti.ImagePath {
							alreadyAdded = true
							break
						}
					}

					if !alreadyAdded {
						filteredImages = append(filteredImages, ti)
					}
				}
			}
		}
	}

	// filter by allow list
	filteredImages = FilterTrustedImagesByPipelinesAllowList(filteredImages, fullRepositoryPath)

	return filteredImages
}

// FilterCredentials returns only credentials used by the trusted images
func FilterCredentials(credentials []*CredentialConfig, trustedImages []*TrustedImageConfig, fullRepositoryPath, branch string) []*CredentialConfig {

	filteredCredentials := []*CredentialConfig{}

	for _, i := range trustedImages {
		credMap := GetCredentialsForTrustedImage(credentials, *i)

		// loop all items in credmap and add to filtered credentials if they haven't been already added
		for _, v := range credMap {
			// filter by allow list
			v = FilterCredentialsByPipelinesAllowList(v, fullRepositoryPath)
			v = FilterCredentialsByBranchesAllowList(v, branch)
			filteredCredentials = AddCredentialsIfNotPresent(filteredCredentials, v)
		}
	}

	return filteredCredentials
}

// AddCredentialsIfNotPresent adds new credentials to source credentials if they're not present yet
func AddCredentialsIfNotPresent(sourceCredentials []*CredentialConfig, newCredentials []*CredentialConfig) []*CredentialConfig {

	for _, c := range newCredentials {

		alreadyAdded := false
		for _, fc := range sourceCredentials {
			if fc.Name == c.Name && fc.Type == c.Type {
				alreadyAdded = true
				break
			}
		}

		if !alreadyAdded {
			sourceCredentials = append(sourceCredentials, c)
		}
	}

	return sourceCredentials
}
