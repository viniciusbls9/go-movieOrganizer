package usecases

import (
	"fmt"
	"net/http"

	"github.com/viniciusbls9/go-movie/pkg/database"
	"github.com/viniciusbls9/go-movie/pkg/models"
	"github.com/viniciusbls9/go-movie/pkg/utils/json"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	var result []models.Movie
	db, err := database.OpenDatabaseConnection()
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't open database connection: %v", err))
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT ID, Title, Genre, Watched FROM movies")
	if err != nil {
		json.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't query DB: %v", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Watched); err != nil {
			json.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't scan rows DB: %v", err))
			return
		}
		result = append(result, movie)
	}
	json.RespondWithJSON(w, http.StatusOK, result)
}
