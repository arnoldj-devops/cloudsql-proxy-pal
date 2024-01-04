[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![Latest release](https://badgen.net/github/release/arnoldj-devops/cloudsql-proxy-pal)](https://github.com/arnoldj-devops/cloudsql-proxy-pal/releases)
[![GitHub license](https://img.shields.io/github/license/arnoldj-devops/cloudsql-proxy-pal.svg)](https://github.com/arnoldj-devops/cloudsql-proxy-pal/blob/master/LICENSE)

# Install cloudsql-proxy-pal

## MacOS

```bash
brew install arnoldj-devops/tools/cloudsql-proxy-pal
```

## Ubuntu

```bash
curl -s https://raw.githubusercontent.com/arnoldj-devops/cloudsql-proxy-pal/master/scripts/install.sh | bash
```

# Usage

**To connect a database instance** <br />

```bash
cloudsql-proxy-pal connect --port=<port-number>
```

**Select the CloudSQL instance to connect**

By default `port=5432` and is optional <br />
<br />
**To disconnect instance** <br />

```bash
cloudsql-proxy-pal disconnect
```

<br />

**For all commands** <br />

```bash
cloudsql-proxy-pal --help
```

 <br />

# Prerequisites:

- [cloud_sql_proxy](https://keyvalue.atlassian.net/wiki/spaces/TECH/pages/263782429/Cloud+SQL+connect+with+SQL+Proxy)
- [gcloud](https://cloud.google.com/sdk/docs/install)

### Install prerequisites in single click (WIP)

```bash
curl -s https://raw.githubusercontent.com/arnoldj-devops/cloudsql-proxy-pal/master/scripts/install_prerequisites.sh | bash
```
