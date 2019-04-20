package main

import (
	"github"
	"html/template"
	"log"
	"net/http"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	result, err := github.SearchIssues([]string{"repo: golang /go ", "is:open", "json", "decoder"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	issueList.Execute(w, result)
	return
}

func main() {

	//if err := issueList.Execute(os.Stdout, result);err!=nil{
	//	log.Fatal(err)
	//}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}
