package openings

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorlui/gopportunities/config"
)

func ListOpeningsHandler(ctx *gin.Context) {
	db = config.GetSQLite()

	listQuery := `SELECT id, role FROM openings`

	rows, err := db.Query(listQuery)

	if err != nil {
		logger.ErrF("erro  fetching opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error fetching opening")
		return
	}

	defer rows.Close()

	var openings []ListOpening

	for rows.Next() {
		var opening ListOpening

		if err := rows.Scan(&opening.ID, &opening.Role); err != nil {
			logger.ErrF("error scanning opening: %v", err.Error())
			sendError(ctx, http.StatusInternalServerError, "error reading opening data")
			return
		}
		openings = append(openings, opening)
	}

	// Verificar se algum item foi encontrado
	if len(openings) == 0 {
		sendError(ctx, http.StatusNotFound, "no openings found")
		return
	}

	sendSucces(ctx, "list openings", openings)
}
