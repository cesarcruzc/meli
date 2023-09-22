# Desafío Teórico

## Procesos, hilos y corrutinas

- Un caso en el que usarías procesos para resolver un problema y por qué.
    - Los procesos se aconseja usarlos cuando las tareas a ejecutar no dependen entre si para completarse, un ejemplo
      practico para este caso sería, un servidor web en el que se pueden ejecutar multiles proceesos para gestionar
      solicitudes de usuarios por separado sin generar dependencia.
- Un caso en el que usarías threads para resolver un problema y por qué.
    - Los hilos los usaría cuando necesito realizar distintas tareas al mismo tiempo, sin que exista dependencia entre
      uno u otra, pero se debe tener cuidado ya que una desventaja de los hilos es su alto consumo de recursos.
- Un caso en el que usarías corrutinas para resolver un problema y por qué.
    - Las corrutinas son buenas en los casos donde se deben realizar bastantes tareas concurrentes, pero con poca
      sobrecarga, un ejemplo útil es el challenge actual donde es necesario consultar el api de items para obtener los
      resultados solicitados, dado que cada solicitud no depende una de la otra, se lograron realizar varias consultas
      simultaneas haciedno el uso de gorutinas, las corrutinas son a comparación de los hilos más eficientes en cuanto a
      recursos ya que ellas se suspenden o reanudan automaticamente cuando esperan sin consumir recursos.

## Optimización de recursos del sistema operativo

- Si tuvieras 1.000.000 de elementos y tuvieras que consultar para cada uno de ellos información en una API HTTP. ¿Cómo
  lo harías? Explicar.
  - Lo haría implementando corrutinas y dividiendo las consultas en chunks más pequeños para aprovechar los recursos del sistema.
  - Dividir en partes iguales el total de elementos para procesar en lotes
  - Crearía un channel para coordinar las corrutinas y por cada peticion al api, ir enviando el resultado al canal.
  - Esta implementación nos ayudaría a procesar las solicitudes al API usando el channel para coordinar y no saturar al API con requests simultaneos.

## Análisis de complejidad

- Dados 4 algoritmos A, B, C y D que cumplen la misma funcionalidad, con complejidades O(n2), O(n3), O(2n) y O(n log n),
  respectivamente, ¿Cuál de los algoritmos favorecerías y cuál descartarías en principio? Explicar por qué.
  - Elegiria el algoritmo D con complejidad O(n log n) porque tiene un crecimiento logartmico, lo que significa que se comporta mejor con entradas de datos grandes.
  - Descartari el algoritmo C con complejidad O(2n) porque tiene crecimiento exponencial, volviéndose muy ineficiente a corto plazo a medida que crece. 
- Asume que dispones de dos bases de datos para utilizar en diferentes problemas a resolver. La primera llamada AlfaDB
  tiene una complejidad de O(1) en consulta y O(n2) en escritura. La segunda llamada BetaDB que tiene una complejidad de
  O(log n) tanto para consulta, como para escritura. ¿Describe en forma sucinta, qué casos de uso podrías atacar con
  cada una?
    - AlfaDB (O(1) consulta, O(n2) escritura)
      - Seria util para consultas rápidas.
      - Util para información que no cambie constantemente.
      - Si hay escrituras masivas continuas, se degradaría radpidamente.
    - BetaDB (O(log n) consulta y escritura)
        - Equilibrada para consultas y escrituras
        - Util para información que cambian con cierta frecuencia.
        - No es tan rápida en consultas cómo AlfaDB O(1) en consultas.
    - Usaría AlfaDB para datos muy consultados y que no cambien constantemente, y BetaDB para datos dinámicos con equilibrio entre consultas y escrituras.