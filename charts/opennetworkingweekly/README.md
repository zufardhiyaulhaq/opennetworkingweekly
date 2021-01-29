# opennetworkingweekly charts
Helm chart for opennetworkingweeklys

### Installing the charts
```
helm repo add zufardhiyaulhaq https://charts.zufardhiyaulhaq.com/
helm install zufardhiyaulhaq/opennetworkingweekly --name-template opennetworkingweekly -f values.yaml
```

### Configuration

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| community | string | `"Open Networking Indonesia Community"` |  |
| cronSchedule | string | `"0 8 * * 4"` |  |
| github.branch | string | `"master"` |  |
| github.organization | string | `"zufardhiyaulhaq"` |  |
| github.repository | string | `"community-ops"` |  |
| github.repository_path | string | `"./manifest/opennetworking-community/"` |  |
| github.token | string | `"your-token"` |  |
| image.name | string | `"opennetworkingweekly"` |  |
| image.repository | string | `"zufardhiyaulhaq/opennetworkingweekly"` |  |
| image.tag | string | `"0.0.1"` |  |
| image_url | string | `"https://opennetworking.org/wp-content/uploads/2020/09/ONF-logo-og.png"` |  |
| jobHistoryLimit | int | `1` |  |
| namespace | string | `"opennetworking-community"` |  |
| tags | string | `"weekly,onf"` |  |

check & modify values.yaml for details
