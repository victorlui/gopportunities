package openings

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorlui/gopportunities/config"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	db = config.GetSQLite()

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "ID not provided")
		return
	}

	query := `SELECT id, role FROM openings WHERE id = ?`

	var opening ListOpening

	err := db.QueryRow(query, id).Scan(&opening.ID, &opening.Role)

	if err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSucces(ctx, "get opening", opening)

}
