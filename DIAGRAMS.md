# Documentación Visual del Proyecto - Golang Palindrome Sandbox

## Tabla de Contenidos
- [Arquitectura del Sistema](#arquitectura-del-sistema)
- [Flujo de Funcionamiento](#flujo-de-funcionamiento)
- [Algoritmo Iterativo](#algoritmo-iterativo)
- [Algoritmo Recursivo](#algoritmo-recursivo)
- [Diagrama de Secuencia](#diagrama-de-secuencia)
- [Estructura de Clases](#estructura-de-clases)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Resumen Técnico](#resumen-técnico)

---

## Arquitectura del Sistema

Este diagrama muestra la arquitectura completa del proyecto, destacando la separación entre el punto de entrada (`main.go`) y la lógica de negocio (`palindrome.go`).

```mermaid
flowchart TB
    subgraph "main.go - Punto de Entrada"
        A[Casos de Prueba<br/>9 strings diferentes]
        B[main función]
        C[Salida formateada]
    end

    subgraph "palindrome.go - Lógica Principal"
        D[IsPalindrome<br/>Iterativo]
        E[IsPalindromeRecursive<br/>Recursivo]
        F[normalizeString<br/>Normalización]
        G[isPalindromeHelper<br/>Función auxiliar]
    end

    subgraph "Procesamiento"
        H{Normalizar texto<br/>lowercase + alphanumeric}
        I[Comparación<br/>dos punteros]
        J[Comparación<br/>recursiva]
    end

    A --> B
    B -->|llama| D
    B -->|llama| E
    D -->|usa| F
    E -->|usa| F
    F --> H
    D --> I
    E --> G
    G --> J
    I --> C
    J --> C

    style A fill:#e1f5ff
    style D fill:#ffe1e1
    style E fill:#e1ffe1
    style F fill:#fff4e1
```

---

## Flujo de Funcionamiento

Diagrama simplificado que muestra cómo fluyen los datos desde los casos de prueba hasta la salida final.

```mermaid
graph TD
    A["Test Cases<br/>9 diferentes strings"] -->|input| B["main.go"]
    B -->|llama| C["IsPalindrome<br/>Iterativo"]
    B -->|llama| D["IsPalindromeRecursive<br/>Recursivo"]
    C -->|usa| E["normalizeString"]
    D -->|usa| E
    E -->|filtra| F["Solo letras y dígitos<br/>En minúsculas"]
    C -->|enfoque de<br/>dos punteros| G["Comparar desde<br/>izquierda y derecha"]
    D -->|enfoque<br/>recursivo| H["isPalindromeHelper"]
    H -->|compara| G
    G -->|resultado| I["Salida<br/>Tabla Formateada"]
    F --> G
```

---

## Algoritmo Iterativo

**Función:** `IsPalindrome(s string) bool`

**Complejidad:**
- Tiempo: O(n)
- Espacio: O(m) donde m es la longitud del string normalizado

```mermaid
flowchart TD
    Start([Inicio: IsPalindrome]) --> Normalize[Normalizar string<br/>normalizeString]
    Normalize --> Init["Inicializar:<br/>left = 0<br/>right = len-1"]
    Init --> Loop{left < right?}
    Loop -->|Sí| Compare{"s[left] == s[right]?"}
    Compare -->|Sí| Increment["left++<br/>right--"]
    Increment --> Loop
    Compare -->|No| ReturnFalse([Return false])
    Loop -->|No| ReturnTrue([Return true])

    style Start fill:#90EE90
    style ReturnTrue fill:#90EE90
    style ReturnFalse fill:#FFB6C1
```

**Ejemplo:**
```
Input: "A man a plan a canal Panama"
Normalizado: "amanaplanacanalpanama"
Proceso: left=0, right=21 → compara 'a' == 'a' ✓ → continúa...
Resultado: true
```

---

## Algoritmo Recursivo

**Función:** `IsPalindromeRecursive(s string) bool`

**Complejidad:**
- Tiempo: O(n)
- Espacio: O(n) debido al stack de llamadas recursivas

```mermaid
flowchart TD
    Start([Inicio: IsPalindromeRecursive]) --> Normalize[Normalizar string<br/>normalizeString]
    Normalize --> CallHelper["Llamar isPalindromeHelper<br/>left=0, right=len-1"]
    CallHelper --> BaseCase{"left >= right?"}
    BaseCase -->|Sí| ReturnTrue1([Return true<br/>Caso base])
    BaseCase -->|No| Compare{"s[left] == s[right]?"}
    Compare -->|Sí| Recurse["Llamada recursiva<br/>left+1, right-1"]
    Recurse --> BaseCase
    Compare -->|No| ReturnFalse([Return false])

    style Start fill:#90EE90
    style ReturnTrue1 fill:#90EE90
    style ReturnFalse fill:#FFB6C1
    style Recurse fill:#FFE4B5
```

**Ejemplo de llamadas recursivas:**
```
Input: "racecar"
Normalizado: "racecar"
helper(0, 6) → 'r' == 'r' ✓
helper(1, 5) → 'a' == 'a' ✓
helper(2, 4) → 'c' == 'c' ✓
helper(3, 3) → left >= right → true
```

---

## Diagrama de Secuencia

Ejemplo de ejecución con el string "Madam" mostrando la interacción entre componentes.

```mermaid
sequenceDiagram
    participant M as main()
    participant IP as IsPalindrome
    participant IPR as IsPalindromeRecursive
    participant NS as normalizeString
    participant Helper as isPalindromeHelper

    M->>IP: "Madam"
    IP->>NS: "Madam"
    NS-->>IP: "madam"
    IP->>IP: Comparar con dos punteros
    IP-->>M: true

    M->>IPR: "Madam"
    IPR->>NS: "Madam"
    NS-->>IPR: "madam"
    IPR->>Helper: ("madam", 0, 4)
    Helper->>Helper: ("madam", 1, 3)
    Helper->>Helper: ("madam", 2, 2)
    Helper-->>IPR: true
    IPR-->>M: true
```

---

## Estructura de Clases

Representación de la estructura del código y sus relaciones.

```mermaid
classDiagram
    class main {
        +main()
    }

    class palindrome {
        +IsPalindrome(s string) bool
        +IsPalindromeRecursive(s string) bool
        -normalizeString(s string) string
        -isPalindromeHelper(s string, left int, right int) bool
    }

    main --> palindrome: uses

    note for palindrome "Dos implementaciones:\n1. Iterativa (dos punteros)\n2. Recursiva (divide y conquista)"
```

**Funciones Públicas:**
- `IsPalindrome` - Implementación iterativa
- `IsPalindromeRecursive` - Implementación recursiva

**Funciones Privadas:**
- `normalizeString` - Limpia y normaliza el string
- `isPalindromeHelper` - Función auxiliar para recursión

---

## Estructura del Proyecto

```mermaid
graph LR
    A[newproject-claude/] --> B[go.mod]
    A --> C[main.go]
    A --> D[palindrome.go]
    A --> E[DIAGRAMS.md]
    A --> F[project-diagram.mmd]

    C --> G[9 casos de prueba]
    D --> H[4 funciones]

    style A fill:#FFE4B5
    style C fill:#E1F5FF
    style D fill:#E1FFE1
    style E fill:#D8BFD8
```

**Archivos del proyecto:**

| Archivo | Descripción | Líneas |
|---------|-------------|--------|
| `go.mod` | Definición del módulo Go | ~3 |
| `main.go` | Punto de entrada con tests | ~40 |
| `palindrome.go` | Lógica de palíndromos | ~60 |
| `DIAGRAMS.md` | Documentación visual | Este archivo |
| `project-diagram.mmd` | Diagramas en formato .mmd | ~200 |

---

## Resumen Técnico

### Casos de Prueba

El proyecto incluye **9 casos de prueba** que cubren diferentes escenarios:

| # | Input | Esperado | Tipo |
|---|-------|----------|------|
| 1 | `"racecar"` | ✅ true | Palíndromo simple |
| 2 | `"A man a plan a canal Panama"` | ✅ true | Con espacios |
| 3 | `"race a car"` | ❌ false | No palíndromo |
| 4 | `"hello"` | ❌ false | No palíndromo |
| 5 | `"Madam"` | ✅ true | Case-insensitive |
| 6 | `"Was it a car or a cat I saw?"` | ✅ true | Con puntuación |
| 7 | `"No 'x' in Nixon"` | ✅ true | Con comillas |
| 8 | `"12321"` | ✅ true | Numérico |
| 9 | `"12345"` | ❌ false | Numérico no palíndromo |

### Comparación de Algoritmos

| Aspecto | Iterativo | Recursivo |
|---------|-----------|-----------|
| **Complejidad Temporal** | O(n) | O(n) |
| **Complejidad Espacial** | O(m) | O(n) |
| **Stack Overflow** | No | Posible con strings muy largos |
| **Legibilidad** | Alta | Muy alta |
| **Performance** | Ligeramente mejor | Overhead de llamadas |
| **Uso de memoria** | Menor | Mayor (stack) |

### Normalización de Strings

La función `normalizeString` realiza las siguientes transformaciones:

```
Input:  "A man, a plan!"
Step 1: "a man, a plan!"     (lowercase)
Step 2: "amanplan"            (solo alphanumeric)
Output: "amanplan"
```

**Proceso:**
1. Convierte todo a minúsculas
2. Filtra solo caracteres alfanuméricos (a-z, 0-9)
3. Elimina espacios, puntuación y caracteres especiales

---

## Conclusión

Este proyecto demuestra:
- ✅ Implementación de dos paradigmas algorítmicos diferentes
- ✅ Manejo correcto de strings en Go
- ✅ Código limpio y modular
- ✅ Cobertura completa de casos de prueba
- ✅ Documentación clara con diagramas visuales

**Ideal para:** Aprendizaje de Go, algoritmos, y buenas prácticas de programación.
