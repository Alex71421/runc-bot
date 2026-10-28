// Package event описывает забеги БС и статус регистрации на них
package event

import (
	"errors"
	"fmt"
	"strings"
)

// Status - статус регистрации на забег
type Status string

const (
	// StatusUnknown - константа для отображения неизвестного статуса
	StatusUnknown Status = "unknown"
	// StatusClosed - константа для отображения статуса закрытой регистрации
	StatusClosed Status = "closed"
	// StatusOpen - константа для отображения статуса открытой регистрации
	StatusOpen Status = "open"
)

// closedPhrases - список фраз событий у которых нет регистрации
var closedPhrases = []string{"регистрация позднее", "регистрация закрыта"}

// openPhrases - список фраз событий у которых открыта регистрация
var openPhrases = []string{"регистрация"}

// Event - конструкция с данными для конкретного события
type Event struct {
	// ID - поддомен события
	ID string
	// Title - название события, как оно написано на сайте
	Title string
	// DateText - дата события
	DateText string
	// Distances - дистанции события
	Distances []string
	// URL - ссылка на событие
	URL string
	// Status - статус события
	Status Status
}

// New - конструктор для структуры Event
func New(id, title, url string) (Event, error) {
	id = strings.TrimSpace(id)
	title = strings.TrimSpace(title)
	if len(id) == 0 {
		return Event{}, errors.New("домент ивента пустой")
	}
	if len(title) == 0 {
		return Event{}, errors.New("название забега пустое")
	}
	if !strings.HasPrefix(url, "https://") {
		return Event{}, fmt.Errorf("некорректная ссылка: %s", url)
	}
	return Event{ID: id, Title: title, URL: url, Status: StatusUnknown}, nil
}

// IsOpen - функция для проверки открыт ли забег для регистрации
func (e Event) IsOpen() bool {
	return e.Status == StatusOpen
}

// String - текстовый вывод информации о событии
func (e Event) String() string {
	return fmt.Sprintf("%s (%s) - %s", e.Title, e.DateText, e.Status.Label())
}

// Label - возвращает строку статуса в понятно для человека виде
func (s Status) Label() string {
	switch s {
	case StatusOpen:
		return "регистрация открыта"
	case StatusClosed:
		return "регистрация закрыта"
	default:
		return "статус неизвестен"
	}
}

// containsAny - проверяет, что строка text содержит одну из подстрок массива
func containsAny(text string, phrases []string) bool {
	for i := 0; i < len(phrases); i++ {
		if strings.Contains(text, phrases[i]) {
			return true
		}
	}
	return false
}

// ParseStatus - получает статус о забеге исходя из текста о нем
func ParseStatus(text string) Status {
	text = strings.TrimSpace(strings.ToLower(text))
	if containsAny(text, closedPhrases) {
		return StatusClosed
	} else if containsAny(text, openPhrases) {
		return StatusOpen
	} else {
		return StatusUnknown
	}
}
