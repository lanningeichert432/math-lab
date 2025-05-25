class Calculator:
    def add(self, num1, num2):
        return num1 + num2

    def subtract(self, num1, num2):
        return num1 - num2

    def multiply(self, num1, num2):
        return num1 * num2

    def divide(self, num1, num2):
        if num2 != 0:
            return num1 / num2
        else:
            raise ValueError("Cannot divide by zero")

class MathLab:
    def start(self):
        print("Welcome to the Math Lab!")
        while True:
            action = input("\nChoose an operation (+, -, *, /): ")
            if action == "quit":
                break

            try:
                num1 = float(input("\nEnter first number: "))
                num2 = float(input("Enter second number: "))

                if action == "+":
                    print(f"{num1} + {num2} = {Calculator().add(num1, num2)}")

                elif action == "-":
                    print(f"{num1} - {num2} = {Calculator().subtract(num1, num2)}")

                elif action == "*":
                    print(f"{num1} * {num2} = {Calculator().multiply(num1, num2)}")

                elif action == "/":
                    try:
                        result = Calculator().divide(num1, num2)
                        print(f"{num1} / {num2} = {result}")
                    except ValueError as e:
                        print(e)

            except ValueError:
                print("Invalid input. Please enter a valid number.")

if __name__ == "__main__":
    lab = MathLab()
