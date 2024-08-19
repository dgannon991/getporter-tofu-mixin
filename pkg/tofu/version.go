package tofu

import (
	"get.porter.sh/mixin/tofu/pkg"
	"get.porter.sh/porter/pkg/mixin"
	"get.porter.sh/porter/pkg/pkgmgmt"
	"get.porter.sh/porter/pkg/porter/version"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	return version.PrintVersion(m.Context, opts, m.Version())
}

func (m *Mixin) Version() mixin.Metadata {
	return mixin.Metadata{
		Name: "tofu",
		VersionInfo: pkgmgmt.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "Porter Authors",
		},
	}
}
