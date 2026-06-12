package aptible

import "runtime/debug"

func version() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == "github.com/aptible/go-deploy" {
				return dep.Version
			}
		}
	}
	return "unknown"
}
