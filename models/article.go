package models

type Article struct {
	ID    string `json:"id" firestore:"id"`
	Image string `json:"image" firestore:"image"`
	Sort  int    `json:"sort" firestore:"sort"`
	Title string `json:"title" firestore:"title"`
	URL   string `json:"url" firestore:"url"`
}
