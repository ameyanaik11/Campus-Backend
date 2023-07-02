package new_exam_results_subscriber

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/TUM-Dev/Campus-Backend/server/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Repository struct {
	DB *gorm.DB
}

func (repository *Repository) FindAllSubscribers() (*[]model.NewExamResultsSubscriber, error) {
	db := repository.DB

	var subscribers []model.NewExamResultsSubscriber

	err := db.Find(&subscribers).Error

	return &subscribers, err
}

func (repository *Repository) NotifySubscriber(subscriber *model.NewExamResultsSubscriber, newGrades *[]model.ExamResultPublished) error {
	url := subscriber.CallbackUrl

	var newExams model.NewExamsPublishedHookPayload

	newExams.PublishedExams = append(newExams.PublishedExams, *newGrades...)

	body, err := json.Marshal(newExams)
	if err != nil {
		log.WithError(err).Errorf("Error while marshalling newGrades")
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")

	if subscriber.ApiKey.Valid {
		req.Header.Set("Authorization", subscriber.ApiKey.String)
	}

	if err != nil {
		log.WithError(err).Errorf("Error while creating request")
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.WithError(err).Errorf("Error while fetching %s", url)
		return err
	}

	subscriber.LastNotifiedAt = sql.NullTime{Time: time.Now(), Valid: true}

	err = repository.DB.Save(subscriber).Error
	if err != nil {
		log.WithError(err).Errorf("Error while saving subscriber")
		return err
	}

	return nil
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
