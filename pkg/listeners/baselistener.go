// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package listeners

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/adibrastegarnia/asn2proto/pkg/asn"
)

// BaseASNListener is a complete listener for a parse tree produced by ASNParser.
type BaseASNListener struct {
	*asn.BaseASNListener
}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseASNListener) VisitTerminal(node antlr.TerminalNode) {
}

// VisitErrorNode is called when an error node is visited.
func (s *BaseASNListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseASNListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseASNListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModules is called when production modules is entered.
func (s *BaseASNListener) EnterModules(ctx *asn.ModulesContext) {}

// ExitModules is called when production modules is exited.
func (s *BaseASNListener) ExitModules(ctx *asn.ModulesContext) {}

// EnterModuleDefinition is called when production moduleDefinition is entered.
func (s *BaseASNListener) EnterModuleDefinition(ctx *asn.ModuleDefinitionContext) {}

// ExitModuleDefinition is called when production moduleDefinition is exited.
func (s *BaseASNListener) ExitModuleDefinition(ctx *asn.ModuleDefinitionContext) {}

// EnterTagDefault is called when production tagDefault is entered.
func (s *BaseASNListener) EnterTagDefault(ctx *asn.TagDefaultContext) {}

// ExitTagDefault is called when production tagDefault is exited.
func (s *BaseASNListener) ExitTagDefault(ctx *asn.TagDefaultContext) {}

// EnterExtensionDefault is called when production extensionDefault is entered.
func (s *BaseASNListener) EnterExtensionDefault(ctx *asn.ExtensionDefaultContext) {}

// ExitExtensionDefault is called when production extensionDefault is exited.
func (s *BaseASNListener) ExitExtensionDefault(ctx *asn.ExtensionDefaultContext) {}

// EnterModuleBody is called when production moduleBody is entered.
func (s *BaseASNListener) EnterModuleBody(ctx *asn.ModuleBodyContext) {}

// ExitModuleBody is called when production moduleBody is exited.
func (s *BaseASNListener) ExitModuleBody(ctx *asn.ModuleBodyContext) {}

// EnterExports is called when production exports is entered.
func (s *BaseASNListener) EnterExports(ctx *asn.ExportsContext) {}

// ExitExports is called when production exports is exited.
func (s *BaseASNListener) ExitExports(ctx *asn.ExportsContext) {}

// EnterSymbolsExported is called when production symbolsExported is entered.
func (s *BaseASNListener) EnterSymbolsExported(ctx *asn.SymbolsExportedContext) {}

// ExitSymbolsExported is called when production symbolsExported is exited.
func (s *BaseASNListener) ExitSymbolsExported(ctx *asn.SymbolsExportedContext) {}

// EnterImports is called when production imports is entered.
func (s *BaseASNListener) EnterImports(ctx *asn.ImportsContext) {}

// ExitImports is called when production imports is exited.
func (s *BaseASNListener) ExitImports(ctx *asn.ImportsContext) {}

// EnterSymbolsImported is called when production symbolsImported is entered.
func (s *BaseASNListener) EnterSymbolsImported(ctx *asn.SymbolsImportedContext) {}

// ExitSymbolsImported is called when production symbolsImported is exited.
func (s *BaseASNListener) ExitSymbolsImported(ctx *asn.SymbolsImportedContext) {}

// EnterSymbolsFromModuleList is called when production symbolsFromModuleList is entered.
func (s *BaseASNListener) EnterSymbolsFromModuleList(ctx *asn.SymbolsFromModuleListContext) {}

// ExitSymbolsFromModuleList is called when production symbolsFromModuleList is exited.
func (s *BaseASNListener) ExitSymbolsFromModuleList(ctx *asn.SymbolsFromModuleListContext) {}

// EnterSymbolsFromModule is called when production symbolsFromModule is entered.
func (s *BaseASNListener) EnterSymbolsFromModule(ctx *asn.SymbolsFromModuleContext) {}

// ExitSymbolsFromModule is called when production symbolsFromModule is exited.
func (s *BaseASNListener) ExitSymbolsFromModule(ctx *asn.SymbolsFromModuleContext) {}

// EnterGlobalModuleReference is called when production globalModuleReference is entered.
func (s *BaseASNListener) EnterGlobalModuleReference(ctx *asn.GlobalModuleReferenceContext) {}

// ExitGlobalModuleReference is called when production globalModuleReference is exited.
func (s *BaseASNListener) ExitGlobalModuleReference(ctx *asn.GlobalModuleReferenceContext) {}

// EnterAssignedIdentifier is called when production assignedIdentifier is entered.
func (s *BaseASNListener) EnterAssignedIdentifier(ctx *asn.AssignedIdentifierContext) {}

// ExitAssignedIdentifier is called when production assignedIdentifier is exited.
func (s *BaseASNListener) ExitAssignedIdentifier(ctx *asn.AssignedIdentifierContext) {}

// EnterSymbolList is called when production symbolList is entered.
func (s *BaseASNListener) EnterSymbolList(ctx *asn.SymbolListContext) {}

// ExitSymbolList is called when production symbolList is exited.
func (s *BaseASNListener) ExitSymbolList(ctx *asn.SymbolListContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseASNListener) EnterSymbol(ctx *asn.SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseASNListener) ExitSymbol(ctx *asn.SymbolContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseASNListener) EnterAssignmentList(ctx *asn.AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseASNListener) ExitAssignmentList(ctx *asn.AssignmentListContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseASNListener) EnterAssignment(ctx *asn.AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseASNListener) ExitAssignment(ctx *asn.AssignmentContext) {}

// EnterSequenceType is called when production sequenceType is entered.
func (s *BaseASNListener) EnterSequenceType(ctx *asn.SequenceTypeContext) {}

// ExitSequenceType is called when production sequenceType is exited.
func (s *BaseASNListener) ExitSequenceType(ctx *asn.SequenceTypeContext) {}

// EnterExtensionAndException is called when production extensionAndException is entered.
func (s *BaseASNListener) EnterExtensionAndException(ctx *asn.ExtensionAndExceptionContext) {}

// ExitExtensionAndException is called when production extensionAndException is exited.
func (s *BaseASNListener) ExitExtensionAndException(ctx *asn.ExtensionAndExceptionContext) {}

// EnterOptionalExtensionMarker is called when production optionalExtensionMarker is entered.
func (s *BaseASNListener) EnterOptionalExtensionMarker(ctx *asn.OptionalExtensionMarkerContext) {}

// ExitOptionalExtensionMarker is called when production optionalExtensionMarker is exited.
func (s *BaseASNListener) ExitOptionalExtensionMarker(ctx *asn.OptionalExtensionMarkerContext) {}

// EnterComponentTypeLists is called when production componentTypeLists is entered.
func (s *BaseASNListener) EnterComponentTypeLists(ctx *asn.ComponentTypeListsContext) {}

// ExitComponentTypeLists is called when production componentTypeLists is exited.
func (s *BaseASNListener) ExitComponentTypeLists(ctx *asn.ComponentTypeListsContext) {}

// EnterRootComponentTypeList is called when production rootComponentTypeList is entered.
func (s *BaseASNListener) EnterRootComponentTypeList(ctx *asn.RootComponentTypeListContext) {}

// ExitRootComponentTypeList is called when production rootComponentTypeList is exited.
func (s *BaseASNListener) ExitRootComponentTypeList(ctx *asn.RootComponentTypeListContext) {}

// EnterComponentTypeList is called when production componentTypeList is entered.
func (s *BaseASNListener) EnterComponentTypeList(ctx *asn.ComponentTypeListContext) {}

// ExitComponentTypeList is called when production componentTypeList is exited.
func (s *BaseASNListener) ExitComponentTypeList(ctx *asn.ComponentTypeListContext) {}

// EnterComponentType is called when production componentType is entered.
func (s *BaseASNListener) EnterComponentType(ctx *asn.ComponentTypeContext) {}

// ExitComponentType is called when production componentType is exited.
func (s *BaseASNListener) ExitComponentType(ctx *asn.ComponentTypeContext) {}

// EnterExtensionAdditions is called when production extensionAdditions is entered.
func (s *BaseASNListener) EnterExtensionAdditions(ctx *asn.ExtensionAdditionsContext) {}

// ExitExtensionAdditions is called when production extensionAdditions is exited.
func (s *BaseASNListener) ExitExtensionAdditions(ctx *asn.ExtensionAdditionsContext) {}

// EnterExtensionAdditionList is called when production extensionAdditionList is entered.
func (s *BaseASNListener) EnterExtensionAdditionList(ctx *asn.ExtensionAdditionListContext) {}

// ExitExtensionAdditionList is called when production extensionAdditionList is exited.
func (s *BaseASNListener) ExitExtensionAdditionList(ctx *asn.ExtensionAdditionListContext) {}

// EnterExtensionAddition is called when production extensionAddition is entered.
func (s *BaseASNListener) EnterExtensionAddition(ctx *asn.ExtensionAdditionContext) {}

// ExitExtensionAddition is called when production extensionAddition is exited.
func (s *BaseASNListener) ExitExtensionAddition(ctx *asn.ExtensionAdditionContext) {}

// EnterExtensionAdditionGroup is called when production extensionAdditionGroup is entered.
func (s *BaseASNListener) EnterExtensionAdditionGroup(ctx *asn.ExtensionAdditionGroupContext) {}

// ExitExtensionAdditionGroup is called when production extensionAdditionGroup is exited.
func (s *BaseASNListener) ExitExtensionAdditionGroup(ctx *asn.ExtensionAdditionGroupContext) {}

// EnterVersionNumber is called when production versionNumber is entered.
func (s *BaseASNListener) EnterVersionNumber(ctx *asn.VersionNumberContext) {}

// ExitVersionNumber is called when production versionNumber is exited.
func (s *BaseASNListener) ExitVersionNumber(ctx *asn.VersionNumberContext) {}

// EnterSequenceOfType is called when production sequenceOfType is entered.
func (s *BaseASNListener) EnterSequenceOfType(ctx *asn.SequenceOfTypeContext) {}

// ExitSequenceOfType is called when production sequenceOfType is exited.
func (s *BaseASNListener) ExitSequenceOfType(ctx *asn.SequenceOfTypeContext) {}

// EnterSizeConstraint is called when production sizeConstraint is entered.
func (s *BaseASNListener) EnterSizeConstraint(ctx *asn.SizeConstraintContext) {}

// ExitSizeConstraint is called when production sizeConstraint is exited.
func (s *BaseASNListener) ExitSizeConstraint(ctx *asn.SizeConstraintContext) {}

// EnterParameterizedAssignment is called when production parameterizedAssignment is entered.
func (s *BaseASNListener) EnterParameterizedAssignment(ctx *asn.ParameterizedAssignmentContext) {}

// ExitParameterizedAssignment is called when production parameterizedAssignment is exited.
func (s *BaseASNListener) ExitParameterizedAssignment(ctx *asn.ParameterizedAssignmentContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseASNListener) EnterParameterList(ctx *asn.ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseASNListener) ExitParameterList(ctx *asn.ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseASNListener) EnterParameter(ctx *asn.ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseASNListener) ExitParameter(ctx *asn.ParameterContext) {}

// EnterParamGovernor is called when production paramGovernor is entered.
func (s *BaseASNListener) EnterParamGovernor(ctx *asn.ParamGovernorContext) {}

// ExitParamGovernor is called when production paramGovernor is exited.
func (s *BaseASNListener) ExitParamGovernor(ctx *asn.ParamGovernorContext) {}

// EnterGovernor is called when production governor is entered.
func (s *BaseASNListener) EnterGovernor(ctx *asn.GovernorContext) {}

// ExitGovernor is called when production governor is exited.
func (s *BaseASNListener) ExitGovernor(ctx *asn.GovernorContext) {}

// EnterObjectClassAssignment is called when production objectClassAssignment is entered.
func (s *BaseASNListener) EnterObjectClassAssignment(ctx *asn.ObjectClassAssignmentContext) {}

// ExitObjectClassAssignment is called when production objectClassAssignment is exited.
func (s *BaseASNListener) ExitObjectClassAssignment(ctx *asn.ObjectClassAssignmentContext) {}

// EnterObjectClass is called when production objectClass is entered.
func (s *BaseASNListener) EnterObjectClass(ctx *asn.ObjectClassContext) {}

// ExitObjectClass is called when production objectClass is exited.
func (s *BaseASNListener) ExitObjectClass(ctx *asn.ObjectClassContext) {}

// EnterDefinedObjectClass is called when production definedObjectClass is entered.
func (s *BaseASNListener) EnterDefinedObjectClass(ctx *asn.DefinedObjectClassContext) {}

// ExitDefinedObjectClass is called when production definedObjectClass is exited.
func (s *BaseASNListener) ExitDefinedObjectClass(ctx *asn.DefinedObjectClassContext) {}

// EnterUsefulObjectClassReference is called when production usefulObjectClassReference is entered.
func (s *BaseASNListener) EnterUsefulObjectClassReference(ctx *asn.UsefulObjectClassReferenceContext) {
}

// ExitUsefulObjectClassReference is called when production usefulObjectClassReference is exited.
func (s *BaseASNListener) ExitUsefulObjectClassReference(ctx *asn.UsefulObjectClassReferenceContext) {}

// EnterExternalObjectClassReference is called when production externalObjectClassReference is entered.
func (s *BaseASNListener) EnterExternalObjectClassReference(ctx *asn.ExternalObjectClassReferenceContext) {
}

// ExitExternalObjectClassReference is called when production externalObjectClassReference is exited.
func (s *BaseASNListener) ExitExternalObjectClassReference(ctx *asn.ExternalObjectClassReferenceContext) {
}

// EnterObjectClassDefn is called when production objectClassDefn is entered.
func (s *BaseASNListener) EnterObjectClassDefn(ctx *asn.ObjectClassDefnContext) {}

// ExitObjectClassDefn is called when production objectClassDefn is exited.
func (s *BaseASNListener) ExitObjectClassDefn(ctx *asn.ObjectClassDefnContext) {}

// EnterWithSyntaxSpec is called when production withSyntaxSpec is entered.
func (s *BaseASNListener) EnterWithSyntaxSpec(ctx *asn.WithSyntaxSpecContext) {}

// ExitWithSyntaxSpec is called when production withSyntaxSpec is exited.
func (s *BaseASNListener) ExitWithSyntaxSpec(ctx *asn.WithSyntaxSpecContext) {}

// EnterSyntaxList is called when production syntaxList is entered.
func (s *BaseASNListener) EnterSyntaxList(ctx *asn.SyntaxListContext) {}

// ExitSyntaxList is called when production syntaxList is exited.
func (s *BaseASNListener) ExitSyntaxList(ctx *asn.SyntaxListContext) {}

// EnterTokenOrGroupSpec is called when production tokenOrGroupSpec is entered.
func (s *BaseASNListener) EnterTokenOrGroupSpec(ctx *asn.TokenOrGroupSpecContext) {}

// ExitTokenOrGroupSpec is called when production tokenOrGroupSpec is exited.
func (s *BaseASNListener) ExitTokenOrGroupSpec(ctx *asn.TokenOrGroupSpecContext) {}

// EnterOptionalGroup is called when production optionalGroup is entered.
func (s *BaseASNListener) EnterOptionalGroup(ctx *asn.OptionalGroupContext) {}

// ExitOptionalGroup is called when production optionalGroup is exited.
func (s *BaseASNListener) ExitOptionalGroup(ctx *asn.OptionalGroupContext) {}

// EnterRequiredToken is called when production requiredToken is entered.
func (s *BaseASNListener) EnterRequiredToken(ctx *asn.RequiredTokenContext) {}

// ExitRequiredToken is called when production requiredToken is exited.
func (s *BaseASNListener) ExitRequiredToken(ctx *asn.RequiredTokenContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseASNListener) EnterLiteral(ctx *asn.LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseASNListener) ExitLiteral(ctx *asn.LiteralContext) {}

// EnterPrimitiveFieldName is called when production primitiveFieldName is entered.
func (s *BaseASNListener) EnterPrimitiveFieldName(ctx *asn.PrimitiveFieldNameContext) {}

// ExitPrimitiveFieldName is called when production primitiveFieldName is exited.
func (s *BaseASNListener) ExitPrimitiveFieldName(ctx *asn.PrimitiveFieldNameContext) {}

// EnterFieldSpec is called when production fieldSpec is entered.
func (s *BaseASNListener) EnterFieldSpec(ctx *asn.FieldSpecContext) {}

// ExitFieldSpec is called when production fieldSpec is exited.
func (s *BaseASNListener) ExitFieldSpec(ctx *asn.FieldSpecContext) {}

// EnterTypeFieldSpec is called when production typeFieldSpec is entered.
func (s *BaseASNListener) EnterTypeFieldSpec(ctx *asn.TypeFieldSpecContext) {}

// ExitTypeFieldSpec is called when production typeFieldSpec is exited.
func (s *BaseASNListener) ExitTypeFieldSpec(ctx *asn.TypeFieldSpecContext) {}

// EnterTypeOptionalitySpec is called when production typeOptionalitySpec is entered.
func (s *BaseASNListener) EnterTypeOptionalitySpec(ctx *asn.TypeOptionalitySpecContext) {}

// ExitTypeOptionalitySpec is called when production typeOptionalitySpec is exited.
func (s *BaseASNListener) ExitTypeOptionalitySpec(ctx *asn.TypeOptionalitySpecContext) {}

// EnterFixedTypeValueFieldSpec is called when production fixedTypeValueFieldSpec is entered.
func (s *BaseASNListener) EnterFixedTypeValueFieldSpec(ctx *asn.FixedTypeValueFieldSpecContext) {}

// ExitFixedTypeValueFieldSpec is called when production fixedTypeValueFieldSpec is exited.
func (s *BaseASNListener) ExitFixedTypeValueFieldSpec(ctx *asn.FixedTypeValueFieldSpecContext) {}

// EnterValueOptionalitySpec is called when production valueOptionalitySpec is entered.
func (s *BaseASNListener) EnterValueOptionalitySpec(ctx *asn.ValueOptionalitySpecContext) {}

// ExitValueOptionalitySpec is called when production valueOptionalitySpec is exited.
func (s *BaseASNListener) ExitValueOptionalitySpec(ctx *asn.ValueOptionalitySpecContext) {}

// EnterVariableTypeValueFieldSpec is called when production variableTypeValueFieldSpec is entered.
func (s *BaseASNListener) EnterVariableTypeValueFieldSpec(ctx *asn.VariableTypeValueFieldSpecContext) {
}

// ExitVariableTypeValueFieldSpec is called when production variableTypeValueFieldSpec is exited.
func (s *BaseASNListener) ExitVariableTypeValueFieldSpec(ctx *asn.VariableTypeValueFieldSpecContext) {}

// EnterFixedTypeValueSetFieldSpec is called when production fixedTypeValueSetFieldSpec is entered.
func (s *BaseASNListener) EnterFixedTypeValueSetFieldSpec(ctx *asn.FixedTypeValueSetFieldSpecContext) {
}

// ExitFixedTypeValueSetFieldSpec is called when production fixedTypeValueSetFieldSpec is exited.
func (s *BaseASNListener) ExitFixedTypeValueSetFieldSpec(ctx *asn.FixedTypeValueSetFieldSpecContext) {}

// EnterValueSetOptionalitySpec is called when production valueSetOptionalitySpec is entered.
func (s *BaseASNListener) EnterValueSetOptionalitySpec(ctx *asn.ValueSetOptionalitySpecContext) {}

// ExitValueSetOptionalitySpec is called when production valueSetOptionalitySpec is exited.
func (s *BaseASNListener) ExitValueSetOptionalitySpec(ctx *asn.ValueSetOptionalitySpecContext) {}

// EnterObject is called when production object is entered.
func (s *BaseASNListener) EnterObject(ctx *asn.ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseASNListener) ExitObject(ctx *asn.ObjectContext) {}

// EnterParameterizedObject is called when production parameterizedObject is entered.
func (s *BaseASNListener) EnterParameterizedObject(ctx *asn.ParameterizedObjectContext) {}

// ExitParameterizedObject is called when production parameterizedObject is exited.
func (s *BaseASNListener) ExitParameterizedObject(ctx *asn.ParameterizedObjectContext) {}

// EnterDefinedObject is called when production definedObject is entered.
func (s *BaseASNListener) EnterDefinedObject(ctx *asn.DefinedObjectContext) {}

// ExitDefinedObject is called when production definedObject is exited.
func (s *BaseASNListener) ExitDefinedObject(ctx *asn.DefinedObjectContext) {}

// EnterObjectSet is called when production objectSet is entered.
func (s *BaseASNListener) EnterObjectSet(ctx *asn.ObjectSetContext) {}

// ExitObjectSet is called when production objectSet is exited.
func (s *BaseASNListener) ExitObjectSet(ctx *asn.ObjectSetContext) {}

// EnterObjectSetSpec is called when production objectSetSpec is entered.
func (s *BaseASNListener) EnterObjectSetSpec(ctx *asn.ObjectSetSpecContext) {}

// ExitObjectSetSpec is called when production objectSetSpec is exited.
func (s *BaseASNListener) ExitObjectSetSpec(ctx *asn.ObjectSetSpecContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseASNListener) EnterFieldName(ctx *asn.FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseASNListener) ExitFieldName(ctx *asn.FieldNameContext) {}

// EnterValueSet is called when production valueSet is entered.
func (s *BaseASNListener) EnterValueSet(ctx *asn.ValueSetContext) {}

// ExitValueSet is called when production valueSet is exited.
func (s *BaseASNListener) ExitValueSet(ctx *asn.ValueSetContext) {}

// EnterElementSetSpecs is called when production elementSetSpecs is entered.
func (s *BaseASNListener) EnterElementSetSpecs(ctx *asn.ElementSetSpecsContext) {}

// ExitElementSetSpecs is called when production elementSetSpecs is exited.
func (s *BaseASNListener) ExitElementSetSpecs(ctx *asn.ElementSetSpecsContext) {}

// EnterRootElementSetSpec is called when production rootElementSetSpec is entered.
func (s *BaseASNListener) EnterRootElementSetSpec(ctx *asn.RootElementSetSpecContext) {}

// ExitRootElementSetSpec is called when production rootElementSetSpec is exited.
func (s *BaseASNListener) ExitRootElementSetSpec(ctx *asn.RootElementSetSpecContext) {}

// EnterAdditionalElementSetSpec is called when production additionalElementSetSpec is entered.
func (s *BaseASNListener) EnterAdditionalElementSetSpec(ctx *asn.AdditionalElementSetSpecContext) {}

// ExitAdditionalElementSetSpec is called when production additionalElementSetSpec is exited.
func (s *BaseASNListener) ExitAdditionalElementSetSpec(ctx *asn.AdditionalElementSetSpecContext) {}

// EnterElementSetSpec is called when production elementSetSpec is entered.
func (s *BaseASNListener) EnterElementSetSpec(ctx *asn.ElementSetSpecContext) {}

// ExitElementSetSpec is called when production elementSetSpec is exited.
func (s *BaseASNListener) ExitElementSetSpec(ctx *asn.ElementSetSpecContext) {}

// EnterUnions is called when production unions is entered.
func (s *BaseASNListener) EnterUnions(ctx *asn.UnionsContext) {}

// ExitUnions is called when production unions is exited.
func (s *BaseASNListener) ExitUnions(ctx *asn.UnionsContext) {}

// EnterExclusions is called when production exclusions is entered.
func (s *BaseASNListener) EnterExclusions(ctx *asn.ExclusionsContext) {}

// ExitExclusions is called when production exclusions is exited.
func (s *BaseASNListener) ExitExclusions(ctx *asn.ExclusionsContext) {}

// EnterIntersections is called when production intersections is entered.
func (s *BaseASNListener) EnterIntersections(ctx *asn.IntersectionsContext) {}

// ExitIntersections is called when production intersections is exited.
func (s *BaseASNListener) ExitIntersections(ctx *asn.IntersectionsContext) {}

// EnterUnionMark is called when production unionMark is entered.
func (s *BaseASNListener) EnterUnionMark(ctx *asn.UnionMarkContext) {}

// ExitUnionMark is called when production unionMark is exited.
func (s *BaseASNListener) ExitUnionMark(ctx *asn.UnionMarkContext) {}

// EnterIntersectionMark is called when production intersectionMark is entered.
func (s *BaseASNListener) EnterIntersectionMark(ctx *asn.IntersectionMarkContext) {}

// ExitIntersectionMark is called when production intersectionMark is exited.
func (s *BaseASNListener) ExitIntersectionMark(ctx *asn.IntersectionMarkContext) {}

// EnterElements is called when production elements is entered.
func (s *BaseASNListener) EnterElements(ctx *asn.ElementsContext) {}

// ExitElements is called when production elements is exited.
func (s *BaseASNListener) ExitElements(ctx *asn.ElementsContext) {}

// EnterObjectSetElements is called when production objectSetElements is entered.
func (s *BaseASNListener) EnterObjectSetElements(ctx *asn.ObjectSetElementsContext) {}

// ExitObjectSetElements is called when production objectSetElements is exited.
func (s *BaseASNListener) ExitObjectSetElements(ctx *asn.ObjectSetElementsContext) {}

// EnterIntersectionElements is called when production intersectionElements is entered.
func (s *BaseASNListener) EnterIntersectionElements(ctx *asn.IntersectionElementsContext) {}

// ExitIntersectionElements is called when production intersectionElements is exited.
func (s *BaseASNListener) ExitIntersectionElements(ctx *asn.IntersectionElementsContext) {}

// EnterSubtypeElements is called when production subtypeElements is entered.
func (s *BaseASNListener) EnterSubtypeElements(ctx *asn.SubtypeElementsContext) {}

// ExitSubtypeElements is called when production subtypeElements is exited.
func (s *BaseASNListener) ExitSubtypeElements(ctx *asn.SubtypeElementsContext) {}

// EnterVariableTypeValueSetFieldSpec is called when production variableTypeValueSetFieldSpec is entered.
func (s *BaseASNListener) EnterVariableTypeValueSetFieldSpec(ctx *asn.VariableTypeValueSetFieldSpecContext) {
}

// ExitVariableTypeValueSetFieldSpec is called when production variableTypeValueSetFieldSpec is exited.
func (s *BaseASNListener) ExitVariableTypeValueSetFieldSpec(ctx *asn.VariableTypeValueSetFieldSpecContext) {
}

// EnterObjectFieldSpec is called when production objectFieldSpec is entered.
func (s *BaseASNListener) EnterObjectFieldSpec(ctx *asn.ObjectFieldSpecContext) {}

// ExitObjectFieldSpec is called when production objectFieldSpec is exited.
func (s *BaseASNListener) ExitObjectFieldSpec(ctx *asn.ObjectFieldSpecContext) {}

// EnterObjectOptionalitySpec is called when production objectOptionalitySpec is entered.
func (s *BaseASNListener) EnterObjectOptionalitySpec(ctx *asn.ObjectOptionalitySpecContext) {}

// ExitObjectOptionalitySpec is called when production objectOptionalitySpec is exited.
func (s *BaseASNListener) ExitObjectOptionalitySpec(ctx *asn.ObjectOptionalitySpecContext) {}

// EnterObjectSetFieldSpec is called when production objectSetFieldSpec is entered.
func (s *BaseASNListener) EnterObjectSetFieldSpec(ctx *asn.ObjectSetFieldSpecContext) {}

// ExitObjectSetFieldSpec is called when production objectSetFieldSpec is exited.
func (s *BaseASNListener) ExitObjectSetFieldSpec(ctx *asn.ObjectSetFieldSpecContext) {}

// EnterObjectSetOptionalitySpec is called when production objectSetOptionalitySpec is entered.
func (s *BaseASNListener) EnterObjectSetOptionalitySpec(ctx *asn.ObjectSetOptionalitySpecContext) {}

// ExitObjectSetOptionalitySpec is called when production objectSetOptionalitySpec is exited.
func (s *BaseASNListener) ExitObjectSetOptionalitySpec(ctx *asn.ObjectSetOptionalitySpecContext) {}

// EnterTypeAssignment is called when production typeAssignment is entered.
func (s *BaseASNListener) EnterTypeAssignment(ctx *asn.TypeAssignmentContext) {}

// ExitTypeAssignment is called when production typeAssignment is exited.
func (s *BaseASNListener) ExitTypeAssignment(ctx *asn.TypeAssignmentContext) {}

// EnterValueAssignment is called when production valueAssignment is entered.
func (s *BaseASNListener) EnterValueAssignment(ctx *asn.ValueAssignmentContext) {}

// ExitValueAssignment is called when production valueAssignment is exited.
func (s *BaseASNListener) ExitValueAssignment(ctx *asn.ValueAssignmentContext) {}

// EnterAsnType is called when production asnType is entered.
func (s *BaseASNListener) EnterAsnType(ctx *asn.AsnTypeContext) {}

// ExitAsnType is called when production asnType is exited.
func (s *BaseASNListener) ExitAsnType(ctx *asn.AsnTypeContext) {}

// EnterBuiltinType is called when production builtinType is entered.
func (s *BaseASNListener) EnterBuiltinType(ctx *asn.BuiltinTypeContext) {}

// ExitBuiltinType is called when production builtinType is exited.
func (s *BaseASNListener) ExitBuiltinType(ctx *asn.BuiltinTypeContext) {}

// EnterObjectClassFieldType is called when production objectClassFieldType is entered.
func (s *BaseASNListener) EnterObjectClassFieldType(ctx *asn.ObjectClassFieldTypeContext) {}

// ExitObjectClassFieldType is called when production objectClassFieldType is exited.
func (s *BaseASNListener) ExitObjectClassFieldType(ctx *asn.ObjectClassFieldTypeContext) {}

// EnterSetType is called when production setType is entered.
func (s *BaseASNListener) EnterSetType(ctx *asn.SetTypeContext) {}

// ExitSetType is called when production setType is exited.
func (s *BaseASNListener) ExitSetType(ctx *asn.SetTypeContext) {}

// EnterSetOfType is called when production setOfType is entered.
func (s *BaseASNListener) EnterSetOfType(ctx *asn.SetOfTypeContext) {}

// ExitSetOfType is called when production setOfType is exited.
func (s *BaseASNListener) ExitSetOfType(ctx *asn.SetOfTypeContext) {}

// EnterReferencedType is called when production referencedType is entered.
func (s *BaseASNListener) EnterReferencedType(ctx *asn.ReferencedTypeContext) {}

// ExitReferencedType is called when production referencedType is exited.
func (s *BaseASNListener) ExitReferencedType(ctx *asn.ReferencedTypeContext) {}

// EnterDefinedType is called when production definedType is entered.
func (s *BaseASNListener) EnterDefinedType(ctx *asn.DefinedTypeContext) {}

// ExitDefinedType is called when production definedType is exited.
func (s *BaseASNListener) ExitDefinedType(ctx *asn.DefinedTypeContext) {}

// EnterConstraint is called when production constraint is entered.
func (s *BaseASNListener) EnterConstraint(ctx *asn.ConstraintContext) {}

// ExitConstraint is called when production constraint is exited.
func (s *BaseASNListener) ExitConstraint(ctx *asn.ConstraintContext) {}

// EnterConstraintSpec is called when production constraintSpec is entered.
func (s *BaseASNListener) EnterConstraintSpec(ctx *asn.ConstraintSpecContext) {}

// ExitConstraintSpec is called when production constraintSpec is exited.
func (s *BaseASNListener) ExitConstraintSpec(ctx *asn.ConstraintSpecContext) {}

// EnterUserDefinedConstraint is called when production userDefinedConstraint is entered.
func (s *BaseASNListener) EnterUserDefinedConstraint(ctx *asn.UserDefinedConstraintContext) {}

// ExitUserDefinedConstraint is called when production userDefinedConstraint is exited.
func (s *BaseASNListener) ExitUserDefinedConstraint(ctx *asn.UserDefinedConstraintContext) {}

// EnterGeneralConstraint is called when production generalConstraint is entered.
func (s *BaseASNListener) EnterGeneralConstraint(ctx *asn.GeneralConstraintContext) {}

// ExitGeneralConstraint is called when production generalConstraint is exited.
func (s *BaseASNListener) ExitGeneralConstraint(ctx *asn.GeneralConstraintContext) {}

// EnterUserDefinedConstraintParameter is called when production userDefinedConstraintParameter is entered.
func (s *BaseASNListener) EnterUserDefinedConstraintParameter(ctx *asn.UserDefinedConstraintParameterContext) {
}

// ExitUserDefinedConstraintParameter is called when production userDefinedConstraintParameter is exited.
func (s *BaseASNListener) ExitUserDefinedConstraintParameter(ctx *asn.UserDefinedConstraintParameterContext) {
}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseASNListener) EnterTableConstraint(ctx *asn.TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseASNListener) ExitTableConstraint(ctx *asn.TableConstraintContext) {}

