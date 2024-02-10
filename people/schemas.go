package people

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Age  int    `json:"age" xml:"age" form:"age"`
}
