#include <strings.h>
#include <stdio.h>
#include <regex.h>

typedef struct {
    const char* IF = "if";
    const char* ELSE = "else";
    const char* OPEN_PARENTHESE = "(";
    const char* CLOSE_PARENTHESE = ")";
    const char* EQUALITY = "==";
    const char* INEQUALITY = "!=";
    const char* LESS = "<";
    const char* EQUAL_LESS = "<=";
    const char* GREAT = ">";
    const char* EQUAL_GREAT = ">=";
    const char* NOT = "!";
    const char* OPEN_BRACE = "{";
    const char* CLOSE_BRACE = "}";
    const char* BREAK_LINE = "\n";
    const char* AND = "&&";
    const char* OR = "||";
    const char* IDENTIFIER = "[a-zA-Z]\w*";
    const char* ATTRIBUTE = "^[a-zA-Z]\w*";
    const char* INT_VALUE = "[0-9]+";
    const char* FLOAT_VALUE = "[0-9]+.[0-9+]";
    const char* STRING_VALUE = "\"[a-zA-Z]\w*\"";
} TokenValue;

typedef enum {
    IF,
    ELSE,
    OPEN_PARENTHESE,
    CLOSE_PARENTHESE,
    EQUALITY,
    INEQUALITY,
    LESS,
    EQUAL_LESS,
    GREAT,
    EQUAL_GREAT,
    NOT,
    OPEN_BRACE,
    CLOSE_BRACE,
    AND,
    OR,
    IDENTIFIER,
    ATTRIBUTE,
    INT_VALUE,
    FLOAT_VALUE
} TokenLabel;