// EnterSimpleTableConstraint is called when production simpleTableConstraint is entered.
func (s *BaseASNListener) EnterSimpleTableConstraint(ctx *asn.SimpleTableConstraintContext) {}

// ExitSimpleTableConstraint is called when production simpleTableConstraint is exited.
func (s *BaseASNListener) ExitSimpleTableConstraint(ctx *asn.SimpleTableConstraintContext) {}

// EnterContentsConstraint is called when production contentsConstraint is entered.
func (s *BaseASNListener) EnterContentsConstraint(ctx *asn.ContentsConstraintContext) {}

// ExitContentsConstraint is called when production contentsConstraint is exited.
func (s *BaseASNListener) ExitContentsConstraint(ctx *asn.ContentsConstraintContext) {}

// EnterComponentPresenceLists is called when production componentPresenceLists is entered.
func (s *BaseASNListener) EnterComponentPresenceLists(ctx *asn.ComponentPresenceListsContext) {}

// ExitComponentPresenceLists is called when production componentPresenceLists is exited.
func (s *BaseASNListener) ExitComponentPresenceLists(ctx *asn.ComponentPresenceListsContext) {}

// EnterComponentPresenceList is called when production componentPresenceList is entered.
func (s *BaseASNListener) EnterComponentPresenceList(ctx *asn.ComponentPresenceListContext) {}

