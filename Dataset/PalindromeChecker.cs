using PostSharp.Aspects;
using System;
[Serializable]
public class LoggingAspect : OnMethodBoundaryAspect {
    public override void OnEntry(MethodExecutionArgs args) {
        Console.WriteLine("Method called: " + args.Method.Name);
    }
}
public class PalindromeChecker {
    [LoggingAspect]
    public bool IsPalindrome(int number) {
        int original = number;
        int reversed = 0;
        while (number != 0) {
            int digit = number % 10;
            reversed = reversed * 10 + digit;
            number /= 10;
        }
        return original == reversed;
    }
    static void Main(string[] args) {
        PalindromeChecker checker = new PalindromeChecker();
        int number = 121;
        bool result = checker.IsPalindrome(number);
        Console.WriteLine("Is " + number + " a palindrome? " + result);
    }
}
