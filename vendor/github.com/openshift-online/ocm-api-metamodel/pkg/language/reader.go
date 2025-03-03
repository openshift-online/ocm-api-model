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

// This file contains the implementation of the object that knows how to read a model from a set of
// source files.

package language

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/openshift-online/ocm-api-metamodel/pkg/annotations"
	"github.com/openshift-online/ocm-api-metamodel/pkg/concepts"
	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
	"github.com/openshift-online/ocm-api-metamodel/pkg/nomenclator"
	"github.com/openshift-online/ocm-api-metamodel/pkg/reporter"
)

// Reader is used to read a model from a set of files. Don't create objects of this type directly,
// use the NewReader function instead.
type Reader struct {
	BaseModelParserListener

	// Object used to report errors and other relevant information.
	reporter *reporter.Reporter

	// Paths of the files and directories to load.
	inputs []string

	// The model, service and version that are currently being loaded:
	model   *concepts.Model
	service *concepts.Service
	version *concepts.Version

	// Concepts that aren't completely defined yet:
	undefinedTypes     map[string]*concepts.Type
	undefinedResources map[string]*concepts.Resource
	undefinedErrors    map[string]*concepts.Error
}

// NewReader creates an object that knows how to read  a model from a set of files. Typical usage is
// to create the object, configure it and then call the Read method. For example, to read a model
// from the files contained in the 'mydir' directory, do the following:
//
//	model, err := language.NewReader().
//		Input("mydir").
//		Load()
//
// Error checking will only happen when the Load method is called. For example, if the 'mydir'
// directory doesn't exist, or isn't readable, the error won't be reported when calling the File
// method, but when calling the Read method.
func NewReader() *Reader {
	return new(Reader)
}

// Reporter sets the object that will be used to report information about the loading process,
// including errors.
func (r *Reader) Reporter(reporter *reporter.Reporter) *Reader {
	r.reporter = reporter
	return r
}

// Input adds the path of a file or directory where the model should be loaded from. If it is a
// directory then all the '.model' files in the directory and all its subdirectories will be loaded.
// Can be called multiple times to load the model from multiple files and directories.
func (r *Reader) Input(value string) *Reader {
	r.inputs = append(r.inputs, value)
	return r
}

// Inputs adds multiple model paths. This is equivalent to calling the Input method multiple times.
func (r *Reader) Inputs(values []string) *Reader {
	r.inputs = append(r.inputs, values...)
	return r
}

// Read reads the model.
func (r *Reader) Read() (model *concepts.Model, err error) {
	// Check the parameters:
	if r.reporter == nil {
		err = fmt.Errorf("parameter 'reporter' is mandatory")
		return
	}

	// Create an empty model:
	r.model = concepts.NewModel()

	// Initialize the indexes of undefined concepts:
	r.undefinedTypes = make(map[string]*concepts.Type)
	r.undefinedResources = make(map[string]*concepts.Resource)
	r.undefinedErrors = make(map[string]*concepts.Error)

	// Load the potentially partial representations of the model from the source files:
	if r.inputs != nil {
		for _, input := range r.inputs {
			r.loadModel(input)
		}
	}

	// Report undefined concepts:
	for _, undefinedType := range r.undefinedTypes {
		r.reporter.Errorf("Type '%s' isn't completely defined", undefinedType.Name())
	}
	for _, undefinedResource := range r.undefinedResources {
		r.reporter.Errorf("Resource '%s' isn't completely defined", undefinedResource.Name())
	}
	for _, undefinedError := range r.undefinedErrors {
		r.reporter.Errorf("Error '%s' isn't completely defined", undefinedError.Name())
	}

	// Create a list type for every type in the model that is not itself a list:
	for _, service := range r.model.Services() {
		for _, version := range service.Versions() {
			for _, typ := range version.Types() {
				if typ.Kind() != concepts.ListType {
					listName := names.Cat(typ.Name(), nomenclator.List)
					listType := version.FindType(listName)
					if listType == nil {
						listType = concepts.NewType()
						listType.SetKind(concepts.ListType)
						listType.SetName(listName)
						listType.SetElement(typ)
						version.AddType(listType)
					}
				}
			}
		}
	}

	// Run checks:
	r.checkModel()

	// Check if there are errors:
	errors := r.reporter.Errors()
	if errors > 0 {
		if errors > 1 {
			err = fmt.Errorf("there were %d errors", errors)
		} else {
			err = fmt.Errorf("there was 1 error")
		}
		return
	}

	// Everything worked, we can return the created model:
	model = r.model

	return
}

