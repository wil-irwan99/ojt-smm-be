package model

type Response struct {
	PrtgVersion string                   `json:"prtg-version"`
	TreeSize    int                      `json:"treesize"`
	HistDatas   []map[string]interface{} `json:"histdata"`
}
