package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/chamzzzzzz/steam"
)

var (
	file = "applist.csv"
)

func main() {
	apps, err := read(file)
	if err != nil {
		log.Fatalf("read file fail. err='%v'", err)
	}

	client := steam.Client{}
	p, err := client.RequestGetAppList()
	if err != nil {
		log.Fatalf("request app list fail. err='%v'", err)
	}

	apps = merge(apps, p.AppList.Apps)
	sort.Slice(apps, func(i, j int) bool {
		return apps[i].AppID < apps[j].AppID
	})

	if err := write(file, apps); err != nil {
		log.Fatalf("write file fail. err='%v'", err)
	}
	log.Printf("write file success. file='%s'", file)
}

func write(file string, apps []*steam.App) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()
	writer.Write([]string{"AppID", "Name"})
	for _, app := range apps {
		if err := writer.Write([]string{strconv.Itoa(app.AppID), app.Name}); err != nil {
			return err
		}
	}
	return nil
}

func read(file string) ([]*steam.App, error) {
	f, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	apps := make([]*steam.App, 0, len(records))
	for _, record := range records[1:] {
		appID, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		apps = append(apps, &steam.App{
			AppID: appID,
			Name:  record[1],
		})
	}
	return apps, nil
}

func merge(apps1, apps2 []*steam.App) []*steam.App {
	m := make(map[int]*steam.App)
	for _, app := range apps1 {
		m[app.AppID] = app
	}
	for _, app2 := range apps2 {
		app1 := m[app2.AppID]
		if app1 == nil {
			apps1 = append(apps1, app2)
		} else {
			app1.Name = app2.Name
		}
	}
	return apps1
}
