package main

func main() {

	r := router()

	err := r.Run(":8080")
	if err != nil {
		//TODO Log error
		panic(err)
	}

}
