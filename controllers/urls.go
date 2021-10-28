package controllers

import (
	"context"
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

func CreateUrlTable(db *bun.DB, ctx context.Context) string {
	_, err := db.NewCreateTable().
		Model((*Urls)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		return err.Error()
	} else {
		return "Table Created"
	}
}

func InsertUrl(db *bun.DB, ctx context.Context, url string) string {
	rand.Seed(time.Now().UnixNano())
	shortenedUrl := ("http://" + utils.RandSeq(6))

	urls := Urls{ID: guuid.New().String(), ShortUrl: shortenedUrl, Url: url}

	_, err := db.NewInsert().Model(&urls).Exec(ctx)
	if err != nil {
		return "Error: Row NOT Created"
	} else {
		return shortenedUrl
	}
}

func GetOriginalUrl(db *bun.DB, ctx context.Context, url string) string {
	u := Urls{}

	err := db.NewSelect().Model((*Urls)(nil)).Where("short_url = ?", url).Scan(ctx, &u)

	if err != nil {
		return "Not Found"
	} else {
		return u.Url
	}
}

func GetShortUrl(db *bun.DB, ctx context.Context, url string) string {
	u := Urls{}

	err := db.NewSelect().Model((*Urls)(nil)).Where("url = ?", url).Scan(ctx, &u)

	if err != nil {
		return "Not Found"
	} else {
		return u.ShortUrl
	}
}
