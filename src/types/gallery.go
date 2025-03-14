package types

type GalleryItem struct {
	Url         string `json:"url"`
	SubmittedBy string `json:"submitted_by"`
}
type GalleryResponse struct {
	Images []GalleryItem `json:"images"`
}
