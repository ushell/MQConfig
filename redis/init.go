package redis

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
)

type MRedis struct {
	IP string
	Port string
	Password string
	Database string
	Obj interface{}
}

//Redis单机
func (this *MRedis)instance() (redis.Client, error) {
	//配置
	node, err := beego.AppConfig.GetSection("redis")

	if err != nil {
		return redis.Client{}, err
	}

	database, _ := strconv.Atoi(node["database"])

	//连接
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", node["ip"], node["port"]),
		Password: node["password"],
		DB:       database,
	})

	pong, err := client.Ping().Result()

	if pong != "PONG" {
		return redis.Client{}, err
	}

	return *client, nil
}

//Redis集群
func (this *MRedis)clusterInstance() (redis.ClusterClient, error){
	nodes := []string{":6379", ":6380", ":6381"}

	options := &redis.ClusterOptions{Addrs:nodes}

	node := redis.NewClusterClient(options)

	err := node.Ping().Err()
	if err != nil {
		log.Fatal("Redis Error:", err)
	}

	return *node, nil
}

//工厂模式
func Factory() *MRedis {
	var object interface{}

	_instance := new(MRedis)

	//集群开关
	isRedisClusters, err := beego.AppConfig.Int("redis_cluster")

	if isRedisClusters == 0 {
		object, err = _instance.instance()
	} else {
		object, err = _instance.clusterInstance()
	}

	if err != nil {
		log.Fatal(err)
	}

	_instance.Obj = object

	return _instance
}

//通用方法
func (this *MRedis)SetAndGet(key, data string, expire time.Duration) string {
	instance, err := this.instance()

	if err != nil {
		panic(err)
	}

	err = instance.Set(key, data, expire).Err()
	if err != nil {
		return ""
	}

	ret, err := instance.Get(key).Result()
	if err == redis.Nil {
		return ""
	}

	if err != nil {
		return ""
	}

	return ret
}

func (this *MRedis)Get(key string) string {

	log.Println("=>", this.Obj)

	return ""
}

func (this *MRedis)MSet(key, data string) string {
	instance, err := this.clusterInstance()

	if err != nil {
		log.Println("=>", err)
	}

	err = instance.Set(key, data, 100).Err()

	if err != nil {
		log.Println(err)
	}

	ret, _ := instance.Get(key).Result()

	return ret
}

