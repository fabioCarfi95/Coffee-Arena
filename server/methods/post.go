package methods

import (
	"fmt"
	"net/http"
)

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "retrieve USER INFO")
	fmt.Println("Endpoint hit: userinfo")
}
