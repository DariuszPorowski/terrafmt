package version

var (
	// Version can be overridden via -ldflags "-X <module>/lib/version.Version=...".
	Version   = "0.5.2"
	GitCommit string
)
