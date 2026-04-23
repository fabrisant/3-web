package main

import (
	"fmt"
	"semana03-taller-relaciones/internal/cafeteria"
)

func main() {
	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()

	cliente1 := cafeteria.Cliente{ID: 1, Nombre: "Ana", Saldo: 15}
	cliente2 := cafeteria.Cliente{ID: 2, Nombre: "Luis", Saldo: 5}
	cliente3 := cafeteria.Cliente{ID: 3, Nombre: "María", Saldo: 3}

	repo.GuardarCliente(cliente1)
	repo.GuardarCliente(cliente2)
	repo.GuardarCliente(cliente3)

	producto1 := cafeteria.Producto{ID: 1, Nombre: "Café", Precio: 1.25, Stock: 20, Categoria: "Bebidas"}
	producto2 := cafeteria.Producto{ID: 2, Nombre: "Té", Precio: 1.00, Stock: 15, Categoria: "Bebidas"}
	producto3 := cafeteria.Producto{ID: 3, Nombre: "Sándwich", Precio: 3.50, Stock: 10, Categoria: "Comida"}

	repo.GuardarProducto(producto1)
	repo.GuardarProducto(producto2)
	repo.GuardarProducto(producto3)

	//btener dicho cliente
	c, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cliente encontrado:", c.Nombre)
	}
	c, err = repo.ObtenerCliente(99)
	if err != nil {
		fmt.Println("Error:", err)
	}

	productos := repo.ListarProductos()

	fmt.Println("\nProductos:")
	for _, p := range productos {
		fmt.Println(p.Nombre, p.Precio)
	}

	// Crear un pedido
	pedido := cafeteria.Pedido{
		ID:       1,
		Cliente:  cliente1,
		Producto: producto1,
		Cantidad: 2,
		Total:    2 * producto1.Precio,
	}

	fmt.Println("\nPedido creado:")
	fmt.Println("Cliente:", pedido.Cliente.Nombre)
	fmt.Println("Producto:", pedido.Producto.Nombre)
	fmt.Println("Cantidad:", pedido.Cantidad)
	fmt.Println("Total:", pedido.Total)

}
// . ¿Tuviste que poner Cliente, Producto y Pedido en el mismo paquete? ¿Por qué sí o por qué no?
// Sí, porque Cliente, Producto y Pedido están relacionados entre sí y forman parte de la misma lógica de negocio de la cafetería.
// y Separarlos en paquetes diferentes podría complicar la gestión de las dependencias y dificultar el mantenimiento del código.
//  Además, al estar en el mismo paquete, se facilita el acceso a los datos y métodos relacionados entre estas entidades, lo que mejora la cohesión del código.

// ¿Qué problema aparecería si intentaras separar Producto en un paquete aparte cuando Pedido lo tiene anidado?
// Si se separara Producto en un paquete aparte, el paquete que contiene Pedido tendría que importar el paquete de Producto para poder usarlo. 
// Esto crearía una dependencia entre los paquetes, lo que podría complicar la estructura del proyecto y hacer que sea más difícil de mantener. 

//Comparando con el Día A (donde usamos IDs): ¿qué ventaja tiene el modelo con IDs para organizar el código en paquetes?
// El modelo con IDs permite una mayor flexibilidad en la organización del código en paquetes, ya que las entidades pueden referenciarse entre sí a través de sus IDs 
// sin necesidad de importar directamente los tipos de datos. Esto facilita la separación de responsabilidades y reduce las dependencias entre paquetes, lo que a su vez mejora la modularidad y mantenibilidad del código. En cambio, el modelo con anidamiento directo de tipos puede generar dependencias más fuertes entre los paquetes, dificultando su organización y mantenimiento.