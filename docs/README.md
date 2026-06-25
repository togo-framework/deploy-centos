# deploy-centos — docs

**CentOS VPS deploy.** SSH/systemd VPS driver for CentOS/RHEL hosts (yum/dnf prerequisites).

## Install

```bash
togo install togo-framework/deploy-centos
```

Registers on the [`deploy`](https://github.com/togo-framework/deploy) base; select it with **deploy.provider in togo.yaml (or DEPLOY_PROVIDER)**, then use **`togo deploy`**.

## Interface

`Deployer` — `Provision`/`Deploy`/`Destroy`/`Status` over a `Spec{App,Dir,BuildCmd,Host,User,Image,Region,Domain}` built from your `togo.yaml`.

## Usage & notes

Same SSH/systemd flow as `deploy-ubuntu` with yum/dnf package install. Config from `togo.yaml`: `host` (required), `user` (default `root`).

## Example

```bash
togo deploy --provider centos --dry-run   # preview the plan
togo deploy --provider centos
```

## Links

- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/deploy-centos)
