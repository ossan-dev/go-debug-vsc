//go:build integration

package integration

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/ossan-dev/go-debug-vsc/internal/inventory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InventoryTestSuite struct {
	suite.Suite
	ctx          context.Context
	container    testcontainers.Container
	gormDbClient *gorm.DB
}

func TestInventorySuite(t *testing.T) {
	suite.Run(t, new(InventoryTestSuite))
}

func (s *InventoryTestSuite) SetupSuite() {
	var err error
	s.ctx = context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "postgres",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	s.container, err = testcontainers.GenericContainer(s.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(s.T(), err)
	endpoint, err := s.container.Endpoint(s.ctx, "")
	require.NoError(s.T(), err)
	s.gormDbClient, err = gorm.Open(postgres.Open(fmt.Sprintf("host=localhost port=%s user=postgres password=postgres dbname=postgres sslmode=disable", strings.Split(endpoint, ":")[1])))
	require.NoError(s.T(), err)
	require.NotNil(s.T(), s.gormDbClient)
	err = s.gormDbClient.AutoMigrate(&inventory.Item{})
	require.NoError(s.T(), err)
}

func (s *InventoryTestSuite) TearDownSuite() {
	if s.container != nil {
		if err := s.container.Terminate(s.ctx); err != nil {
			s.T().Fatal(err)
		}
	}
}

func (s *InventoryTestSuite) TestAddItemGivesErrWhenQtyLessThanZero() {
	err := inventory.AddItem(s.gormDbClient, &inventory.Item{Name: "mobile phone", Quantity: -4})
	require.Error(s.T(), err)
}

func (s *InventoryTestSuite) TestAddItemGivesErrWhenNameIsEmpty() {
	err := inventory.AddItem(s.gormDbClient, &inventory.Item{Name: "", Quantity: 4})
	require.Error(s.T(), err)
}

func (s *InventoryTestSuite) TestAddItemHappyPath() {
	item := inventory.Item{Name: "mobile phone", Quantity: 4}
	err := inventory.AddItem(s.gormDbClient, &item)
	require.NoError(s.T(), err)
	assert.NotZero(s.T(), item.ID)
}
