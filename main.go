package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var data []Antrian

// Antrian describe queue
type Antrian struct {
	ID     string `json:"id"`
	Status bool   `json:"status"`
}

func main() {
	router := gin.Default()
	router.POST("/api/v1/antrian", AddAntrianHandler)
	router.GET("/api/v1/antrian/status", GetAntrianHandler)
	router.PUT("api/v1/antrian/id/:idAntrian", UpdateAntrianHandler)
	router.DELETE("api/v1/antrian/id/:idAntrian/delete", DeleteAntrianHandler)
	router.Run(":8080")
}

func getAntrian() (bool, []Antrian, error) {
	return true, data, nil
}

func addAntrian() (bool, error) {
	_, dataAntrian, _ := getAntrian()
	var ID string

	if dataAntrian == nil {
		ID = fmt.Sprintf("B-0")
	} else {
		ID = fmt.Sprintf("B-%d", len(dataAntrian))
	}

	data = append(data, Antrian{
		ID:     ID,
		Status: false,
	})

	return true, nil
}

func updateAntrian(idAntrian string) (bool, error) {
	for i := range data {
		if data[i].ID == idAntrian {
			data[i].Status = true
			break
		}
	}

	return true, nil
}

func deleteAntrian(idAntrian string) (bool, error) {
	for i := range data {
		if data[i].ID == idAntrian {
			data = append(data[:i], data[i+1:]...)
		}
	}

	return true, nil
}

// AddAntrianHandler is a function to add new antrian
func AddAntrianHandler(c *gin.Context) {
	flag, err := addAntrian()
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// GetAntrianHandler is a function to get all antrian
func GetAntrianHandler(c *gin.Context) {
	flag, data, err := getAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   data,
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// UpdateAntrianHandler is a function to update an antrian
func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := updateAntrian(idAntrian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// DeleteAntrianHandler is a function to delete an antrian
func DeleteAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := deleteAntrian(idAntrian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}
