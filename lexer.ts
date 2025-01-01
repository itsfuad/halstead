

type TokenKind = "OPERATOR" | "OPERAND";

class Token {
    Kind: TokenKind;
    Value: string;
}

const LanguageDefinition: { [key: string]: { skippable: RegExp[], operators: RegExp[], operands: RegExp[] } }
= {
    "cpp": {
        //regexes
        "skippable": [/\s/, /\/\/.*/, /\/\*.*\*\//], //whitespace, comments (single line, multi line)
        "operators": [/\+/, /-/, /\*/, /\//, /%/, /\^/, /=/, /&/, /\|/, /~/, /!/, /</, />/, /\?/, /:/, /\./, /,/, /;/, /\(/, /\)/, /\[/, /\]/, /\{/, /\}/],
        "operands": [/[a-zA-Z_]\w*/, /\d+/, /#include(["<])\w+([">])/, /"[a-zA-Z0-9_\s]*"/, /[a-zA-Z_]\(.*\)/], //identifiers, numbers, includes, strings, functions
    }
}

function tokenize(src: string, lang: string): Token[] {
    const tokens: Token[] = [];
    const def = LanguageDefinition[lang];
    const regexes = [...def.skippable, ...def.operators, ...def.operands];

    let i = 0;

    while (i < src.length) {
        //when we find a match update the position and add the token to the list and continue from the next position
        for (const regex of regexes) {
            const match = regex.exec(src.slice(i));
            if (match && match.index === 0) {
                const token = new Token();
                token.Value = match[0];
                if (def.skippable.includes(regex)) {
                    i += match[0].length;
                    break;
                } else if (def.operators.includes(regex)) {
                    token.Kind = "OPERATOR";
                    tokens.push(token);
                    i += match[0].length;
                    break;
                } else if (def.operands.includes(regex)) {
                    token.Kind = "OPERAND";
                    tokens.push(token);
                    i += match[0].length;
                    break;
                }
            }
        }
    }

    return tokens;
}


//palindrome checker in c++
const src = `
#include <iostream>
#include <string>
using namespace std;

int main() {
    string str, rev;
    cout << "Enter a string: ";
    cin >> str;
    rev = str;
    reverse(rev.begin(), rev.end());
    if (str == rev) {
        cout << str << " is a palindrome";
        } else {
            cout << str << " is not a palindrome";
        }
    }
    return 0;
}
`;

const tokens = tokenize(src, "cpp");
console.log(tokens);