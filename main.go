package main

import (
	"log"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/lufred/Proxy_Crawler/config"
	"github.com/lufred/Proxy_Crawler/filter"
	"github.com/lufred/Proxy_Crawler/getter"
	"github.com/lufred/Proxy_Crawler/storage"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go func() {
		for {
			getter.KDL()
			time.Sleep(10 * time.Minute)
		}

	}()

	for {
		filter.Run()
		time.Sleep(5 * time.Minute)
	}

	//swg.Wait()

	log.Println("Done")

}
func MongoTest() {

	var myConfig = config.NewConfig()
	var conn = storage.NewStorage(myConfig.Mongo.Addr, myConfig.Mongo.DB, myConfig.Mongo.Collection)

	ip := storage.NewIP()
	ip.IP = "101.6.53.811"
	err := ip.FindOne(conn, bson.M{"ip": ip.IP})
	if err != nil {
		log.Println(err)
	}
}
func HttpTest() {

	var proxy = "http://119.254.84.90:80"
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		log.Println(err)
	}

	transport := http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	myClient := &http.Client{Transport: &transport}
	myClient.Timeout = 10 * time.Second
	const checkURL = "https://movie.douban.com/tag/"
	log.Println("_______________________________________")
	resp, err := myClient.Get(checkURL)
	if err != nil {
		log.Println("+++++++++++++++++++++++++++++++++++++")
		log.Println(err)

	}
	log.Printf("开始清洗%s,结果为：%d", proxy, resp.StatusCode)
}
