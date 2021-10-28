package controllers

import (
	"context"
	"testing"

	"github.com/alexdang1993374/short-link-engine/config"
)

func TestCreateUrlTable(t *testing.T) {
	db := config.Connect()
	ctx := context.Background()

	if CreateUrlTable(db, ctx) != "Table Created" {
		t.Errorf("Table NOT Created")
	}
}

func TestInsertUrl(t *testing.T) {
	db := config.Connect()
	ctx := context.Background()

	shortUrl := InsertUrl(db, ctx, "http://test.com")

	if shortUrl == "Error: Row NOT Created" {
		t.Errorf("Row NOT Created")
	} else {
		db.NewDelete().Model((*Urls)(nil)).Where("short_url = ?", shortUrl).Exec(ctx)
	}
}

func TestGetOriginalUrl(t *testing.T) {
	db := config.Connect()
	ctx := context.Background()

	shortUrl := InsertUrl(db, ctx, "http://test.com")

	if GetOriginalUrl(db, ctx, shortUrl) != "http://test.com" {
		t.Errorf("Url NOT Found")
	} else {
		db.NewDelete().Model((*Urls)(nil)).Where("short_url = ?", shortUrl).Exec(ctx)
	}
}

func TestGetShortUrl(t *testing.T) {
	db := config.Connect()
	ctx := context.Background()

	shortUrl := InsertUrl(db, ctx, "http://test.com")

	if GetShortUrl(db, ctx, "http://test.com") != shortUrl {
		t.Errorf("Url NOT Found")
	} else {
		db.NewDelete().Model((*Urls)(nil)).Where("short_url = ?", shortUrl).Exec(ctx)
	}
}
