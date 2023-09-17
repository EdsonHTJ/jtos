package test

type Friends struct {
	Id   int32
	Name string
}

type Data struct {
	About         string
	Address       string
	Age           int32
	Balance       string
	Company       string
	Email         string
	EyeColor      string
	FavoriteFruit string
	Friends       []Friends
	Gender        string
	Greeting      string
	Guid          string
	Id            string
	Index         int32
	IsActive      bool
	Latitude      float64
	Longitude     float64
	Name          string
	Phone         string
	Picture       string
	Registered    string
	Tags          []string
}

type Person struct {
	Data []Data
}
