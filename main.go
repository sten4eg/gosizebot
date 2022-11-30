package main

import (
	"github.com/google/uuid"
	"log"
	"tgbot/features"
	"tgbot/tg"
	"time"
)

func main() {
	updateId := int64(0)

	for {
		log.Println(time.Now().Unix())
		u, err := tg.GetUpdates(updateId)
		if err != nil {
			log.Fatal(err)
		}
		if !u.Ok {
			log.Fatal(*u.Description, ": ", *u.ErrorCode)
		}

		if len(*u.Result) != 0 {
			for _, update := range *u.Result {
				updateId = update.Id + 1

				if update.InlineQuery != nil {
					//TODO
					go ServeInline(update.InlineQuery.Id, []tg.InlineQueryResultArticle{
						{
							Id:    uuid.New().String(),
							Type:  "article",
							Title: "Узнать размер...",
							InputMessageContent: tg.InputMessageContent{
								MessageText: features.GoSize(),
							},
						},
					})
				}
			}
		} else {
			log.Println("Updates is nil")
		}
	}
}

func ServeInline(qid string, articles []tg.InlineQueryResultArticle) {
	err := tg.SendInlineReplyArticle(qid, articles)
	if err != nil {
		log.Println(err)
	}
}
