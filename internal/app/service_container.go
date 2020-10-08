package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"strconv"
	"sync"
	"untitled/internal/infrastructure"
	"untitled/internal/interfaces"
	"untitled/internal/usecase"
)

type serviceManager struct{}

var (
	sm            *serviceManager
	containerOnce sync.Once
	connectorOnce sync.Once
	pgConnector *infrastructure.PostgresConnector
)

type ServiceContainer interface {
	PgConnectService() *infrastructure.PostgresConnector
	AccountWebServiceFactory() interfaces.AccountWebService
}

func NewServiceContainer() ServiceContainer {
	if sm == nil {
		containerOnce.Do(func() {
			sm = &serviceManager{}
		})
	}
	return sm
}

func (sm *serviceManager) PgConnectService() *infrastructure.PostgresConnector {
	if pgConnector == nil {
		connectorOnce.Do(func() {
			var err error
			pgPort, err := strconv.Atoi(os.Getenv("PG_POST"))
			if err != nil {
				panic(err)
			}
			psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				os.Getenv("PG_HOST"),
				pgPort,
				os.Getenv("PG_USER"),
				os.Getenv("PG_PASS"),
				os.Getenv("PG_DB"))
			pgConnector = &infrastructure.PostgresConnector{}
			pgConnector.Pool, err = pgxpool.Connect(context.Background(), psqlInfo)
			if err != nil {
				panic(err)
			}
		})
	}

	return pgConnector
}

func (sm *serviceManager) AccountWebServiceFactory() interfaces.AccountWebService {

	pgConnectorInstance := sm.PgConnectService()

	accountRepositoryInstance := &interfaces.AccountSQLRepository{DB: pgConnectorInstance}
	accountServiceInstance := &usecase.AccountService{AccountRepository: accountRepositoryInstance}

	return interfaces.AccountWebService{AccountService: accountServiceInstance}
}

