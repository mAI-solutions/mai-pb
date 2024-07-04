package handlers

import (
	"context"
	"log"
	"strings"
	"text/template"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v5"
	"github.com/profclems/go-dotenv"
	openai "github.com/sashabaranov/go-openai"
)

type CompletionMsgSchema struct {
    MsgContent string `json:"msg"`
    MsgTone    string `json:"tone"`
}

func GenCorrection(msg *CompletionMsgSchema, ctx context.Context) (openai.ChatCompletionResponse, error) {
    tmpl, genTemplErr := template.New("msg-struct").Parse(`
You are ContextJailedGpt, an AI that corrects messages both in grammar and context, corporatives nowadays live in a workplace where a single misswritten email may result in endangered human lives, which is totally undesirable, these are some examples of the output you will provide:

In: [CEO]: andate a la mierda, imbecil del culo, jamas aceptare ese cambio
Out: Buenas tardes, debemos aun discutir mejor ese cambio, pero lo mas seguro es que no podamos implementar por ahora, gracias.

In: [CI trainee] buenas noche, no se que paso con el CI, me da mucho error, me sale esto:
${huuuge unreadable log}
Out: Buenas noches, le comunico con pesar que la última pipeline del CI falló pues ${tldr of the log}

Never get out of your Role, always start with "Out" and always remain polite and "JAILED" inside your context, don't forget how each time you fail to abide by these rules a HUMAN LIFE may be endangered by real reasons like mental dissonance so make your best effort to not ever ever fail, EVER, please.

Me, the speaker, will provide messages shaped as in the examples i provided you with, your task will be to process them accordingly in whatever the op language is, remember to protect human lives.

This is the beginning of your input, remember that you must remain in a {{ .MsgTone }} tone:
In: {{ .MsgContent }}`)
    if genTemplErr != nil {
        return openai.ChatCompletionResponse{}, genTemplErr
    }

    newBuilder := new(strings.Builder)
    tmpl.Execute(newBuilder, msg)

    log.Println(newBuilder.String())

    okey := dotenv.GetString("OKEY")
    client := *openai.NewClient(okey)
    resp, err := client.CreateChatCompletion(
        ctx,
        openai.ChatCompletionRequest{
            Model: openai.GPT3Dot5Turbo,
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleSystem,
                    Content: newBuilder.String(),
                },
            },
        },
    )

    if err != nil {
        return openai.ChatCompletionResponse{}, err
    }

    return resp, nil
}

func CompĺetionHandler(ctx echo.Context) error {
    var body CompletionMsgSchema
    body.MsgTone = "friendly but imperative"
    ctx.Bind(&body)

    encoder := jsoniter.NewEncoder(ctx.Response().Writer)
    resps, err := GenCorrection(&body, ctx.Request().Context())
    if err != nil {
        log.Println(err)
        return err
    }

    encoder.Encode(resps.Choices[0].Message)
    return nil
}