// ExitComponentPresenceList is called when production componentPresenceList is exited.
func (s *BaseASNListener) ExitComponentPresenceList(ctx *asn.ComponentPresenceListContext) {}

// EnterComponentPresence is called when production componentPresence is entered.
func (s *BaseASNListener) EnterComponentPresence(ctx *asn.ComponentPresenceContext) {}

// ExitComponentPresence is called when production componentPresence is exited.
func (s *BaseASNListener) ExitComponentPresence(ctx *asn.ComponentPresenceContext) {}

// EnterSubtypeConstraint is called when production subtypeConstraint is entered.
func (s *BaseASNListener) EnterSubtypeConstraint(ctx *asn.SubtypeConstraintContext) {}

// ExitSubtypeConstraint is called when production subtypeConstraint is exited.
func (s *BaseASNListener) ExitSubtypeConstraint(ctx *asn.SubtypeConstraintContext) {}

// EnterValue is called when production value is entered.
func (s *BaseASNListener) EnterValue(ctx *asn.ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseASNListener) ExitValue(ctx *asn.ValueContext) {}

// EnterBuiltinValue is called when production builtinValue is entered.
func (s *BaseASNListener) EnterBuiltinValue(ctx *asn.BuiltinValueContext) {}

// ExitBuiltinValue is called when production builtinValue is exited.
func (s *BaseASNListener) ExitBuiltinValue(ctx *asn.BuiltinValueContext) {}

