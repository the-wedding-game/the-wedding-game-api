package models

import (
	"errors"
	"testing"
	test "the-wedding-game-api/_tests"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
)

var (
	testGalleryItem1 = types.GalleryItem{Url: "https://example.com/image1.jpg", SubmittedBy: "user1"}
	testGalleryItem2 = types.GalleryItem{Url: "https://example.com/image3.jpg", SubmittedBy: "user3"}
)

func TestGetGalleryImages(t *testing.T) {
	test.SetupMockDb()

	gallery, err := GetGalleryImages()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	if len(gallery) != 2 {
		t.Errorf("expected %d but got %d", 2, len(gallery))
	}

	if gallery[0].Url != testGalleryItem1.Url {
		t.Errorf("expected %s but got %s", testGalleryItem1.Url, gallery[0].Url)
	}

	if gallery[0].SubmittedBy != testGalleryItem1.SubmittedBy {
		t.Errorf("expected %s but got %s", testGalleryItem1.SubmittedBy, gallery[0].SubmittedBy)
	}

	if gallery[1].Url != testGalleryItem2.Url {
		t.Errorf("expected %s but got %s", testGalleryItem2.Url, gallery[1].Url)
	}

	if gallery[1].SubmittedBy != testGalleryItem2.SubmittedBy {
		t.Errorf("expected %s but got %s", testGalleryItem2.SubmittedBy, gallery[1].SubmittedBy)
	}
}

func TestGetGalleryImagesError(t *testing.T) {
	mockDb := test.SetupMockDb()
	mockDb.Error = errors.New("test_error")

	_, err := GetGalleryImages()
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsDatabaseError(err) {
		t.Errorf("expected true but got false")
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}
