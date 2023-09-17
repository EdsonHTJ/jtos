package test

type Friends struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type Data struct {
	About         string    `json:"about"`
	Address       string    `json:"address"`
	Age           int32     `json:"age"`
	Balance       string    `json:"balance"`
	Company       string    `json:"company"`
	Email         string    `json:"email"`
	EyeColor      string    `json:"eyeColor"`
	FavoriteFruit string    `json:"favoriteFruit"`
	Friends       []Friends `json:"friends"`
	Gender        string    `json:"gender"`
	Greeting      string    `json:"greeting"`
	Guid          string    `json:"guid"`
	Id            string    `json:"_id"`
	Index         int32     `json:"index"`
	IsActive      bool      `json:"isActive"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	Picture       string    `json:"picture"`
	Registered    string    `json:"registered"`
	Tags          []string  `json:"tags"`
}

type Person struct {
	Data []Data `json:"data"`
}
