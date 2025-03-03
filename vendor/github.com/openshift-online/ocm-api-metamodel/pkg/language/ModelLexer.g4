/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

lexer grammar ModelLexer;

options {
  language=Go;
}

@members {
  func (l *ModelLexer) comment() {
    comments[l.GetLine()] = l.GetText()
  }
}

// Keywords:
ATTRIBUTE: 'attribute';
CLASS: 'class';
CODE: 'code';
ENUM: 'enum';
ERROR: 'error';
FALSE: 'false';
IN: 'in';
LINK: 'link';
LOCATOR: 'locator';
METHOD: 'method';
OUT: 'out';
PARAMETER: 'parameter';
RESOURCE: 'resource';
STRUCT: 'struct';
TARGET: 'target';
TRUE: 'true';
VALUE: 'value';
VARIABLE: 'variable';

// Punctuation:
AT_SIGN: '@';
LEFT_CURLY_BRACKET: '{';
RIGHT_CURLY_BRACKET: '}';
LEFT_SQUARE_BRACKET: '[';
RIGHT_SQUARE_BRACKET: ']';
LEFT_PARENTHESIS: '(';
RIGHT_PARENTHESIS: ')';

// Operators:
EQUALS_SIGN: '=';

// Integer literals:
INTEGER_LITERAL:
  [0-9]+
;

// String literals:
STRING_LITERAL:
  SHORT_STRING
| LONG_STRING
;

fragment SHORT_STRING:
  '"' .*? '"'
;

fragment LONG_STRING:
  '`' .*? '`'
;

// Identifiers:
IDENTIFIER:
  [a-zA-Z_][a-zA-Z0-9_]*
;

// Comments:
LINE_COMMENT:
  '//' ~[\r\n]* { l.comment() } -> skip
;

BLOCK_COMMENT:
  '/*' .*? '*/' -> skip
;

WS:
  [ \t\r\n]+ -> skip
;
