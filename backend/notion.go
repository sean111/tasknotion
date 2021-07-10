package backend

import (
	notion "github.com/oyekanmiayo/go-notion/notion/version1"
	"log"
	"net/http"
	"time"
)
var (
	accessToken string
	notionClient *notion.Client
	databaseId string
)

func init() {
	LoadConfig()
	accessToken = config.APIKey
	databaseId = config.DatabaseID
	notionClient = notion.NewClient(http.DefaultClient, accessToken)
}

type Task struct {
	ID int
	Name string
	Status string
	DueDate time.Time
	Priority string
	Notes string
	Modified time.Time
}

func Test() {
	records, err := GetRecords()
	for _, record := range records {
		if err != nil {
			log.Fatal(err)
		}
		log.Println(record.Properties)
	}
}

func GetDatabase() (*notion.Database, error) {
	database, _, err := notionClient.Databases.RetrieveDatabase(databaseId)
	if err != nil {
		log.Fatal(err)
	}
	return database, err
}

func GetRecords() ([]notion.Page, error) {
	var params = &notion.QueryDatabaseBodyParams{
		Sorts: []notion.SortCriteria{
			{
				Property: "ID",
				Direction: "ascending",
			},
		},
	}
	results, _, err := notionClient.Databases.QueryDatabase(databaseId, params)
	if err != nil {
		log.Fatal(err)
	}
	return results.Results, err
}
