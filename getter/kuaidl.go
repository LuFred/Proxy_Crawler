package getter

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/lufred/Proxy_Crawler/config"
	"github.com/lufred/Proxy_Crawler/storage"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/mgo.v2/bson"
)

var myConfig = config.NewConfig()
var conn = storage.NewStorage(myConfig.Mongo.Addr, myConfig.Mongo.DB, myConfig.Mongo.Collection)

func KDL() {
	pullURL := "http://www.kuaidaili.com/free/inha/"
	for i := 1; i < 50; i++ {
		currentPullURL := pullURL + strconv.Itoa(i)
		resp, err := http.Get(currentPullURL)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(currentPullURL + `statusCode` + resp.Status)
		go analysisHTML(resp)
		time.Sleep(1 * time.Second)
	}

}

func analysisHTML(res *http.Response) {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Println(err)
	}
	selection := doc.Find("#list tbody tr")
	selection.Each(func(i int, s *goquery.Selection) {
		td := s.Find("td")
		port, err := strconv.Atoi(td.Nodes[1].FirstChild.Data)
		if err != nil {
			return
		}
		var ipModel = storage.NewIP()
		ipModel.IP = td.Nodes[0].FirstChild.Data
		ipModel.Port = port
		ipModel.Protocol = td.Nodes[3].FirstChild.Data
		log.Println(ipModel)
		saveIpToMongo(ipModel)

	})
}
func saveIpToMongo(ipModel *storage.IP) {

	err := ipModel.Save(conn, bson.M{"ip": ipModel.IP})
	if err != nil {
		log.Println(err)
	}

}
