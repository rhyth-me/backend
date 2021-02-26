package main

import (
	"context"
	"log"

	"github.com/ScoreMarket/backend/domain/model"
	"github.com/ScoreMarket/backend/utils"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	firestore := utils.InitFirestore()

	uidObj, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	recode := model.Item{
		ID: uidObj.String(),
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
	}

	firestore.Collection("items").Doc(recode.ID).Set(ctx, recode)
}
