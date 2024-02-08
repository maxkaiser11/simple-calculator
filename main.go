package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func main() {
	err := ui.Main(func() {
		db, err := sql.Open("sqlite3", "./users.db")
		if err != nil {
			log.Fatal(err)
		}

		sqlStmt := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			hashed_password TEXT NOT NULL
		);
		`

		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Fatal(err)
		}

		const passwordCost = 14

		usernameEntry := ui.NewEntry()
		passwordEntry := ui.NewPasswordEntry()
		loginButton := ui.NewButton("Login")
		registerButton := ui.NewButton("Register")
		calculationEntry := ui.NewEntry()
		calculationEntry.SetReadOnly(true)
		var currentExpression string

		buttons := [][]string{
			{"7", "8", "9", "/"},
			{"4", "5", "6", "*"},
			{"1", "2", "3", "-"},
			{"0", ".", "=", "+"},
		}

		buttonGrid := ui.NewGrid()

		for i, row := range buttons {
			for j, label := range row {
				button := ui.NewButton(label)
				buttonGrid.Append(button, i, j, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
				button.OnClicked(func(*ui.Button) {
					handleButtonClick(label, &currentExpression, calculationEntry)
				})
			}
		}

		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Login Form"), false)
		box.Append(usernameEntry, false)
		box.Append(passwordEntry, false)

		buttonBox := ui.NewHorizontalBox()
		buttonBox.Append(loginButton, false)
		buttonBox.Append(registerButton, false)
		box.Append(buttonBox, false)

		vbox := ui.NewVerticalBox()
		vbox.Append(ui.NewLabel("Calculator"), false)
		vbox.Append(calculationEntry, false)
		vbox.Append(buttonGrid, false)

		box.Append(vbox, false)

		window := ui.NewWindow("Login App", 400, 300, false)
		window.SetChild(box)

		loginButton.OnClicked(func(*ui.Button) {
			username := usernameEntry.Text()
			password := passwordEntry.Text()

			var hashedPasswordDB string
			sqlStmt := `SELECT hashed_password FROM users WHERE username = ?`
			err = db.QueryRow(sqlStmt, username).Scan(&hashedPasswordDB)
			if err != nil {
				log.Println("Invalid username:", err)
				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(hashedPasswordDB), []byte(password))
			if err != nil {
				log.Println("Invalid password:", err)
				return
			}

			log.Println("Successfully Logged In!")
		})

		registerButton.OnClicked(func(*ui.Button) {
			// Create a new window for registration
			registrationWindow := ui.NewWindow("Register", 300, 200, false)

			// UI elements for registration
			registrationUsernameEntry := ui.NewEntry()
			registrationPasswordEntry := ui.NewPasswordEntry()
			registerConfirmButton := ui.NewButton("Register")

			registrationBox := ui.NewVerticalBox()
			registrationBox.Append(ui.NewLabel("Registration Form"), false)
			registrationBox.Append(registrationUsernameEntry, false)
			registrationBox.Append(registrationPasswordEntry, false)
			registrationBox.Append(registerConfirmButton, false)

			registrationWindow.SetChild(registrationBox)

			registerConfirmButton.OnClicked(func(*ui.Button) {
				// Get username and password from the registration form
				newUsername := registrationUsernameEntry.Text()
				newPassword := registrationPasswordEntry.Text()

				// Hash the password before storing it in the database
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), passwordCost)
				if err != nil {
					log.Println("Error hashing password:", err)
					return
				}

				// Insert the new user into the database
				insertStmt := `INSERT INTO users (username, hashed_password) VALUES (?, ?)`
				_, err = db.Exec(insertStmt, newUsername, hashedPassword)
				if err != nil {
					log.Println("Error inserting user:", err)
					return
				}

				log.Println("User registered successfully!")
				// Close the registration window after successful registration
				registrationWindow.Destroy()
			})

			// Show the registration window
			registrationWindow.Show()
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()

	})
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

}

func handleButtonClick(button string, currentExpression *string, calculationEntry *ui.Entry) {
	switch button {
	case "=":
		result, err := evaluateExpression(*currentExpression)
		if err != nil {
			log.Println("Error in calculation:", err)
			return
		}
		*currentExpression = strconv.FormatFloat(result, 'f', -1, 64)
	case ".":
		// Handle decimal point separately
		if len(*currentExpression) > 0 && (*currentExpression)[len(*currentExpression)-1] != '.' {
			*currentExpression += button
		}
	default:
		*currentExpression += button
	}

	calculationEntry.SetText(*currentExpression)
}

func evaluateExpression(expression string) (float64, error) {
	// Your existing calculation functions

	if len(expression) == 0 {
		return 0, nil
	}

	switch expression[len(expression)-1] {
	case '+':
		return addition(expression[:len(expression)-1])
	case '-':
		return subtraction(expression[:len(expression)-1])
	case '*':
		return multiplication(expression[:len(expression)-1])
	case '/':
		return division(expression[:len(expression)-1])
	default:
		return 0, fmt.Errorf("invalid expression")
	}
}

func addition(expression string) (float64, error) {
	operands := parseOperands(expression)
	if len(operands) != 2 {
		return 0, fmt.Errorf("invalid addition expression")
	}
	return operands[0] + operands[1], nil
}

func subtraction(expression string) (float64, error) {
	operands := parseOperands(expression)
	if len(operands) != 2 {
		return 0, fmt.Errorf("invalid subtraction expression")
	}
	return operands[0] - operands[1], nil
}

func multiplication(expression string) (float64, error) {
	operands := parseOperands(expression)
	if len(operands) != 2 {
		return 0, fmt.Errorf("invalid multiplication expression")
	}
	return operands[0] * operands[1], nil
}

func division(expression string) (float64, error) {
	operands := parseOperands(expression)
	if len(operands) != 2 || operands[1] == 0 {
		return 0, fmt.Errorf("invalid division expression")
	}
	return operands[0] / operands[1], nil
}

func parseOperands(expression string) []float64 {
	// Split the expression into operands
	// and convert them to float64
	var operands []float64
	for _, operandStr := range splitOperands(expression) {
		operand, _ := strconv.ParseFloat(operandStr, 64)
		operands = append(operands, operand)
	}
	return operands
}

func splitOperands(expression string) []string {
	// Split the expression by operators
	// to extract the operands
	var operands []string
	currentOperand := ""

	// Add a check for the first character to handle the case where it is an operator
	for i, char := range expression {
		if i == 0 && (char == '+' || char == '-' || char == '*' || char == '/') {
			currentOperand = string(char)
			continue
		}

		if char == '+' || char == '-' || char == '*' || char == '/' {
			if currentOperand != "" {
				operands = append(operands, currentOperand)
				currentOperand = ""
			}
		}
		currentOperand += string(char)
	}

	if currentOperand != "" {
		operands = append(operands, currentOperand)
	}

	return operands
}