func (r *Reader) loadModel(dir string) {
	// Check that the directory exists:
	dirInfo, err := os.Stat(dir)
	if os.IsNotExist(err) {
		r.reporter.Errorf("Model directory '%s' doesn't exist", dir)
		return
	}
	if err != nil {
		r.reporter.Errorf("Can't check if model directory '%s' exists: %v", dir, err)
		return
	}
	if !dirInfo.IsDir() {
		r.reporter.Errorf("Model directory '%s' isn't actually a directory", dir)
		return
	}

	// List the services:
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		r.reporter.Errorf("Can't list model directory '%s': %v", dir, err)
		return
	}

	// Load the services:
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			r.loadService(filepath.Join(dir, fileInfo.Name()))
		} else {
			r.reporter.Infof(
				"Model directory '%s' should only contain service "+
					"sub-directories, but it contains file '%s', "+
					"it will be ignored",
				dir, fileInfo.Name(),
			)
		}
	}
}

func (r *Reader) loadService(dir string) {
	// Check that the directory exists:
	dirInfo, err := os.Stat(dir)
	if os.IsNotExist(err) {
		r.reporter.Errorf("Service directory '%s' doesn't exist", dir)
		return
	}
	if err != nil {
		r.reporter.Errorf("Can't check if service directory '%s' exists: %v", dir, err)
		return
	}
	if !dirInfo.IsDir() {
		r.reporter.Errorf("Service directory '%s' isn't actually a directory", dir)
		return
	}

	// Create the service, if needed:
	name := names.ParseUsingSeparator(dirInfo.Name(), "_")
	r.service = r.model.FindService(name)
	if r.service == nil {
		r.service = concepts.NewService()
		r.service.SetName(name)
		r.model.AddService(r.service)
	}

	// List the versions:
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		r.reporter.Errorf("Can't list service directory '%s': %v", dir, err)
		return
	}

	// Load the versions:
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			r.loadVersion(filepath.Join(dir, fileInfo.Name()))
		} else {
			r.reporter.Infof(
				"Service directory '%s' should only contain version "+
					"sub-directories, but it contains file '%s',"+
					" it will be ignored",
				dir, fileInfo.Name(),
			)
		}
	}
}

func (r *Reader) loadVersion(dir string) {
	// Check that the directory exists:
	dirInfo, err := os.Stat(dir)
	if os.IsNotExist(err) {
		r.reporter.Errorf("Version directory '%s' doesn't exist", dir)
		return
	}
	if err != nil {
		r.reporter.Errorf("Can't check if version directory '%s' exists: %v", dir, err)
		return
	}
	if !dirInfo.IsDir() {
		r.reporter.Errorf("Version directory '%s' isn't actually a directory", dir)
		return
	}

	// Create the version, if needed:
	name := names.ParseUsingSeparator(dirInfo.Name(), "_")
	r.version = r.service.FindVersion(name)
	if r.version == nil {
		r.version = concepts.NewVersion()
		r.version.SetName(name)
		r.service.AddVersion(r.version)
	}

	// List the files:
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		r.reporter.Errorf("Can't list version directory '%s': %v", dir, err)
		return
	}

	// Load the files:
	for _, fileInfo := range fileInfos {
		if fileInfo.Mode().IsRegular() {
			r.loadFile(filepath.Join(dir, fileInfo.Name()))
		} else {
			r.reporter.Infof(
				"Version directory '%s' should only model files but it "+
					"contains '%s' which isn't a regular file, it will "+
					"be ignored",
				dir, fileInfo.Name(),
			)
		}
	}
}

func (r *Reader) loadFile(path string) {
	// Inform that the file is being loaded:
	r.reporter.Infof("Loading file '%s'", path)

	// Clear the map of comment lines:
	comments = map[int]string{}

	// Open the source file:
	input, err := antlr.NewFileStream(path)
	if err != nil {
		r.reporter.Errorf("Can't open source file '%s': %s", path, err.Error())
		return
	}

	// Create the error listener:
	listener, err := NewErrorListener().
		Reporter(r.reporter).
		Build()
	if err != nil {
		r.reporter.Errorf("Can't create error listener: %v", err)
		return
	}

	// Create the lexer:
	lexer := NewModelLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	// Create the parser:
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewModelParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)
	tree := parser.File()
	antlr.ParseTreeWalkerDefault.Walk(r, tree)
}

