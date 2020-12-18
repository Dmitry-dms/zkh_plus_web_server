package models

import (
	"errors"
)

type DataVolume struct {
	Id          int     `json:"-" db:"id"`
	UserId      int     `json:"-" db:"user_id"`
	Electricity *string `json:"electricity"  db:"el_volume"`
	Gas         *string `json:"gas"  db:"gas_volume"`
	HotWater    *string `json:"hot_water"  db:"hot_w_volume"`
	ColdWater   *string `json:"cold_water" db:"cold_w_volume"`
	Warming     *string `json:"warming" db:"warming_volume"`
	FullDate    string  `json:"date_full" db:"date_full"`
	Year        string  `json:"date_year" db:"date_year"`
	Month       string  `json:"date_month" db:"date_month"`
	Day         string  `json:"date_day" db:"date_day"`
}
type VolumeResponse struct {
	Electricity float32 `json:"electricity"`
	Gas         float32 `json:"gas"`
	HotWater    float32 `json:"hot_water"`
	ColdWater   float32 `json:"cold_water"`
	Warming     float32 `json:"warming"`
	Summ        float32 `json:"summ"`
}

func (d DataVolume) Validate() error {
	if d.Electricity == nil && d.Gas == nil && d.HotWater == nil && d.ColdWater == nil && d.Warming == nil {
		return errors.New("insert values has no values")
	}
	return nil
}