// EnterObjectIdentifierValue is called when production objectIdentifierValue is entered.
func (s *BaseASNListener) EnterObjectIdentifierValue(ctx *asn.ObjectIdentifierValueContext) {}

// ExitObjectIdentifierValue is called when production objectIdentifierValue is exited.
func (s *BaseASNListener) ExitObjectIdentifierValue(ctx *asn.ObjectIdentifierValueContext) {}

// EnterObjIdComponentsList is called when production objIdComponentsList is entered.
func (s *BaseASNListener) EnterObjIdComponentsList(ctx *asn.ObjIdComponentsListContext) {}

// ExitObjIdComponentsList is called when production objIdComponentsList is exited.
func (s *BaseASNListener) ExitObjIdComponentsList(ctx *asn.ObjIdComponentsListContext) {}

// EnterObjIdComponents is called when production objIdComponents is entered.
func (s *BaseASNListener) EnterObjIdComponents(ctx *asn.ObjIdComponentsContext) {}

// ExitObjIdComponents is called when production objIdComponents is exited.
func (s *BaseASNListener) ExitObjIdComponents(ctx *asn.ObjIdComponentsContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseASNListener) EnterIntegerValue(ctx *asn.IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseASNListener) ExitIntegerValue(ctx *asn.IntegerValueContext) {}