func (r *Reader) ExitIdentifier(ctx *IdentifierContext) {
	text := ctx.GetId().GetText()
	ctx.SetResult(names.ParseUsingCase(text))
	ctx.SetText(text)
}

func (r *Reader) ExitEnumDecl(ctx *EnumDeclContext) {
	// Check if there is an existing type:
	name := ctx.GetName().GetResult()
	typ := r.version.FindType(name)
	if typ == nil {
		typ = concepts.NewType()
		typ.SetKind(concepts.EnumType)
		typ.SetName(name)
		r.version.AddType(typ)
	} else if r.isUndefinedType(typ) {
		typ.SetKind(concepts.EnumType)
		r.removeUndefinedType(typ)
	} else {
		r.reporter.Errorf("Type '%s' is already defined", name)
		return
	}

	// Add the documentation:
	doc := r.getDoc(ctx.GetStart())
	if doc != "" {
		typ.SetDoc(doc)
	}

	// Add the annotations:
	r.addAnnotations(typ, ctx.GetAnnotations())

	// Add the values:
	memberCtxs := ctx.GetMembers()
	if len(memberCtxs) > 0 {
		for _, memberCtx := range ctx.GetMembers() {
			typ.AddValue(memberCtx.GetResult())
		}
	}
}

func (r *Reader) ExitEnumMemberDecl(ctx *EnumMemberDeclContext) {
	// Create the type and set the basic properties:
	value := concepts.NewEnumValue()
	value.SetName(ctx.GetName().GetResult())

	// Add the documentation:
	doc := r.getDoc(ctx.GetStart())
	if doc != "" {
		value.SetDoc(doc)
	}

	// Add the annotations:
	r.addAnnotations(value, ctx.GetAnnotations())

	// Return the value:
	ctx.SetResult(value)
}

func (r *Reader) ExitClassDecl(ctx *ClassDeclContext) {
	// Check if there is an existing type:
	name := ctx.GetName().GetResult()
	typ := r.version.FindType(name)
	if typ == nil {
		typ = concepts.NewType()
		typ.SetKind(concepts.ClassType)
		typ.SetName(name)
		r.version.AddType(typ)
	} else if r.isUndefinedType(typ) {
		typ.SetKind(concepts.ClassType)
		r.removeUndefinedType(typ)
	} else {
		r.reporter.Errorf("Type '%s' is already defined", name)
		return
	}

	// Add the documentation:
	doc := r.getDoc(ctx.GetStart())
	if doc != "" {
		typ.SetDoc(doc)
	}

	// Add the annotations:
	r.addAnnotations(typ, ctx.GetAnnotations())

	// Add the attributes:
	memberCtxs := ctx.GetMembers()
	if len(memberCtxs) > 0 {
		for _, memberCtx := range ctx.GetMembers() {
			typ.AddAttribute(memberCtx.GetResult())
		}
	}

	if path := annotations.ReferencePath(typ); path != "" {
		// If the type has a reference set ExplicitDeclared
		typ.SetExplicitDeclared(true)
		r.removeLinkedAttributes(typ)
		r.handleClassRef(typ, path)
	}
}

func (r *Reader) removeLinkedAttributes(typ *concepts.Type) {
	for _, types := range r.version.Types() {
		for _, attribute := range types.Attributes() {
			// It could be that the type is already in an attribute of the service
			// It needs to add the SetExplicitDeclared to it. {
			if attribute.Type().Name() == typ.Name() {
				attribute.SetLinkOwner(nil)
				attribute.Type().SetExplicitDeclared(true)
			}
			if attribute.Type().IsList() || attribute.Type().IsMap() {
				if attribute.Type().Element().Name().String() == typ.Name().String() {
					attribute.SetLinkOwner(nil)
					attribute.Type().SetExplicitDeclared(true)
					attribute.Type().Element().SetExplicitDeclared(true)
					attribute.Type().Element().SetOwner(typ.Owner())
				}
			}
		}
	}
}

