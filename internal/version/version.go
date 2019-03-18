package version

var (
	// Version is current version of app. It should be set
	// explicitly with ldflags while compiling.
	Version = "unset"

	// Commit is the latest commit made in the repository. It
	// should be set explicitly with ldflags while compiling.
	Commit = "unset"

	// BuildTime is the time and date when application was
	// build. It should be set explicitly with ldflags while
	// compiling.
	BuildTime = "unset"
)
