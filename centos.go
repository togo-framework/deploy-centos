// Package centos is a VPS deploy driver for togo deploy targeting CentOS / RHEL /
// Rocky / Alma hosts. Same SSH + systemd flow as deploy-ubuntu, but installs
// prerequisites with yum/dnf. Select with DEPLOY_PROVIDER=centos.
package centos

import (
	"github.com/togo-framework/deploy"
	ubuntu "github.com/togo-framework/deploy-ubuntu"
	"github.com/togo-framework/togo"
)

func init() {
	deploy.RegisterDriver("centos", func(k *togo.Kernel) (deploy.Deployer, error) {
		return ubuntu.New("centos", "(yum install -y ca-certificates rsync || dnf install -y ca-certificates rsync)"), nil
	})
}