func (r *Reader) handleClassRef(typ *concepts.Type, path string) {
	if len(r.inputs) > 1 {
		panic("referenced service with multiple inputs in undefined")
	}
	input := r.inputs[0]
	path = strings.TrimPrefix(path, "/")
	components := strings.Split(path, "/")
	referencedServiceName := components[0]
	referencedVersion := components[1]
	referencedTypeName := components[2]

	// Create an ad-hoc reader and model for the specific referenced service.
	refReader := NewReader().
		Reporter(r.reporter)
	refReader.model = concepts.NewModel()

	// Initialize the indexes of undefined concepts:
	refReader.undefinedTypes = make(map[string]*concepts.Type)
	refReader.undefinedResources = make(map[string]*concepts.Resource)
	refReader.undefinedErrors = make(map[string]*concepts.Error)

	// load the ad-hoc service and version referenced and find the correct type.
	refReader.loadService(fmt.Sprintf("%s/%s", input, referencedServiceName))
	refVersion := refReader.service.FindVersion(names.ParseUsingSeparator(referencedVersion, "_"))
	// Once loading the service, we find the reference type
	// then recursively iterate the type tree and add the types to the current version.
	if referencedType := refVersion.FindType(names.ParseUsingSeparator(referencedTypeName, "_")); referencedType != nil {
		r.recursivelyAddTypeToVersion(typ, referencedType)
	}
}

// A helper function to recursively add types to a version
func (r *Reader) recursivelyAddTypeToVersion(currType *concepts.Type,
	referencedType *concepts.Type) {
	if referencedType.IsBasicType() {
		return
	}
	for _, attribute := range referencedType.Attributes() {
		// If attribute is explicitDeclared, it needs to change it's owner to be the same as currType
		if attribute.Link() {
			// If the attribute is a Link set the LinkOwner
			// to the attribute type Owner, if it is also a list it should change the owner of the
			// element type
			existingAttType := r.version.FindType(attribute.Type().Name())
			if existingAttType != nil && existingAttType.ExplicitDeclared() {
				attribute.Type().SetOwner(r.version)
				// Stop the iteration over the attributes of this type
				continue
			} else if attribute.Type().IsList() || attribute.Type().IsMap() {
				existingAttType = r.version.FindType(attribute.Type().Element().Name())
				if existingAttType != nil && existingAttType.ExplicitDeclared() {
					attribute.Type().SetOwner(r.version)
					attribute.Type().Element().SetOwner(r.version)
					continue
				}
			}
			// Otherwise add the setLinkOwner to the attribute
			attribute.SetLinkOwner(attribute.Type().Owner())
		} else if attribute.Type().IsList() || attribute.Type().IsMap() {
			r.version.AddType(attribute.Type())
			r.recursivelyAddTypeToVersion(currType, attribute.Type().Element())
		} else {
			r.recursivelyAddTypeToVersion(currType, attribute.Type())
		}
	}

	existingType := r.version.FindType(referencedType.Name())
	if existingType != nil && existingType.ExplicitDeclared() {
		referencedType.SetExplicitDeclared(true)
	}
	r.version.AddType(referencedType)

}

func (r *Reader) ExitStructDecl(ctx *StructDeclContext) {
	// Check if there is an existing type:
	name := ctx.GetName().GetResult()
	typ := r.version.FindType(name)
	if typ == nil {
		typ = concepts.NewType()
		typ.SetKind(concepts.StructType)
		typ.SetName(name)
		r.version.AddType(typ)
	} else if r.isUndefinedType(typ) {
		typ.SetKind(concepts.StructType)
		r.removeUndefinedType(typ)
	} else {
		r.reporter.Errorf("Type '%s' is already defined", name)
		return
	}

	// Add the documentation:
	doc := r.getDoc(ctx.GetStart())
	if doc != "" {
		typ.SetDoc(doc)
	}

	// Add the annotations:
	r.addAnnotations(typ, ctx.GetAnnotations())

	// Add the attributes:
	memberCtxs := ctx.GetMembers()
	if len(memberCtxs) > 0 {
		for _, memberCtx := range ctx.GetMembers() {
			typ.AddAttribute(memberCtx.GetResult())
		}
	}
}

