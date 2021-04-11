package jobs

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Service - our comment service
type Service struct {
	DB *gorm.DB
}

// Job Search - struct
type Job struct {
	gorm.Model
	Slug    string
	Body    string
	Author  string
	Created time.Time
}

// Jobervice - Definition 
type JobsService interface {
	GetJob(ID uint) (Job, error)
	GetAllJobs() ([]Job, error)
	PostJob(job Job) (Job, error)
	UpdateJob(ID uint, newJob Job) (Job error)
	DeleteJob(ID uint) error
}

// GetAllJobs - get all jobs from the db
func (s *Service)GetAllJobs()([]Job, error){
	var jobs []Job
	if result := s.DB.Find(&jobs); result.Error != nil {
		return jobs, result.Error
	}
	return jobs, nil
}

// GetJob- retrieves jobss by their ID from the database
func (s *Service) GetJob(ID uint) (Job, error) {
	var job Job
	if result := s.DB.First(&job, ID); result.Error != nil {
		return Job{}, result.Error
	}
	return job, nil
}

// GetCommentsBySlug - retrieves all comments by slug (path - /article/name/)
// func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
// 	var comments []Comment
// 	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
// 		return []Comment{}, result.Error
// 	}
// 	return comments, nil
// }

// JobService - implementation
// PostJob- adds a new job to the database
func (s *Service) PostJob(job Job) (Job, error) {
	if result := s.DB.Save(&job); result.Error != nil {
		return Job{}, result.Error
	}
	return job, nil
}

// UpdateJob- updates a job by ID with new comment info
func (s *Service) UpdateJob(ID uint, newJob Job) (Job, error) {
	job, err := s.GetJob(ID)
	if err != nil {
		return Job{}, err
	}

	if result := s.DB.Model(&job).Updates(newJob); result.Error != nil {
		return Job{}, result.Error
	}

	return job, nil
}

// DeleteJob - deletes a job from the database by ID
func (s *Service) DeleteJob(ID uint) error {
	if result := s.DB.Delete(&Job{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}


// NewService - returns a newjobs service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}