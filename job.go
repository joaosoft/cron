package cron

type ListJobs []*Job
type MapJobs map[string]*Job

type Job struct {
	Name     string `json:"name" db.read:"name"`
	Key      string `json:"key" db.read:"key"`
	Running  bool   `json:"running" db.read:"running"`
	Settings string `json:"settings" db.read:"settings"`
	Active   bool   `json:"active" db.read:"active"`
	Tasks    ListTasks
}

func NewJob(name string, key string, settings string, tasks ...ITask) *Job {
	return &Job{
		Name:     name,
		Key:      key,
		Settings: settings,
		Running:  false,
		Active:   true,
		Tasks:    tasks,
	}
}
