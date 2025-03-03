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

parser grammar ModelParser;

options {
  language=Go;
  tokenVocab=ModelLexer;
}

@header {
  import (
    "github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
    "github.com/openshift-online/ocm-api-metamodel/pkg/names"
  )
}

file:
  declaration*
;

declaration:
  typeDecl
| resourceDecl
;

typeDecl returns[result: *concepts.Type]:
  enumDecl
| classDecl
| structDecl
| errorDecl
;

enumDecl returns[result: *concepts.Type]:
  annotations += annotation*
  'enum' name = identifier '{'
    members += enumMemberDecl*
  '}'
;

enumMemberDecl returns[result: *concepts.EnumValue]:
  annotations += annotation*
  'value'? name = identifier
;

classDecl returns[result: *concepts.Type]:
  annotations += annotation*
  'class' name = identifier '{'
    members += structMemberDecl*
  '}'
;

structDecl returns[result: *concepts.Type]:
  annotations += annotation*
  'struct' name = identifier '{'
    members += structMemberDecl*
  '}'
;

structMemberDecl returns[result: *concepts.Attribute]:
  annotations += annotation*
  kind = attributeKind? name = identifier reference = typeReference
;

attributeKind returns[result: int]:
  'attribute'
| 'link'
;

typeReference returns[result: *concepts.Type]:
  plain = plainTypeReference
| list = listTypeReference
| mp = mapTypeReference
;

plainTypeReference returns[result: *concepts.Type]:
  name = identifier
;

listTypeReference returns[result: *concepts.Type]:
  '[' ']' element = identifier
;

mapTypeReference returns[result: *concepts.Type]:
  '[' index = identifier ']' element = identifier
;

resourceDecl returns[result: *concepts.Resource]:
  annotations += annotation*
  'resource' name = identifier '{'
    members += resourceMemberDecl*
  '}'
;

resourceMemberDecl returns[result: interface{}]:
  methodDecl
| locatorDecl
;

methodDecl returns[result: *concepts.Method]:
  annotations += annotation*
  'method'? name = identifier '{'
    members += methodMemberDecl*
  '}'
;

methodMemberDecl returns[result: interface{}]:
  methodParameterDecl
;

methodParameterDecl returns[result: *concepts.Parameter]:
  annotations += annotation*
  'parameter'?
  directions += parameterDirection*
  name = identifier reference =
  typeReference
  ( '=' dflt = literal )?
;

parameterDirection:
  in = 'in'
| out = 'out'
;

locatorDecl returns[result: *concepts.Locator]:
  annotations += annotation*
  'locator' name = identifier '{'
    members += locatorMemberDecl*
  '}'
;

locatorMemberDecl returns[result: interface{}]:
  locatorTargetDecl
| locatorVariableDecl
;

locatorTargetDecl returns[result: *concepts.Resource]:
  'target' reference = resourceReference
;

locatorVariableDecl returns[result: *names.Name]:
  'variable' name = identifier
;

resourceReference returns[result: *concepts.Resource]:
  name = identifier
;

errorDecl returns[result: *concepts.Error]:
  annotations += annotation*
  'error' name = identifier '{'
    members += errorMemberDecl*
  '}'
;

errorMemberDecl returns[result: interface{}]:
  errorCodeDecl
;

errorCodeDecl returns[result: int]:
  'code' code = INTEGER_LITERAL
;

annotation returns[result: *concepts.Annotation]:
  '@' name = identifier parameters = annotationParameters?
;

annotationParameters returns[result: map[string]interface{}]:
  '(' parameters += annotationParameter* ')'
;

annotationParameter returns[name: string, value: interface{}]:
  identifier '=' literal
;

literal returns[result: interface{}]:
  booleanLiteral
| integerLiteral
| stringLiteral
;

booleanLiteral returns[result: bool]:
  'true'
| 'false'
;

integerLiteral returns[result: int]:
  INTEGER_LITERAL
;

stringLiteral returns[result: string]:
  STRING_LITERAL
;

identifier returns[result: *names.Name, text: string]:
  id = IDENTIFIER
;
