package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// データの型を定義
// JSONにparseされた時のkeyの名前も指定
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// サンプルデータ
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	// ":8080"
	router.Run()
}

// 全ての情報を取得
func getAlbums(c *gin.Context) {
	// 構造体をJSONにシリアライズ
	c.IndentedJSON(http.StatusOK, albums)
}

// 情報を追加
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// albumsスライスに追加
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// 特定のアルバムを取得
func getAlbumByID(c *gin.Context) {
	id := c.Params.ByName("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
