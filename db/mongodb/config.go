package mongodb

import (
	"fmt"
	"strings"
)

// Config mongodb deployment 连接基础配置：mongodb单实例、副本集(可有读写分离)、分片集群
type Config struct {
	//该连接池可操作的数据库（可选配置）
	DBNames []string `yaml:"dbnames"`
	//初始化mongodb连接管理器时 给定默认数据库名（必须配置，否则业务方需注意不调用Manager.DB()函数）
	DefaultDBName string `yaml:"default_dbname"`

	//副本集的各个实例地址，此处配置多个 仅为防止初始化时单点失败
	//连接后由driver监控副本集状态 做自动故障转移，读写分离可由连接池配置ReadPreference决定
	Hosts []string `yaml:"hosts"`
	//副本集名字：指定该名字 必须和mongodb部署时相应节点配置的名字一致，client会将相同名字的节点视为在同一副本集内 而忽略其他节点，默认配置空
	ReplicaSet string `yaml:"replica_set"`
	//读首选项配置：决定mongodb driver如何将读操作 路由到副本集的成员，需要读写分离可在此配置
	//若配置文件不指定下述配置，则解析配置时 按照client driver给定的默认值primary模式进行配置
	ReadPreference struct {
		//PrimaryMode 1, PrimaryPreferredMode 2, SecondaryMode 3, SecondaryPreferredMode 4, NearestMode 5
		Mode int `yaml:"mode"`
		//在1个副本集中，从次节点读的 允许的最大复制延迟(单位s)
		MaxStaleness int `yaml:"max_staleness"`
		//Name-Value 标签
		Tags map[string]string `yaml:"tags"`
	} `yaml:"read_preference"`

	//认证：支持用户名密码、AWS、GSSAPI/SSPI、LDAP、SCRAM。X509还需tls配置
	//详见 mongo/options/clientoptions.go Credential注释
	Auth struct {
		Mechanism           string            `yaml:"mechanism"`
		MechanismProperties map[string]string `yaml:"mechanism_properties"`
		Source              string            `yaml:"source"`
		Username            string            `yaml:"username"`
		Password            string            `yaml:"password"`
		PasswordSet         bool              `yaml:"password_set"`
	} `yaml:"auth"`

	//创建连接时的连接超时时间(单位s)，0值表示没有timeout。不调用设值时 driver默认30s
	ConnectTimeout int `yaml:"connect_timeout"`
	//1个连接在1个连接池中 保持空闲状态的最长时长(单位s)。超过该时长 该连接会被移出连接池并关闭。driver默认0值 连接可保持无限期空闲。
	MaxConnIdleTime int `yaml:"max_conn_idle_time"`
	//到每1个server的允许的最大连接数，到达后新请求会阻塞。0值会被driver设置为math.MaxInt64，driver默认100
	MaxPoolSize int `yaml:"max_pool_size"`
	//到每1个server的允许的最小连接数。非0时 若连接均长时空闲 driver会在后台维持最小连接数，driver默认0
	MinPoolSize int `yaml:"min_pool_size"`
}

// String 打印可输出的配置
func (c *Config) String() string {
	var str strings.Builder
	fmt.Fprintln(&str, "mongodb config:")
	fmt.Fprintln(&str, "dbnames:", c.DBNames)
	fmt.Fprintln(&str, "default_dbname:", c.DefaultDBName)

	fmt.Fprintln(&str, "hosts:", c.Hosts)
	fmt.Fprintln(&str, "replica_set:", c.ReplicaSet)
	fmt.Fprintln(&str, "read_preference:", c.ReadPreference)

	fmt.Fprintln(&str, "auth: username:", c.Auth.Username, "source:", c.Auth.Source)

	fmt.Fprintln(&str, "connect_timeout:", c.ConnectTimeout)
	fmt.Fprintln(&str, "max_conn_idle_time:", c.MaxConnIdleTime)
	fmt.Fprintln(&str, "max_pool_size:", c.MaxPoolSize)
	fmt.Fprintln(&str, "min_pool_size:", c.MinPoolSize)

	return str.String()
}