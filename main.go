package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/docker/cli/cli/version"
	"github.com/nats-io/jwt/v2"
	"net/http"
)

func main() {
	fmt.Println(version.Version)
	_, _ = jwt.DecorateJWT("token")
	JWTAuthTransport := jira.JWTAuthTransport{nil, "debricked", &jira.BasicAuthTransport{}}
	_, _ = JWTAuthTransport.RoundTrip(nil)
	fs := http.FileServer(http.Dir("./web/public"))
	http.Handle("/", fs)
}
