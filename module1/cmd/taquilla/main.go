package main

import (
	"bufio"
	"fmt"
	"os"
	"reservations"
	"strconv"
	"strings"
)

func main() {
	tstore := reservations.NewStore()
	var n int
	var err error
	var reserv *reservations.Reservation

	tstore.AddProduct(&reservations.Product{
		ID:    "p1",
		Name:  "Product",
		Stock: 10,
	})
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		i := scanner.Text()
		parts := strings.Fields(i)

		if len(parts) == 0 {
			continue
		}

		if parts[0] == "quit" {
			break
		}
		switch parts[0] {
		case "reserve":
			if len(parts) != 3 {
				fmt.Println("error: Invalid parameters")
			} else if n, err = strconv.Atoi(parts[2]); err != nil {
				fmt.Println("error: Invalid number")
			} else {
				reserv, err = tstore.Reserve(parts[1], n)
				if err != nil {
					fmt.Println("error: " + err.Error())
				} else {
					fmt.Println(reserv.ID + " " + reserv.Status.String())
				}
			}
		case "stock":
			if len(parts) != 2 {
				fmt.Println("error: Invalid parameters")
			} else {
				n, err = tstore.GetStock(parts[1])
				if err != nil {
					fmt.Println("error: " + err.Error())
				} else {
					fmt.Println(strconv.Itoa(n))
				}
			}
		case "confirm":
			if len(parts) != 2 {
				fmt.Println("error: Invalid parameters")
			} else {
				err = tstore.Confirm(parts[1])
				if err != nil {
					fmt.Println("error: " + err.Error())
				} else {
					fmt.Println("ok")
				}
			}
		case "cancel":
			if len(parts) != 2 {
				fmt.Println("error: Invalid parameters")
			} else {
				err = tstore.Cancel(parts[1])
				if err != nil {
					fmt.Println("error: " + err.Error())
				} else {
					fmt.Println("ok")
				}
			}
		default:
			fmt.Println("error: Command not found")
		}
	}
}
