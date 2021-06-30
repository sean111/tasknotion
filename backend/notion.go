package backend

import (
	_ "github.com/joho/godotenv/autoload"
	notion "github.com/oyekanmiayo/go-notion/notion/version1"
	"log"
	"net/http"
	"os"
)
var (
	accessToken = os.Getenv("NOTION_ACCESS_TOKEN")
	notionClient = notion.NewClient(http.DefaultClient, accessToken)
	databaseId = os.Getenv("DATABASE_ID")
)

func GetDatabase() (*notion.Database, error) {
	database, _, err := notionClient.Databases.RetrieveDatabase(databaseId)
	if err != nil {
		log.Fatal(err)
	}
	return database, err
}


