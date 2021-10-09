package exception

type Duplicate struct {
	Error interface{}
}

func PanicDuplicate(id int, msg string) {
	if id > 0 {
		panic(Duplicate{
			Error: msg,
		})
	}
}