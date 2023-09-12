package usecases

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/viniciusbls9/go-movie/pkg/database"
	"github.com/viniciusbls9/go-movie/pkg/utils/json"
)

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	db, err := database.OpenDatabaseConnection()
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to connect to database: %v", err))
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM movies WHERE ID = ?")
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to DELETE movie: %v", err))
	}
	defer stmt.Close()

	_, err = stmt.Exec(movieID)
	if err != nil {
	}

}
