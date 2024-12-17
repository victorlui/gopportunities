package openings

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorlui/gopportunities/config"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	logger = config.GetLogger("update")
	db = config.GetSQLite()
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "ID not provided")
		return
	}

	request := UpdateOpeningRequest{}

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

	// Query para atualizar o item
	updateQuery := `UPDATE openings SET role = ?, company = ?, location = ?, remote = ?, link = ?, salary = ? WHERE id = ?`

	// Executar a query
	result, err := db.Exec(updateQuery, request.Role, request.Company, request.Location, request.Remote, request.Link, request.Salary, id)

	if err != nil {
		logger.ErrF("error updating opening with id %v: %v", id, err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.ErrF("error retrieving rows affected: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "could not verify update")
		return
	}

	if rowsAffected == 0 {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSucces(ctx, "update opening", request)

}