func (r *Reader) ExitStructMemberDecl(ctx *StructMemberDeclContext) {
	// Create the attribute and set the basic properties:
	attribute := concepts.NewAttribute()
	attribute.SetName(ctx.GetName().GetResult())
	attribute.SetType(ctx.GetReference().GetResult())

	// Add the documentation:
	doc := r.getDoc(ctx.GetStart())
	if doc != "" {
		attribute.SetDoc(doc)
	}

	// Add the annotations:
	r.addAnnotations(attribute, ctx.GetAnnotations())

	// Set the link flag:
	kind := ctx.GetKind()
	if kind != nil {
		attribute.SetLink(kind.GetResult() == ModelLexerLINK)
	}

	// Return the attribute:
	ctx.SetResult(attribute)
}

func (r *Reader) ExitAttributeKind(ctx *AttributeKindContext) {
	ctx.SetResult(ctx.GetStart().GetTokenType())
}

func (r *Reader) ExitTypeReference(ctx *TypeReferenceContext) {
	var typ *concepts.Type
	if ctx.GetPlain() != nil {
		typ = ctx.GetPlain().GetResult()
	}
	if ctx.GetList() != nil {
		typ = ctx.GetList().GetResult()
	}
	if ctx.GetMp() != nil {
		typ = ctx.GetMp().GetResult()
	}
	ctx.SetResult(typ)
}

func (r *Reader) ExitPlainTypeReference(ctx *PlainTypeReferenceContext) {
	// Try to find an existing type, or else create a new one that is only partially defined:
	name := ctx.GetName().GetResult()
	typ := r.version.FindType(name)
	if typ == nil {
		typ = concepts.NewType()
		typ.SetName(name)
		r.version.AddType(typ)
		r.addUndefinedType(typ)
	}

	// Return the type:
	ctx.SetResult(typ)
}

func (r *Reader) ExitListTypeReference(ctx *ListTypeReferenceContext) {
	// Try to find an existing element type, or else create a new one that is only partially
	// defined:
	elementName := ctx.GetElement().GetResult()
	elementType := r.version.FindType(elementName)
	if elementType == nil {
		elementType = concepts.NewType()
		elementType.SetName(elementName)
		r.version.AddType(elementType)
		r.addUndefinedType(elementType)
	}

	// Try to find an existing list type, or else create a new one:
	listName := names.Cat(elementName, nomenclator.List)
	listType := r.version.FindType(listName)
	if listType == nil {
		listType = concepts.NewType()
		listType.SetKind(concepts.ListType)
		listType.SetName(listName)
		listType.SetElement(elementType)
		r.version.AddType(listType)
	}

	// Return the list type:
	ctx.SetResult(listType)
}

func (r *Reader) ExitMapTypeReference(ctx *MapTypeReferenceContext) {
	// Try to find existing index and element types, or else create new ones that are only
	// partially defined:
	indexName := ctx.GetIndex().GetResult()
	indexType := r.version.FindType(indexName)
	if indexType == nil {
		indexType = concepts.NewType()
		indexType.SetName(indexName)
		r.version.AddType(indexType)
		r.addUndefinedType(indexType)
	}
	elementName := ctx.GetElement().GetResult()
	elementType := r.version.FindType(elementName)
	if elementType == nil {
		elementType = concepts.NewType()
		elementType.SetName(elementName)
		r.version.AddType(elementType)
		r.addUndefinedType(elementType)
	}

	// Try to find an existing map type, or else create a new one:
	mapName := names.Cat(indexName, elementName, nomenclator.Map)
	mapType := r.version.FindType(mapName)
	if mapType == nil {
		mapType = concepts.NewType()
		mapType.SetKind(concepts.MapType)
		mapType.SetName(mapName)
		mapType.SetIndex(indexType)
		mapType.SetElement(elementType)
		r.version.AddType(mapType)
	}

	// Return the map type:
	ctx.SetResult(mapType)
}

func (r *Reader) ExitResourceReference(ctx *ResourceReferenceContext) {
	name := ctx.GetName().GetResult()
	resource := r.version.FindResource(name)
	if resource == nil {
		resource = concepts.NewResource()
		resource.SetName(name)
		r.version.AddResource(resource)
		r.addUndefinedResource(resource)
	}
	ctx.SetResult(resource)
}

