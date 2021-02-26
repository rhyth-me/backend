package main

import (
	"context"

	"github.com/ScoreMarket/backend/domain/model"
	"github.com/ScoreMarket/backend/utils"
)

func main() {
	ctx := context.Background()
	firestore := utils.InitFirestore()

	var itemID string = utils.RandomString(8)

	recode := model.Item{
		ID: itemID,
		Snippet: model.ItemSnippet{
			Title:        "【プロセカ】『夜に駆ける』【創作譜面】",
			ThumbnailURL: "http://img.youtube.com/vi/zMxR1jZg--U/maxresdefault.jpg",
			MusicTitle:   "夜に駆ける",
			Price:        100,
			Tags:         []string{"プロセカ", "創作譜面"},
		},
		Statistics: model.ItemStatistics{
			Rating:        3.7,
			DownloadCount: 123,
		},
		Author: model.User{
			ID:              "milk_choco",
			DisplayName:     "みるくちょこ",
			ProfileImageURL: "https://pbs.twimg.com/profile_images/1364780894595076099/AQFVWEBa_400x400.png",
			StatusMessage:   "創作譜面",
		},
	}

	firestore.Collection("items").Add(ctx, recode)
}
