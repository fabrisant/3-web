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
