package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-redis/redis"
	"mosn.io/mosn/pkg/log"
	"mosn.io/mosn/pkg/plugin"
	"mosn.io/mosn/pkg/plugin/proto"
)

type checker struct {
	config map[string]string
}

func (c *checker) Call(request *proto.Request) (*proto.Response, error) {
	header := request.Header
	for k, v := range c.config {
		value, ok := header[k]
		if !ok || value != v {
			return &proto.Response{
				Status: -1,
			}, nil
		}
	}

	//req path
	//X-Mosn-Path

	reqPath := header["X-Mosn-Path"]
	reqPath = strings.ReplaceAll(reqPath, "/v2", "")
	if len(reqPath) > 0 {
		//
		namespace := "lan-k8s"
		if strings.Contains(reqPath, namespace) {
			reqPath = strings.ReplaceAll(reqPath, namespace+"/", "")

			repoTag := strings.Split(reqPath, "/")
			if len(repoTag) > 0 {
				if repoTag[0] != "" {
					//从redis获取repo的token
					token := getToken(repoTag[0])
					if token != "" {
						h := make(map[string]string)
						h["Authorization"] = token
						return &proto.Response{
							Status: 1,
							Header: h,
						}, nil
					}
				}
			} else {
				//直接返回
			}
		} else {
			//错误的namespace 直接返回
			return &proto.Response{
				Status: -1,
			}, nil
		}

	} else {
		//v2请求 可以直接返回结果
		return &proto.Response{
			Status: -1,
		}, nil
	}

	// for k, v := range header {
	// 	log.DefaultLogger.Infof("header.key: %s value: %s", k, v)
	// }

	return &proto.Response{}, nil
}

func getToken(repo string) string {
	if redisClient != nil {

		val, err := redisClient.Get("key").Result()
		if err != nil {
			log.DefaultLogger.Errorf("redis.get.valu.failed;%s", err)
		} else {
			return val
			log.DefaultLogger.Infof("hello.key.value:%s", val)
		}

	} else {
		log.DefaultLogger.Errorf("connect.redis.failed!")
	}
	return ""
}

var (
	redisClient *redis.Client
)

func init() {

	redisAddr := os.Getenv("MOSN_REDIS_ADDR")

	redisPasswd := os.Getenv("MOSN_REDIS_PASSWD")

	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPasswd, // no password set
		// DB:       0,           // use default DB
	})

}

func main() {

	conf := flag.String("c", "", "-c config.json")
	flag.Parse()
	m := make(map[string]string)
	b, err := ioutil.ReadFile(*conf)
	if err == nil {
		json.Unmarshal(b, &m)
	}
	log.DefaultLogger.Infof("configfile: %s get config: %v", *conf, m)

	if redisClient != nil {
		str, err := redisClient.Ping().Result()
		if err != nil {
			log.DefaultLogger.Errorf("ping.result.failed: %s", err)
			// os.Exit(-1)
			// panic(err)
		}
		log.DefaultLogger.Infof("ping.result.success: %s", str)
	} else {
		log.DefaultLogger.Errorf("connect.redis.failed!")
	}

	plugin.Serve(&checker{m})

}
