package goperset_test

import (
	"context"
	"fmt"
	"github.com/justmike1/goperset"
	"os"
	"testing"
)

var (
	basePath = getEnv("SUPERSET_BASE_PATH", "https://superset.domain.net/")
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func initTest() (*goperset.Goperset, goperset.ClientToken) {
	ctx := context.Background()
	client := goperset.NewClient(ctx, basePath)
	tokens, err := goperset.GetAccessTokens(client, "admin", "admin")
	if err != nil {
		return client, goperset.ClientToken{}
	}
	return client, tokens
}

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	client := goperset.NewClient(ctx, basePath)
	if client == nil {
		t.Error("Client is nil")
	}
}

func TestCreateDatabase(t *testing.T) {
	client, tokens := initTest()
	databaseUri := fmt.Sprintf("postgresql+psycopg2://%s:%s@%s:%s/database", "username", "password", "postgresql", "5432")
	payload := goperset.DatabasePayload{
		ConfigurationMethod: goperset.StringP("sqlalchemy_form"),
		Engine:              goperset.StringP("postgresql+psycopg2"),
		DatabaseName:        goperset.StringP("postgresql"),
		SQLAlchemyURI:       goperset.StringP(databaseUri),
	}
	_, err := goperset.CreateDatabase(client, tokens, payload)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAccessTokens(t *testing.T) {
	ctx := context.Background()
	client := goperset.NewClient(ctx, basePath)
	tokens, err := goperset.GetAccessTokens(client, "admin", "admin")
	if err != nil {
		t.Error(err)
		return
	}
	if tokens.AccessToken == "" {
		t.Error("Access token is empty")
	}
	if tokens.CsrfToken == "" {
		t.Error("CSRF token is empty")
	}
}
