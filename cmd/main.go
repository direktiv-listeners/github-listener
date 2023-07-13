package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	cloudevents "github.com/cloudevents/sdk-go"
	cehttp "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v53/github"
	"github.com/google/uuid"
)

const (
	ENV_DEBUG          = "DIREKTIV_GITHUB_DEBUG"
	ENV_GITHUB_TOKEN   = "DIREKTIV_GITHUB_SECRET"
	ENV_DIREKTIV_TOKEN = "DIREKTIV_GITHUB_TOKEN"
	ENV_ENDPOINT       = "DIREKTIV_GITHUB_ENDPOINT"
	ENV_INSECURE       = "DIREKTIV_GITHUB_INSECURE_TLS"
	ENV_PATH           = "DIREKTIV_GITHUB_PATH"

	HEADER_GITHUB_EVENT = "X-Github-Event"
	HEADER_GITHUB_UUID  = "X-GitHub-Delivery"
)

var (
	localGitHubToken, endpoint string
)

func startServer() error {

	gin.SetMode(gin.ReleaseMode)

	// set logging
	debug := os.Getenv(ENV_DEBUG)
	if debug != "" {
		gin.SetMode(gin.DebugMode)
	}

	endpoint = os.Getenv(ENV_ENDPOINT)
	if os.Getenv("K_SINK") != "" {
		endpoint = os.Getenv("K_SINK")
	}

	if endpoint == "" {
		log.Fatal("endpoint for receiver not set")
	}

	log.Printf("using endpoint %s", endpoint)

	localGitHubToken = os.Getenv(ENV_GITHUB_TOKEN)

	path := os.Getenv(ENV_PATH)
	if path == "" {
		path = "/github"
	}

	log.Printf("serving %s", path)

	r := gin.Default()
	r.POST(path, handleRequest)
	return r.Run()

}

func main() {

	log.Println("starting github listener")

	err := startServer()
	if err != nil {
		log.Fatalf("can not start server: %s", err.Error())
	}

}

func handleRequest(c *gin.Context) {

	payload, err := github.ValidatePayload(c.Request, []byte(localGitHubToken))
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		log.Printf("can not validate payload: %v.", err)
		return
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(payload, &data)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		log.Printf("invalid payload: %v.", err)
		return
	}

	eventType := c.GetHeader(HEADER_GITHUB_EVENT)
	log.Printf("event type: %s", eventType)

	id, err := uuid.Parse(c.GetHeader(HEADER_GITHUB_UUID))
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		log.Printf("invalid event ID: %v.", err)
		return
	}
	log.Printf("event id: %s", id)

	ce := cloudevents.NewEvent()

	pr, ok := data["repository"]
	if !ok {
		c.Writer.WriteHeader(http.StatusBadRequest)
		log.Printf("missing repository field")
		return
	}
	project := pr.(map[string]interface{})
	source := fmt.Sprint(project["full_name"])

	nodeid := fmt.Sprint(project["node_id"])

	ce.SetID(id.String())
	ce.SetType(eventType)
	ce.SetSource(source)
	ce.SetTime(time.Now())
	ce.SetDataContentType("application/json")
	ce.SetExtension("nodeid", nodeid)
	err = ce.SetData(data)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		log.Printf("failed to set cloudevent data: %v", err)
		return
	}

	c.Writer.WriteHeader(200)
	go sendEvent(ce, endpoint)

}

func sendEvent(event cloudevents.Event, endpoint string) {

	skipTLS := false
	if os.Getenv(ENV_INSECURE) != "" {
		skipTLS = true
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipTLS},
	}

	options := []cehttp.Option{
		cloudevents.WithTarget(endpoint),
		cloudevents.WithStructuredEncoding(),
		cloudevents.WithHTTPTransport(tr),
	}

	if len(os.Getenv(ENV_DIREKTIV_TOKEN)) > 0 {
		options = append(options,
			cehttp.WithHeader("Direktiv-Token", os.Getenv(ENV_DIREKTIV_TOKEN)))
	}

	t, err := cloudevents.NewHTTPTransport(
		options...,
	)
	if err != nil {
		log.Printf("unable to create transport: %s", err.Error())
		return
	}

	c, err := cloudevents.NewClient(t)
	if err != nil {
		log.Printf("unable to create client: %s", err.Error())
		return
	}

	_, _, err = c.Send(context.Background(), event)
	if err != nil {
		log.Printf("unable to send cloudevent: %s", err.Error())
		return
	}

}
