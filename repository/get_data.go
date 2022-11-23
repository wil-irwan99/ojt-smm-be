package repository

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project-ojt/model"
)

type GetDataRepository interface {
	GetJson(id_sensor string, ip string, user string, password string, sdate string, edate string, stime string, etime string) (model.Response, error)
}

type getDataRepository struct {
	dummy string
}

func (g *getDataRepository) GetJson(id_sensor string, ip string, user string, password string, sdate string, edate string, stime string, etime string) (model.Response, error) {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transCfg}
	url_api := fmt.Sprintf("https://%s/api/historicdata.json?id=%s&avg=3600&sdate=%s-%s&edate=%s-%s&usecaption=1&username=%s&password=%s", ip, id_sensor, sdate, stime, edate, etime, user, password)
	resp, err := client.Get(url_api)
	if err != nil {
		return model.Response{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Response{}, err
	}

	var data model.Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func NewGetDataRepository(dummy string) GetDataRepository {
	repo := new(getDataRepository)
	repo.dummy = dummy
	return repo
}
