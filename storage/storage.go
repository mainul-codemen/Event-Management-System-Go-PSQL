package storage

import "time"

type (
	User struct {
		ID        int32     `db:"id"`
		FirstName string    `db:"first_name"`
		LastName  string    `db:"last_name"`
		Username  string    `db:"username"`
		Email     string    `db:"email"`
		IsActive  bool      `db:"is_active"`
		IsAdmin   bool      `db:"is_admin"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	EventType struct {
		ID            int32     `db:"id"`
		EventTypeName string    `db:"event_type_name"`
		CreatedAt     time.Time `db:"created_at"`
		UpdatedAt     time.Time `db:"updated_at"`
	}

	Speakers struct {
		ID        int32     `db:"id"`
		FirstName string    `db:"first_name"`
		LastName  string    `db:"last_name"`
		Phone     string    `db:"phone"`
		Username  string    `db:"username"`
		Email     string    `db:"email"`
		Address   string    `db:"address"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	Events struct {
		ID             int32      `db:"id"`
		EventTypeId    int32      `db:"event_type_id"`
		NumberOfGuest  int32      `db:"number_of_guest"`
		StartTime      time.Time  `db:"start_time"`
		EndTime        time.Time  `db:"end_time"`
		EventDate      time.Month `db:"event_date"`
		PerPersonPrice int32      `db:"per_person_price"`
		SpeakerId      int32      `db:"speaker_id"`
		Status         bool       `db:"address"`
		CreatedAt      time.Time  `db:"created_at"`
		UpdatedAt      time.Time  `db:"updated_at"`
	}
)
