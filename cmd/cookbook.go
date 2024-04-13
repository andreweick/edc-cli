package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RecipeSummary []struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Slug           string   `json:"slug"`
	Image          string   `json:"image"`
	Description    string   `json:"description"`
	RecipeCategory []string `json:"recipeCategory"`
	Tags           []string `json:"tags"`
	Rating         any      `json:"rating"`
	DateAdded      string   `json:"dateAdded"`
	DateUpdated    string   `json:"dateUpdated"`
}

type RecipeSummarySlug struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug"`
	DateAdded string `json:"dateAdded"`
}

func SyncCookbook(cookbookApiKey string, connStr string) error {
	fmt.Println("syncing cookbook")
	panic(1)

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://cookbook.eick.com/api/recipes/summary", nil)
	req.Header.Set("Authorization", "Bearer "+cookbookApiKey)

	resp, err := client.Do(req)
	CheckError(err)

	defer resp.Body.Close()

	fmt.Println("Response code: " + resp.Status)

	//responseBody, err := io.ReadAll(resp.Body)

	CheckError(err)

	var recipeSummarySlugArr []RecipeSummarySlug

	//json.Unmarshal(responseBody, &recipeSummarySlugArr)
	json.NewDecoder(resp.Body).Decode(&recipeSummarySlugArr)

	// write it to the database

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Printf("Error opening database: %s\n", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Could not ping database: %s\n", err)
	}

	for _, s := range recipeSummarySlugArr {
		url := "https://cookbook.eick.com/api/media/recipes/" + s.Slug + "/images/original.webp"

		fmt.Printf("Trying to get %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error %s getting url:%s\n", err, url)
			panic(1)
		}

		defer resp.Body.Close()

		outputFileName := "./out/" + s.Slug + ".webp"
		fmt.Printf("Trying to create: %s\n", outputFileName)
		out, err := os.Create(outputFileName)
		if err != nil {
			fmt.Printf("Cannot create file %s\n", outputFileName)
			panic(1)
		}

		defer out.Close()

		_, err1 := io.Copy(out, resp.Body)
		if err1 != nil {
			fmt.Printf("couldn't write file %s\n", outputFileName)
			panic(1)
		}

	}
	//for _, s := range recipeSummarySlugArr {
	//	url := "https://cookbook.eick.com/api/recipes/" + s.Slug
	//	req, err := http.NewRequest("GET", url, nil)
	//	req.Header.Set("Authorization", "Bearer "+cookbookApiKey)
	//
	//	if err != nil {
	//		fmt.Printf("Error is %s\n", err)
	//	}
	//
	//	resp, err := client.Do(req)
	//
	//	if err != nil {
	//		fmt.Printf("Do Error: %s", err)
	//	}
	//
	//	b, err := io.ReadAll(resp.Body)
	//	if err != nil {
	//		fmt.Printf("error reading body: %s\n", err)
	//	}
	//
	//	_, err1 := db.Exec("INSERT INTO recipes (mealie_id, slug, recipe_json, published_at, date_added) VALUES($1, $2, $3, $4) ON CONFLICT(slug) DO NOTHING;", s.ID, s.Slug, string(b), s.DateAdded, s.DateAdded)
	//
	//	//fmt.Printf("Recipe JSON: %s\n", string(b))
	//
	//	if err1 != nil {
	//		fmt.Printf("Database error trying to Exec: %s\n\n", err1)
	//		fmt.Printf("ID: %d, Slug: %s\n", s.ID, s.Slug)
	//		panic(1)
	//	}
	//
	//	fmt.Printf("ID: %d, Slug: %s, Date: %x\n", s.ID, s.Slug, s.DateAdded)
	//}

	return nil
}
