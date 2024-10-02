package main

import (
	

	"github.com/dinhminh0307/Shop-Dev-Ecommerce/internal/routers"

)

func main() {
  r := routers.NewRouter(); // call the function from go
  r.Run(":0307") // listen and serve on 0.0.0.0:0307 (for windows "localhost:0307")
}


