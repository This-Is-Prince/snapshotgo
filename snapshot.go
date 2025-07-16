package snapshotgo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type SnapshotHub struct {
	Url          string
	IsLimited    bool
	InitialBurst int
	RateLimit    rate.Limit
	Limiter      *rate.Limiter
	Client       *http.Client
}

func NewSnapshot() *SnapshotHub {
	initialBurst := 1
	rateLimit := rate.Every(2 * time.Second)
	limiter := rate.NewLimiter(rateLimit, initialBurst)
	client := &http.Client{Timeout: 5 * time.Second}

	return &SnapshotHub{
		Url:          "https://hub.snapshot.org/graphql",
		IsLimited:    true,
		InitialBurst: initialBurst,
		RateLimit:    rateLimit,
		Limiter:      limiter,
		Client:       client,
	}
}

func Query[T any](snapshothub *SnapshotHub, args Args, data *T) error {
	if snapshothub.IsLimited {
		err := snapshothub.Limiter.Wait(context.Background())
		if err != nil {
			return err
		}
	}

	reqBody := map[string]any{
		"query":     args.Query,
		"variables": args.Variables,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil
	}

	res, err := snapshothub.Client.Post(snapshothub.Url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("failed to fetch data: " + res.Status)
	}

	_data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var graphqlResponse GraphQLResponse[T]
	err = json.Unmarshal(_data, &graphqlResponse)
	if err != nil {
		return err
	}

	if len(graphqlResponse.Errors) > 0 {
		return errors.New("graphql error: " + graphqlResponse.Errors[0].Message)
	}

	*data = graphqlResponse.Data

	return nil
}
