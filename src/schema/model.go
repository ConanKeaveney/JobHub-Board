package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID       string    `bson:"task_id" json:"id"`
	Title    string    `json:"title"`
	DueDate  time.Time `json:"due_date"`
	Complete bool      `json:"complete"`
}

type JobDetails struct {
	Company     string    `json:"company"`
	Title       string    `json:"title"`
	Location    string    `json:"location"`
	Category    string    `json:"category"`
	PostDate    time.Time `json:"post_date"`
	Description string    `json:"description"`
	Experience  string    `json:"experience"`
	URL         string    `json:"url"`
	DateAdded   string    `json:"date_added"`
	Salary      int       `json:"salary"`
	Tasks       []Task
}

type Job struct {
	ID         string `bson:"job_id" json:"id"`
	JobDetails JobDetails
}

type Category struct {
	ID    int    `bson:"category_id" json:"id"`
	Title string `json:"title"`
	Jobs  []Job
}

type Board struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OwnerID    string             `bson:"owner_id,omitempty" json:"ownerid"`
	Title      string             `json:"title"`
	Categories []Category
}
