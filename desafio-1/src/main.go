package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", StartServer)
	http.ListenAndServe(":8000", nil)
}

func StartServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
	<head>
		<title>Desafio 1 - Imersão Full Cycle</title>
		<style>
		    body {
				background-color: #000000;  
				background-image: url(https://www.meioemensagem.com.br/wp-content/uploads/2017/08/background-cinema-filme1600.jpg);
				background-repeat: no-repeat;
				background-position: center center;
				background-size: cover; 
			}

			h1 {
				color: #FFFFFF;
				text-align:center;
				font-size: 42px;
			}
		</style>
    </head>
	<body>
		<h1>Imersão Full Cycle</h1>
	</body>
</html>
	`)
}
