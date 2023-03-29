package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf = JsonConfigNodes()

var vp *viper.Viper

func JsonConfigNodes() ConfigNodes {
	//intial := JsonInitialConfig()
	vp = viper.New()
	var config ConfigNodes

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config_files")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		log.Println("Error while reading init_config.json")
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		log.Println("Error Unmarrsall", err)
	}
	return config
}

func DBConfig() (string, string, string, string, string, string) {
	data := JsonConfigNodes()
	postgresqlHost := data.ConfigDB.PostgreHost
	sslmode := data.ConfigDB.PostgreSSLMode
	postgresqlUsername := data.ConfigDB.PostgreUser
	postgresqlPassword := data.ConfigDB.PostgrePassword
	postgresqlPort := data.ConfigDB.PostgrePort
	postgresqlSchema := data.ConfigDB.PostgreSchema

	return sslmode, postgresqlUsername, postgresqlPassword, postgresqlHost, postgresqlPort, postgresqlSchema
}

func GRPCConfig1() (string, string, bool) {
	data := JsonConfigNodes()
	grpc_host := data.GRPCServer1.GRPCHost
	grpc_port := data.GRPCServer1.GRPCPort
	grpc_tls := data.GRPCServer1.GRPCTls

	return grpc_host, grpc_port, grpc_tls
}

func GRPCConfig2() (string, string, bool) {
	data := JsonConfigNodes()
	grpc_host := data.GRPCServer2.GRPCHost
	grpc_port := data.GRPCServer2.GRPCPort
	grpc_tls := data.GRPCServer2.GRPCTls

	return grpc_host, grpc_port, grpc_tls
}

func GRPCConfig3() (string, string, bool) {
	data := JsonConfigNodes()
	grpc_host := data.GRPCServer3.GRPCHost
	grpc_port := data.GRPCServer3.GRPCPort
	grpc_tls := data.GRPCServer3.GRPCTls

	return grpc_host, grpc_port, grpc_tls
}

func GRPCConfig4() (string, string, bool) {
	data := JsonConfigNodes()
	grpc_host := data.GRPCServer4.GRPCHost
	grpc_port := data.GRPCServer4.GRPCPort
	grpc_tls := data.GRPCServer4.GRPCTls

	return grpc_host, grpc_port, grpc_tls
}

func GRPCConfig5() (string, string, bool) {
	data := JsonConfigNodes()
	grpc_host := data.GRPCServer5.GRPCHost
	grpc_port := data.GRPCServer5.GRPCPort
	grpc_tls := data.GRPCServer5.GRPCTls

	return grpc_host, grpc_port, grpc_tls
}

func GRPCPathsConfig() (string, string) {
	data := JsonConfigNodes()
	cert_File := data.GRPCPaths.CertFile
	key_File := data.GRPCPaths.CertFile

	return cert_File, key_File
}

func SuperUseConfig() (string, string, bool, bool) {
	data := JsonConfigNodes()
	user := data.SuperAdmin.UserName
	password := data.SuperAdmin.Password
	role := data.SuperAdmin.Role
	permission := data.SuperAdmin.Permissions

	return user, password, role, permission
}
