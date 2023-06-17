package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../../..")
)

const ConfName = "userScoreConf"

var Config config

type config struct {
	ServerIP string `yaml:"serverip"`

	RpcRegisterIP string `yaml:"rpcRegisterIP"`
	ListenIP      string `yaml:"listenIP"`

	ServerVersion string `yaml:"serverversion"`
	Api           struct {
		GinPort  []int  `yaml:"apiPort"`
		ListenIP string `yaml:"listenIP"`
	}

	Mysql struct {
		DBAddress      []string `yaml:"dbMysqlAddress"`
		DBUserName     string   `yaml:"dbMysqlUserName"`
		DBPassword     string   `yaml:"dbMysqlPassword"`
		DBDatabaseName string   `yaml:"dbMysqlDatabaseName"`
		DBTableName    string   `yaml:"DBTableName"`
		DBMsgTableNum  int      `yaml:"dbMsgTableNum"`
		DBMaxOpenConns int      `yaml:"dbMaxOpenConns"`
		DBMaxIdleConns int      `yaml:"dbMaxIdleConns"`
		DBMaxLifeTime  int      `yaml:"dbMaxLifeTime"`
		LogLevel       int      `yaml:"logLevel"`
		SlowThreshold  int      `yaml:"slowThreshold"`
	}
	Mongo struct {
		DBUri                string   `yaml:"dbUri"`
		DBAddress            []string `yaml:"dbAddress"`
		DBDirect             bool     `yaml:"dbDirect"`
		DBTimeout            int      `yaml:"dbTimeout"`
		DBDatabase           string   `yaml:"dbDatabase"`
		DBSource             string   `yaml:"dbSource"`
		DBUserName           string   `yaml:"dbUserName"`
		DBPassword           string   `yaml:"dbPassword"`
		DBMaxPoolSize        int      `yaml:"dbMaxPoolSize"`
		DBRetainChatRecords  int      `yaml:"dbRetainChatRecords"`
		ChatRecordsClearTime string   `yaml:"chatRecordsClearTime"`
	}
	Redis struct {
		DBAddress     []string `yaml:"dbAddress"`
		DBMaxIdle     int      `yaml:"dbMaxIdle"`
		DBMaxActive   int      `yaml:"dbMaxActive"`
		DBIdleTimeout int      `yaml:"dbIdleTimeout"`
		DBUserName    string   `yaml:"dbUserName"`
		DBPassWord    string   `yaml:"dbPassWord"`
		EnableCluster bool     `yaml:"enableCluster"`
	}
	RpcPort struct {
		OpenImMessageGatewayPort []int `yaml:"openImMessageGatewayPort"`
		UserScorePort            []int `yaml:"userScorePort"`
		ChainUpPort              []int `yaml:"chainUpPort"`
	}
	RpcRegisterName struct {
		UserScoreName   string `yaml:"userScoreName"`
		ChainUpName     string `yaml:"chainUpName"`
		OpenImRelayName string `yaml:"openImRelayName"`
	}
	Etcd struct {
		EtcdSchema        string   `yaml:"etcdSchema"`
		EtcdStorageSchema string   `yaml:"etcdStorageSchema"`
		EtcdAddr          []string `yaml:"etcdAddr"`
		UserName          string   `yaml:"userName"`
		Password          string   `yaml:"password"`
		Secret            string   `yaml:"secret"`
	}
	Log struct {
		StorageLocation       string   `yaml:"storageLocation"`
		RotationTime          int      `yaml:"rotationTime"`
		RemainRotationCount   uint     `yaml:"remainRotationCount"`
		RemainLogLevel        uint     `yaml:"remainLogLevel"`
		ElasticSearchSwitch   bool     `yaml:"elasticSearchSwitch"`
		ElasticSearchAddr     []string `yaml:"elasticSearchAddr"`
		ElasticSearchUser     string   `yaml:"elasticSearchUser"`
		ElasticSearchPassword string   `yaml:"elasticSearchPassword"`
	}
	Kafka struct {
		SASLUserName string `yaml:"SASLUserName"`
		SASLPassword string `yaml:"SASLPassword"`
		BusinessTop  struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
			Group string   `yaml:"group"`
		} `yaml:"businessTop"`
		Ws2mschat struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		}
		LikesAction struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		} `yaml:"likesAction"`
		AnnouncementAction struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		} `yaml:"announcementAction"`
		//Ws2mschatOffline struct {
		//	Addr  []string `yaml:"addr"`
		//	Topic string   `yaml:"topic"`
		//}
		MsgToMongo struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		}
		Ms2pschat struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		}
		ConsumerGroupID struct {
			MsgToRedis    string `yaml:"msgToTransfer"`
			MsgToMongo    string `yaml:"msgToMongo"`
			MsgToMySql    string `yaml:"msgToMySql"`
			MsgToPush     string `yaml:"msgToPush"`
			MsgToLike     string `yaml:"msgToLike"`
			MsgToAnnounce string `yaml:"msgToAnnounce"`
		}
	}

	Prometheus struct {
		Enable                  bool  `yaml:"enable"`
		UserScorePrometheusPort []int `yaml:"userScorePrometheusPort"`
		ChainUpPrometheusPort   []int `yaml:"chainUpPrometheusPort"`
	} `yaml:"prometheus"`

	IsPublicEnv bool `yaml:"ispublicenv"`

	WithdrawTransferOwnerAddress       string `yaml:"withdrawTransferOwnerAddress"`
	WithdrawTransferOwnerPrivateKeyHex string `yaml:"withdrawTransferOwnerPrivateKeyHex"`
	MinWithdrawAmount                  int64  `yaml:"minWithdrawAmount"`
	WithdrawScoreCommission            uint64 `yaml:"withdrawScoreCommission"`

	CoinInfoMap map[string]*CoinInfo `yaml:"coinInfoMap"`
}

type CoinInfo struct {
	Decimal  int    `yaml:"decimal"`
	Rate     int64  `yaml:"rate"`
	Contract string `yaml:"contract"`
	Endpoint string `yaml:"endpoint"`
}

func unmarshalConfig(config interface{}, configName string) {
	var env string
	if configName == "config.yaml" {
		env = "CONFIG_NAME"
	}
	cfgName := os.Getenv(env)
	if len(cfgName) != 0 {
		bytes, err := ioutil.ReadFile(filepath.Join(cfgName, "config", configName))
		if err != nil {
			bytes, err = ioutil.ReadFile(filepath.Join(Root, "config", configName))
			if err != nil {
				panic(err.Error() + " config: " + filepath.Join(cfgName, "config", configName))
			}
		} else {
			Root = cfgName
		}
		if err = yaml.Unmarshal(bytes, config); err != nil {
			panic(err.Error())
		}
	} else {
		bytes, err := ioutil.ReadFile(fmt.Sprintf("../config/%s", configName))
		if err != nil {
			bytes, err = ioutil.ReadFile(filepath.Join(Root, "config", configName))
			if err != nil {
				panic(err.Error() + " config: " + filepath.Join(cfgName, "config", configName))
			}
		}
		if err = yaml.Unmarshal(bytes, config); err != nil {
			panic(err.Error())
		}
	}
}

func init() {
	unmarshalConfig(&Config, "config.yaml")
}
