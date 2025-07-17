//go:build wireinject
// +build wireinject

package inject

import "github.com/google/wire"

//var conncache = wire.NewSet(GetConnectionCache)

var cacheSet = wire.NewSet(
	wire.ProviderFunc(GetConnectionCache),
	wire.ProviderFunc(NewCacheRepository),
	wire.Bind(new(CacheRepository), new(*CacheRepositoryImpl)),
)

//var connsql = wire.NewSet(GetConnectionSql)

var sqlSet = wire.NewSet(
	wire.ProviderFunc(GetConnectionSql),
	wire.ProviderFunc(NewSqlRepository),
	wire.Bind(new(SqlRepository), new(*SqlRepositoryImpl)),
)

var repositorySet = wire.NewSet(
	wire.ProviderFunc(NewRepository),
	wire.Bind(new(Repository), new(*RepositoryImpl)),
)

// Update injector to use AppContext
func InitializeEvent() (*AppContext, error) {
	wire.Build(wire.ProviderFunc(NewAppContext), cacheSet, sqlSet, repositorySet)
	return nil, nil
}
