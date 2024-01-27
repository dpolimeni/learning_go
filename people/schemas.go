package people

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}
