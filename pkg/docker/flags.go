package docker

// Flags defines the options avaialble to override the semaphorge docker sub-commands
type Flags struct {
	ImageTag         string
	AppType          string
	UploadPath       string
	ComposeFile      string
	RepoName         string
	HealthCheckPorts string
	NetworkBridge    string
	TestCollection   string
	EnvironmentSpec  string
	CurrentImage     string
	NewImage         string
}
