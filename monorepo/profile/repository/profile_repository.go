package repository

import (
	"database/sql"
	"v1/profile/model"
)

type ProfileRepository interface {
	GetProfile(userID uint64) (model.UserProfile, error)
	AddEducation(userID uint64, e model.Education) (uint64, error)
	UpdateEducation(userID, id uint64, e model.Education) error
	DeleteEducation(userID, id uint64) error
	AddProject(userID uint64, p model.Project) (uint64, error)
	UpdateProject(userID, id uint64, p model.Project) error
	DeleteProject(userID, id uint64) error
	AddCourse(userID uint64, c model.Course) (uint64, error)
	UpdateCourse(userID, id uint64, c model.Course) error
	DeleteCourse(userID, id uint64) error
}

type profileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) GetProfile(userID uint64) (model.UserProfile, error) {
	var p model.UserProfile
	row := r.db.QueryRow(
		`SELECT id, name, nick, IFNULL(image_url,'') FROM users WHERE id = ?`, userID,
	)
	if err := row.Scan(&p.ID, &p.Name, &p.Nick, &p.ImageURL); err != nil {
		return p, err
	}

	edRows, err := r.db.Query(
		`SELECT id, institution, degree, field_of_study, start_year, end_year, description
		 FROM user_education WHERE user_id = ? ORDER BY start_year DESC`, userID,
	)
	if err != nil {
		return p, err
	}
	defer edRows.Close()
	for edRows.Next() {
		var e model.Education
		if err := edRows.Scan(&e.ID, &e.Institution, &e.Degree, &e.FieldOfStudy,
			&e.StartYear, &e.EndYear, &e.Description); err != nil {
			return p, err
		}
		p.Education = append(p.Education, e)
	}
	if p.Education == nil {
		p.Education = []model.Education{}
	}

	prRows, err := r.db.Query(
		`SELECT id, title, description, url, year
		 FROM user_projects WHERE user_id = ? ORDER BY year DESC, id DESC`, userID,
	)
	if err != nil {
		return p, err
	}
	defer prRows.Close()
	for prRows.Next() {
		var pr model.Project
		if err := prRows.Scan(&pr.ID, &pr.Title, &pr.Description, &pr.URL, &pr.Year); err != nil {
			return p, err
		}
		p.Projects = append(p.Projects, pr)
	}
	if p.Projects == nil {
		p.Projects = []model.Project{}
	}

	coRows, err := r.db.Query(
		`SELECT id, title, institution, year, url
		 FROM user_courses WHERE user_id = ? ORDER BY year DESC, id DESC`, userID,
	)
	if err != nil {
		return p, err
	}
	defer coRows.Close()
	for coRows.Next() {
		var c model.Course
		if err := coRows.Scan(&c.ID, &c.Title, &c.Institution, &c.Year, &c.URL); err != nil {
			return p, err
		}
		p.Courses = append(p.Courses, c)
	}
	if p.Courses == nil {
		p.Courses = []model.Course{}
	}

	return p, nil
}

func (r *profileRepository) AddEducation(userID uint64, e model.Education) (uint64, error) {
	res, err := r.db.Exec(
		`INSERT INTO user_education (user_id, institution, degree, field_of_study, start_year, end_year, description)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		userID, e.Institution, e.Degree, e.FieldOfStudy, e.StartYear, e.EndYear, e.Description,
	)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return uint64(id), nil
}

func (r *profileRepository) UpdateEducation(userID, id uint64, e model.Education) error {
	_, err := r.db.Exec(
		`UPDATE user_education SET institution=?, degree=?, field_of_study=?, start_year=?, end_year=?, description=?
		 WHERE id=? AND user_id=?`,
		e.Institution, e.Degree, e.FieldOfStudy, e.StartYear, e.EndYear, e.Description, id, userID,
	)
	return err
}

func (r *profileRepository) DeleteEducation(userID, id uint64) error {
	_, err := r.db.Exec(`DELETE FROM user_education WHERE id=? AND user_id=?`, id, userID)
	return err
}

func (r *profileRepository) AddProject(userID uint64, p model.Project) (uint64, error) {
	res, err := r.db.Exec(
		`INSERT INTO user_projects (user_id, title, description, url, year) VALUES (?, ?, ?, ?, ?)`,
		userID, p.Title, p.Description, p.URL, p.Year,
	)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return uint64(id), nil
}

func (r *profileRepository) UpdateProject(userID, id uint64, p model.Project) error {
	_, err := r.db.Exec(
		`UPDATE user_projects SET title=?, description=?, url=?, year=? WHERE id=? AND user_id=?`,
		p.Title, p.Description, p.URL, p.Year, id, userID,
	)
	return err
}

func (r *profileRepository) DeleteProject(userID, id uint64) error {
	_, err := r.db.Exec(`DELETE FROM user_projects WHERE id=? AND user_id=?`, id, userID)
	return err
}

func (r *profileRepository) AddCourse(userID uint64, c model.Course) (uint64, error) {
	res, err := r.db.Exec(
		`INSERT INTO user_courses (user_id, title, institution, year, url) VALUES (?, ?, ?, ?, ?)`,
		userID, c.Title, c.Institution, c.Year, c.URL,
	)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return uint64(id), nil
}

func (r *profileRepository) UpdateCourse(userID, id uint64, c model.Course) error {
	_, err := r.db.Exec(
		`UPDATE user_courses SET title=?, institution=?, year=?, url=? WHERE id=? AND user_id=?`,
		c.Title, c.Institution, c.Year, c.URL, id, userID,
	)
	return err
}

func (r *profileRepository) DeleteCourse(userID, id uint64) error {
	_, err := r.db.Exec(`DELETE FROM user_courses WHERE id=? AND user_id=?`, id, userID)
	return err
}
