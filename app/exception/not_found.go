package exception

type NotFound struct {
	Error interface{}
}

func PanicNotFound(id int) {
	if id == 0 {
		panic(NotFound{
			Error: "Data Tidak Ditemukan",
		})
	}
}
