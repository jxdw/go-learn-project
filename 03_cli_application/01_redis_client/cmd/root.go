package cmd

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
	//"os"
)
var RedisServer string

var listCmd=&cobra.Command{
	Use:"list",
	Short:"",
	Long:"",
	Run: func(cmd *cobra.Command,args []string) {
		RedisClient := redis.NewClient(&redis.Options{
			Addr:RedisServer,
			Password:"",
			DB:0,
			PoolSize: 1,
		})
		fmt.Println("ping ",RedisServer,"result:",RedisClient.Ping().Val());
		redisKey:=RedisClient.Keys(args[0]).Val()
		fmt.Println(redisKey)
	},
}

var deleteCmd=&cobra.Command{
	Use:"delete",
	Short:"",
	Long:"",
	Run: func(cmd *cobra.Command,args []string) {
		RedisClient := redis.NewClient(&redis.Options{
			Addr:RedisServer,
			Password:"",
			DB:0,
			PoolSize: 1,
		})
		fmt.Println("ping ",RedisServer,"result:",RedisClient.Ping().Val());
		redisKey:=RedisClient.Keys(args[0]).Val()
		RedisClient.Del(redisKey...)
	},
}

var RootCmd=&cobra.Command{
	Use: "",
	Short:"",
	Long:"",
}
func init(){
	RootCmd.AddCommand(deleteCmd,listCmd)
	deleteCmd.Flags().StringVarP(&RedisServer,"redisaddress","r","","redis server and port")
	listCmd.Flags().StringVarP(&RedisServer,"redisaddress","r","","redis server and port")
}
