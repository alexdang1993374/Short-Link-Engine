package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/alexdang1993374/short-link-engine/config"
	"github.com/alexdang1993374/short-link-engine/utils"
	guuid "github.com/google/uuid"
)

type Url struct {
	ID        string
	ShortUrl  string
	Url       string
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func CreateUrlTable() {
	ctx := context.Background()
	db := config.Connect()

	_, err := db.NewCreateTable().
		Model((*Url)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func InsertUrl(url string) {
	ctx := context.Background()
	db := config.Connect()

	rand.Seed(time.Now().UnixNano())
	shortenedUrl := utils.RandSeq(6)

	urls := Url{ID: guuid.New().String(), ShortUrl: "host://" + shortenedUrl, Url: url}

	_, err := db.NewInsert().Model(&urls).Exec(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("shortened url: host://" + shortenedUrl)
}
