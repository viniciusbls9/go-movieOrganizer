package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/viniciusbls9/go-movie/pkg/database"
	"github.com/viniciusbls9/go-movie/pkg/models"
	jsonUtils "github.com/viniciusbls9/go-movie/pkg/utils/json"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		jsonUtils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid JSON: %v", err))
		return
	}

	db, err := database.OpenDatabaseConnection()
	if err != nil {
		jsonUtils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to connect to database: %v", err))
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO movies(id, title, genre, watched) VALUES($1, $2, $3, $4)")
	if err != nil {
		jsonUtils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't prepare statement: %v", err))
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid.New(), movie.Title, movie.Genre, false)
	if err != nil {
		jsonUtils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't execute statement: %v", err))
		return
	}

	jsonUtils.RespondWithJSON(w, http.StatusCreated, movie)
}