// EnterChoiceValue is called when production choiceValue is entered.
func (s *BaseASNListener) EnterChoiceValue(ctx *asn.ChoiceValueContext) {}

// ExitChoiceValue is called when production choiceValue is exited.
func (s *BaseASNListener) ExitChoiceValue(ctx *asn.ChoiceValueContext) {}

// EnterEnumeratedValue is called when production enumeratedValue is entered.
func (s *BaseASNListener) EnterEnumeratedValue(ctx *asn.EnumeratedValueContext) {}

// ExitEnumeratedValue is called when production enumeratedValue is exited.
func (s *BaseASNListener) ExitEnumeratedValue(ctx *asn.EnumeratedValueContext) {}

// EnterSignedNumber is called when production signedNumber is entered.
func (s *BaseASNListener) EnterSignedNumber(ctx *asn.SignedNumberContext) {}

// ExitSignedNumber is called when production signedNumber is exited.
func (s *BaseASNListener) ExitSignedNumber(ctx *asn.SignedNumberContext) {}

// EnterChoiceType is called when production choiceType is entered.
func (s *BaseASNListener) EnterChoiceType(ctx *asn.ChoiceTypeContext) {}

// ExitChoiceType is called when production choiceType is exited.
func (s *BaseASNListener) ExitChoiceType(ctx *asn.ChoiceTypeContext) {}

