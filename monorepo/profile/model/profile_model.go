package model

type Education struct {
	ID           uint64  `json:"id"`
	UserID       uint64  `json:"userId,omitempty"`
	Institution  string  `json:"institution"`
	Degree       string  `json:"degree"`
	FieldOfStudy string  `json:"fieldOfStudy"`
	StartYear    int     `json:"startYear"`
	EndYear      *int    `json:"endYear"`
	Description  *string `json:"description"`
}

type Project struct {
	ID          uint64  `json:"id"`
	UserID      uint64  `json:"userId,omitempty"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         *string `json:"url"`
	Year        *int    `json:"year"`
}

type Course struct {
	ID          uint64  `json:"id"`
	UserID      uint64  `json:"userId,omitempty"`
	Title       string  `json:"title"`
	Institution string  `json:"institution"`
	Year        *int    `json:"year"`
	URL         *string `json:"url"`
}

type UserProfile struct {
	ID        uint64      `json:"id"`
	Name      string      `json:"name"`
	Nick      string      `json:"nick"`
	ImageURL  string      `json:"imageUrl"`
	Education []Education `json:"education"`
	Projects  []Project   `json:"projects"`
	Courses   []Course    `json:"courses"`
}
