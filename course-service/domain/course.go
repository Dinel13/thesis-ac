package domain

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Teacher     int    `json:"teacher"`
	Capacity    int    `json:"capacity"`
	Remain      int    `json:"remain"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
