package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

//Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.
//
//
//В рамках задания необходимо:
//Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
//Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
//Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
//Реализовать middleware для логирования запросов
//
//
//Методы API:
//POST /create_event
//POST /update_event
//POST /delete_event
//GET /events_for_day
//GET /events_for_week
//GET /events_for_month
//
//
//Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
//В GET методах параметры передаются через queryString, в POST через тело запроса.
//В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."}
//в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.
//
//В рамках задачи необходимо:
//Реализовать все методы.
//Бизнес логика НЕ должна зависеть от кода HTTP сервера.
//В случае ошибки бизнес-логики сервер должен возвращать HTTP 503.
//В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
//В случае остальных ошибок сервер должен возвращать HTTP 500.
//Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.


const (
	noSuchEvent = "no event with such id"
	wrongMethod = "wrong method"
	invalidData = "invalid data"
	invalidDate = "invalid date"
	invalidUser = "invalid user"
	nuSuchEvent = "no such events"
)

// Event событие
type Event struct {
	ID   int       `json:"id"`
	UserID int		`json:"user_id"`
	Data string 	`json:"data"`
	Date time.Time `json:"date"`
}

// repo данные
type repo struct {
	data map[int]*Event
	mu sync.Mutex
}

// Repository интерфейс для работы с repo
type Repository interface {
	// EventsForNDays получение данных по событиям
	EventsForNDays(date time.Time, days int) ([]*Event, error)

	// Обработка событий

	CreateEvent(event *Event) error
	UpdateEvent(event *Event) error
	DeleteEvent(event *Event) error
}

// Respond ответ сервера
type Respond struct {
	Result string `json:"result"`
}

// Error ответ сервера (ошибка)
type Error struct {
	Error string `json:"error"`
}

func (r *repo) EventsForNDays(date time.Time, days int) ([]*Event, error) {
	events := make([]*Event, 0)

	// ищем нужные события
	for _, event := range r.data {
		if event.Date.Sub(date) >= time.Duration(days*time.Now().Day()) {
			events = append(events, event)
		}
	}

	// проверяем есть ли такие события
	if len(events) == 0 {
		return events, errors.New(nuSuchEvent)
	}

	return events, nil
}


func (r *repo) DeleteEvent(event *Event) error {
	e, ok := r.data[event.ID]
	if !ok {
		return errors.New(noSuchEvent)
	}
	if e.UserID != event.UserID {
		return errors.New(invalidUser)
	}

	// удаляем event
	r.mu.Lock()
	delete(r.data, event.ID)
	r.mu.Unlock()

	return nil
}
func (r *repo) CreateEvent(event *Event) error {
	// уникальный ID
	event.ID = len(r.data) + 1

	// добавляем event
	r.mu.Lock()
	r.data[event.ID] = event
	r.mu.Unlock()

	return nil
}
func (r *repo) UpdateEvent(event *Event) error {
	if _, ok := r.data[event.ID]; !ok {
		return errors.New(noSuchEvent)
	}

	// обновляем event
	r.mu.Lock()
	r.data[event.ID] = event
	r.mu.Unlock()

	return nil
}


// ServiceHandler обработка запросов и работа сервиса
type ServiceHandler struct {
	repo Repository
}

func (s *ServiceHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Response(w, Error{Error: wrongMethod}, http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.PostFormValue("user_id"))
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	data := r.PostFormValue("data")
	if len(data) == 0 {
		Response(w, Error{Error:invalidData}, http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", r.PostFormValue("date"))
	if err != nil {
		Response(w, Error{Error:invalidDate}, http.StatusBadRequest)
		return
	}

	if err = s.repo.CreateEvent(&Event{
		UserID: userID,
		Data:   data,
		Date:   date,
	}); err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusServiceUnavailable)
		return
	}

	Response(w, Respond{Result: "event created"}, http.StatusCreated)
}

// UpdateEvent обновить событие
func (s *ServiceHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Response(w, Error{Error:wrongMethod}, http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	eventID, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.PostFormValue("user_id"))
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	data := r.PostFormValue("data")
	if len(data) == 0 {
		Response(w, Error{Error:invalidData}, http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", r.PostFormValue("date"))
	if err != nil {
		Response(w, Error{Error:invalidDate}, http.StatusBadRequest)
		return
	}

	if err = s.repo.UpdateEvent(&Event{
		ID: eventID,
		UserID: userID,
		Data:   data,
		Date:   date,
	}); err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusServiceUnavailable)
		return
	}

	Response(w, Respond{Result: "event updated"}, http.StatusOK)
}

// DeleteEvent удалить событие
func (s *ServiceHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Response(w, Error{Error:wrongMethod}, http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	eventID, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.PostFormValue("user_id"))
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	if err = s.repo.DeleteEvent(&Event{
		ID: eventID,
		UserID: userID,
	}); err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusServiceUnavailable)
		return
	}

	Response(w, Respond{Result: "event deleted"}, http.StatusOK)
}

func (s *ServiceHandler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Response(w, Error{Error:wrongMethod}, http.StatusBadRequest)
		return
	}
	date, err := handleEventsForNDays(r)
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	events, err := s.repo.EventsForNDays(date, 0)
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusServiceUnavailable)
		return
	}
	Response(w, events, http.StatusOK)
}

func (s *ServiceHandler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Response(w, Error{Error:wrongMethod}, http.StatusBadRequest)
		return
	}
	date, err := handleEventsForNDays(r)
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	events, err := s.repo.EventsForNDays(date, 7)
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusServiceUnavailable)
		return
	}
	Response(w, events, http.StatusOK)
}

func (s *ServiceHandler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Response(w, Error{Error:wrongMethod}, http.StatusBadRequest)
		return
	}
	date, err := handleEventsForNDays(r)
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusBadRequest)
		return
	}

	events, err := s.repo.EventsForNDays(date, 30)
	if err != nil {
		Response(w, Error{Error:err.Error()}, http.StatusServiceUnavailable)
		return
	}
	Response(w, events, http.StatusOK)
}

func handleEventsForNDays(r *http.Request) (time.Time, error) {
	date, err := time.Parse(
		"2006-01-02",
		r.URL.Query().Get("date"),
	)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}


// loggingMiddleware логгирование
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func Response(w http.ResponseWriter, data interface{}, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
	}
}


func main() {
	repository := &repo{
		data: make(map[int]*Event),
		mu:   sync.Mutex{},
	}
	serviceHandler := &ServiceHandler{repo: repository}

	serv := http.DefaultServeMux

	// обработка событий
	serv.HandleFunc("/create_event", serviceHandler.CreateEvent)
	serv.HandleFunc("/update_event", serviceHandler.UpdateEvent)
	serv.HandleFunc("/delete_event", serviceHandler.DeleteEvent)

	// получение данных по событиям
	serv.HandleFunc("/events_for_day", serviceHandler.EventsForDay)
	serv.HandleFunc("/events_for_week", serviceHandler.EventsForWeek)
	serv.HandleFunc("/events_for_month", serviceHandler.EventsForMonth)

	handler := loggingMiddleware(serv)

	portConfig := ":8080"

	err := http.ListenAndServe(portConfig, handler)
	if err != nil {
		log.Fatalln("Failed listen and serve", err)
	}
}