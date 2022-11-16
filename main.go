package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/joho/sqltocsv"
	_ "github.com/lib/pq"
	"gitlab.com/eiprice/crawlers/james/crawler"
	"gitlab.com/eiprice/crawlers/james/utils"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func run(db *gorm.DB, lat string, lng string, segment string, store string, category string, scan string) error {
	api := crawler.Crawler{db, segment, store, category, scan, lat, lng}

	log.Printf("Start Crawler from", lat, lng)

	err := api.Start()

	if err != nil {
		return err
	}

	DB, err := sql.Open("postgres", os.Getenv("dsn"))

	if err != nil {
		panic(err)
	}

	defer DB.Close()
	result, err := DB.Query(`Select * from james.products p where p.scan = '` + scan + `'`)

	if err != nil {
		return err
	}

	year, month, day := time.Now().Date()

	csvConverter := sqltocsv.New(result)
	csvConverter.Delimiter = ';'
	csvConverter.WriteFile(`files/` + scan + `_` + lat + `_` + lng + `.csv`)
	key := fmt.Sprint(year) + `/` + fmt.Sprint(int(month)) + `/james/` + fmt.Sprint(day) + `_` + scan + `_` + lat + `_` + lng + `.csv`
	filename := `files/` + scan + `_` + lat + `_` + lng + `.csv`
	utils.UploadCSV(key, filename)

	return nil
}

func main() {

	var lat string
	var lng string
	var segment string
	var store string
	var category string
	var scan string
	var drop string
	var file string
	//900020401
	var sc int64

	flag.StringVar(&lat, "lat", "-23.584293", "Set Lat")
	flag.StringVar(&lng, "lng", "-46.674584", "Set Lng")
	flag.StringVar(&segment, "segment", "", "Set Segment")
	flag.StringVar(&store, "store", "", "Set Store")
	flag.StringVar(&category, "category", "", "Set Category")
	flag.StringVar(&scan, "scan", "1", "Set Scan")
	flag.StringVar(&drop, "drop", "", "Set Drop")
	flag.StringVar(&file, "file", "", "Set File")
	flag.Parse()

	db := utils.ConnectDB(drop)
	if drop == "all" {
		os.Exit(0)
	}

	start := time.Now()
	//Crawler get data from James website

	if file != "" {
		Coords := utils.Read(file)

		for _, item := range Coords {
			sc = sc + 1
			err := run(db, item.Lat, item.Lng, segment, store, category, fmt.Sprint(sc))

			if err != nil {
				continue
			}
		}
		os.Exit(0)
	}

	err := run(db, lat, lng, segment, store, category, scan)
	fmt.Println(`Error:`, err)
	elapsed := time.Since(start)
	log.Printf("Finish Crawler Took %s", elapsed)

}
