package models

import (
	MRedis "MQConfig/redis"
	"log"
)

type RedisModel struct {

}

func (this *RedisModel)Demo() string {

	instance := new(MRedis.MRedis)

	return instance.SetAndGet("aa", "admin098765", 0)
}

func (this *RedisModel)Demo2() string {
	instance := MRedis.Factory()

	//instance.SetAndGet("aaa", "111111", 10)

	data := instance.Get("aaa")

	log.Println(instance.Obj)

	return data
}

func (this *RedisModel)Demo3() string {
	instance := new(MRedis.MRedis)

	data := instance.MSet("beeooooo", "34567890-")

	return data
}