func (r *Reader) ExitResourceMemberDecl(ctx *ResourceMemberDeclContext) {
	if ctx.MethodDecl() != nil {
		ctx.SetResult(ctx.MethodDecl().GetResult())
		return
	}
	if ctx.LocatorDecl() != nil {
		ctx.SetResult(ctx.LocatorDecl().GetResult())
		return
	}
}

func (r *Reader) ExitResourceDecl(ctx *ResourceDeclContext) {
	// Check if there is an existing resource:
	name := ctx.GetName().GetResult()
	resource := r.version.FindResource(name)
	if resource == nil {
		resource = concepts.NewResource()
		resource.SetName(name)
		r.version.AddResource(resource)
	} else if r.isUndefinedResource(resource) {
		r.removeUndefinedResource(resource)
	} else {
		r.reporter.Errorf("Resource '%s' is already defined", name)
		return
	}

	// Add the documentation:
	doc := r.getDoc(ctx.GetStart())
	if doc != "" {
		resource.SetDoc(doc)
	}

	// Add the annotations:
	r.addAnnotations(resource, ctx.GetAnnotations())

	// Add the attributes:
	memberCtxs := ctx.GetMembers()
	if len(memberCtxs) > 0 {
		for _, memberCtx := range ctx.GetMembers() {
			switch member := memberCtx.GetResult().(type) {
			case *concepts.Method:
				resource.AddMethod(member)
			case *concepts.Locator:
				resource.AddLocator(member)
			}
		}
	}
}

func (r *Reader) ExitMethodDecl(ctx *MethodDeclContext) {
	// Create the method and set the basic properties:
	method := concepts.NewMethod()
	method.SetName(ctx.GetName().GetResult())

	// Add the documentation:
	doc := r.getDoc(ctx.GetStart())
	if doc != "" {
		method.SetDoc(doc)
	}

	// Add the annotations:
	r.addAnnotations(method, ctx.GetAnnotations())

	// Add the membmers:
	membersCtxs := ctx.GetMembers()
	if len(membersCtxs) > 0 {
		for _, memberCtx := range membersCtxs {
			switch member := memberCtx.GetResult().(type) {
			case *concepts.Parameter:
				method.AddParameter(member)
			}
		}
	}

	// Return the method:
	ctx.SetResult(method)
}

func (r *Reader) ExitMethodMemberDecl(ctx *MethodMemberDeclContext) {
	if ctx.MethodParameterDecl() != nil {
		ctx.SetResult(ctx.MethodParameterDecl().GetResult())
		return
	}
}

func (r *Reader) ExitMethodParameterDecl(ctx *MethodParameterDeclContext) {
	// Create the parameter and set the basic properties:
	parameter := concepts.NewParameter()
	parameter.SetName(ctx.GetName().GetResult())
	parameter.SetType(ctx.GetReference().GetResult())

	// Add the documentation:
	doc := r.getDoc(ctx.GetStart())
	if doc != "" {
		parameter.SetDoc(doc)
	}

	// Add the annotations:
	r.addAnnotations(parameter, ctx.GetAnnotations())

	// Set the direction flags:
	directionsCtxs := ctx.GetDirections()
	if len(directionsCtxs) > 0 {
		for _, directionCtx := range directionsCtxs {
			if directionCtx.GetIn() != nil {
				parameter.SetIn(true)
			}
			if directionCtx.GetOut() != nil {
				parameter.SetOut(true)
			}
		}
	} else {
		parameter.SetOut(true)
	}

	// Set the default value:
	dfltCtx := ctx.GetDflt()
	if dfltCtx != nil {
		parameter.SetDefault(dfltCtx.GetResult())
	}

	// Return the parameter:
	ctx.SetResult(parameter)
}

func (r *Reader) ExitLocatorDecl(ctx *LocatorDeclContext) {
	// Create the locator and set the basic properties:
	locator := concepts.NewLocator()
	locator.SetName(ctx.GetName().GetResult())

	// Add the documentation:
	doc := r.getDoc(ctx.GetStart())
	if doc != "" {
		locator.SetDoc(doc)
	}

	// Add the annotations:
	r.addAnnotations(locator, ctx.GetAnnotations())

	// Add the members:
	membersCtxs := ctx.GetMembers()
	if len(membersCtxs) > 0 {
		for _, memberCtx := range membersCtxs {
			switch member := memberCtx.GetResult().(type) {
			case *names.Name:
				locator.SetVariable(true)
			case *concepts.Resource:
				locator.SetTarget(member)
			}
		}
	}

	// Check that the target resource has been provided:
	if locator.Target() == nil {
		r.reporter.Errorf("Locator '%s' doesn't have a target", locator.Name())
	}

	// Return the parameter:
	ctx.SetResult(locator)
}

