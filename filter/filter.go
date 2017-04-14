package filter

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/lufred/Proxy_Crawler/config"
	"github.com/lufred/Proxy_Crawler/storage"

	"gopkg.in/mgo.v2/bson"
)

var myConfig = config.NewConfig()
var conn = storage.NewStorage(myConfig.Mongo.Addr, myConfig.Mongo.DB, myConfig.Mongo.Collection)
var checkURL = myConfig.ProxyCheckURL

func Run() {
	ipList, err := getAllUnUsableIp()
	if err != nil {
		log.Println(err)
	}
	log.Printf("共:%d", len(ipList))
	for _, ip := range ipList {
		log.Printf("开始分析:%s", ip.IP)
		go CheckIP(ip)

	}

}

func CheckIP(ip *storage.IP) {
	client, err := GetHttpProxyClient(ip)
	sum := 0
	checkResult := false
	for sum < 3 {

		if err != nil {
			log.Println(err)
		}
		resp, err := client.Get(checkURL)
		if err != nil {
			//log.Printf("222开始第%d次清洗%s:%d,结果为：%s", sum, ip.IP, ip.Port, "超时")
		} else {
			if resp.StatusCode == 200 {
				checkResult = true
				break
			}
		}
		time.Sleep(5 * time.Second)
		sum += 1
	}
	if !checkResult {
		go RemoveIP(ip)
	} else {
		go UpdateIP(ip)
	}
	log.Printf("%d次清洗%s:%d,结果为：%t", sum, ip.IP, ip.Port, checkResult)

}

func getAllUnUsableIp() ([]*storage.IP, error) {
	var ip storage.IP
	log.Println(ip)
	var ipList []*storage.IP
	var err error
	ipList, err = storage.Find(conn, bson.M{"usable": false})
	if err != nil {
		return nil, err
	}
	return ipList, nil
}
func GetHttpProxyClient(ip *storage.IP) (*http.Client, error) {
	if ip != nil {
		var proxy = ip.Protocol + "://" + ip.IP + ":" + strconv.Itoa(ip.Port)

		proxyURL, err := url.Parse(proxy)
		if err != nil {
			log.Println(err)
		}
		transport := http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		pClient := &http.Client{Transport: &transport}
		pClient.Timeout = 10 * time.Second
		return pClient, nil
	} else {
		return nil, fmt.Errorf("Agent address cannot be empty")
	}

}

func RemoveIP(ip *storage.IP) {
	err := ip.Remove(conn, bson.M{"_id": ip.ID})
	if err != nil {
		log.Println(err)
	}
}
func UpdateIP(ip *storage.IP) {
	ip.Usable = true
	err := ip.Update(conn, bson.M{"_id": ip.ID})
	if err != nil {
		log.Println(err)
	}
}
