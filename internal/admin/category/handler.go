package category

import (
	"bd-backend/internal/appresult"
	"bd-backend/internal/handlers"
	"bd-backend/pkg/logging"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	testURL = "/"
)

type handler struct {
	repository Repository
	logger     *logging.Logger
}

var myMap = map[int]string{
	1: "Matematika",
	2: "Fizika",
	3: "Inlis dili",
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	router.GET(testURL, h.GetAll)
	router.POST(testURL, h.Create)

}

func (h *handler) GetAll(c *gin.Context) {

	limit := c.Query("limit")
	result, err := h.repository.GetAllData(context.Background(), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrInternalServer)
		return
	}

	successResult := appresult.Success
	successResult.Data = result
	c.JSON(http.StatusOK, successResult)

	return
}

func (h *handler) Create(c *gin.Context) {
	var (
		err error
	)

	body, errBody := io.ReadAll(c.Request.Body)
	if errBody != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	reqDTO := ReqDTO{}
	err = json.Unmarshal(body, &reqDTO)

	dt := testFilter(reqDTO.Value)
	if len(dt) == 0 {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}
	err = h.repository.AddData(context.Background(), dt)

	if err != nil {

		return
	}

	successResult := appresult.Success
	successResult.Data = dt
	c.JSON(http.StatusOK, successResult)
	return
}

func testFilter(values []int) []interface{} {
	var dataslice []interface{}
	m := make(map[int]string)

	for _, x := range values {
		if myMap[x] != "" {
			m[x] = myMap[x]
		}
	}
	if len(m) != 0 {
		dataslice = append(dataslice, m)
	}
	return dataslice

}
