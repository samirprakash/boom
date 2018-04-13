package maven

// Flags defines the flags available for overriding semaphorge maven sub-commands
type Flags struct {
	RunIntegrationTests bool
	RunUnitTests        bool
	SkipTests           bool
	RepoID              string
}
