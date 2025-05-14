package handlers

func FindNull(bytes []byte) int {
	for index, val := range bytes {
		if val == 0 {
			return index
		}
	}
	return len(bytes)
}

// Explicación paso a paso:
// Parámetro:

// bytes []byte: Es un slice de bytes que se pasa como entrada a la función.
// Bucle for:

// La función recorre el slice bytes usando un bucle for con el rango (range).
// index: Es el índice actual del byte en el slice.
// val: Es el valor del byte en la posición actual.
// Condición if val == 0:

// Comprueba si el valor del byte actual (val) es igual a 0 (un byte nulo).
// Si encuentra un byte nulo, devuelve el índice (index) donde se encuentra.
// Si no encuentra un byte nulo:

// Si el bucle termina sin encontrar un byte nulo, la función devuelve len(bytes), que es la longitud del slice.
// Propósito:
// Esta función es útil para encontrar el separador nulo (\0) en estructuras binarias, como las entradas de un objeto de tipo árbol (tree) en Git. En estas estructuras, el byte nulo separa los metadatos (como el modo y el nombre) del contenido (como el hash).

// Ejemplo:
// Supongamos que el slice bytes contiene:

// La función recorre los valores:

// Índice 0: val = 100 → No es nulo.
// Índice 1: val = 101 → No es nulo.
// Índice 2: val = 102 → No es nulo.
// Índice 3: val = 0 → Es nulo.
// Devuelve 3, que es el índice del primer byte nulo.

// Si el slice no contiene un byte nulo, por ejemplo:

// La función devolverá len(bytes), que en este caso es 5.

// Uso típico:
// En el contexto de Git, esta función puede usarse para:

// Encontrar el byte nulo que separa el modo y el nombre de un archivo en un objeto de tipo árbol.
// Procesar datos binarios estructurados donde los campos están delimitados por bytes nulos.
