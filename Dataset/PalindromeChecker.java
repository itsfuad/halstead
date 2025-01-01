public aspect LoggingAspect {
    pointcut methodCalls(): execution(* *(..));
    before(): methodCalls() {
        System.out.println("Method called: " + thisJoinPoint.getSignature());
    }
}
public class PalindromeChecker {
    public boolean isPalindrome(int number) {
        int original = number;
        int reversed = 0;
        while (number != 0) {
            int digit = number % 10;
            reversed = reversed * 10 + digit;
            number /= 10;
        }
        return original == reversed;
    }
    public static void main(String[] args) {
        PalindromeChecker checker = new PalindromeChecker();
        int number = 121;
        boolean result = checker.isPalindrome(number);
        System.out.println("Is " + number + " a palindrome? " + result);
    }
}
