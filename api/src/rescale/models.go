package rescale

// CoreType represents Rescale Compute environments
type CoreType struct {
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	IsDefault       bool    `json:"isDefault"`
	MustBeRequested bool    `json:"mustBeRequested"`
	ProcessorInfo   string  `json:"processorInfo"`
	Cores           []int   `json:"cores"`
	GPUCounts       []int   `json:"gpuCounts"`
	Compute         *string `json:"compute"`
	BaseClockSpeed  *string `json:"baseClockSpeed"`
	IO              string  `json:"io"`
	Storage         int     `json:"storage"`
	Memory          int     `json:"memory"`
	DisplayOrder    int     `json:"displayOrder"`
}

// Application represents Rescale supported application
type Application struct {
	Code                string               `json:"code"`
	Versions            []ApplicationVersion `json:"versions"`
	HasRescaleLicense   bool                 `json:"hasRescaleLicense"`
	HasOnDemandLicense  bool                 `json:"hasOnDemandLicense"`
	HasShortTermLicense bool                 `json:"hasShortTermLicense"`
	LicenseSettings     []ApplicationLicense `json:"licenseSettings"`
}

// ApplicationVersion represents a version of an application
type ApplicationVersion struct {
	ID              string   `json:"id"`
	Version         string   `json:"version"`
	VersionCode     string   `json:"versionCode"`
	Type            string   `json:"type"`
	OS              []string `json:"oses"`
	CoreTypes       []string `json:"allowedCoreTypes"`
	MustBeRequested bool     `json:"mustBeRequested"`
}

// ApplicationLicense represents license settings of an application
type ApplicationLicense struct { // nolint:maligned
	Name                 string `json:"name"`
	Required             bool   `json:"required"`
	LicenseType          string `json:"licenseType"`
	Label                string `json:"label"`
	IsServer             bool   `json:"isServer"`
	CanCheckAvailability bool   `json:"canCheckAvailability"`
}

// UploadedFile represents an uploaded file
type UploadedFile struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	FilePath    string              `json:"path"`
	DownloadURL string              `json:"downloadUrl"`
	MD5         string              `json:"md5"`
	Owner       string              `json:"owner"`
	Storage     UploadedFileStorage `json:"storage"`
}

// UploadedFileStorage represents a storage which contains the uploaded file
type UploadedFileStorage struct {
	ID string `json:"id"`
}

// JobInput defines Rescale job conditions
type JobInput struct {
	Name             string            `json:"name"`
	JobAnalyses      []JobInputAnalyse `json:"jobanalyses"`
	JobVariables     []string          `json:"jobvariables"`
	IsLowPriority    bool              `json:"isLowPriority"`
	IsTemplateDryRun bool              `json:"isTemplateDryRun"`
}

// JobInputAnalyse defines Rescale job infrastructure
type JobInputAnalyse struct {
	Command    string         `json:"command"`
	InputFiles []JobInputFile `json:"inputFiles"`
	Analysis   JobAnalyse     `json:"analysis"`
	Hardware   JobHardware    `json:"hardware"`
}

// JobInputFile defines files to be used in the job
type JobInputFile struct {
	ID         string `json:"id"`
	Decompress bool   `json:"decompress"`
}

// JobAnalyse defines Rescale job software
type JobAnalyse struct {
	Code    string `json:"code"`
	Version string `json:"version"`
}

// JobHardware defines Rescale job hardware
type JobHardware struct {
	Type         string      `json:"type"`
	CoreType     JobCoreType `json:"coreType"`
	Slots        int         `json:"slots"`
	CoresPerSlot int         `json:"coresPerSlot"`
	WallTime     int         `json:"walltime"`
}

// JobCoreType defines Rescale job core type
type JobCoreType struct {
	Code string `json:"code"`
}

// JobStatus defines the status of a Rescale job
type JobStatus struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
