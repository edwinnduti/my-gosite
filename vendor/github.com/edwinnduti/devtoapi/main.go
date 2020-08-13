/*
Author :Edwin Nduti
Date   : Aug 2020

Description :
	 A dev.to consuming API code.
*/

package devtoapi

//imports
import (
	"time"
	"net/http"
	"encoding/json"
	"log"
)

//constants
const (
	baseURL string = "https://dev.to/api/articles?"
	Username string = "username="
)

var client = &http.Client{}


//Get username
func GetUserName(name string) (interface{},error){
	//join to form a FQDN
	URL := baseURL+Username+name

	//make request
	req,err := http.NewRequest("GET",URL,nil)
	Check(err)

	//Get response
	resp,err := client.Do(req)
	Check(err)

	//Close Conn stream
	defer resp.Body.Close()

	//variable to decode to resp to
	var result interface{}

	//return Interface that contains json data
	json.NewDecoder(resp.Body).Decode(&result)

	return result,nil


}

//indented format
func GetUserNameAndIndent(name string) (string,error) {
	//Get all user's data
	dataN,err := GetUserName(name)
	Check(err)

	//Indent it
	fullData,err := json.MarshalIndent(dataN,"","   ")
	Check(err)

	//response
	return string(fullData),nil
}

//Give titles
func GetTitles(name string) ([]string,[]string,error) {

	//Call the get data function
	data,err := GetUserNameAndIndent(name)
	Check(err)

	var extract []Article

        //unmarshal to extract
	err = json.Unmarshal([]byte(data),&extract)
	Check(err)

	var title,description []string
        //traverse the all the responses given
	for _,data := range extract{
		title = append(title,data.Title)
		description  = append(description,data.Description)
	}
	return title,description,nil
}

//err handler
func Check(e error){
	if e!= nil{
		log.Fatalln(e)
	}
}

// Articles is a list of articles.
type Articles struct {
	TypeOf                 string       `json:"type_of"`
	ID                     int          `json:"id"`
	Title                  string       `json:"title"`
	Description            string       `json:"description"`
	CoverImage             string       `json:"cover_image"`
	Published              bool         `json:"published"`
	PublishedAt            time.Time    `json:"published_at"`
	TagList                []string     `json:"tag_list"`
	Tags                   string       `json:"tags"`
	Slug                   string       `json:"slug"`
	Path                   string       `json:"path"`
	URL                    string       `json:"url"`
	CanonicalURL           string       `json:"canonical_url"`
	CommentsCount          int          `json:"comments_count"`
	PositiveReactionsCount int          `json:"positive_reactions_count"`
	PublishedTimestamp     time.Time    `json:"published_timestamp"`
	User                   User         `json:"user"`
	Organization           Organization `json:"organization,omitempty"`
	FlareTag               FlareTag     `json:"flare_tag,omitempty"`
}

// Article just an article
type Article struct {
	TypeOf                 string       `json:"type_of"`
	ID                     int          `json:"id"`
	Title                  string       `json:"title"`
	Description            string       `json:"description"`
	CoverImage             string       `json:"cover_image"`
	Published              bool         `json:"published"`
	PublishedAt            time.Time    `json:"published_at"`
	TagList                []string       `json:"tag_list"`
	Tags                   string     `json:"tags"`
	Slug                   string       `json:"slug"`
	Path                   string       `json:"path"`
	URL                    string       `json:"url"`
	CanonicalURL           string       `json:"canonical_url"`
	CommentsCount          int          `json:"comments_count"`
	PositiveReactionsCount int          `json:"positive_reactions_count"`
	PublishedTimestamp     time.Time    `json:"published_timestamp"`
	User                   User         `json:"user"`
	Organization           Organization `json:"organization,omitempty"`
	FlareTag               FlareTag     `json:"flare_tag,omitempty"`
}

// User struct contains information about user
type User struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	WebsiteURL      string `json:"website_url"`
	ProfileImage    string `json:"profile_image"`    // 640x640
	ProfileImage90  string `json:"profile_image_90"` // 90x90
}

// Organization struct contains information about organization (may be empty)
type Organization struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Slug           string `json:"slug"`
	ProfileImage   string `json:"profile_image"`    // 640x640
	ProfileImage90 string `json:"profile_image_90"` // 90x90
}

// FlareTag struct contains information about flare tag (may be empty)
type FlareTag struct {
	Name         string `json:"name"`
	BgColorHex   string `json:"bg_color_hex"`
	TextColorHex string `json:"text_color_hex"`
}

