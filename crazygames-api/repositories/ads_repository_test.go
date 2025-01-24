package repositories

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"testing"

	"crazygames.io/entities"
	"crazygames.io/handler/request"
	"github.com/stretchr/testify/assert"
)

func createAds(i int) (*entities.Ads, error) {
	category, err := createCategory("forgame"+strconv.Itoa(i), true)
	if err != nil {
		log.Fatalf("failed to create category for test: %v", err)
		return nil, err
	}
	game, _ := createGame(&entities.Game{
		GameTitle:     "game title",
		Description:   "description",
		Developer:     "developer",
		ThumbnailURL:  "http://www.thumbnail.url",
		ReleaseDate:   nil,
		Technology:    "technology",
		Rating:        4.5,
		HoverVideoUrl: "http://www.hover.video.url",
		GameURL:       "http://www.game.url",
		PlayCount:     0,
	}, strconv.Itoa(int(category.ID)))

	ads := &entities.Ads{
		ImageUrl: "http://www.image.url",
		Position: uint(rand.Intn(100) + 1),
		GameId:   game.ID,
	}
	err = adsRepo.Create(context.Background(), ads)
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func Test_GetAdsById(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE ads;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	ads, err := createAds(1)
	assert.NoError(t, err, "failed to create ads for test")

	t.Run("get ads by valid ID should succeed", func(t *testing.T) {
		ads, err := adsRepo.GetById(context.Background(), ads.ID)
		assert.NoError(t, err, "failed to fetch ads by ID")
		assert.Equal(t, ads.ImageUrl, "http://www.image.url")
	})

	t.Run("get ads by invalid ID should fail", func(t *testing.T) {
		_, err := adsRepo.GetById(context.Background(), uint(100_000))
		assert.Error(t, err, "failed to fetch ads by ID")
	})
}

func Test_GetAllAds(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE ads;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	for i := 1; i <= 10; i++ {
		_, err := createAds(i)
		assert.NoError(t, err, "failed to create ads for test")
	}

	var query = request.AdsRequestQuery{
		PageNumber: 1,
		PageSize:   10,
	}
	ads, _, err := adsRepo.GetAll(context.Background(), query)
	assert.NoError(t, err, "failed to get all ads for test")

	assert.Equal(t, len(ads), 10)
}

func Test_CreateAds(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE ads;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	t.Run("create ads should succeed", func(t *testing.T) {
		ads, err := createAds(1)
		assert.NoError(t, err, "failed to create ads for test")
		assert.Equal(t, ads.GameId, uint(1))
	})

	t.Run("create ads without GameId should fail", func(t *testing.T) {
		ads := &entities.Ads{
			ImageUrl: "http://www.image.url",
			Position: uint(rand.Intn(100)),
		}
		err := adsRepo.Create(context.Background(), ads)
		assert.Error(t, err, "failed to create ads for test")
	})

	t.Run("create ads without ImageUrl should fail", func(t *testing.T) {
		ads := &entities.Ads{
			GameId:   1,
			Position: uint(rand.Intn(100)),
		}
		err := adsRepo.Create(context.Background(), ads)
		assert.Error(t, err, "failed to create ads for test")
	})

	t.Run("create ads without Position should fail", func(t *testing.T) {
		ads := &entities.Ads{
			GameId:   1,
			ImageUrl: "http://www.image.url",
		}
		err := adsRepo.Create(context.Background(), ads)
		assert.Error(t, err, "failed to create ads for test")
	})
}

func Test_UpdateAds(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE ads;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	ads, err := createAds(1)
	assert.NoError(t, err, "failed to create ads for test")

	t.Run("update ads by valid ID should succeed", func(t *testing.T) {
		ads, err := adsRepo.Update(context.Background(), &entities.Ads{
			ID:       ads.ID,
			ImageUrl: "update image url",
			Position: 10,
		})
		assert.NoError(t, err, "failed to update ads for test")
		assert.Equal(t, ads.ImageUrl, "update image url")
		assert.Equal(t, ads.Position, uint(10))
	})

	t.Run("update ads by invalid ID should fail", func(t *testing.T) {
		_, err := adsRepo.Update(context.Background(), &entities.Ads{
			ID:       uint(100_000),
			ImageUrl: "update image url",
			Position: 10,
		})
		assert.NoError(t, err, "failed to update ads for test")
	})
}

func Test_DeleteAds(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE ads;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	ads, err := createAds(1)
	assert.NoError(t, err, "failed to create ads for test")

	t.Run("delete ads by valid ID should succeed", func(t *testing.T) {
		err := adsRepo.Delete(context.Background(), ads.ID)
		assert.NoError(t, err, "failed to delete ads for test")
	})

	t.Run("delete ads by invalid ID should succeed", func(t *testing.T) {
		err := adsRepo.Delete(context.Background(), uint(100_000))
		assert.NoError(t, err, "failed to delete ads for test")
	})
}
