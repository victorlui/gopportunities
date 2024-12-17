package openings

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorlui/gopportunities/config"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	db = config.GetSQLite()

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "ID not found")
		return
	}

	deleteQuery := `DELETE FROM openings WHERE id = ?`

	// Executando a query de delete
	result, err := db.Exec(deleteQuery, id)
	if err != nil {
		logger.ErrF("error deleting opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error deleting opening from database")
		return
	}

	// Verificando se alguma linha foi afetada
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.ErrF("error retrieving rows affected: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "could not verify deletion")
		return
	}

	if rowsAffected == 0 {
		logger.WarningF("no opening found with id: %v", id)
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSucces(ctx, "delete opening", fmt.Sprintf("opening with id %v deleted successfully", id))
}
