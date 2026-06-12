package aptible

import "runtime/debug"

func version() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == "github.com/aptible/go-deploy" {
				return dep.Version
			}
		}
		if info.Main.Version != "" {
			return info.Main.Version
		}
	}
	return "aptible/go-deploy/unknown"
}
