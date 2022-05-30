package mongo

import (
	"context"
	"crypto/tls"
	"fmt"
	"runtime"
	"time"

	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	"github.com/qiniu/qmgo"
	qmgoOptions "github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const (
	//mgTimeout       = 11 * time.Second
	mgSocketTimeout = int64(10 * time.Second)
	//mgoPoolLimit    = 2048
)

func ConnectToMongo(config Config) (*qmgo.Client, error) {
	ctx := context.Background()
	uri := fmt.Sprintf("%s://%s:%s@%s/?authSource=%s", config.Scheme, config.Username, config.Password, config.Host, config.AuthDBName)
	dbname := config.Database
	connSetting, err := connstring.Parse(uri)
	if err != nil {
		fmt.Printf("ERROR: uri: %s, error err: %s\n", uri, err)
		return nil, err
	}

	socketTimeout := mgSocketTimeout
	conf := qmgo.Config{
		Uri:             uri,
		Database:        dbname,
		SocketTimeoutMS: &socketTimeout,
	}
	oriOpts := options.ClientOptions{}

	// TODO: connect with production or stage
	// If AuthSource is not set in connstring, set it to authDBName
	if connSetting.Username != "" && connSetting.AuthSource == "" {
		conf.Auth = &qmgo.Credential{
			AuthMechanism: connSetting.AuthMechanism,
			Username:      connSetting.Username,
			Password:      connSetting.Password,
			PasswordSet:   connSetting.PasswordSet,
			AuthSource:    config.AuthDBName,
		}
	}

	// total connection pool size
	// NumCPU will get all available CPUs on host,
	// use GOMAXPROCS to get available CPU resource
	poolSize := int(float64(runtime.GOMAXPROCS(0)) * config.PoolSizeMultiplier)
	// because each host has its own connection pool,
	// if we set poolSize directly, it generate too many connections,
	// we set each host's pool size by divide the number of hosts
	poolSize = (poolSize + len(connSetting.Hosts) - 1) / len(connSetting.Hosts)
	minPoolSize := uint64(poolSize / 4)
	maxPoolSize := uint64(poolSize)
	conf.MaxPoolSize = &maxPoolSize
	conf.MinPoolSize = &minPoolSize

	fmt.Printf("mongo driver pool size: %d\n", poolSize)

	// TODO: production mongoConfig
	//poolMonitor := event.PoolMonitor{Event: func(evt *event.PoolEvent) {
	//	met.BumpSum("poolevent", 1.0, "type", evt.Type, "reason", evt.Reason)
	//}}
	//clientOpts.SetPoolMonitor(&poolMonitor)

	if config.SSL {
		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true
		oriOpts.TLSConfig = tlsConfig
	}

	if config.SetSafe {
		// Force the server to wait for a majority ofv members of a replica set to return
		oriOpts.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	}

	opts := qmgoOptions.ClientOptions{ClientOptions: &oriOpts}
	client, err := qmgo.NewClient(ctx, &conf, opts)
	if err != nil {
		fmt.Printf("mongoURI: %s, dbname: %s, err: %s, fail to create mongo clients\n", uri, dbname, err)
		return nil, err
	}

	//fmt.Printf("mongoURI: %s, dbname: %s, connected to mongo\n", uri, dbname)

	return client, nil
}
