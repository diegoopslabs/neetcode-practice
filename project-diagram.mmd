# Diagramas del Proyecto Golang Palindrome Sandbox

## Diagrama de Arquitectura del Sistema

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

## Flujo de Funcionamiento

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

## Diagrama de Flujo - Algoritmo Iterativo

```mermaid
flowchart TD
    Start([Inicio: IsPalindrome]) --> Normalize[Normalizar string<br/>normalizeString]
    Normalize --> Init[Inicializar:<br/>left = 0<br/>right = len-1]
    Init --> Loop{left < right?}
    Loop -->|Sí| Compare{s[left] == s[right]?}
    Compare -->|Sí| Increment[left++<br/>right--]
    Increment --> Loop
    Compare -->|No| ReturnFalse([Return false])
    Loop -->|No| ReturnTrue([Return true])

    style Start fill:#90EE90
    style ReturnTrue fill:#90EE90
    style ReturnFalse fill:#FFB6C1
```

## Diagrama de Flujo - Algoritmo Recursivo

```mermaid
flowchart TD
    Start([Inicio: IsPalindromeRecursive]) --> Normalize[Normalizar string<br/>normalizeString]
    Normalize --> CallHelper[Llamar isPalindromeHelper<br/>left=0, right=len-1]
    CallHelper --> BaseCase{left >= right?}
    BaseCase -->|Sí| ReturnTrue1([Return true<br/>Caso base])
    BaseCase -->|No| Compare{s[left] == s[right]?}
    Compare -->|Sí| Recurse[Llamada recursiva<br/>left+1, right-1]
    Recurse --> BaseCase
    Compare -->|No| ReturnFalse([Return false])

    style Start fill:#90EE90
    style ReturnTrue1 fill:#90EE90
    style ReturnFalse fill:#FFB6C1
    style Recurse fill:#FFE4B5
```

## Diagrama de Secuencia - Ejemplo de Ejecución

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

## Diagrama de Clases/Estructura

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

    note for palindrome "Dos implementaciones:<br/>1. Iterativa (dos punteros)<br/>2. Recursiva (divide y conquista)"
```

## Estructura del Proyecto

```mermaid
graph LR
    A[newproject-claude/] --> B[go.mod]
    A --> C[main.go]
    A --> D[palindrome.go]
    A --> E[project-diagram.mmd]

    C --> F[9 casos de prueba]
    D --> G[4 funciones]

    style A fill:#FFE4B5
    style C fill:#E1F5FF
    style D fill:#E1FFE1
```
