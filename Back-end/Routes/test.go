package Routes

import (
	"net/http"
	"package/Controllers"
)

//ROUTES FUNCTION
func LoadTests() {

	//TESTS ROUTE
	http.HandleFunc("/teste", Controllers.Something)

}
