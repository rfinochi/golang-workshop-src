# Web API

## Hello World API

Para poder crear nuestra primera REST API, necesitamos contar con los siguientes tres elementos:

El primer elemento es un `Handler`. Si venimos de algun entorno MVC, podemos pensar en los handlers como si fueran `controllers`. Son responsables de ejecutar la lógica de la aplicación y de escribir HTTP headers y responses.

El segundo elemento es un enrutador (o `servemux` en la terminología de Go). El `servemux` almacena un mapeo entre los patrones de URL de la aplicación (rutas) y los handlers correspondientes. Go provee un servermux basico, pero en los ejemplos vamos a utilizar el paquete `gorilla/mux` 

Lo último que necesitamos es un servidor web. Una de las mejores cosas de Go es que puede establecer un servidor web y escuchar las solicitudes entrantes como parte de su propia aplicación. No necesita un servidor externo de terceros como Nginx o Apache. 

Con estos tres componentes podemos generar una aplicación funcional:

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//handler
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	//servermux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloWorld)

	//webserver
	log.Fatal(http.ListenAndServe(":8080", router))
}

```

El handler `helloWorld` es una función normal de Go con dos parámetros. El parámetro `http.ResponseWriter` proporciona métodos para construir una respuesta HTTP y enviarla al usuario, y el parámetro `* http.Request` es un puntero a una estructura que contiene información sobre el request actual (por ejemplo el método HTTP y la URL solicitada). 

Cuando se ejecute este código, se va a iniciar un servidor web escuchando en el puerto 8080 de `localhost`. Cada vez que el servidor recibe una nueva solicitud HTTP, pasará la solicitud al servemux y, a su vez, el servemux verificará la ruta de la URL y enviará la solicitud al handler correspondiente. 

Mientras el servidor se está ejecutando, podemos abrir un navegador web y visitar `http://localhost:8080`, para poder obtener la respuesta.

El ejemplo completo se encuentra en la carpeta `helloworld-api`

## REST API

Siguiendo la misma logica anterior, en la carpeta `rest-api`, vamos a encontrar un ejemplo mas complejo, con las siguientes particularidades:

- Usamos el API de `GorillaMux` para enrutar a distintos handlers utilizando tanto la URL como el Verbo HTTP

```go
//request con vervo POST y url /api se ejecutan en el handler create
router.HandleFunc("/api", create).Methods("POST")
//request con vervo GET y url /api se ejecutan en el handler getAll
router.HandleFunc("/api", getAll).Methods("GET")
```

- Utilizamos el paquete `encoding/json` para poder leer el payload del request en JSON y transformarlo a una instancia de nuestra estructura de datos. El paquete `io/ioutil` nos provee el metodo `ReadAll` que nos permite leer todo el body del request, para luego poder manejarlo como un JSON

```go
type data struct {
	Num  int    `json:"Num"`
	Text string `json:"Text"`
}

func create(w http.ResponseWriter, r *http.Request) {
	var newData data
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	json.Unmarshal(reqBody, &newData)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newData)
}
```

## Middlewares

Cuando se crea una aplicación web, probablemente exista alguna funcionalidad compartida que se desee reutilizar para muchos (o todos) los `requests HTTP`.
Por ejemplo, es posible que se desee registrar cada `request`, comprimir cada respuesta o verificar un caché antes de pasar el `request` a sus `handlers`.

Una forma habitual de organizar esta funcionalidad compartida es implementarla como `middlewares`. Esencialmente, un `middleware` se trata de código autónomo que actúa de forma independiente sobre un `request` antes o después de los `handlers` de aplicaciones normales. 

El patrón estándar para crear un `middleware` custom es el siguiente:

```go
func myMiddleware(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        // TODO: Execute our middleware logic here...
        next.ServeHTTP(w, r)
    }

    return http.HandlerFunc(fn)
}
```

- La función `myMiddleware()` es esencialmente un `wrapper` sobre el `handler` `next` (El cual representa el siguiente `handler` de la cadena).

- Establece una función `fn` que se cierra sobre el `handler` `next` para formar un `closure`. Cuando se ejecuta `fn`, ejecuta nuestra lógica de middleware y luego transfiere el control al `handler` `next` llamando al método `ServeHTTP()`.

- Luego convertimos este `closure` en `http.Handler` y lo devolvemos, usando el adaptador `http.HandlerFunc()`.

Un ajuste a este patrón, es usar una función anónima: 

```go
func myMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // TODO: Execute our middleware logic here...
        next.ServeHTTP(w, r)
    })
}
```

Es importante explicar que la ubicación del `middleware` en la cadena de `handlers` afectará el comportamiento de la aplicación.

Si coloca el middleware antes del `servemux` en la cadena, actuará en cada solicitud que reciba la aplicación. 

```
myMiddleware → servemux → application handler
```

Un buen ejemplo de dónde esto sería útil es el `middleware` para registrar `requests`, ya que normalmente es algo que se quiere hacer para todas los `requests`.

Tambien se puede colocar el `middleware` después del `servemux` en la cadena, envolviendo un `handler` de aplicación específico. Esto hace que el middleware solo se ejecute para rutas específicas. 

```
servemux → myMiddleware → application handler
```

Un ejemplo de esto sería algo así como el `middleware` de autorización, que es posible que solo se desee ejecutar en rutas específicas. 