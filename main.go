package main

import (
	"encoding/json"
	"errors"
	"fmt"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/jinzhu/configor"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)
type OpenIP struct {
	Query string `json:"query"`
}
type Info struct {
	Access_key_id string
	Access_key_secret string
	RecordId string
	Domain_prefix string
	Path string
}
var Config = struct {
	Info
}{}

func main() {
	configor.Load(&Config, GetAppPath()+"/config.yml")
	err := _main(Config.Info)
	fmt.Println(err)

}
func _main (info Info) (_err error) {
	LocalIP :=getOpenIP()
	if(checkIP(LocalIP)==true){
		client, _err := CreateClient(tea.String(info.Access_key_id), tea.String(info.Access_key_secret))
		if _err != nil {
			return _err
		}

		updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
			RecordId: tea.String(info.RecordId),
			RR: tea.String(info.Domain_prefix),
			Type: tea.String("A"),
			Value: tea.String(LocalIP),
		}
		_, _err = client.UpdateDomainRecord(updateDomainRecordRequest)
		if _err != nil {
			return _err
		}
		UpdateIP(LocalIP)
		return errors.New("   NEW LOCAL IP IS:"+LocalIP)
	}else{
		return errors.New("   IP NOT CHANGE :"+LocalIP)
	}

}


func getOpenIP() string{
	API_URL:="http://ip-api.com/json/"
	jsonstr := Get(API_URL)
	var Ip OpenIP
	json.Unmarshal(jsonstr, &Ip)
	return Ip.Query
}
func checkIP(ip string)bool{
	content ,_ :=ioutil.ReadFile(GetAppPath()+"/ip.log")
	fmt.Println(ip+"==="+string(content))
	if(ip==string(content)){
		return false
	}else{
		//ioutil.WriteFile("ip.log", []byte(ip), 0666)
		return true
	}
}

func UpdateIP(ip string){
      ioutil.WriteFile(GetAppPath()+"/ip.log", []byte(ip),0666)
}
func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *alidns20150109.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String("alidns.cn-hangzhou.aliyuncs.com")
	_result = &alidns20150109.Client{}
	_result, _err = alidns20150109.NewClient(config)
	return _result, _err
}

func randUA() string {
	var ua = make(map[int]string)
	ua[0] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36"
	ua[1] = "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50"
	ua[2] = "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50"
	ua[3] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:2.0.1) Gecko/20100101 Firefox/4.0.1"
	ua[4] = "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0"
	ua[5] = "Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11"
	ua[6] = "Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1"
	ua[7] = "Mozilla/5.0 (Linux; Android 4.1.2; Nexus 7 Build/JZ054K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19"
	ua[8] = "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_4 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) CriOS/27.0.1453.10 Mobile/10B350 Safari/8536.25"
	ua[9] = "Mozilla/5.0 (compatible; WOW64; MSIE 10.0; Windows NT 6.2)"
	return ua[rand.Intn(10)]
}
func Get(url string) []byte {
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	userAgent := randUA()
	fmt.Println(userAgent)
	req.Header.Set("user-agent", userAgent)
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	req.Header.Add("Referer", url)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body

}
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
