package models

import (
	"the-wedding-game-api/types"
	"the-wedding-game-api/utils"
)

func GetGalleryImages() ([]types.GalleryItem, error) {
	conn := GetConnection()
	gallery, err := conn.GetGallery()
	if err != nil {
		return nil, err
	}

	var validGallery = make([]types.GalleryItem, 0)
	for i := range gallery {
		if utils.IsURLStrict(gallery[i].Url) {
			validGallery = append(validGallery, gallery[i])
		}
	}

	return validGallery, nil
}
