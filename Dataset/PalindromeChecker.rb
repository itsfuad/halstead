module LoggingAspect
    def is_palindrome(number)
      puts "Method called: is_palindrome"
      super
    end
  end
  class PalindromeChecker
    prepend LoggingAspect
    def is_palindrome(number)
      original = number
      reversed = 0
      while number != 0
        digit = number % 10
        reversed = reversed * 10 + digit
        number /= 10
      end
      original == reversed
    end
  end
  checker = PalindromeChecker.new
  number = 121
  result = checker.is_palindrome(number)
  puts "Is #{number} a palindrome? #{result}"
  