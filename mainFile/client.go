package web

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func DoClientIndex(s string) {
	client := http.Client{}
	response, err := client.Get("http://localhost:8000/index?data="+s)
	if err != nil {
		log.Println("get err:", err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ReadAll err:", err)
		return
	}
	log.Println("get data:", string(data))
}

func DoClientSend(s string) (r string) {
	client := http.Client{}
	postData := new(sendData)
	postData.Data = s
	marshal, err := json.Marshal(postData)
	if err != nil {
		log.Println("Marshal err:", err)
		return
	}
	reader := bytes.NewReader(marshal)
	response, err := client.Post("http://localhost:8000/send", "application/json", reader)
	if err != nil {
		log.Println("get err:", err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ReadAll err:", err)
		return
	}
	res := new(respData)
	if err = json.Unmarshal(data, res); err != nil {
		log.Println("unmarshal err:", err.Error())
		return
	}
	if res.Data != nil {
		r = res.Data.(string)
	}
	return r
}