func (r *Reader) ExitLocatorMemberDecl(ctx *LocatorMemberDeclContext) {
	if ctx.LocatorTargetDecl() != nil {
		ctx.SetResult(ctx.LocatorTargetDecl().GetResult())
		return
	}
	if ctx.LocatorVariableDecl() != nil {
		ctx.SetResult(ctx.LocatorVariableDecl().GetResult())
		return
	}
}

func (r *Reader) ExitLocatorTargetDecl(ctx *LocatorTargetDeclContext) {
	ctx.SetResult(ctx.GetReference().GetResult())
}

func (r *Reader) ExitLocatorVariableDecl(ctx *LocatorVariableDeclContext) {
	ctx.SetResult(ctx.GetName().GetResult())
}

func (r *Reader) ExitErrorDecl(ctx *ErrorDeclContext) {
	// Check if there is an existing error:
	name := ctx.GetName().GetResult()
	err := r.version.FindError(name)
	if err == nil {
		err = concepts.NewError()
		err.SetName(name)
		r.version.AddError(err)
	} else if r.isUndefinedError(err) {
		r.removeUndefinedError(err)
	} else {
		r.reporter.Errorf("Error '%s' is already defined", name)
		return
	}

	// Add the documentation:
	doc := r.getDoc(ctx.GetStart())
	if doc != "" {
		err.SetDoc(doc)
	}

	// Add the annotations:
	r.addAnnotations(err, ctx.GetAnnotations())

	// Process the members:
	memberCtxs := ctx.GetMembers()
	if len(memberCtxs) > 0 {
		for _, memberCtx := range ctx.GetMembers() {
			switch member := memberCtx.GetResult().(type) {
			case int:
				err.SetCode(member)
			}
		}
	}
}

func (r *Reader) ExitErrorMemberDecl(ctx *ErrorMemberDeclContext) {
	if ctx.ErrorCodeDecl() != nil {
		ctx.SetResult(ctx.ErrorCodeDecl().GetResult())
		return
	}
}

func (r *Reader) ExitErrorCodeDecl(ctx *ErrorCodeDeclContext) {
	text := ctx.GetCode().GetText()
	code, err := strconv.Atoi(text)
	if err != nil {
		r.reporter.Errorf("Error code '%s' isn't a valid integer", text)
	}
	ctx.SetResult(code)
}

func (r *Reader) ExitBooleanLiteral(ctx *BooleanLiteralContext) {
	if ctx.TRUE() != nil {
		ctx.SetResult(true)
		return
	}
	if ctx.FALSE() != nil {
		ctx.SetResult(false)
		return
	}
	r.reporter.Errorf("Unknown boolean literal '%s'", ctx.GetText())
}

func (r *Reader) ExitIntegerLiteral(ctx *IntegerLiteralContext) {
	text := ctx.INTEGER_LITERAL().GetText()
	value, err := strconv.Atoi(text)
	if err != nil {
		r.reporter.Errorf("Incorrect integer literal '%s': %v", text, err)
		return
	}
	ctx.SetResult(value)
}

func (r *Reader) ExitStringLiteral(ctx *StringLiteralContext) {
	text := ctx.STRING_LITERAL().GetText()
	text, err := strconv.Unquote(text)
	if err != nil {
		r.reporter.Errorf("Incorrect string literal '%s': %v", text, err)
		return
	}
	ctx.SetResult(text)
}

func (r *Reader) ExitLiteral(ctx *LiteralContext) {
	switch {
	case ctx.BooleanLiteral() != nil:
		ctx.SetResult(ctx.BooleanLiteral().GetResult())
	case ctx.IntegerLiteral() != nil:
		ctx.SetResult(ctx.IntegerLiteral().GetResult())
	case ctx.StringLiteral() != nil:
		ctx.SetResult(ctx.StringLiteral().GetResult())
	default:
		r.reporter.Errorf("Unknown kind of literal '%s'", ctx.GetText())
	}
}

