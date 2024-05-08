package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model        // Это встроенная структура GORM, которая автоматически добавляет поля ID,
	Question   string `json:"question" gorm:"text;not null;default:null"` // поля в бд будут представлены ввиде текстовых столбцов,
	Answer     string `json:"answer" gorm:"text;not null;default:null"`   // ни в одном стоблце не допускается null
}
