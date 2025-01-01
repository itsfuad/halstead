import aspectlib
@aspectlib.Aspect
def logging_aspect(cutpoint):
    print(f"Method called: {cutpoint.__name__}")
    yield aspectlib.Proceed
@logging_aspect
def is_palindrome(number):
    original = number
    reversed_number = 0
    while number != 0:
        digit = number % 10
        reversed_number = reversed_number * 10 + digit
        number //= 10
    return original == reversed_number
if __name__ == "__main__":
    number = 121  # Example number
    result = is_palindrome(number)
    print(f"Is {number} a palindrome? {result}")
