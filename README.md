# Guía:  Backtracking

## Ejercicios

Usando los conceptos de _backtracking_ resuelva los siguientes ejercicios.

1. [N Reinas](./nreinas): Modificar la implementación de las N Reinas para que devuelva un arreglo de **todas** las soluciones encontradas. También se deberán actualizar los tests.

2. [Cambio de monedas](./cambio): Implementar una solución para el problema del cambio de monedas usando backtracking. El problema consiste en entregar la menor cantidad de monedas posibles para un monto dado.

3. [Problema de los dados](./dados): Se tienen $n$ dados de $k$ caras cada uno, se desea saber la cantidad de formas de obtener una suma de $x$ puntos al lanzar los $n$ dados.

   Ejemplo: Si tenemos $n=3$ dados de $k=6$ caras y queremos obtener el valor $x=7$ la solución es:

$$\begin{align*}
[(1, 1, 5), (1, 2, 4), (1, 3, 3), (1, 4, 2), (1, 5, 1), \\
(2, 1, 4), (2, 2, 3), (2, 3, 2), (2, 4, 1), (3, 1, 3), \\
(3, 2, 2), (3, 3, 1), (4, 1, 2), (4, 2, 1), (5, 1, 1)]
\end{align*}$$

   En total 15 variantes distintas.

4. [Sudoku](./sudoku): El sudoku es un rompecabezas de lógica que consiste en llenar una cuadrícula de 9x9 con números del 1 al 9, de tal manera que cada fila, cada columna y cada una de las nueve subcuadrículas de 3x3 contengan todos los dígitos del 1 al 9 sin repetir. En general se inicia de un tablero parcialmente lleno.

   Se pide implementar la función `possible` (Factible) del TDA `Sudoku`.

   Para ejecutar el solucionador visual se debe ejecutar:

   ```console
   go run ./sudoku/visual/main.go
   ```

5. [Problema de la mochila (Knapsack)](./mochila/mochila5.go): Este es un problema [NP completo](https://es.wikipedia.org/wiki/NP-completo). Consiste en un problema de optimización combinatoria, donde se espera poder llenar una "mochila" con un peso limitado, por una cantidad de objetos, cada uno con un peso y valor específico, máximizando el valor total almacenado. Los objetos no se pueden fraccionar.

6. [Problema de la mochila modificado](./mochila/mochila6.go): Modificar el problema anterior para que devuelva ademas la lista de objetos a incluir en la mochila.
