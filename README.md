<!-- togo-header -->
<div align="center">
  <picture><source media="(prefers-color-scheme: dark)" srcset=".github/assets/togo-mark-dark.svg" /><img src=".github/assets/togo-mark.svg" alt="ToGO" height="64" /></picture>
  <h1>togo-framework/deploy-centos</h1>
  <p><a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1F8A99" /></a> <a href="https://pkg.go.dev/github.com/togo-framework/deploy-centos"><img src="https://pkg.go.dev/badge/github.com/togo-framework/deploy-centos.svg" /></a> <img src="https://img.shields.io/badge/license-MIT-blue" /></p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/deploy
togo install togo-framework/deploy-centos
```
<!-- /togo-header -->

CentOS / RHEL / Rocky / Alma VPS deploy driver — same SSH + systemd flow as deploy-ubuntu, installing prerequisites with yum/dnf.

Configure in `togo.yaml`:

```yaml
deploy:
  provider: centos
  host: 1.2.3.4
  user: root
```

<!-- togo-sponsors -->
---
<div align="center"><h3>Premium sponsors</h3><p><a href="https://id8media.com"><strong>ID8 Media</strong></a> · <a href="https://one-studio.co"><strong>One Studio</strong></a></p><p><sub><a href="https://github.com/sponsors/fadymondy">Become a sponsor</a>.</sub></p></div>
<!-- /togo-sponsors -->