// EnterAlternativeTypeLists is called when production alternativeTypeLists is entered.
func (s *BaseASNListener) EnterAlternativeTypeLists(ctx *asn.AlternativeTypeListsContext) {}

// ExitAlternativeTypeLists is called when production alternativeTypeLists is exited.
func (s *BaseASNListener) ExitAlternativeTypeLists(ctx *asn.AlternativeTypeListsContext) {}

// EnterExtensionAdditionAlternatives is called when production extensionAdditionAlternatives is entered.
func (s *BaseASNListener) EnterExtensionAdditionAlternatives(ctx *asn.ExtensionAdditionAlternativesContext) {
}

// ExitExtensionAdditionAlternatives is called when production extensionAdditionAlternatives is exited.
func (s *BaseASNListener) ExitExtensionAdditionAlternatives(ctx *asn.ExtensionAdditionAlternativesContext) {
}

// EnterExtensionAdditionAlternativesList is called when production extensionAdditionAlternativesList is entered.
func (s *BaseASNListener) EnterExtensionAdditionAlternativesList(ctx *asn.ExtensionAdditionAlternativesListContext) {
}

// ExitExtensionAdditionAlternativesList is called when production extensionAdditionAlternativesList is exited.
func (s *BaseASNListener) ExitExtensionAdditionAlternativesList(ctx *asn.ExtensionAdditionAlternativesListContext) {
}

