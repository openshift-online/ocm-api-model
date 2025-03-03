// Code generated from ModelParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package language // ModelParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ModelParserListener is a complete listener for a parse tree produced by ModelParser.
type ModelParserListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterEnumDecl is called when entering the enumDecl production.
	EnterEnumDecl(c *EnumDeclContext)

	// EnterEnumMemberDecl is called when entering the enumMemberDecl production.
	EnterEnumMemberDecl(c *EnumMemberDeclContext)

	// EnterClassDecl is called when entering the classDecl production.
	EnterClassDecl(c *ClassDeclContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterStructMemberDecl is called when entering the structMemberDecl production.
	EnterStructMemberDecl(c *StructMemberDeclContext)

	// EnterAttributeKind is called when entering the attributeKind production.
	EnterAttributeKind(c *AttributeKindContext)

	// EnterTypeReference is called when entering the typeReference production.
	EnterTypeReference(c *TypeReferenceContext)

	// EnterPlainTypeReference is called when entering the plainTypeReference production.
	EnterPlainTypeReference(c *PlainTypeReferenceContext)

	// EnterListTypeReference is called when entering the listTypeReference production.
	EnterListTypeReference(c *ListTypeReferenceContext)

	// EnterMapTypeReference is called when entering the mapTypeReference production.
	EnterMapTypeReference(c *MapTypeReferenceContext)

	// EnterResourceDecl is called when entering the resourceDecl production.
	EnterResourceDecl(c *ResourceDeclContext)

	// EnterResourceMemberDecl is called when entering the resourceMemberDecl production.
	EnterResourceMemberDecl(c *ResourceMemberDeclContext)

	// EnterMethodDecl is called when entering the methodDecl production.
	EnterMethodDecl(c *MethodDeclContext)

	// EnterMethodMemberDecl is called when entering the methodMemberDecl production.
	EnterMethodMemberDecl(c *MethodMemberDeclContext)

	// EnterMethodParameterDecl is called when entering the methodParameterDecl production.
	EnterMethodParameterDecl(c *MethodParameterDeclContext)

	// EnterParameterDirection is called when entering the parameterDirection production.
	EnterParameterDirection(c *ParameterDirectionContext)

	// EnterLocatorDecl is called when entering the locatorDecl production.
	EnterLocatorDecl(c *LocatorDeclContext)

	// EnterLocatorMemberDecl is called when entering the locatorMemberDecl production.
	EnterLocatorMemberDecl(c *LocatorMemberDeclContext)

	// EnterLocatorTargetDecl is called when entering the locatorTargetDecl production.
	EnterLocatorTargetDecl(c *LocatorTargetDeclContext)

	// EnterLocatorVariableDecl is called when entering the locatorVariableDecl production.
	EnterLocatorVariableDecl(c *LocatorVariableDeclContext)

	// EnterResourceReference is called when entering the resourceReference production.
	EnterResourceReference(c *ResourceReferenceContext)

	// EnterErrorDecl is called when entering the errorDecl production.
	EnterErrorDecl(c *ErrorDeclContext)

	// EnterErrorMemberDecl is called when entering the errorMemberDecl production.
	EnterErrorMemberDecl(c *ErrorMemberDeclContext)

	// EnterErrorCodeDecl is called when entering the errorCodeDecl production.
	EnterErrorCodeDecl(c *ErrorCodeDeclContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterAnnotationParameters is called when entering the annotationParameters production.
	EnterAnnotationParameters(c *AnnotationParametersContext)

	// EnterAnnotationParameter is called when entering the annotationParameter production.
	EnterAnnotationParameter(c *AnnotationParameterContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitEnumDecl is called when exiting the enumDecl production.
	ExitEnumDecl(c *EnumDeclContext)

	// ExitEnumMemberDecl is called when exiting the enumMemberDecl production.
	ExitEnumMemberDecl(c *EnumMemberDeclContext)

	// ExitClassDecl is called when exiting the classDecl production.
	ExitClassDecl(c *ClassDeclContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitStructMemberDecl is called when exiting the structMemberDecl production.
	ExitStructMemberDecl(c *StructMemberDeclContext)

	// ExitAttributeKind is called when exiting the attributeKind production.
	ExitAttributeKind(c *AttributeKindContext)

	// ExitTypeReference is called when exiting the typeReference production.
	ExitTypeReference(c *TypeReferenceContext)

	// ExitPlainTypeReference is called when exiting the plainTypeReference production.
	ExitPlainTypeReference(c *PlainTypeReferenceContext)

	// ExitListTypeReference is called when exiting the listTypeReference production.
	ExitListTypeReference(c *ListTypeReferenceContext)

	// ExitMapTypeReference is called when exiting the mapTypeReference production.
	ExitMapTypeReference(c *MapTypeReferenceContext)

	// ExitResourceDecl is called when exiting the resourceDecl production.
	ExitResourceDecl(c *ResourceDeclContext)

	// ExitResourceMemberDecl is called when exiting the resourceMemberDecl production.
	ExitResourceMemberDecl(c *ResourceMemberDeclContext)

	// ExitMethodDecl is called when exiting the methodDecl production.
	ExitMethodDecl(c *MethodDeclContext)

	// ExitMethodMemberDecl is called when exiting the methodMemberDecl production.
	ExitMethodMemberDecl(c *MethodMemberDeclContext)

	// ExitMethodParameterDecl is called when exiting the methodParameterDecl production.
	ExitMethodParameterDecl(c *MethodParameterDeclContext)

	// ExitParameterDirection is called when exiting the parameterDirection production.
	ExitParameterDirection(c *ParameterDirectionContext)

	// ExitLocatorDecl is called when exiting the locatorDecl production.
	ExitLocatorDecl(c *LocatorDeclContext)

	// ExitLocatorMemberDecl is called when exiting the locatorMemberDecl production.
	ExitLocatorMemberDecl(c *LocatorMemberDeclContext)

	// ExitLocatorTargetDecl is called when exiting the locatorTargetDecl production.
	ExitLocatorTargetDecl(c *LocatorTargetDeclContext)

	// ExitLocatorVariableDecl is called when exiting the locatorVariableDecl production.
	ExitLocatorVariableDecl(c *LocatorVariableDeclContext)

	// ExitResourceReference is called when exiting the resourceReference production.
	ExitResourceReference(c *ResourceReferenceContext)

	// ExitErrorDecl is called when exiting the errorDecl production.
	ExitErrorDecl(c *ErrorDeclContext)

	// ExitErrorMemberDecl is called when exiting the errorMemberDecl production.
	ExitErrorMemberDecl(c *ErrorMemberDeclContext)

	// ExitErrorCodeDecl is called when exiting the errorCodeDecl production.
	ExitErrorCodeDecl(c *ErrorCodeDeclContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitAnnotationParameters is called when exiting the annotationParameters production.
	ExitAnnotationParameters(c *AnnotationParametersContext)

	// ExitAnnotationParameter is called when exiting the annotationParameter production.
	ExitAnnotationParameter(c *AnnotationParameterContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
