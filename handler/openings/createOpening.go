package openings

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victorlui/gopportunities/config"
	"github.com/victorlui/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningResponse true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	logger = config.GetLogger("main")
	db = config.GetSQLite()

	request := CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.ErrF("binding error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "invalid request format")
		return
	}

	if err := request.Validate(); err != nil {
		logger.ErrF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.OpeningResponse{
		Role:      request.Role,
		Company:   request.Company,
		Location:  request.Location,
		Remote:    *request.Remote,
		Link:      request.Link,
		Salary:    request.Salary,
		CreatedAt: time.Now(),
	}

	insertQuery := `INSERT INTO openings (role, company, location, remote, link, salary, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

	// Passando os campos individualmente
	_, err := db.Exec(
		insertQuery,
		opening.Role,
		opening.Company,
		opening.Location,
		opening.Remote,
		opening.Link,
		opening.Salary,
		opening.CreatedAt,
	)

	if err != nil {
		logger.ErrF("error creting opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSucces(ctx, "create opening", opening)

}
