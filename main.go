package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("https://mikerybka.com/resume.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var resume Resume
	err = json.Unmarshal(b, &resume)
	if err != nil {
		fmt.Println(err)
		return
	}

	flag.Parse()
	if flag.Arg(0) == "json" {
		resume.WriteJSON(os.Stdout)
	}
}

type Resume struct {
	Basics struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Location struct {
			Address string `json:"address"`
		} `json:"location"`
		Website string `json:"website"`
	} `json:"basics"`
	Links struct {
		GitHub   string `json:"github"`
		Twitter  string `json:"twitter"`
		LinkedIn string `json:"linkedin"`
		Email    string `json:"email"`
	} `json:"links"`
	Education []struct {
		Institution string `json:"institution"`
		Area        string `json:"area"`
		Date        string `json:"date"`
		StudyType   string `json:"studyType"`
	} `json:"education"`
	Work []struct {
		Highlights   []string `json:"highlights"`
		Company      string   `json:"company"`
		Position     string   `json:"position"`
		Date         string   `json:"date"`
		Technologies []string `json:"technologies"`
	} `json:"work"`
	Skills []struct {
		Name     string   `json:"name"`
		Keywords []string `json:"keywords"`
	} `json:"skills"`
	Projects []struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		URL         string   `json:"url"`
		Keywords    []string `json:"keywords"`
	} `json:"projects"`
}

func (r *Resume) WriteJSON(w io.Writer) error {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte{'\n'})
	return err
}
