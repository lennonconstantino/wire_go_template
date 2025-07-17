package inject

import (
	"testing"
)

func TestCacheRepositoryImpl_GetConnection(t *testing.T) {
	conn, err := GetConnectionCache()
	if err != nil {
		t.Fatalf("failed to get cache connection: %v", err)
	}
	repo, err := NewCacheRepository(conn)
	if err != nil {
		t.Fatalf("failed to create cache repository: %v", err)
	}
	repo.GetConnection() // Should print connection string
}

func TestSqlRepositoryImpl_GetConnection(t *testing.T) {
	conn, err := GetConnectionSql()
	if err != nil {
		t.Fatalf("failed to get sql connection: %v", err)
	}
	repo, err := NewSqlRepository(conn)
	if err != nil {
		t.Fatalf("failed to create sql repository: %v", err)
	}
	repo.GetConnection() // Should print connection string
}

func TestRepositoryImpl_GetCacheAndSql(t *testing.T) {
	cacheConn, _ := GetConnectionCache()
	cacheRepo, _ := NewCacheRepository(cacheConn)
	sqlConn, _ := GetConnectionSql()
	sqlRepo, _ := NewSqlRepository(sqlConn)
	repo, err := NewRepository(sqlRepo, cacheRepo)
	if err != nil {
		t.Fatalf("failed to create repository: %v", err)
	}
	if repo.GetCache() == nil || repo.GetSql() == nil {
		t.Error("repository should return non-nil cache and sql repositories")
	}
}

func TestAppContextInitialization(t *testing.T) {
	cacheConn, _ := GetConnectionCache()
	cacheRepo, _ := NewCacheRepository(cacheConn)
	sqlConn, _ := GetConnectionSql()
	sqlRepo, _ := NewSqlRepository(sqlConn)
	repo, _ := NewRepository(sqlRepo, cacheRepo)
	appCtx, err := NewAppContext(cacheRepo, sqlRepo, repo)
	if err != nil {
		t.Fatalf("failed to create app context: %v", err)
	}
	if appCtx.CacheRepository == nil || appCtx.SqlRepository == nil || appCtx.Repository == nil {
		t.Error("app context should have all repositories initialized")
	}
}

func TestWireInitializeEvent(t *testing.T) {
	appCtx, err := InitializeEvent()
	if err != nil {
		t.Fatalf("wire InitializeEvent failed: %v", err)
	}
	if appCtx == nil {
		t.Error("wire InitializeEvent returned nil app context")
	}
}
