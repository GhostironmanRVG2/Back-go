package Routes

import (
	"net/http"
	"package/Controllers"
)

//FUNCAO PARA DAR LOAD DESTA ROTA
func ImagemRoute() {

	//IMAGEM ROUTE
	http.HandleFunc("/Imagem", Controllers.ImagemPost)

}
