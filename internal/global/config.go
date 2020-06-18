package global

import (
	"github.com/profzone/eden-framework/pkg/courier/transport_grpc"
	"github.com/profzone/eden-framework/pkg/courier/transport_http"
	"longhorn/proxy-admin/internal/constants/enum"
	"time"
)

type SnowflakeConfig struct {
	Epoch      int64
	BaseNodeID int64
	NodeCount  int64
	NodeBits   uint8
	StepBits   uint8
}

type DBConfig struct {
	DBType             enum.DbType
	Endpoints          []string
	UserName           string
	Password           string
	DatabaseName       string
	ConnectionTimeout  time.Duration
	ClusterPrefix      string
	ServerPrefix       string
	BindPrefix         string
	ApiPrefix          string
	RouterPrefix       string
	OrganizationPrefix string
}

var Config = struct {
	// administrator
	GRPCServer transport_grpc.ServeGRPC
	HTTPServer transport_http.ServeHTTP

	// db
	DBConfig

	// id generation
	SnowflakeConfig
}{
	GRPCServer: transport_grpc.ServeGRPC{
		Port: 8900,
	},
	HTTPServer: transport_http.ServeHTTP{
		Port:     8001,
		WithCORS: true,
	},

	DBConfig: DBConfig{
		DBType:             enum.DB_TYPE__MONGODB,
		Endpoints:          []string{"127.0.0.1:27017"},
		ClusterPrefix:      "clusters",
		ServerPrefix:       "servers",
		BindPrefix:         "binds",
		ApiPrefix:          "apis",
		RouterPrefix:       "routers",
		OrganizationPrefix: "organizations",
	},

	SnowflakeConfig: SnowflakeConfig{
		Epoch:      1288351723598,
		BaseNodeID: 1,
		NodeCount:  100,
		NodeBits:   10,
		StepBits:   12,
	},
}
