package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/alexdang1993374/short-link-engine/utils"
	guuid "github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Urls struct {
	ID        string
	ShortUrl  string
	Url       string
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func CreateUrlTable(db *bun.DB, ctx context.Context) {
	_, err := db.NewCreateTable().
		Model((*Urls)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func InsertUrl(db *bun.DB, ctx context.Context, url string) {
	rand.Seed(time.Now().UnixNano())
	shortenedUrl := utils.RandSeq(6)

	urls := Urls{ID: guuid.New().String(), ShortUrl: "http://" + shortenedUrl, Url: url}

	_, err := db.NewInsert().Model(&urls).Exec(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("shortened url: http://" + shortenedUrl)
}

func GetOriginalUrl(db *bun.DB, ctx context.Context, url string) string {
	u := Urls{}

	err := db.NewSelect().Model((*Urls)(nil)).Where("short_url = ?", url).Scan(ctx, &u)

	if err != nil {
		return "not found"
	}

	return u.Url
}

func GetShortUrl(db *bun.DB, ctx context.Context, url string) string {
	u := Urls{}

	err := db.NewSelect().Model((*Urls)(nil)).Where("url = ?", url).Scan(ctx, &u)

	if err != nil {
		return "not found"
	}

	return u.ShortUrl
}
