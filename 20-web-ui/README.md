# Web UI

## Handlers para Obtener HTML

Desde los handlers de Go, podemos utilizar el paquete `html/template` para emitir como respuesta HTML conformado utilizando templates de las paginas y datos de la aplicación

La funcion `ReadFile` del paquete `ioutil` permite leer informacion.

El siguiente ejemplo lee el contenido del archivo `sample.txt`, chequea que si se producieron errores durante la operacion de lectura, e imprime el contenido del archivo.

```go

type toDoItem struct {
	ID      int
	Content string
	IsDone  bool
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

Primero parseamos los archivos que se van a utilizar para generar la salida HTML, en este caso la `home.page` y el `base.layout`. 
Luego parseamos los archivos utilizando el método `ParseFiles` del paquete `html/template`.
En el caso de ocurrir un error en el parseo, lo obtenemos en la variable `err`, y el resultado del parseo va a estar en la variable `ts` del tipo `*template.Template`

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

Con nuestros templates listos en `ts`, el último paso es ejecutar el metodo `Execute`, al cual le pasamos la instancia actual de `http.ResponseWriter` donde se va a escribir el HTML de respuesta, y los datos que querramos renderizar en el template.

## HTML Templates

Los templates HTML en Go utilizan la nomenclatura {name}.page.tmpl y contienen el markup HTML para describir la página. Dentro de estos templates HTML, cualquier dato dinámico que pase está representado por el caracter `.` carácter (denominado `dot`).

En este caso específico, el tipo de punto subyacente será una estructura `toDoItem`. Cuando el tipo de punto subyacente es una estructura, puede representar  el valor de cualquier campo exportado colocando un punto con el nombre del campo. Entonces, debido a que nuestra estructura `toDoItem` tiene un campo `Content`, podríamos mostrar el contenido escribiendo `{{.Content}}` en nuestros templates.

```go
{{template "base" .}}

{{define "title"}}ToDo Item #{{.ID}}{{end}}

{{define "main"}}
<div>
    <div>
        <strong>{{.Content}}</strong>
    </div>
    <pre><code>Done: {{.IsDone}}</code></pre>
</div>
{{end}}
```

Este template declara la estructura HTML de los grupos `title` y `main`, los  cuales estan definidos en `base.layout.tmpl`, template que tambien se suma al grupo de templates para parsear y generar la salida HTML, el cual provee la estructura que se utiliza en todo el sitio. El código de este template se puede encontrar en la carpeta `views` para poder entender la estructura y como se complementan ambos archivos para generar la salida HTML.

#### Templates anidados

Cuando invoca una plantilla desde otra plantilla, el punto debe pasarse  explícitamente a la plantilla que se invoca.

```go
{{template "base" .}}
{{template "main" .}}
{{template "footer" .}}
{{block "sidebar" .}}{{end}}
```

#### Invocación de metodos en los templates

Si el objeto que está obteniendo tiene métodos definidos, puede llamarlos (siempre que se exporten y devuelvan solo un valor, o un solo valor y un error).

Por ejemplo, si .Created tiene el tipo subyacente time.Time podría representar el nombre del día de la semana llamando a su método Weekday()

```go
<span>{{.Created.Weekday}}</span>
```
También puede pasar parámetros a métodos. Por ejemplo, podría usar el método AddDate() para agregar seis meses a la fecha

```go
<span>{{.Snippet.Created.AddDate 0 6 0}}</span>
```

#### Acciones y funciones de los Templates

Ya hemos hablado sobre algunas de las acciones ({{define}}, {{template}} y {{block}}, pero hay tres más que puede usar para controlar la visualización de datos dinámicos: {{if}} , {{with}} y {{range}}. 

Acción  | Descripción
------------- | -------------
{{if .Foo}} C1 {{else}} C2 {{end}}  | Si .Foo no está vacío, renderice el contenido C1, de lo contrario renderice el contenido C2
{{with .Foo}} C1 {{else}} C2 {{end}}  | Si .Foo no está vacío, establezca el punto en el valor de .Foo y renderice el contenido C1; de lo contrario, renderice el contenido C2
{{range .Foo}} C1 {{else}} C2 {{end}}  | If the length of .Foo is greater than zero then loop over each element, setting dot to the value of each element and rendering the content C1. If the length of .Foo is zero then render the content C2

Hay algunas cosas sobre estas acciones para señalar:

- Para las tres acciones, la cláusula {{else}} es opcional. Por ejemplo, puede escribir {{if .Foo}} C1 {{end}} si no hay contenido C2 que desee renderizar.

- Los valores vacíos son falso, 0, cualquier puntero nulo o valor de interfaz y cualquier matriz, sector, mapa o cadena de longitud cero.

 - Es importante comprender que las acciones `with` y `range` cambian el valor del punto. Una vez que se comience a usarlo, lo que representa el punto puede ser diferente dependiendo de dónde se encuentre en la plantilla y de lo que esté haciendo. 

El paquete `html/template` también proporciona algunas funciones que se pueden usar para agregar lógica adicional a los templates y controlar lo que se representa en tiempo de ejecución.

Acción  | Descripción
------------- | -------------
{{eq .Foo .Bar}}  | Verdadero si .Foo es igual a .Bar
{{ne .Foo .Bar}}  | Verdadero si .Foo no es igual a .Bar 
{{not .Foo}} | la negación booleana de .Foo 
{{or .Foo .Bar}} | .Foo si .Foo no está vacío; de lo contrario, da como resultado .Bar
{{index .Foo i}} | el valor de .Foo en el índice i. El tipo subyacente de .Foo debe ser un mapa, un sector o una matriz
{{printf "%s-%s" .Foo .Bar}} | Una cadena formateada que contiene los valores .Foo y .Bar. Funciona de la misma forma que fmt.Sprintf ()
{{len .Foo}} | la longitud de .Foo como un número entero
{{$bar := len .Foo}} | Asignar la longitud de .Foo a la variable de plantilla $ bar 
