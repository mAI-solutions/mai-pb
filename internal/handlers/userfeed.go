package handlers

import (
	"log"
	"pb-mockup/internal/feed"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v5"
)

type FeedReqSchema struct {
    Timeout int64    `json:"timeout"`
    Url     []string `json:"url"`
}

func FeedHandler(ctx echo.Context) error {
    var body FeedReqSchema
    ctx.Bind(&body)

    posts, err := feed.GetPost(body.Url, body.Timeout, ctx.Request().Context())
    if err != nil {
        log.Println(err)
        return err
    }

    encoder := jsoniter.NewEncoder(ctx.Response().Writer)
    encoder.Encode(posts)
    return nil
}




