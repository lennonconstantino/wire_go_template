package inject

import "fmt"

/////

type CacheConnection struct {
	connection string
}

func GetConnectionCache() (CacheConnection, error) {
	return CacheConnection{
		connection: "cacheConnection",
	}, nil
}

type CacheRepository interface {
	GetConnection()
	GetCache()
}

type CacheRepositoryImpl struct {
	connection CacheConnection
}

func NewCacheRepository(connection CacheConnection) (*CacheRepositoryImpl, error) {
	return &CacheRepositoryImpl{
		connection: connection,
	}, nil
}

// Change method receivers to pointer receivers
func (c *CacheRepositoryImpl) GetConnection() {
	fmt.Println(c.connection.connection)
}

func (c *CacheRepositoryImpl) GetCache() {
	fmt.Println("GetCache")
}

// /
type SqlConnection struct {
	connection string
}

func GetConnectionSql() (SqlConnection, error) {
	return SqlConnection{
		connection: "sqlConnection",
	}, nil
}

type SqlRepository interface {
	GetConnection()
	GetSql()
}

type SqlRepositoryImpl struct {
	connection SqlConnection
}

func NewSqlRepository(connection SqlConnection) (*SqlRepositoryImpl, error) {
	return &SqlRepositoryImpl{
		connection: connection,
	}, nil
}

func (s *SqlRepositoryImpl) GetConnection() {
	fmt.Println(s.connection.connection)
}

func (s *SqlRepositoryImpl) GetSql() {
	fmt.Println("GetSql")
}

// /
type Repository interface {
	GetCache() CacheRepository
	GetSql() SqlRepository
}

type RepositoryImpl struct {
	Cache CacheRepository
	Sql   SqlRepository
}

func NewRepository(sql SqlRepository, cache CacheRepository) (*RepositoryImpl, error) {
	return &RepositoryImpl{
		Sql:   sql,
		Cache: cache,
	}, nil
}

func (r *RepositoryImpl) GetCache() CacheRepository {
	fmt.Println("GetCache")
	return r.Cache
}

func (r *RepositoryImpl) GetSql() SqlRepository {
	fmt.Println("GetSql")
	return r.Sql
}

// Rename Initialization to AppContext and use interfaces for fields

type AppContext struct {
	CacheRepository CacheRepository
	SqlRepository   SqlRepository
	Repository      Repository
}

// Update constructor to use interfaces
func NewAppContext(cacheRepository CacheRepository, sqlRepository SqlRepository, repository Repository) (*AppContext, error) {
	return &AppContext{
		CacheRepository: cacheRepository,
		SqlRepository:   sqlRepository,
		Repository:      repository,
	}, nil
}
