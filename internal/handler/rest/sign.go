package rest

import (
	"certification/internal/consts"
	"certification/internal/handler/responce"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateSign(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		responce.NewErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	file.Read(buffer)

	err = h.sc.Sign.CreateSign(buffer, userId.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CheckMySign(c *gin.Context) {
	userId, ok := c.Get(consts.UserId)
	if !ok {
		responce.NewErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	file.Read(buffer)

	check, err := h.sc.Sign.CheckSign(buffer, userId.(int))
	if !check {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CheckUserSign(c *gin.Context) {
	userIdString := c.Param(consts.UserId)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	file.Read(buffer)

	check, err := h.sc.Sign.CheckSign(buffer, userId)
	if !check {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteSign(c *gin.Context) {
	userId, ok := c.Get(consts.UserId)
	if !ok {
		responce.NewErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer file.Close()

	buffer := make([]byte, fileHeader.Size)
	file.Read(buffer)

	err = h.sc.Sign.DeleteSign(buffer, userId.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
