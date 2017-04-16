/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-zoo/bone"
	"github.com/mainflux/mainflux-core/models"
	"github.com/mainflux/mainflux-mongodb-reader/db"
	"gopkg.in/mgo.v2/bson"
)

// getMessage function
func getMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	cid := bone.GetValue(r, "channel_id")

	if err := Db.C("channels").Find(bson.M{"id": cid}).One(nil); err != nil {
		w.WriteHeader(http.StatusNotFound)
		str := `{"response": "Channel not found", "id": "` + cid + `"}`
		io.WriteString(w, str)
		return
	}

	// Get fileter values from parameters:
	// - start_time = messages from this moment. UNIX time format.
	// - end_time = messages to this moment. UNIX time format.
	var st float64
	var et float64
	var err error
	var s string
	s = r.URL.Query().Get("start_time")
	if len(s) == 0 {
		st = 0
	} else {
		st, err = strconv.ParseFloat(s, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			str := `{"response": "wrong start_time format"}`
			io.WriteString(w, str)
			return
		}
	}
	s = r.URL.Query().Get("end_time")
	if len(s) == 0 {
		et = float64(time.Now().Unix())
	} else {
		et, err = strconv.ParseFloat(s, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			str := `{"response": "wrong end_time format"}`
			io.WriteString(w, str)
			return
		}
	}

	results := []models.Message{}
	if err := Db.C("messages").Find(bson.M{"channel": cid, "time": bson.M{"$gt": st, "$lt": et}}).
		All(&results); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		str := `{"response": "not found", "id": "` + cid + `"}`
		io.WriteString(w, str)
		return
	}

	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(results)
	if err != nil {
		log.Print(err)
	}
	io.WriteString(w, string(res))
}
