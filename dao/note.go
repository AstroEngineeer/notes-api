package dao

import (
	"strings"

	"github.com/vigneshganesan008/notes-api/models"
)

func ListNotes(userId uint64) (notes []models.Note, err error) {

	rows, err := db.Query("SELECT id, user_id, title, content FROM notes WHERE user_id=$1", userId)
	if rows.Err() != nil || err != nil {
		return notes, err
	}

	for rows.Next() {
		var note = models.Note{}
		if err := rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content); err != nil {
			return notes, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func GetNote(id uint64) (note models.Note, err error) {

	if err := db.QueryRow(
		"SELECT id, user_id, title, content FROM notes WHERE id=$2",
		id,
	).Scan(note.ID, note.UserID, note.Title, note.Content); err != nil {
		return note, err
	}

	return note, nil
}

func CreateNote(note models.Note) (id uint64, err error) {

	if err = db.QueryRow(
		"INSERT INTO notes(user_id, title, content) VALUES($1, $2, $3) RETURNING id;",
		note.UserID,
		note.Title,
		note.Content,
	).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func UpdateNote(note models.Note) (affrows int64, err error) {
	queryBegin := "UPDATE notes SET "
	var queryMid []string
	var queryArgs []interface{}
	queryEnd := " WHERE id=?"

	if note.Title != "" {
		queryMid = append(queryMid, "title=?")
		queryArgs = append(queryArgs, note.Title)
	}
	if note.Content != "" {
		queryMid = append(queryMid, "content=?")
		queryArgs = append(queryArgs, note.Content)
	}

	queryArgs = append(queryArgs, note.ID)
	query := queryBegin + strings.Join(queryMid, ",") + queryEnd

	res, err := db.Exec(query, queryArgs...)
	if err != nil {
		return affrows, err
	}

	return res.RowsAffected()
}

func DeleteNote(id uint64) (affrows int64, err error) {

	res, err := db.Exec("DELETE FROM notes WHERE id=$1;", id)
	if err != nil {
		return affrows, err
	}

	return res.RowsAffected()
}
