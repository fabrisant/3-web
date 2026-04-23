// =============================================================================
// ARCHIVO DE INICIO — Taller Semana 3, Día B
// Aplicaciones Web II (TDI-601) — Ing. John Cevallos Macías, Mg.
//
// Dominio: Mini-Cafetería Universitaria
//
// Este archivo es el punto de partida para TODOS los estudiantes.
// Es un programa monolítico que funciona correctamente con IDs.
//
// TU MISIÓN HOY:
//   1. Cambiar las relaciones de IDs a structs anidados
//   2. Separar el código en paquetes (internal/cafeteria/)
//   3. Definir una interfaz Repository
//   4. Implementar errores personalizados
//
// Para correr este archivo tal cual:
//   go run main.go
//
// =============================================================================

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// =============================================================================
// ENTIDADES (con IDs — tu trabajo es cambiarlas a relaciones anidadas)
// =============================================================================

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	ClienteID  int // <-- HOY esto cambia a: Cliente Cliente
	ProductoID int // <-- HOY esto cambia a: Producto Producto
	Cantidad   int
	Total      float64
	Fecha      string
}

// =============================================================================
// DATOS INICIALES
// =============================================================================

var clientes []Cliente
var productos []Producto
var pedidos []Pedido
var nextPedidoID = 1

func cargarDatos() {
	clientes = []Cliente{
		{ID: 1, Nombre: "Ana Torres", Carrera: "TI", Saldo: 15.00},
		{ID: 2, Nombre: "Luis Vera", Carrera: "Civil", Saldo: 8.50},
		{ID: 3, Nombre: "María Paz", Carrera: "TI", Saldo: 3.00},
	}

	productos = []Producto{
		{ID: 1, Nombre: "Café", Precio: 1.25, Stock: 20, Categoria: "Bebidas"},
		{ID: 2, Nombre: "Empanada", Precio: 0.75, Stock: 15, Categoria: "Snacks"},
		{ID: 3, Nombre: "Almuerzo del día", Precio: 3.50, Stock: 10, Categoria: "Almuerzos"},
		{ID: 4, Nombre: "Jugo natural", Precio: 1.00, Stock: 12, Categoria: "Bebidas"},
		{ID: 5, Nombre: "Galletas", Precio: 0.50, Stock: 25, Categoria: "Snacks"},
	}
}

// =============================================================================
// FUNCIONES DE BÚSQUEDA (retornan índice, -1 si no existe)
// =============================================================================

func buscarClientePorID(id int) int {
	for i, c := range clientes {
		if c.ID == id {
			return i
		}
	}
	return -1
}

func buscarProductoPorID(id int) int {
	for i, p := range productos {
		if p.ID == id {
			return i
		}
	}
	return -1
}

// =============================================================================
// FUNCIONES DE MUTACIÓN (reciben puntero)
// =============================================================================

func descontarSaldo(cliente *Cliente, monto float64) {
	cliente.Saldo -= monto
}

func descontarStock(producto *Producto, cantidad int) {
	producto.Stock -= cantidad
}

// =============================================================================
// REGISTRAR PEDIDO (cruza las 3 entidades)
// =============================================================================

func registrarPedido(clienteID, productoID, cantidad int, fecha string) bool {
	// Buscar cliente
	idxCliente := buscarClientePorID(clienteID)
	if idxCliente == -1 {
		fmt.Println("Error: cliente no encontrado")
		return false
	}

	// Buscar producto
	idxProducto := buscarProductoPorID(productoID)
	if idxProducto == -1 {
		fmt.Println("Error: producto no encontrado")
		return false
	}

	// Validar stock
	if productos[idxProducto].Stock < cantidad {
		fmt.Println("Error: stock insuficiente")
		return false
	}

	// Calcular total
	total := productos[idxProducto].Precio * float64(cantidad)

	// Validar saldo
	if clientes[idxCliente].Saldo < total {
		fmt.Println("Error: saldo insuficiente")
		return false
	}

	// Descontar
	descontarSaldo(&clientes[idxCliente], total)
	descontarStock(&productos[idxProducto], cantidad)

	// Crear pedido
	pedido := Pedido{
		ID:         nextPedidoID,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      total,
		Fecha:      fecha,
	}
	nextPedidoID++
	pedidos = append(pedidos, pedido)

	fmt.Printf("Pedido #%d registrado: Cliente %d -> Producto %d x%d ($%.2f)\n",
		pedido.ID, pedido.ClienteID, pedido.ProductoID, pedido.Cantidad, pedido.Total)
	return true
}

// =============================================================================
// FUNCIONES DE LISTADO
// =============================================================================

func listarClientes() {
	fmt.Println("\n--- CLIENTES ---")
	for _, c := range clientes {
		fmt.Printf("  [%d] %s (%s) - Saldo: $%.2f\n", c.ID, c.Nombre, c.Carrera, c.Saldo)
	}
}

func listarProductos() {
	fmt.Println("\n--- PRODUCTOS ---")
	for _, p := range productos {
		fmt.Printf("  [%d] %s - $%.2f (stock: %d) [%s]\n",
			p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria)
	}
}

func listarPedidos() {
	fmt.Println("\n--- PEDIDOS ---")
	if len(pedidos) == 0 {
		fmt.Println("  No hay pedidos registrados.")
		return
	}
	for _, ped := range pedidos {
		// Con IDs, necesitamos buscar los nombres manualmente
		nombreCliente := "Desconocido"
		idx := buscarClientePorID(ped.ClienteID)
		if idx != -1 {
			nombreCliente = clientes[idx].Nombre
		}

		nombreProducto := "Desconocido"
		idx = buscarProductoPorID(ped.ProductoID)
		if idx != -1 {
			nombreProducto = productos[idx].Nombre
		}

		fmt.Printf("  #%d: %s -> %s x%d ($%.2f) [%s]\n",
			ped.ID, nombreCliente, nombreProducto,
			ped.Cantidad, ped.Total, ped.Fecha)
	}
}

// =============================================================================
// MENÚ INTERACTIVO
// =============================================================================

func leerEntero(reader *bufio.Reader, mensaje string) int {
	fmt.Print(mensaje)
	texto, _ := reader.ReadString('\n')
	texto = strings.TrimSpace(texto)
	num, err := strconv.Atoi(texto)
	if err != nil {
		return -1
	}
	return num
}

func main() {
	cargarDatos()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n========================================")
		fmt.Println("    MINI-CAFETERÍA UNIVERSITARIA")
		fmt.Println("========================================")
		fmt.Println("  1. Listar clientes")
		fmt.Println("  2. Listar productos")
		fmt.Println("  3. Registrar pedido")
		fmt.Println("  4. Ver pedidos")
		fmt.Println("  5. Salir")
		fmt.Println("========================================")

		opcion := leerEntero(reader, "Elige una opción: ")

		switch opcion {
		case 1:
			listarClientes()
		case 2:
			listarProductos()
		case 3:
			listarClientes()
			clienteID := leerEntero(reader, "ID del cliente: ")
			listarProductos()
			productoID := leerEntero(reader, "ID del producto: ")
			cantidad := leerEntero(reader, "Cantidad: ")
			if clienteID > 0 && productoID > 0 && cantidad > 0 {
				registrarPedido(clienteID, productoID, cantidad, "2026-04-20")
			} else {
				fmt.Println("Datos inválidos.")
			}
		case 4:
			listarPedidos()
		case 5:
			fmt.Println("¡Hasta pronto!")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}
