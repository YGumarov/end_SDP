package main

import (
	"fmt"
	"strings"
)

type FoodItem interface {
	Price() float64
	Description() string
}

type Pizza struct {
	Type string
}

type Burger struct {
	Type string
}

// Strategy
func (p *Pizza) Price() float64 {
	switch p.Type {
	case "Margherita":
		return 10.0
	case "Pepperoni":
		return 12.0
	default:
		return 0.0
	}
}

func (p *Pizza) Description() string {
	return p.Type + " Pizza"
}

// Strategy
func (b *Burger) Price() float64 {
	switch b.Type {
	case "Cheeseburger":
		return 5.0
	case "Chicken Burger":
		return 6.0
	default:
		return 0.0
	}
}

func (b *Burger) Description() string {
	return b.Type + " Burger"
}

// Decorator
type ToppingDecorator struct {
	foodItem FoodItem
	topping  string
	price    float64
}

func (t *ToppingDecorator) Price() float64 {
	return t.foodItem.Price() + t.price
}

func (t *ToppingDecorator) Description() string {
	return t.foodItem.Description() + " with " + t.topping
}

// Factory Method
func CreateFoodItemFactory(foodType string) FoodItem {
	switch foodType {
	case "Pizza":
		return &Pizza{}
	case "Burger":
		return &Burger{}
	default:
		return nil
	}
}

type OrderManager struct {
	orders []FoodItem
}

// Singleton
var instance *OrderManager

func NewOrderManager() *OrderManager {
	if instance == nil {
		instance = &OrderManager{}
	}
	return instance
}

func (om *OrderManager) PlaceOrder(foodItem FoodItem) {
	om.orders = append(om.orders, foodItem)
}

func main() {
	fmt.Println("Welcome to the Food Market!")
	fmt.Println("Choose your food:")
	fmt.Println("1. Pizza")
	fmt.Println("2. Burger")

	var choice int
	if _, err := fmt.Scan(&choice); err != nil {
		fmt.Println("Invalid choice:", err)
		return
	}

	var foodType string
	var foodItem FoodItem

	switch choice {
	case 1:
		foodType = "Pizza"
		fmt.Println("Choose your pizza:")
		fmt.Println("1. Margherita")
		fmt.Println("2. Pepperoni")
		var pizzaChoice int
		if _, err := fmt.Scan(&pizzaChoice); err != nil {
			fmt.Println("Invalid pizza choice:", err)
			return
		}

		switch pizzaChoice {
		case 1:
			foodItem = CreateFoodItemFactory(foodType).(*Pizza)
			foodItem.(*Pizza).Type = "Margherita"
		case 2:
			foodItem = CreateFoodItemFactory(foodType).(*Pizza)
			foodItem.(*Pizza).Type = "Pepperoni"
		default:
			fmt.Println("Invalid pizza choice")
			return
		}
	case 2:
		foodType = "Burger"
		fmt.Println("Choose your burger:")
		fmt.Println("1. Cheese burger")
		fmt.Println("2. Chicken Burger")
		var burgerChoice int
		if _, err := fmt.Scan(&burgerChoice); err != nil {
			fmt.Println("Invalid burger choice:", err)
			return
		}

		switch burgerChoice {
		case 1:
			foodItem = CreateFoodItemFactory(foodType).(*Burger)
			foodItem.(*Burger).Type = "Cheese"
		case 2:
			foodItem = CreateFoodItemFactory(foodType).(*Burger)
			foodItem.(*Burger).Type = "Chicken"
		default:
			fmt.Println("Invalid burger choice")
			return
		}
	default:
		fmt.Println("Invalid choice")
		return
	}

	orderManager := NewOrderManager()

	fmt.Println("Do you want any extra topping? (Y/N)")
	var wantTopping string
	if _, err := fmt.Scan(&wantTopping); err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	wantTopping = strings.TrimSpace(wantTopping)

	if wantTopping == "Y" || wantTopping == "y" {
		fmt.Print("Enter the topping: ")
		var topping string
		if _, err := fmt.Scan(&topping); err != nil {
			fmt.Println("Invalid topping input:", err)
			return
		}
		topping = strings.TrimSpace(topping)

		fmt.Print("Enter the topping price: ")
		var toppingPrice float64
		if _, err := fmt.Scan(&toppingPrice); err != nil {
			fmt.Println("Invalid topping price input:", err)
			return
		}

		foodItemWithTopping := &ToppingDecorator{foodItem, topping, toppingPrice}
		fmt.Printf("You ordered a %s for %f.\n", foodItemWithTopping.Description(), foodItemWithTopping.Price())
		orderManager.PlaceOrder(foodItemWithTopping)
	} else {
		fmt.Printf("You ordered a %s for %f.\n", foodItem.Description(), foodItem.Price())
	}
}
