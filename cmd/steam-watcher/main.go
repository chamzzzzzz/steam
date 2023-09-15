package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/chamzzzzzz/steam"
)

func main() {
	client := steam.Client{}
	p, err := client.RequestGetAppList()
	if err != nil {
		log.Fatalf("request app list fail. err='%v'", err)
	}

	sort.Slice(p.AppList.Apps, func(i, j int) bool {
		return p.AppList.Apps[i].AppID < p.AppList.Apps[j].AppID
	})

	file := "applist.csv"
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("create file fail. err='%v'", err)
		return
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()
	writer.Write([]string{"AppID", "Name"})
	for _, app := range p.AppList.Apps {
		writer.Write([]string{strconv.Itoa(app.AppID), app.Name})
	}
	log.Printf("write file success. file='%s'", file)
}
