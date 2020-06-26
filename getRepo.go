package main

import (
  "github.com/levigross/getrequests"
  "log"
  "os"
)

var GITHUB_TOKEN= os.Getenv("GITHUB_TOKEN")
var requestsOptions = &grequests.RequestOptions{Auth:[]string{GITHUB_TOKEN,"x-oauth-basics"}}

type Repo struct {
  ID int `json:"id"`
  Name String `json:"full_name"`
  Forks int `json:"forks"`
  Private bool `json:"private"`
}

func getStats(url string) *grequests.Response {
  resp,err :=grequests.Get(url,requestsOptions)

  if err != nil {
    log.Fatalln("Unable to make request: ",err)
  }
  return resp
}
func main(){
  var repos []Repo
  var repoURL="http://api.github.com/users/ausaafnabi/repos"
  resp :=getStats(repoUrl)
  resp.JSON(&repos)
  log.Println(repos)
}
