# goperset


Golang Apache Superset API Package

This project is mentioned in the [Apache Superset](https://superset.apache.org/) [official wiki](https://github.com/apache/superset/wiki/Community-Resource-Library#third-party-libraries) as a Golang API Package.

This project is inspired from: [gocloak](https://github.com/Nerzal/gocloak)

For Questions either raise an issue, or come to the [gopher-slack](https://invite.slack.golangbridge.org/) into the channel [#goperset](https://gophers.slack.com/app_redirect?channel=goperset)

__SOON__: Benchmarks can be found [here](https://justmike1.github.io/goperset/dev/bench/)

## Contribution

(WIP) <https://github.com/justmike1/goperset/wiki/Contribute>

## Changelog

For release notes please consult the specific releases [here](https://github.com/justmike1/goperset/releases)

For tags please consult the specific tags [here](https://github.com/justmike1/goperset/tags)


## Usage

### Installation

```shell
go install github.com/justmike1/goperset@latest
```

### Importing

```go
 import "github.com/justmike1/goperset"
```

### Connecting to Superset Client

```go
func main() {
    ctx := context.Background()
    client := goperset.NewClient(ctx, "https://superset.domain.net/")
    tokens, err := goperset.GetAccessTokens(client, "admin", "admin")
    if err != nil {
        t.Errorf("Error getting access tokens: %v", err)
        return
    }
}
```

### Example of Creating a new database connection

```go
func createDatabase(client, tokens string) error {
    databaseUri := fmt.Sprintf("postgresql+psycopg2://%s:%s@%s:%s/database", "username", "password", "postgresql", "5432")
    payload := goperset.DatabasePayload{
        ConfigurationMethod: goperset.StringP("sqlalchemy_form"),
        Engine:              goperset.StringP("postgresql+psycopg2"),
        DatabaseName:        goperset.StringP("postgresql"),
        SQLAlchemyURI:       goperset.StringP(databaseUri),
    }
    _, err := goperset.CreateDatabase(client, tokens, payload)
    if err != nil {
        return fmt.Errorf("error sending request: %v", err)
    }
    return nil
}
```

## Features

```go
// GoPerset holds all methods a client should fulfill
type GoPerset interface {
	
 // Client
 NewClient(basePath string) &Goperset
 GetAccessTokens(client *Goperset, username string, password string) (ClientToken, error)
 ClientResty(client *Goperset, token string, csrfToken string, contentType string, method string, endpoint string, payload interface{}) ([]byte, error)
 
 // Database
 CreateDatabase(client *Goperset, tokens ClientToken, payload DatabasePayload) ([]byte, error)
 
 // Dashboard
 CreateDashboard(client *Goperset, tokens ClientToken, payload DashboardPayload) ([]byte, error)
 GetDashboard(client *Goperset, tokens ClientToken, params GetDashboardParams) ([]byte, error)
 GetDashboardInfo(client *Goperset, tokens ClientToken, params DashboardInfoParams) ([]byte, error)

 // Dataset
 CreateDataset(client *Goperset, tokens ClientToken, payload DatasetPayload) ([]byte, error)
 
 // Chart
 CreateChart(client *Goperset, tokens ClientToken, payload DatasetPayload) ([]byte, error)
}
```

### Unit Testing

```shell
SUPERSET_BASE_PATH=https://superset.domain.net go test client_test.go -v
```
