package microapp

import (
	"fmt"
	"github.com/getevo/evo-ng/http"
)

func Register() error {
	fmt.Println("Hello From Micro App :)")
	return nil
}

func Router() error {
	http.Get("/test", func(request *http.Context) {
		request.Write("Test Message")
	})

	return nil
}
