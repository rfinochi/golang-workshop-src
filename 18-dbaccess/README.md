# Db Access

## SQL Databases

Para poder conectarnos a `SQL Databases` nuestra aplicacion de GO debe tener referencia al paquete `database/sql`, y tenemos que instalar el paquete del driver especifico para el tipo de base de datos que nos vamos a conectar. En nuestros ejemplos utilizamos `mysql`, por lo cual el driver que vamos a instalar es el siguiente `github.com/go-sql-driver/mysql`

Una vez instalado y referenciado los paquetes necesarios, el siguiente paso es crear un pool de conexiones a la base de datos. Para hacer esto, necesitamos la función `sql.Open` del paquete `database/sql`, que utiliza de la siguente manera:

```go
db, err := sql.Open("mysql", "web:pass@/snippetbox?parseTime=true")
if err != nil {
    ...
}

defer db.Close()
```

Algunos puntos a tener en cuenta:

- El primer parámetro de `sql.Open` es el nombre del driver y el segundo parámetro es el nombre de la fuente de datos (a veces también llamado cadena de conexión o DSN) que describe cómo conectarse a la base de datos.

- El formato del nombre de la fuente de datos dependerá de la base de datos y el driver que esté utilizando. Normalmente, se puede encontrar información y ejemplos en la documentación del driver específico.

- La parte `parseTime=true` del DSN anterior es un parámetro específico del driver de `mysql` que indica que se conviertan los campos `SQL TIME` y `DATE` en objetos Go `time.Time`.

- La función `sql.Open` devuelve un objeto `sql.DB`. Esta no es una conexión de base de datos, es un conjunto de muchas conexiones (pool). Ésta es una diferencia importante de comprender. Go administra estas conexiones según sea necesario, abriendo y cerrando automáticamente conexiones a través del driver.

- El pool de conexiones es seguro para el acceso simultáneo.

- El pool de conexiones está destinado a ser de larga duración. En una aplicación es normal inicializar el pool de conexiones en su función `main` y luego pasar el pool de conexiones a sus diferentes funciones. Utilizarlo de otra forma sería un desperdicio de memoria y recursos de red. 

- Incluir la ejecución de  `db.Close` utilizando `defer` debajo de la inicialización del pool de conexiones es un buen hábito, permitiendo que la aplicación cierre sus recursos de manera ordenada.

### Testeando la conexión 

El método `Ping` nos permite probar la conexión. Si todo esta bien, debería poder crear una conexión sin errores. Si en cambio hay un error, puede ser un problema con el DSN o la disponibilidad de la base de datos.

```go
if err = db.Ping(); err != nil {
	return nil, err
}
```

### Ejecución de sentencias SQL 

El paquete `database/sql` de Go proporciona tres métodos diferentes para ejecutar consultas de base de datos:

- `DB.Query` se utiliza para consultas SELECT que devuelven varias filas.
- `DB.QueryRow` se usa para consultas SELECT que devuelven una sola fila.
- `DB.Exec` se usa para declaraciones que no devuelven filas (como INSERT y DELETE).

Las sentencias de SQL a ejecutar se crean utilizando variables de tipo `string` con `?` como placeholders para los valores. 

```sql
INSERT INTO items (title, isDone)
VALUES(?, ?)
```

Si buscamos realizar el Insert, entonces, en nuestro caso, la herramienta más apropiada para el trabajo es el método `DB.Exec`

```go
func CreateItem(newItem Item) int {
	db, err := connnect()

	if err != nil {
		errorLog.Fatal(err)
		return 0
	}

	stmt := `INSERT INTO items (title, isDone)
    VALUES(?, ?)`

	result, err := db.Exec(stmt, newItem.Title, newItem.IsDone)

	if err != nil {
		errorLog.Fatal(err)
		return 0
	}

	id, err := result.LastInsertId()

	if err != nil {
		errorLog.Fatal(err)
		return 0
	}

	return int(id)
}
```

La interfaz `sql.Result` devuelta por `DB.Exec` proporciona dos métodos:

- `LastInsertId`: devuelve el entero generado por la base de datos en respuesta a un comando. Por lo general, esto será de una columna de "incremento automático" al insertar una nueva fila.

- `RowsAffected`: devuelve el número de filas afectadas por la ejecucion de la sentencia. 

Si buscamor realizar un Select, el patrón es diferente. Utilizaremos el método `QueryRow` que nos permite ejecutar consultas de una sola fila con una sentencia SQL select a la cual le vamos a pasar el Id del elemento a obtener, y nos devuelve un puntero a `sql.Row` con el resultado de la consulta.

```go
func GetItem(id int) (item Item) {
	db, err := connnect()

	if err != nil {
		errorLog.Fatal(err)
		return Item{}
	}

	stmt := `SELECT id, title, isDone FROM items
	where id = ?`

	i := Item{}

	err = db.QueryRow(stmt, id).Scan(&i.ID, &i.Title, &i.IsDone)

	if err != nil {
		errorLog.Fatal(err)
		return Item{}
	}
	return
}
```

`sql.Row` posee el metodo `Scan` que se utiliza para copiar los valores de cada campo en `sql.Row` al campo correspondiente en la estructura de destino. Los argumentos que se le pasan a `Scan` son punteros a cada lugar en el que se desea copiar los datos, y el número de argumentos debe ser exactamente el mismo que el número de columnas devueltas por su declaración. 
