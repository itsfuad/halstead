function logMethodCalls(target) {
    return new Proxy(target, {
        apply: function(target, thisArg, argumentsList) {
            console.log(`Method called: ${target.name}`);
            return target.apply(thisArg, argumentsList);
        }
    });
}
function isPalindrome(number) {
    let original = number;
    let reversed = 0;
    while (number != 0) {
        let digit = number % 10;
        reversed = reversed * 10 + digit;
        number = Math.floor(number / 10);
    }
    return original === reversed;
}
isPalindrome = logMethodCalls(isPalindrome);
let number = 121;
let result = isPalindrome(number);
console.log(`Is ${number} a palindrome? ${result}`);