// EnterExtensionAdditionAlternative is called when production extensionAdditionAlternative is entered.
func (s *BaseASNListener) EnterExtensionAdditionAlternative(ctx *asn.ExtensionAdditionAlternativeContext) {
}

// ExitExtensionAdditionAlternative is called when production extensionAdditionAlternative is exited.
func (s *BaseASNListener) ExitExtensionAdditionAlternative(ctx *asn.ExtensionAdditionAlternativeContext) {
}

// EnterExtensionAdditionAlternativesGroup is called when production extensionAdditionAlternativesGroup is entered.
func (s *BaseASNListener) EnterExtensionAdditionAlternativesGroup(ctx *asn.ExtensionAdditionAlternativesGroupContext) {
}

// ExitExtensionAdditionAlternativesGroup is called when production extensionAdditionAlternativesGroup is exited.
func (s *BaseASNListener) ExitExtensionAdditionAlternativesGroup(ctx *asn.ExtensionAdditionAlternativesGroupContext) {
}

// EnterRootAlternativeTypeList is called when production rootAlternativeTypeList is entered.
func (s *BaseASNListener) EnterRootAlternativeTypeList(ctx *asn.RootAlternativeTypeListContext) {}

// ExitRootAlternativeTypeList is called when production rootAlternativeTypeList is exited.
func (s *BaseASNListener) ExitRootAlternativeTypeList(ctx *asn.RootAlternativeTypeListContext) {}

// EnterAlternativeTypeList is called when production alternativeTypeList is entered.
func (s *BaseASNListener) EnterAlternativeTypeList(ctx *asn.AlternativeTypeListContext) {}

// ExitAlternativeTypeList is called when production alternativeTypeList is exited.
func (s *BaseASNListener) ExitAlternativeTypeList(ctx *asn.AlternativeTypeListContext) {}

// EnterNamedType is called when production namedType is entered.
func (s *BaseASNListener) EnterNamedType(ctx *asn.NamedTypeContext) {}

// ExitNamedType is called when production namedType is exited.
func (s *BaseASNListener) ExitNamedType(ctx *asn.NamedTypeContext) {}

// EnterEnumeratedType is called when production enumeratedType is entered.
func (s *BaseASNListener) EnterEnumeratedType(ctx *asn.EnumeratedTypeContext) {}

// ExitEnumeratedType is called when production enumeratedType is exited.
func (s *BaseASNListener) ExitEnumeratedType(ctx *asn.EnumeratedTypeContext) {}

// EnterEnumerations is called when production enumerations is entered.
func (s *BaseASNListener) EnterEnumerations(ctx *asn.EnumerationsContext) {}

// ExitEnumerations is called when production enumerations is exited.
func (s *BaseASNListener) ExitEnumerations(ctx *asn.EnumerationsContext) {}

// EnterRootEnumeration is called when production rootEnumeration is entered.
func (s *BaseASNListener) EnterRootEnumeration(ctx *asn.RootEnumerationContext) {}

// ExitRootEnumeration is called when production rootEnumeration is exited.
func (s *BaseASNListener) ExitRootEnumeration(ctx *asn.RootEnumerationContext) {}

// EnterEnumeration is called when production enumeration is entered.
func (s *BaseASNListener) EnterEnumeration(ctx *asn.EnumerationContext) {}

// ExitEnumeration is called when production enumeration is exited.
func (s *BaseASNListener) ExitEnumeration(ctx *asn.EnumerationContext) {}

// EnterEnumerationItem is called when production enumerationItem is entered.
func (s *BaseASNListener) EnterEnumerationItem(ctx *asn.EnumerationItemContext) {}

// ExitEnumerationItem is called when production enumerationItem is exited.
func (s *BaseASNListener) ExitEnumerationItem(ctx *asn.EnumerationItemContext) {}

// EnterNamedNumber is called when production namedNumber is entered.
func (s *BaseASNListener) EnterNamedNumber(ctx *asn.NamedNumberContext) {}

// ExitNamedNumber is called when production namedNumber is exited.
func (s *BaseASNListener) ExitNamedNumber(ctx *asn.NamedNumberContext) {}

// EnterDefinedValue is called when production definedValue is entered.
func (s *BaseASNListener) EnterDefinedValue(ctx *asn.DefinedValueContext) {}

// ExitDefinedValue is called when production definedValue is exited.
func (s *BaseASNListener) ExitDefinedValue(ctx *asn.DefinedValueContext) {}

