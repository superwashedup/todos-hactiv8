package entity

import (
	"time"

	"github.com/alvingxv/todos-kelompok5/dto"
)

type Todo struct {
	Id        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"not null;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Todo) TodoToTodoResponses() dto.GetAllTodoResponse {
	return dto.GetAllTodoResponse{
		Id:        t.Id,
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
