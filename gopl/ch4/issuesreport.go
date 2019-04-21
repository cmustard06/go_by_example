package main

import (
	"github"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range.Items}}------------------------
Number: {{.Number}}
User:{{.User.Login}}
Title:{{.Title|printf "%.64s"}}
Age: {{.CreateAt|daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)) //向template中映射函数

func noMust() {
	report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	result, err := github.SearchIssues([]string{"repo: golang /go ", "is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}

func main() {
	result, err := github.SearchIssues([]string{"repo: golang /go ", "is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}
