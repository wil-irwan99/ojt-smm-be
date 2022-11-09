package repository

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project-ojt/config"
	"project-ojt/model"
)

type GetDataRepository interface {
	GetJson(id_sensor string, sdate string, edate string, stime string, etime string) (model.Response, error)
}

type getDataRepository struct {
	config config.Config
}

func (g *getDataRepository) GetJson(id_sensor string, sdate string, edate string, stime string, etime string) (model.Response, error) {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transCfg}
	url_api := fmt.Sprintf("https://%s/api/historicdata.json?id=%s&avg=3600&sdate=%s-%s&edate=%s-%s&usecaption=1&username=%s&password=%s", g.config.Ip, id_sensor, sdate, stime, edate, etime, g.config.User, g.config.Password)
	resp, err := client.Get(url_api)
	if err != nil {
		return model.Response{}, err
	}

	//4755
	//3330

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

func NewGetDataRepository(config config.Config) GetDataRepository {
	repo := new(getDataRepository)
	repo.config = config
	return repo
}
