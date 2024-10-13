package control

func Main() {
	// in go:
	// -- if - else statements are available
	// -- switch (statement) is available
	// -- for (...) is available
	// -- there is no while/do-while loop
	// -- no parentheses are required around the condition

	// there is a special scope for variables declared in the if condition
	// they are only available in the if block and the else block but not outside
	// if (x := 42; x == 42) {...}
	// so x is declared & the condition is if x equals 42

	var user = "admin"
	// switch statement
	switch user {
	case "admin":
		print("You are an admin\n")
	case "user":
		print("You are a user\n")
	default:
		print("Who are you?\n")
	}

	// switch can also not even have a condition
	switch {
	case user == "admin":
		print("You are an admin\n")
	case user == "user": // you can check for anything here not only user
		print("You are a user\n")
	default:
		print("Who are you?\n")
	}

	// for loop (the only loop here)
	// it can does anything literally

	// normal stuff
	for i := 0; i < 5; i++ {
		print(i, " ")
	}

	collection := []int{1, 2, 3, 4, 5}

	// range loop
	// it's like a foreach loop (u receive the index and the value)
	for i, v := range collection {
		print(i, ":", v, " ")
	}

	// u can use it as a while loop
	gameEnded := false
	for !gameEnded {
		// do something ...
		gameEnded = true
	}

	// any condition
	x := 5
	for x < 10 {
		x++ // will go till 10
	}

	// infinite loop
	// for {
	// 	// this is gonna run forever ...
	// }
}
