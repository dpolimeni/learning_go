package people

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
	Age  int    `json:"age" xml:"age" form:"age"`
}
