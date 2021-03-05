# Web UI

## Handlers para Obtener HTML

Desde los handlers de Go, podemos utilizar el paquete `html/template` para emitir como respuesta HTML conformado utilizando templates de las paginas y datos de la aplicación

```go

type toDoItem struct {
    ID      int
    Content string
    IsDone  bool
    Created time.Time
}

var toDoItems = make([]toDoItem, 0, 0)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/home.page.tmpl",
		"./views/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		errorLog.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, toDoItems)
	if err != nil {
		errorLog.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
```
Los pasos a seguir para generar la respuesta HTML con un template y datos de la aplicacion son los siguientes:

Primero parseamos los archivos que se van a utilizar para generar la salida HTML, en este caso la `home.page` y el `base.layout`, utilizando el método `ParseFiles` del paquete `html/template`.
En el caso de ocurrir un error, lo obtendremos en la variable `err`, y el resultado del parseo va a estar en la variable `ts` del tipo `*template.Template`

```go
files := []string{
    "./views/home.page.tmpl",
    "./views/base.layout.tmpl",
}

ts, err := template.ParseFiles(files...)
```

```go
var toDoItems = make([]toDoItem, 0, 0)

err = ts.Execute(w, toDoItems)
```

Con nuestros templates listos en `ts`, el último paso es invocar el metodo `Execute`, al cual le pasamos la instancia actual de `http.ResponseWriter` donde se va a escribir el HTML de respuesta, y los datos que querramos renderizar en el template.

## HTML Templates

Los templates HTML en Go, por convención, utilizan la nomenclatura `{name}.page.tmpl` y contienen el markup HTML para describir la página. Dentro de estos templates HTML, cualquier dato dinámico que pase está representado por el caracter `.` (denominado `dot`).

En este caso específico, el tipo de `dot` será una estructura `toDoItem`. Cuando el tipo `dot` es una estructura, se puede representar  el valor de cualquier campo exportado colocando un punto con el nombre del campo. Por jemjemplo, como la estructura `toDoItem` tiene un campo `Content`, podemos mostrar el contenido escribiendo `{{.Content}}` en nuestros templates.

```go
{{template "base" .}}

{{define "title"}}ToDo Item #{{.ID}}{{end}}

{{define "main"}}
<div>
    <div>
        <strong>{{.Content}}</strong>
    </div>
</div>
{{end}}
```

Este template declara la estructura HTML de las secciones `title` y `main`, las cuales estan definidos en el template `base.layout.tmpl` (el cual provee la estructura que se utiliza en todo el sitio), que tambien se suma al grupo de templates para parsear y generar la salida HTML. El código de este template se puede encontrar en la carpeta `views`, para poder entender la estructura y como se complementan ambos archivos para generar una única salida HTML.

#### Templates anidados

Cuando invoca un template desde otro template, `dot` debe pasarse explícitamente al template que se invoca.

```go
{{template "base" .}}
{{template "main" .}}
{{template "footer" .}}
{{block "sidebar" .}}{{end}}
```

#### Invocación de métodos en los templates

Si el elemento que está representado en `dot` tiene métodos definidos, se pueden invocar (siempre que se exporten y devuelvan solo un valor, o un solo valor y un error).

Por ejemplo, si `.Created` tiene el tipo subyacente `time.Time` se puede representar el nombre del día de la semana llamando a su método `Weekday()`

```go
<span>{{.Created.Weekday}}</span>
```

También puede pasar parámetros a los métodos en los templates. Por ejemplo, al usar el método `AddDate()` para agregar seis meses a la fecha

```go
<span>{{.Created.AddDate 0 6 0}}</span>
```

#### Acciones y funciones de los Templates

Ya hemos hablado sobre algunas de las acciones (`{{define}}`, `{{template}}` y `{{block}}`, pero hay tres más que puede usar para controlar la visualización de datos dinámicos: `{{if}}` , `{{with}}` y `{{range}}`. 

Acción  | Descripción
------------- | -------------
{{if .Foo}} C1 {{else}} C2 {{end}}  | Si .Foo no está vacío, renderiza el contenido C1, de lo contrario renderiza el contenido C2
{{with .Foo}} C1 {{else}} C2 {{end}}  | Si .Foo no está vacío, establece dot en el valor de .Foo y renderiza el contenido C1; de lo contrario, renderiza el contenido C2
{{range .Foo}} C1 {{else}} C2 {{end}}  | Si la longitud de .Foo es mayor que cero, recorre cada elemento, estableciendo dot en el valor de cada elemento y renderizando el contenido C1. Si la longitud de .Foo es cero, renderiza el contenido C2 

Algunos comentarios sobre estas acciones:

- Para las tres acciones, la cláusula `{{else}}` es opcional. Por ejemplo, se puede escribir `{{if .Foo}} C1 {{end}}` si no hay contenido C2 que se desee renderizar.

- Los valores vacíos son falso, 0, cualquier puntero nulo o valor de interfaz y cualquier matriz, sector, mapa o cadena de longitud cero.

 - Es importante comprender que las acciones `with` y `range` cambian el valor de `dot`. Una vez que se comience a usarlo, lo que representa `dot` puede ser diferente dependiendo de dónde se encuentre en el template y de lo que se esté haciendo. 

El paquete `html/template` también proporciona algunas funciones que se pueden usar para agregar lógica adicional a los templates y controlar lo que se representa en tiempo de ejecución.

Acción  | Descripción
------------- | -------------
{{eq .Foo .Bar}}  | Verdadero si .Foo es igual a .Bar
{{ne .Foo .Bar}}  | Verdadero si .Foo no es igual a .Bar 
{{not .Foo}} | La negación booleana de .Foo 
{{or .Foo .Bar}} | .Foo si .Foo no está vacío; de lo contrario, da como resultado .Bar
{{index .Foo i}} | El valor de .Foo en el índice i. El tipo subyacente de .Foo debe ser un mapa, un sector o una matriz
{{printf "%s-%s" .Foo .Bar}} | Una cadena formateada que contiene los valores .Foo y .Bar. Funciona de la misma forma que fmt.Sprintf ()
{{len .Foo}} | La longitud de .Foo como un número entero
{{$bar := len .Foo}} | Asignar la longitud de .Foo a la variable de template $bar 
