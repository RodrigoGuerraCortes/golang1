package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AggregateDataHandler(c *gin.Context) {
	m1Resp, err := http.Get("http://localhost:8081/api/v1/m1")
	if err != nil {
		log.Println("Error calling M1:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call M1"})
		return
	}
	defer m1Resp.Body.Close()
	m1Data, _ := ioutil.ReadAll(m1Resp.Body)

	m2Resp, err := http.Get("http://localhost:8082/api/v1/m2")
	if err != nil {
		log.Println("Error calling M2:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call M2"})
		return
	}
	defer m2Resp.Body.Close()
	m2Data, _ := ioutil.ReadAll(m2Resp.Body)

	c.JSON(http.StatusOK, gin.H{
		"m1": string(m1Data),
		"m2": string(m2Data),
	})
}
