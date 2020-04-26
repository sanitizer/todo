package model

import (
	"encoding/json"
	"time"
)

type Todo struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status 		string    `json:"status"`
	CreateDt    time.Time `json:"createDate"`
	UpdateDt    time.Time `json:"updateDate"`
	IsDeleted   bool 	  `json:"isDeleted"`
}

func (this *Todo) String() string {
	data, _ := this.ToJson()
	return string(data)
}

func (this *Todo) ToJson() ([]byte, error) {
	return json.Marshal(this)
}

func (this *Todo) FromJson(data []byte) error {
	return json.Unmarshal(data, this)
}
