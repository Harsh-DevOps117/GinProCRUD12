package service

import (
	"errors"

	"github.com/harshdevops117/models"
	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func NewNotesService(db *gorm.DB) *NotesService {
	return &NotesService{db: db}
}

func (s *NotesService) CreateNote(userID uint, title, content string) error {
	note := models.Notes{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	return s.db.Create(&note).Error
}

func (s *NotesService) GetNotes(userID uint) ([]models.Notes, error) {
	var notes []models.Notes

	err := s.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&notes).Error

	return notes, err
}

func (s *NotesService) GetNoteByID(userID, noteID uint) (*models.Notes, error) {
	var note models.Notes

	err := s.db.Where("id = ? AND user_id = ?", noteID, userID).
		First(&note).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("note not found")
	}

	return &note, err
}

func (s *NotesService) UpdateNote(userID, noteID uint, title, content string) error {
	result := s.db.Model(&models.Notes{}).
		Where("id = ? AND user_id = ?", noteID, userID).
		Updates(map[string]interface{}{
			"title":   title,
			"content": content,
		})

	if result.RowsAffected == 0 {
		return errors.New("note not found")
	}

	return result.Error
}

func (s *NotesService) DeleteNote(userID, noteID uint) error {
	result := s.db.Where("id = ? AND user_id = ?", noteID, userID).
		Delete(&models.Notes{})

	if result.RowsAffected == 0 {
		return errors.New("note not found")
	}

	return result.Error
}