func (r *Reader) ExitAnnotation(ctx *AnnotationContext) {
	// Create the annotation and set the basic properties:
	annotation := concepts.NewAnnotation()
	annotation.SetName(ctx.Identifier().GetText())

	// Set the parameters:
	if ctx.GetParameters() != nil {
		for name, value := range ctx.GetParameters().GetResult() {
			annotation.AddParameter(name, value)
		}
	}

	// Set the result:
	ctx.SetResult(annotation)
}

func (r *Reader) ExitAnnotationParameters(ctx *AnnotationParametersContext) {
	result := map[string]interface{}{}
	for _, parameterCtx := range ctx.GetParameters() {
		name := parameterCtx.GetName()
		value := parameterCtx.GetValue()
		result[name] = value
	}
	ctx.SetResult(result)
}

func (r *Reader) ExitAnnotationParameter(ctx *AnnotationParameterContext) {
	ctx.SetName(ctx.Identifier().GetText())
	ctx.SetValue(ctx.Literal().GetResult())
}

func (r *Reader) isUndefinedType(typ *concepts.Type) bool {
	key := typ.Name().String()
	_, ok := r.undefinedTypes[key]
	return ok
}

func (r *Reader) addUndefinedType(typ *concepts.Type) {
	key := typ.Name().String()
	r.undefinedTypes[key] = typ
}

func (r *Reader) removeUndefinedType(typ *concepts.Type) {
	key := typ.Name().String()
	delete(r.undefinedTypes, key)
}

func (r *Reader) isUndefinedResource(resource *concepts.Resource) bool {
	key := resource.Name().String()
	_, ok := r.undefinedResources[key]
	return ok
}

func (r *Reader) addUndefinedResource(resource *concepts.Resource) {
	key := resource.Name().String()
	r.undefinedResources[key] = resource
}

func (r *Reader) removeUndefinedResource(typ *concepts.Resource) {
	key := typ.Name().String()
	delete(r.undefinedResources, key)
}

func (r *Reader) isUndefinedError(err *concepts.Error) bool {
	key := err.Name().String()
	_, ok := r.undefinedErrors[key]
	return ok
}

func (r *Reader) addUndefinedError(err *concepts.Error) {
	key := err.Name().String()
	r.undefinedErrors[key] = err
}

func (r *Reader) removeUndefinedError(typ *concepts.Error) {
	key := typ.Name().String()
	delete(r.undefinedErrors, key)
}

// comments stores the line comments that haven't yet been converted into documentation comments:
var comments map[int]string

// getDoc tries to extract the documentation comment that may appear before the given token.
func (r *Reader) getDoc(token antlr.Token) string {
	// Get all the consecutive comments that are right before the token:
	var lines []string
	for index := token.GetLine() - 1; index >= 0; index-- {
		line, ok := comments[index]
		if !ok {
			break
		}
		lines = append([]string{line}, lines...)
		delete(comments, index)
	}

	// Remove the comment mark:
	for i, line := range lines {
		if strings.HasPrefix(line, "//") {
			lines[i] = line[2:]
		} else {
			r.reporter.Errorf(
				"Documentation comment at line %d doesn't start with '//'",
				token.GetLine()-len(lines)+i,
			)
			r.reporter.Errorf("Offending line is '%s'", line)
			return ""
		}
	}

	// Remove from all lines the number of leading spaces of the first line:
	if len(lines) > 0 {
		first := lines[0]
		index := strings.IndexFunc(first, func(r rune) bool {
			return !unicode.IsSpace(r)
		})
		prefix := first[0:index]
		for i, line := range lines {
			if line != "" {
				if strings.HasPrefix(line, prefix) {
					lines[i] = strings.TrimPrefix(line, prefix)
				} else {
					r.reporter.Errorf(
						"Documentation comment at line %d has different "+
							"prefix than first line",
						token.GetLine()-len(lines)+i,
					)
					r.reporter.Errorf("First line is '%s'", first)
					r.reporter.Errorf("Offending line is '%s'", line)
					return ""
				}
			}
		}
	}

	// Return the lines joined:
	return strings.Join(lines, "\n")
}

func (r *Reader) addAnnotations(object concepts.Annotated, ctxs []IAnnotationContext) {
	for _, ctx := range ctxs {
		object.AddAnnotation(ctx.GetResult())
	}
}
