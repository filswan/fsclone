package service

import "fmt"

const (
	MajorVersion = 2
	MinorVersion = 5
	FixVersion   = 0
	CommitHash   = ""
)

func GetVersion() string {
	if CommitHash == "" {
		return fmt.Sprintf("migration-v%v.%v.%v", MajorVersion, MinorVersion, FixVersion)
	} else {
		return fmt.Sprintf("migration-v%v.%v.%v-%s", MajorVersion, MinorVersion, FixVersion, CommitHash)
	}
}
