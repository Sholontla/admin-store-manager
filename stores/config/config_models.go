package config

type ConfigNodes struct {
	ConfigPath  Paths       `mapstructure:"config_path"`
	ConfigDB    DB          `mapstructure:"config_db"`
	GRPCServer1 GRPCServer1 `mapstructure:"grpc_server1"`
	GRPCServer2 GRPCServer2 `mapstructure:"grpc_server2"`
	GRPCServer3 GRPCServer3 `mapstructure:"grpc_server3"`
	GRPCServer4 GRPCServer4 `mapstructure:"grpc_server4"`
	GRPCServer5 GRPCServer5 `mapstructure:"grpc_server5"`
	GRPCPaths   GRPCPaths   `mapstructure:"config_grpc_path"`
}

type Paths struct {
	Product_id string `mapstructure:"product_id"`
	Quantity   int    `mapstructure:"quantity"`
	Workera    string `mapstructure:"workera"`
	Workerb    string `mapstructure:"workerb"`
	Workerc    string `mapstructure:"workerc"`
	Workerd    string `mapstructure:"workerd"`
}

type DB struct {
	PostgreUser     string `mapstructure:"postgre_user"`
	PostgrePassword string `mapstructure:"postgre_password"`
	PostgreHost     string `mapstructure:"postgre_host"`
	PostgrePort     string `mapstructure:"postgre_port"`
	PostgreSchema   string `mapstructure:"postgre_schema"`
	PostgreSSLMode  string `mapstructure:"postgre_sslmode"`
}

type GRPCServer1 struct {
	GRPCHost string `mapstructure:"host"`
	GRPCPort string `mapstructure:"port"`
	GRPCTls  bool   `mapstructure:"tls"`
}

type GRPCServer2 struct {
	GRPCHost string `mapstructure:"host"`
	GRPCPort string `mapstructure:"port"`
	GRPCTls  bool   `mapstructure:"tls"`
}

type GRPCServer3 struct {
	GRPCHost string `mapstructure:"host"`
	GRPCPort string `mapstructure:"port"`
	GRPCTls  bool   `mapstructure:"tls"`
}

type GRPCServer4 struct {
	GRPCHost string `mapstructure:"host"`
	GRPCPort string `mapstructure:"port"`
	GRPCTls  bool   `mapstructure:"tls"`
}

type GRPCServer5 struct {
	GRPCHost string `mapstructure:"host"`
	GRPCPort string `mapstructure:"port"`
	GRPCTls  bool   `mapstructure:"tls"`
}

type GRPCPaths struct {
	CertFile string `mapstructure:"cert_File"`
	KeyFile  string `mapstructure:"key_File"`
}