// EnterParameterizedValue is called when production parameterizedValue is entered.
func (s *BaseASNListener) EnterParameterizedValue(ctx *asn.ParameterizedValueContext) {}

// ExitParameterizedValue is called when production parameterizedValue is exited.
func (s *BaseASNListener) ExitParameterizedValue(ctx *asn.ParameterizedValueContext) {}

// EnterSimpleDefinedValue is called when production simpleDefinedValue is entered.
func (s *BaseASNListener) EnterSimpleDefinedValue(ctx *asn.SimpleDefinedValueContext) {}

// ExitSimpleDefinedValue is called when production simpleDefinedValue is exited.
func (s *BaseASNListener) ExitSimpleDefinedValue(ctx *asn.SimpleDefinedValueContext) {}

// EnterActualParameterList is called when production actualParameterList is entered.
func (s *BaseASNListener) EnterActualParameterList(ctx *asn.ActualParameterListContext) {}

// ExitActualParameterList is called when production actualParameterList is exited.
func (s *BaseASNListener) ExitActualParameterList(ctx *asn.ActualParameterListContext) {}

// EnterActualParameter is called when production actualParameter is entered.
func (s *BaseASNListener) EnterActualParameter(ctx *asn.ActualParameterContext) {}

// ExitActualParameter is called when production actualParameter is exited.
func (s *BaseASNListener) ExitActualParameter(ctx *asn.ActualParameterContext) {}

// EnterExceptionSpec is called when production exceptionSpec is entered.
func (s *BaseASNListener) EnterExceptionSpec(ctx *asn.ExceptionSpecContext) {}

// ExitExceptionSpec is called when production exceptionSpec is exited.
func (s *BaseASNListener) ExitExceptionSpec(ctx *asn.ExceptionSpecContext) {}

// EnterExceptionIdentification is called when production exceptionIdentification is entered.
func (s *BaseASNListener) EnterExceptionIdentification(ctx *asn.ExceptionIdentificationContext) {}

// ExitExceptionIdentification is called when production exceptionIdentification is exited.
func (s *BaseASNListener) ExitExceptionIdentification(ctx *asn.ExceptionIdentificationContext) {}

// EnterAdditionalEnumeration is called when production additionalEnumeration is entered.
func (s *BaseASNListener) EnterAdditionalEnumeration(ctx *asn.AdditionalEnumerationContext) {}

// ExitAdditionalEnumeration is called when production additionalEnumeration is exited.
func (s *BaseASNListener) ExitAdditionalEnumeration(ctx *asn.AdditionalEnumerationContext) {}

// EnterIntegerType is called when production integerType is entered.
func (s *BaseASNListener) EnterIntegerType(ctx *asn.IntegerTypeContext) {}

// ExitIntegerType is called when production integerType is exited.
func (s *BaseASNListener) ExitIntegerType(ctx *asn.IntegerTypeContext) {}

// EnterNamedNumberList is called when production namedNumberList is entered.
func (s *BaseASNListener) EnterNamedNumberList(ctx *asn.NamedNumberListContext) {}

// ExitNamedNumberList is called when production namedNumberList is exited.
func (s *BaseASNListener) ExitNamedNumberList(ctx *asn.NamedNumberListContext) {}

// EnterObjectidentifiertype is called when production objectidentifiertype is entered.
func (s *BaseASNListener) EnterObjectidentifiertype(ctx *asn.ObjectidentifiertypeContext) {}

// ExitObjectidentifiertype is called when production objectidentifiertype is exited.
func (s *BaseASNListener) ExitObjectidentifiertype(ctx *asn.ObjectidentifiertypeContext) {}

// EnterComponentRelationConstraint is called when production componentRelationConstraint is entered.
func (s *BaseASNListener) EnterComponentRelationConstraint(ctx *asn.ComponentRelationConstraintContext) {
}

// ExitComponentRelationConstraint is called when production componentRelationConstraint is exited.
func (s *BaseASNListener) ExitComponentRelationConstraint(ctx *asn.ComponentRelationConstraintContext) {
}

// EnterAtNotation is called when production atNotation is entered.
func (s *BaseASNListener) EnterAtNotation(ctx *asn.AtNotationContext) {}

// ExitAtNotation is called when production atNotation is exited.
func (s *BaseASNListener) ExitAtNotation(ctx *asn.AtNotationContext) {}

// EnterLevel is called when production level is entered.
func (s *BaseASNListener) EnterLevel(ctx *asn.LevelContext) {}

// ExitLevel is called when production level is exited.
func (s *BaseASNListener) ExitLevel(ctx *asn.LevelContext) {}

// EnterComponentIdList is called when production componentIdList is entered.
func (s *BaseASNListener) EnterComponentIdList(ctx *asn.ComponentIdListContext) {}

// ExitComponentIdList is called when production componentIdList is exited.
func (s *BaseASNListener) ExitComponentIdList(ctx *asn.ComponentIdListContext) {}

// EnterOctetStringType is called when production octetStringType is entered.
func (s *BaseASNListener) EnterOctetStringType(ctx *asn.OctetStringTypeContext) {}

// ExitOctetStringType is called when production octetStringType is exited.
func (s *BaseASNListener) ExitOctetStringType(ctx *asn.OctetStringTypeContext) {}

// EnterBitStringType is called when production bitStringType is entered.
func (s *BaseASNListener) EnterBitStringType(ctx *asn.BitStringTypeContext) {}

// ExitBitStringType is called when production bitStringType is exited.
func (s *BaseASNListener) ExitBitStringType(ctx *asn.BitStringTypeContext) {}

// EnterNamedBitList is called when production namedBitList is entered.
func (s *BaseASNListener) EnterNamedBitList(ctx *asn.NamedBitListContext) {}

// ExitNamedBitList is called when production namedBitList is exited.
func (s *BaseASNListener) ExitNamedBitList(ctx *asn.NamedBitListContext) {}

// EnterNamedBit is called when production namedBit is entered.
func (s *BaseASNListener) EnterNamedBit(ctx *asn.NamedBitContext) {}

// ExitNamedBit is called when production namedBit is exited.
func (s *BaseASNListener) ExitNamedBit(ctx *asn.NamedBitContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseASNListener) EnterBooleanValue(ctx *asn.BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseASNListener) ExitBooleanValue(ctx *asn.BooleanValueContext) {}
