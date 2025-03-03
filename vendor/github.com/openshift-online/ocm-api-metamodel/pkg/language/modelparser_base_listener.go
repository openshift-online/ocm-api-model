// Code generated from ModelParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package language // ModelParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseModelParserListener is a complete listener for a parse tree produced by ModelParser.
type BaseModelParserListener struct{}

var _ ModelParserListener = &BaseModelParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseModelParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseModelParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseModelParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseModelParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseModelParserListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseModelParserListener) ExitFile(ctx *FileContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseModelParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseModelParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseModelParserListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseModelParserListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterEnumDecl is called when production enumDecl is entered.
func (s *BaseModelParserListener) EnterEnumDecl(ctx *EnumDeclContext) {}

// ExitEnumDecl is called when production enumDecl is exited.
func (s *BaseModelParserListener) ExitEnumDecl(ctx *EnumDeclContext) {}

// EnterEnumMemberDecl is called when production enumMemberDecl is entered.
func (s *BaseModelParserListener) EnterEnumMemberDecl(ctx *EnumMemberDeclContext) {}

// ExitEnumMemberDecl is called when production enumMemberDecl is exited.
func (s *BaseModelParserListener) ExitEnumMemberDecl(ctx *EnumMemberDeclContext) {}

// EnterClassDecl is called when production classDecl is entered.
func (s *BaseModelParserListener) EnterClassDecl(ctx *ClassDeclContext) {}

// ExitClassDecl is called when production classDecl is exited.
func (s *BaseModelParserListener) ExitClassDecl(ctx *ClassDeclContext) {}

// EnterStructDecl is called when production structDecl is entered.
func (s *BaseModelParserListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production structDecl is exited.
func (s *BaseModelParserListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterStructMemberDecl is called when production structMemberDecl is entered.
func (s *BaseModelParserListener) EnterStructMemberDecl(ctx *StructMemberDeclContext) {}

// ExitStructMemberDecl is called when production structMemberDecl is exited.
func (s *BaseModelParserListener) ExitStructMemberDecl(ctx *StructMemberDeclContext) {}

// EnterAttributeKind is called when production attributeKind is entered.
func (s *BaseModelParserListener) EnterAttributeKind(ctx *AttributeKindContext) {}

// ExitAttributeKind is called when production attributeKind is exited.
func (s *BaseModelParserListener) ExitAttributeKind(ctx *AttributeKindContext) {}

// EnterTypeReference is called when production typeReference is entered.
func (s *BaseModelParserListener) EnterTypeReference(ctx *TypeReferenceContext) {}

// ExitTypeReference is called when production typeReference is exited.
func (s *BaseModelParserListener) ExitTypeReference(ctx *TypeReferenceContext) {}

// EnterPlainTypeReference is called when production plainTypeReference is entered.
func (s *BaseModelParserListener) EnterPlainTypeReference(ctx *PlainTypeReferenceContext) {}

// ExitPlainTypeReference is called when production plainTypeReference is exited.
func (s *BaseModelParserListener) ExitPlainTypeReference(ctx *PlainTypeReferenceContext) {}

// EnterListTypeReference is called when production listTypeReference is entered.
func (s *BaseModelParserListener) EnterListTypeReference(ctx *ListTypeReferenceContext) {}

// ExitListTypeReference is called when production listTypeReference is exited.
func (s *BaseModelParserListener) ExitListTypeReference(ctx *ListTypeReferenceContext) {}

// EnterMapTypeReference is called when production mapTypeReference is entered.
func (s *BaseModelParserListener) EnterMapTypeReference(ctx *MapTypeReferenceContext) {}

// ExitMapTypeReference is called when production mapTypeReference is exited.
func (s *BaseModelParserListener) ExitMapTypeReference(ctx *MapTypeReferenceContext) {}

// EnterResourceDecl is called when production resourceDecl is entered.
func (s *BaseModelParserListener) EnterResourceDecl(ctx *ResourceDeclContext) {}

// ExitResourceDecl is called when production resourceDecl is exited.
func (s *BaseModelParserListener) ExitResourceDecl(ctx *ResourceDeclContext) {}

// EnterResourceMemberDecl is called when production resourceMemberDecl is entered.
func (s *BaseModelParserListener) EnterResourceMemberDecl(ctx *ResourceMemberDeclContext) {}

// ExitResourceMemberDecl is called when production resourceMemberDecl is exited.
func (s *BaseModelParserListener) ExitResourceMemberDecl(ctx *ResourceMemberDeclContext) {}

// EnterMethodDecl is called when production methodDecl is entered.
func (s *BaseModelParserListener) EnterMethodDecl(ctx *MethodDeclContext) {}

// ExitMethodDecl is called when production methodDecl is exited.
func (s *BaseModelParserListener) ExitMethodDecl(ctx *MethodDeclContext) {}

// EnterMethodMemberDecl is called when production methodMemberDecl is entered.
func (s *BaseModelParserListener) EnterMethodMemberDecl(ctx *MethodMemberDeclContext) {}

// ExitMethodMemberDecl is called when production methodMemberDecl is exited.
func (s *BaseModelParserListener) ExitMethodMemberDecl(ctx *MethodMemberDeclContext) {}

// EnterMethodParameterDecl is called when production methodParameterDecl is entered.
func (s *BaseModelParserListener) EnterMethodParameterDecl(ctx *MethodParameterDeclContext) {}

// ExitMethodParameterDecl is called when production methodParameterDecl is exited.
func (s *BaseModelParserListener) ExitMethodParameterDecl(ctx *MethodParameterDeclContext) {}

// EnterParameterDirection is called when production parameterDirection is entered.
func (s *BaseModelParserListener) EnterParameterDirection(ctx *ParameterDirectionContext) {}

// ExitParameterDirection is called when production parameterDirection is exited.
func (s *BaseModelParserListener) ExitParameterDirection(ctx *ParameterDirectionContext) {}

// EnterLocatorDecl is called when production locatorDecl is entered.
func (s *BaseModelParserListener) EnterLocatorDecl(ctx *LocatorDeclContext) {}

// ExitLocatorDecl is called when production locatorDecl is exited.
func (s *BaseModelParserListener) ExitLocatorDecl(ctx *LocatorDeclContext) {}

// EnterLocatorMemberDecl is called when production locatorMemberDecl is entered.
func (s *BaseModelParserListener) EnterLocatorMemberDecl(ctx *LocatorMemberDeclContext) {}

// ExitLocatorMemberDecl is called when production locatorMemberDecl is exited.
func (s *BaseModelParserListener) ExitLocatorMemberDecl(ctx *LocatorMemberDeclContext) {}

// EnterLocatorTargetDecl is called when production locatorTargetDecl is entered.
func (s *BaseModelParserListener) EnterLocatorTargetDecl(ctx *LocatorTargetDeclContext) {}

// ExitLocatorTargetDecl is called when production locatorTargetDecl is exited.
func (s *BaseModelParserListener) ExitLocatorTargetDecl(ctx *LocatorTargetDeclContext) {}

// EnterLocatorVariableDecl is called when production locatorVariableDecl is entered.
func (s *BaseModelParserListener) EnterLocatorVariableDecl(ctx *LocatorVariableDeclContext) {}

// ExitLocatorVariableDecl is called when production locatorVariableDecl is exited.
func (s *BaseModelParserListener) ExitLocatorVariableDecl(ctx *LocatorVariableDeclContext) {}

// EnterResourceReference is called when production resourceReference is entered.
func (s *BaseModelParserListener) EnterResourceReference(ctx *ResourceReferenceContext) {}

// ExitResourceReference is called when production resourceReference is exited.
func (s *BaseModelParserListener) ExitResourceReference(ctx *ResourceReferenceContext) {}

// EnterErrorDecl is called when production errorDecl is entered.
func (s *BaseModelParserListener) EnterErrorDecl(ctx *ErrorDeclContext) {}

// ExitErrorDecl is called when production errorDecl is exited.
func (s *BaseModelParserListener) ExitErrorDecl(ctx *ErrorDeclContext) {}

// EnterErrorMemberDecl is called when production errorMemberDecl is entered.
func (s *BaseModelParserListener) EnterErrorMemberDecl(ctx *ErrorMemberDeclContext) {}

// ExitErrorMemberDecl is called when production errorMemberDecl is exited.
func (s *BaseModelParserListener) ExitErrorMemberDecl(ctx *ErrorMemberDeclContext) {}

// EnterErrorCodeDecl is called when production errorCodeDecl is entered.
func (s *BaseModelParserListener) EnterErrorCodeDecl(ctx *ErrorCodeDeclContext) {}

// ExitErrorCodeDecl is called when production errorCodeDecl is exited.
func (s *BaseModelParserListener) ExitErrorCodeDecl(ctx *ErrorCodeDeclContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseModelParserListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseModelParserListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterAnnotationParameters is called when production annotationParameters is entered.
func (s *BaseModelParserListener) EnterAnnotationParameters(ctx *AnnotationParametersContext) {}

// ExitAnnotationParameters is called when production annotationParameters is exited.
func (s *BaseModelParserListener) ExitAnnotationParameters(ctx *AnnotationParametersContext) {}

// EnterAnnotationParameter is called when production annotationParameter is entered.
func (s *BaseModelParserListener) EnterAnnotationParameter(ctx *AnnotationParameterContext) {}

// ExitAnnotationParameter is called when production annotationParameter is exited.
func (s *BaseModelParserListener) ExitAnnotationParameter(ctx *AnnotationParameterContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseModelParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseModelParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseModelParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseModelParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseModelParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseModelParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseModelParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseModelParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseModelParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseModelParserListener) ExitIdentifier(ctx *IdentifierContext) {}
