// Code generated from ASN_3gpp.g4 by ANTLR 4.7.1. DO NOT EDIT.

package asn // ASN_3gpp
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseASN_3gppListener is a complete listener for a parse tree produced by ASN_3gppParser.
type BaseASN_3gppListener struct{}

var _ ASN_3gppListener = &BaseASN_3gppListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseASN_3gppListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseASN_3gppListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseASN_3gppListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseASN_3gppListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModules is called when production modules is entered.
func (s *BaseASN_3gppListener) EnterModules(ctx *ModulesContext) {}

// ExitModules is called when production modules is exited.
func (s *BaseASN_3gppListener) ExitModules(ctx *ModulesContext) {}

// EnterModuleDefinition is called when production moduleDefinition is entered.
func (s *BaseASN_3gppListener) EnterModuleDefinition(ctx *ModuleDefinitionContext) {}

// ExitModuleDefinition is called when production moduleDefinition is exited.
func (s *BaseASN_3gppListener) ExitModuleDefinition(ctx *ModuleDefinitionContext) {}

// EnterTagDefault is called when production tagDefault is entered.
func (s *BaseASN_3gppListener) EnterTagDefault(ctx *TagDefaultContext) {}

// ExitTagDefault is called when production tagDefault is exited.
func (s *BaseASN_3gppListener) ExitTagDefault(ctx *TagDefaultContext) {}

// EnterExtensionDefault is called when production extensionDefault is entered.
func (s *BaseASN_3gppListener) EnterExtensionDefault(ctx *ExtensionDefaultContext) {}

// ExitExtensionDefault is called when production extensionDefault is exited.
func (s *BaseASN_3gppListener) ExitExtensionDefault(ctx *ExtensionDefaultContext) {}

// EnterModuleBody is called when production moduleBody is entered.
func (s *BaseASN_3gppListener) EnterModuleBody(ctx *ModuleBodyContext) {}

// ExitModuleBody is called when production moduleBody is exited.
func (s *BaseASN_3gppListener) ExitModuleBody(ctx *ModuleBodyContext) {}

// EnterExports is called when production exports is entered.
func (s *BaseASN_3gppListener) EnterExports(ctx *ExportsContext) {}

// ExitExports is called when production exports is exited.
func (s *BaseASN_3gppListener) ExitExports(ctx *ExportsContext) {}

// EnterSymbolsExported is called when production symbolsExported is entered.
func (s *BaseASN_3gppListener) EnterSymbolsExported(ctx *SymbolsExportedContext) {}

// ExitSymbolsExported is called when production symbolsExported is exited.
func (s *BaseASN_3gppListener) ExitSymbolsExported(ctx *SymbolsExportedContext) {}

// EnterImports is called when production imports is entered.
func (s *BaseASN_3gppListener) EnterImports(ctx *ImportsContext) {}

// ExitImports is called when production imports is exited.
func (s *BaseASN_3gppListener) ExitImports(ctx *ImportsContext) {}

// EnterSymbolsImported is called when production symbolsImported is entered.
func (s *BaseASN_3gppListener) EnterSymbolsImported(ctx *SymbolsImportedContext) {}

// ExitSymbolsImported is called when production symbolsImported is exited.
func (s *BaseASN_3gppListener) ExitSymbolsImported(ctx *SymbolsImportedContext) {}

// EnterSymbolsFromModuleList is called when production symbolsFromModuleList is entered.
func (s *BaseASN_3gppListener) EnterSymbolsFromModuleList(ctx *SymbolsFromModuleListContext) {}

// ExitSymbolsFromModuleList is called when production symbolsFromModuleList is exited.
func (s *BaseASN_3gppListener) ExitSymbolsFromModuleList(ctx *SymbolsFromModuleListContext) {}

// EnterSymbolsFromModule is called when production symbolsFromModule is entered.
func (s *BaseASN_3gppListener) EnterSymbolsFromModule(ctx *SymbolsFromModuleContext) {}

// ExitSymbolsFromModule is called when production symbolsFromModule is exited.
func (s *BaseASN_3gppListener) ExitSymbolsFromModule(ctx *SymbolsFromModuleContext) {}

// EnterGlobalModuleReference is called when production globalModuleReference is entered.
func (s *BaseASN_3gppListener) EnterGlobalModuleReference(ctx *GlobalModuleReferenceContext) {}

// ExitGlobalModuleReference is called when production globalModuleReference is exited.
func (s *BaseASN_3gppListener) ExitGlobalModuleReference(ctx *GlobalModuleReferenceContext) {}

// EnterAssignedIdentifier is called when production assignedIdentifier is entered.
func (s *BaseASN_3gppListener) EnterAssignedIdentifier(ctx *AssignedIdentifierContext) {}

// ExitAssignedIdentifier is called when production assignedIdentifier is exited.
func (s *BaseASN_3gppListener) ExitAssignedIdentifier(ctx *AssignedIdentifierContext) {}

// EnterSymbolList is called when production symbolList is entered.
func (s *BaseASN_3gppListener) EnterSymbolList(ctx *SymbolListContext) {}

// ExitSymbolList is called when production symbolList is exited.
func (s *BaseASN_3gppListener) ExitSymbolList(ctx *SymbolListContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseASN_3gppListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseASN_3gppListener) ExitSymbol(ctx *SymbolContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseASN_3gppListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseASN_3gppListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseASN_3gppListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseASN_3gppListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterSequenceType is called when production sequenceType is entered.
func (s *BaseASN_3gppListener) EnterSequenceType(ctx *SequenceTypeContext) {}

// ExitSequenceType is called when production sequenceType is exited.
func (s *BaseASN_3gppListener) ExitSequenceType(ctx *SequenceTypeContext) {}

// EnterExtensionAndException is called when production extensionAndException is entered.
func (s *BaseASN_3gppListener) EnterExtensionAndException(ctx *ExtensionAndExceptionContext) {}

// ExitExtensionAndException is called when production extensionAndException is exited.
func (s *BaseASN_3gppListener) ExitExtensionAndException(ctx *ExtensionAndExceptionContext) {}

// EnterOptionalExtensionMarker is called when production optionalExtensionMarker is entered.
func (s *BaseASN_3gppListener) EnterOptionalExtensionMarker(ctx *OptionalExtensionMarkerContext) {}

// ExitOptionalExtensionMarker is called when production optionalExtensionMarker is exited.
func (s *BaseASN_3gppListener) ExitOptionalExtensionMarker(ctx *OptionalExtensionMarkerContext) {}

// EnterComponentTypeLists is called when production componentTypeLists is entered.
func (s *BaseASN_3gppListener) EnterComponentTypeLists(ctx *ComponentTypeListsContext) {}

// ExitComponentTypeLists is called when production componentTypeLists is exited.
func (s *BaseASN_3gppListener) ExitComponentTypeLists(ctx *ComponentTypeListsContext) {}

// EnterRootComponentTypeList is called when production rootComponentTypeList is entered.
func (s *BaseASN_3gppListener) EnterRootComponentTypeList(ctx *RootComponentTypeListContext) {}

// ExitRootComponentTypeList is called when production rootComponentTypeList is exited.
func (s *BaseASN_3gppListener) ExitRootComponentTypeList(ctx *RootComponentTypeListContext) {}

// EnterComponentTypeList is called when production componentTypeList is entered.
func (s *BaseASN_3gppListener) EnterComponentTypeList(ctx *ComponentTypeListContext) {}

// ExitComponentTypeList is called when production componentTypeList is exited.
func (s *BaseASN_3gppListener) ExitComponentTypeList(ctx *ComponentTypeListContext) {}

// EnterComponentType is called when production componentType is entered.
func (s *BaseASN_3gppListener) EnterComponentType(ctx *ComponentTypeContext) {}

// ExitComponentType is called when production componentType is exited.
func (s *BaseASN_3gppListener) ExitComponentType(ctx *ComponentTypeContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseASN_3gppListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseASN_3gppListener) ExitTag(ctx *TagContext) {}

// EnterNeedTag is called when production needTag is entered.
func (s *BaseASN_3gppListener) EnterNeedTag(ctx *NeedTagContext) {}

// ExitNeedTag is called when production needTag is exited.
func (s *BaseASN_3gppListener) ExitNeedTag(ctx *NeedTagContext) {}

// EnterCondTag is called when production condTag is entered.
func (s *BaseASN_3gppListener) EnterCondTag(ctx *CondTagContext) {}

// ExitCondTag is called when production condTag is exited.
func (s *BaseASN_3gppListener) ExitCondTag(ctx *CondTagContext) {}

// EnterExtensionAdditions is called when production extensionAdditions is entered.
func (s *BaseASN_3gppListener) EnterExtensionAdditions(ctx *ExtensionAdditionsContext) {}

// ExitExtensionAdditions is called when production extensionAdditions is exited.
func (s *BaseASN_3gppListener) ExitExtensionAdditions(ctx *ExtensionAdditionsContext) {}

// EnterExtensionAdditionList is called when production extensionAdditionList is entered.
func (s *BaseASN_3gppListener) EnterExtensionAdditionList(ctx *ExtensionAdditionListContext) {}

// ExitExtensionAdditionList is called when production extensionAdditionList is exited.
func (s *BaseASN_3gppListener) ExitExtensionAdditionList(ctx *ExtensionAdditionListContext) {}

// EnterExtensionAddition is called when production extensionAddition is entered.
func (s *BaseASN_3gppListener) EnterExtensionAddition(ctx *ExtensionAdditionContext) {}

// ExitExtensionAddition is called when production extensionAddition is exited.
func (s *BaseASN_3gppListener) ExitExtensionAddition(ctx *ExtensionAdditionContext) {}

// EnterExtensionAdditionGroup is called when production extensionAdditionGroup is entered.
func (s *BaseASN_3gppListener) EnterExtensionAdditionGroup(ctx *ExtensionAdditionGroupContext) {}

// ExitExtensionAdditionGroup is called when production extensionAdditionGroup is exited.
func (s *BaseASN_3gppListener) ExitExtensionAdditionGroup(ctx *ExtensionAdditionGroupContext) {}

// EnterVersionNumber is called when production versionNumber is entered.
func (s *BaseASN_3gppListener) EnterVersionNumber(ctx *VersionNumberContext) {}

// ExitVersionNumber is called when production versionNumber is exited.
func (s *BaseASN_3gppListener) ExitVersionNumber(ctx *VersionNumberContext) {}

// EnterSequenceOfType is called when production sequenceOfType is entered.
func (s *BaseASN_3gppListener) EnterSequenceOfType(ctx *SequenceOfTypeContext) {}

// ExitSequenceOfType is called when production sequenceOfType is exited.
func (s *BaseASN_3gppListener) ExitSequenceOfType(ctx *SequenceOfTypeContext) {}

// EnterSizeConstraint is called when production sizeConstraint is entered.
func (s *BaseASN_3gppListener) EnterSizeConstraint(ctx *SizeConstraintContext) {}

// ExitSizeConstraint is called when production sizeConstraint is exited.
func (s *BaseASN_3gppListener) ExitSizeConstraint(ctx *SizeConstraintContext) {}

// EnterParameterizedAssignment is called when production parameterizedAssignment is entered.
func (s *BaseASN_3gppListener) EnterParameterizedAssignment(ctx *ParameterizedAssignmentContext) {}

// ExitParameterizedAssignment is called when production parameterizedAssignment is exited.
func (s *BaseASN_3gppListener) ExitParameterizedAssignment(ctx *ParameterizedAssignmentContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseASN_3gppListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseASN_3gppListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseASN_3gppListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseASN_3gppListener) ExitParameter(ctx *ParameterContext) {}

// EnterParamGovernor is called when production paramGovernor is entered.
func (s *BaseASN_3gppListener) EnterParamGovernor(ctx *ParamGovernorContext) {}

// ExitParamGovernor is called when production paramGovernor is exited.
func (s *BaseASN_3gppListener) ExitParamGovernor(ctx *ParamGovernorContext) {}

// EnterGovernor is called when production governor is entered.
func (s *BaseASN_3gppListener) EnterGovernor(ctx *GovernorContext) {}

// ExitGovernor is called when production governor is exited.
func (s *BaseASN_3gppListener) ExitGovernor(ctx *GovernorContext) {}

// EnterObjectClassAssignment is called when production objectClassAssignment is entered.
func (s *BaseASN_3gppListener) EnterObjectClassAssignment(ctx *ObjectClassAssignmentContext) {}

// ExitObjectClassAssignment is called when production objectClassAssignment is exited.
func (s *BaseASN_3gppListener) ExitObjectClassAssignment(ctx *ObjectClassAssignmentContext) {}

// EnterObjectClass is called when production objectClass is entered.
func (s *BaseASN_3gppListener) EnterObjectClass(ctx *ObjectClassContext) {}

// ExitObjectClass is called when production objectClass is exited.
func (s *BaseASN_3gppListener) ExitObjectClass(ctx *ObjectClassContext) {}

// EnterDefinedObjectClass is called when production definedObjectClass is entered.
func (s *BaseASN_3gppListener) EnterDefinedObjectClass(ctx *DefinedObjectClassContext) {}

// ExitDefinedObjectClass is called when production definedObjectClass is exited.
func (s *BaseASN_3gppListener) ExitDefinedObjectClass(ctx *DefinedObjectClassContext) {}

// EnterUsefulObjectClassReference is called when production usefulObjectClassReference is entered.
func (s *BaseASN_3gppListener) EnterUsefulObjectClassReference(ctx *UsefulObjectClassReferenceContext) {
}

// ExitUsefulObjectClassReference is called when production usefulObjectClassReference is exited.
func (s *BaseASN_3gppListener) ExitUsefulObjectClassReference(ctx *UsefulObjectClassReferenceContext) {
}

// EnterExternalObjectClassReference is called when production externalObjectClassReference is entered.
func (s *BaseASN_3gppListener) EnterExternalObjectClassReference(ctx *ExternalObjectClassReferenceContext) {
}

// ExitExternalObjectClassReference is called when production externalObjectClassReference is exited.
func (s *BaseASN_3gppListener) ExitExternalObjectClassReference(ctx *ExternalObjectClassReferenceContext) {
}

// EnterObjectClassDefn is called when production objectClassDefn is entered.
func (s *BaseASN_3gppListener) EnterObjectClassDefn(ctx *ObjectClassDefnContext) {}

// ExitObjectClassDefn is called when production objectClassDefn is exited.
func (s *BaseASN_3gppListener) ExitObjectClassDefn(ctx *ObjectClassDefnContext) {}

// EnterWithSyntaxSpec is called when production withSyntaxSpec is entered.
func (s *BaseASN_3gppListener) EnterWithSyntaxSpec(ctx *WithSyntaxSpecContext) {}

// ExitWithSyntaxSpec is called when production withSyntaxSpec is exited.
func (s *BaseASN_3gppListener) ExitWithSyntaxSpec(ctx *WithSyntaxSpecContext) {}

// EnterSyntaxList is called when production syntaxList is entered.
func (s *BaseASN_3gppListener) EnterSyntaxList(ctx *SyntaxListContext) {}

// ExitSyntaxList is called when production syntaxList is exited.
func (s *BaseASN_3gppListener) ExitSyntaxList(ctx *SyntaxListContext) {}

// EnterTokenOrGroupSpec is called when production tokenOrGroupSpec is entered.
func (s *BaseASN_3gppListener) EnterTokenOrGroupSpec(ctx *TokenOrGroupSpecContext) {}

// ExitTokenOrGroupSpec is called when production tokenOrGroupSpec is exited.
func (s *BaseASN_3gppListener) ExitTokenOrGroupSpec(ctx *TokenOrGroupSpecContext) {}

// EnterOptionalGroup is called when production optionalGroup is entered.
func (s *BaseASN_3gppListener) EnterOptionalGroup(ctx *OptionalGroupContext) {}

// ExitOptionalGroup is called when production optionalGroup is exited.
func (s *BaseASN_3gppListener) ExitOptionalGroup(ctx *OptionalGroupContext) {}

// EnterRequiredToken is called when production requiredToken is entered.
func (s *BaseASN_3gppListener) EnterRequiredToken(ctx *RequiredTokenContext) {}

// ExitRequiredToken is called when production requiredToken is exited.
func (s *BaseASN_3gppListener) ExitRequiredToken(ctx *RequiredTokenContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseASN_3gppListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseASN_3gppListener) ExitLiteral(ctx *LiteralContext) {}

// EnterPrimitiveFieldName is called when production primitiveFieldName is entered.
func (s *BaseASN_3gppListener) EnterPrimitiveFieldName(ctx *PrimitiveFieldNameContext) {}

// ExitPrimitiveFieldName is called when production primitiveFieldName is exited.
func (s *BaseASN_3gppListener) ExitPrimitiveFieldName(ctx *PrimitiveFieldNameContext) {}

// EnterFieldSpec is called when production fieldSpec is entered.
func (s *BaseASN_3gppListener) EnterFieldSpec(ctx *FieldSpecContext) {}

// ExitFieldSpec is called when production fieldSpec is exited.
func (s *BaseASN_3gppListener) ExitFieldSpec(ctx *FieldSpecContext) {}

// EnterTypeFieldSpec is called when production typeFieldSpec is entered.
func (s *BaseASN_3gppListener) EnterTypeFieldSpec(ctx *TypeFieldSpecContext) {}

// ExitTypeFieldSpec is called when production typeFieldSpec is exited.
func (s *BaseASN_3gppListener) ExitTypeFieldSpec(ctx *TypeFieldSpecContext) {}

// EnterTypeOptionalitySpec is called when production typeOptionalitySpec is entered.
func (s *BaseASN_3gppListener) EnterTypeOptionalitySpec(ctx *TypeOptionalitySpecContext) {}

// ExitTypeOptionalitySpec is called when production typeOptionalitySpec is exited.
func (s *BaseASN_3gppListener) ExitTypeOptionalitySpec(ctx *TypeOptionalitySpecContext) {}

// EnterFixedTypeValueFieldSpec is called when production fixedTypeValueFieldSpec is entered.
func (s *BaseASN_3gppListener) EnterFixedTypeValueFieldSpec(ctx *FixedTypeValueFieldSpecContext) {}

// ExitFixedTypeValueFieldSpec is called when production fixedTypeValueFieldSpec is exited.
func (s *BaseASN_3gppListener) ExitFixedTypeValueFieldSpec(ctx *FixedTypeValueFieldSpecContext) {}

// EnterValueOptionalitySpec is called when production valueOptionalitySpec is entered.
func (s *BaseASN_3gppListener) EnterValueOptionalitySpec(ctx *ValueOptionalitySpecContext) {}

// ExitValueOptionalitySpec is called when production valueOptionalitySpec is exited.
func (s *BaseASN_3gppListener) ExitValueOptionalitySpec(ctx *ValueOptionalitySpecContext) {}

// EnterVariableTypeValueFieldSpec is called when production variableTypeValueFieldSpec is entered.
func (s *BaseASN_3gppListener) EnterVariableTypeValueFieldSpec(ctx *VariableTypeValueFieldSpecContext) {
}

// ExitVariableTypeValueFieldSpec is called when production variableTypeValueFieldSpec is exited.
func (s *BaseASN_3gppListener) ExitVariableTypeValueFieldSpec(ctx *VariableTypeValueFieldSpecContext) {
}

// EnterFixedTypeValueSetFieldSpec is called when production fixedTypeValueSetFieldSpec is entered.
func (s *BaseASN_3gppListener) EnterFixedTypeValueSetFieldSpec(ctx *FixedTypeValueSetFieldSpecContext) {
}

// ExitFixedTypeValueSetFieldSpec is called when production fixedTypeValueSetFieldSpec is exited.
func (s *BaseASN_3gppListener) ExitFixedTypeValueSetFieldSpec(ctx *FixedTypeValueSetFieldSpecContext) {
}

// EnterValueSetOptionalitySpec is called when production valueSetOptionalitySpec is entered.
func (s *BaseASN_3gppListener) EnterValueSetOptionalitySpec(ctx *ValueSetOptionalitySpecContext) {}

// ExitValueSetOptionalitySpec is called when production valueSetOptionalitySpec is exited.
func (s *BaseASN_3gppListener) ExitValueSetOptionalitySpec(ctx *ValueSetOptionalitySpecContext) {}

// EnterObject is called when production object is entered.
func (s *BaseASN_3gppListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseASN_3gppListener) ExitObject(ctx *ObjectContext) {}

// EnterParameterizedObject is called when production parameterizedObject is entered.
func (s *BaseASN_3gppListener) EnterParameterizedObject(ctx *ParameterizedObjectContext) {}

// ExitParameterizedObject is called when production parameterizedObject is exited.
func (s *BaseASN_3gppListener) ExitParameterizedObject(ctx *ParameterizedObjectContext) {}

// EnterDefinedObject is called when production definedObject is entered.
func (s *BaseASN_3gppListener) EnterDefinedObject(ctx *DefinedObjectContext) {}

// ExitDefinedObject is called when production definedObject is exited.
func (s *BaseASN_3gppListener) ExitDefinedObject(ctx *DefinedObjectContext) {}

// EnterObjectSet is called when production objectSet is entered.
func (s *BaseASN_3gppListener) EnterObjectSet(ctx *ObjectSetContext) {}

// ExitObjectSet is called when production objectSet is exited.
func (s *BaseASN_3gppListener) ExitObjectSet(ctx *ObjectSetContext) {}

// EnterObjectSetSpec is called when production objectSetSpec is entered.
func (s *BaseASN_3gppListener) EnterObjectSetSpec(ctx *ObjectSetSpecContext) {}

// ExitObjectSetSpec is called when production objectSetSpec is exited.
func (s *BaseASN_3gppListener) ExitObjectSetSpec(ctx *ObjectSetSpecContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseASN_3gppListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseASN_3gppListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterValueSet is called when production valueSet is entered.
func (s *BaseASN_3gppListener) EnterValueSet(ctx *ValueSetContext) {}

// ExitValueSet is called when production valueSet is exited.
func (s *BaseASN_3gppListener) ExitValueSet(ctx *ValueSetContext) {}

// EnterElementSetSpecs is called when production elementSetSpecs is entered.
func (s *BaseASN_3gppListener) EnterElementSetSpecs(ctx *ElementSetSpecsContext) {}

// ExitElementSetSpecs is called when production elementSetSpecs is exited.
func (s *BaseASN_3gppListener) ExitElementSetSpecs(ctx *ElementSetSpecsContext) {}

// EnterRootElementSetSpec is called when production rootElementSetSpec is entered.
func (s *BaseASN_3gppListener) EnterRootElementSetSpec(ctx *RootElementSetSpecContext) {}

// ExitRootElementSetSpec is called when production rootElementSetSpec is exited.
func (s *BaseASN_3gppListener) ExitRootElementSetSpec(ctx *RootElementSetSpecContext) {}

// EnterAdditionalElementSetSpec is called when production additionalElementSetSpec is entered.
func (s *BaseASN_3gppListener) EnterAdditionalElementSetSpec(ctx *AdditionalElementSetSpecContext) {}

// ExitAdditionalElementSetSpec is called when production additionalElementSetSpec is exited.
func (s *BaseASN_3gppListener) ExitAdditionalElementSetSpec(ctx *AdditionalElementSetSpecContext) {}

// EnterElementSetSpec is called when production elementSetSpec is entered.
func (s *BaseASN_3gppListener) EnterElementSetSpec(ctx *ElementSetSpecContext) {}

// ExitElementSetSpec is called when production elementSetSpec is exited.
func (s *BaseASN_3gppListener) ExitElementSetSpec(ctx *ElementSetSpecContext) {}

// EnterUnions is called when production unions is entered.
func (s *BaseASN_3gppListener) EnterUnions(ctx *UnionsContext) {}

// ExitUnions is called when production unions is exited.
func (s *BaseASN_3gppListener) ExitUnions(ctx *UnionsContext) {}

// EnterExclusions is called when production exclusions is entered.
func (s *BaseASN_3gppListener) EnterExclusions(ctx *ExclusionsContext) {}

// ExitExclusions is called when production exclusions is exited.
func (s *BaseASN_3gppListener) ExitExclusions(ctx *ExclusionsContext) {}

// EnterIntersections is called when production intersections is entered.
func (s *BaseASN_3gppListener) EnterIntersections(ctx *IntersectionsContext) {}

// ExitIntersections is called when production intersections is exited.
func (s *BaseASN_3gppListener) ExitIntersections(ctx *IntersectionsContext) {}

// EnterUnionMark is called when production unionMark is entered.
func (s *BaseASN_3gppListener) EnterUnionMark(ctx *UnionMarkContext) {}

// ExitUnionMark is called when production unionMark is exited.
func (s *BaseASN_3gppListener) ExitUnionMark(ctx *UnionMarkContext) {}

// EnterIntersectionMark is called when production intersectionMark is entered.
func (s *BaseASN_3gppListener) EnterIntersectionMark(ctx *IntersectionMarkContext) {}

// ExitIntersectionMark is called when production intersectionMark is exited.
func (s *BaseASN_3gppListener) ExitIntersectionMark(ctx *IntersectionMarkContext) {}

// EnterElements is called when production elements is entered.
func (s *BaseASN_3gppListener) EnterElements(ctx *ElementsContext) {}

// ExitElements is called when production elements is exited.
func (s *BaseASN_3gppListener) ExitElements(ctx *ElementsContext) {}

// EnterObjectSetElements is called when production objectSetElements is entered.
func (s *BaseASN_3gppListener) EnterObjectSetElements(ctx *ObjectSetElementsContext) {}

// ExitObjectSetElements is called when production objectSetElements is exited.
func (s *BaseASN_3gppListener) ExitObjectSetElements(ctx *ObjectSetElementsContext) {}

// EnterIntersectionElements is called when production intersectionElements is entered.
func (s *BaseASN_3gppListener) EnterIntersectionElements(ctx *IntersectionElementsContext) {}

// ExitIntersectionElements is called when production intersectionElements is exited.
func (s *BaseASN_3gppListener) ExitIntersectionElements(ctx *IntersectionElementsContext) {}

// EnterSubtypeElements is called when production subtypeElements is entered.
func (s *BaseASN_3gppListener) EnterSubtypeElements(ctx *SubtypeElementsContext) {}

// ExitSubtypeElements is called when production subtypeElements is exited.
func (s *BaseASN_3gppListener) ExitSubtypeElements(ctx *SubtypeElementsContext) {}

// EnterVariableTypeValueSetFieldSpec is called when production variableTypeValueSetFieldSpec is entered.
func (s *BaseASN_3gppListener) EnterVariableTypeValueSetFieldSpec(ctx *VariableTypeValueSetFieldSpecContext) {
}

// ExitVariableTypeValueSetFieldSpec is called when production variableTypeValueSetFieldSpec is exited.
func (s *BaseASN_3gppListener) ExitVariableTypeValueSetFieldSpec(ctx *VariableTypeValueSetFieldSpecContext) {
}

// EnterObjectFieldSpec is called when production objectFieldSpec is entered.
func (s *BaseASN_3gppListener) EnterObjectFieldSpec(ctx *ObjectFieldSpecContext) {}

// ExitObjectFieldSpec is called when production objectFieldSpec is exited.
func (s *BaseASN_3gppListener) ExitObjectFieldSpec(ctx *ObjectFieldSpecContext) {}

// EnterObjectOptionalitySpec is called when production objectOptionalitySpec is entered.
func (s *BaseASN_3gppListener) EnterObjectOptionalitySpec(ctx *ObjectOptionalitySpecContext) {}

// ExitObjectOptionalitySpec is called when production objectOptionalitySpec is exited.
func (s *BaseASN_3gppListener) ExitObjectOptionalitySpec(ctx *ObjectOptionalitySpecContext) {}

// EnterObjectSetFieldSpec is called when production objectSetFieldSpec is entered.
func (s *BaseASN_3gppListener) EnterObjectSetFieldSpec(ctx *ObjectSetFieldSpecContext) {}

// ExitObjectSetFieldSpec is called when production objectSetFieldSpec is exited.
func (s *BaseASN_3gppListener) ExitObjectSetFieldSpec(ctx *ObjectSetFieldSpecContext) {}

// EnterObjectSetOptionalitySpec is called when production objectSetOptionalitySpec is entered.
func (s *BaseASN_3gppListener) EnterObjectSetOptionalitySpec(ctx *ObjectSetOptionalitySpecContext) {}

// ExitObjectSetOptionalitySpec is called when production objectSetOptionalitySpec is exited.
func (s *BaseASN_3gppListener) ExitObjectSetOptionalitySpec(ctx *ObjectSetOptionalitySpecContext) {}

// EnterTypeAssignment is called when production typeAssignment is entered.
func (s *BaseASN_3gppListener) EnterTypeAssignment(ctx *TypeAssignmentContext) {}

// ExitTypeAssignment is called when production typeAssignment is exited.
func (s *BaseASN_3gppListener) ExitTypeAssignment(ctx *TypeAssignmentContext) {}

// EnterValueAssignment is called when production valueAssignment is entered.
func (s *BaseASN_3gppListener) EnterValueAssignment(ctx *ValueAssignmentContext) {}

// ExitValueAssignment is called when production valueAssignment is exited.
func (s *BaseASN_3gppListener) ExitValueAssignment(ctx *ValueAssignmentContext) {}

// EnterAsnType is called when production asnType is entered.
func (s *BaseASN_3gppListener) EnterAsnType(ctx *AsnTypeContext) {}

// ExitAsnType is called when production asnType is exited.
func (s *BaseASN_3gppListener) ExitAsnType(ctx *AsnTypeContext) {}

// EnterBuiltinType is called when production builtinType is entered.
func (s *BaseASN_3gppListener) EnterBuiltinType(ctx *BuiltinTypeContext) {}

// ExitBuiltinType is called when production builtinType is exited.
func (s *BaseASN_3gppListener) ExitBuiltinType(ctx *BuiltinTypeContext) {}

// EnterObjectClassFieldType is called when production objectClassFieldType is entered.
func (s *BaseASN_3gppListener) EnterObjectClassFieldType(ctx *ObjectClassFieldTypeContext) {}

// ExitObjectClassFieldType is called when production objectClassFieldType is exited.
func (s *BaseASN_3gppListener) ExitObjectClassFieldType(ctx *ObjectClassFieldTypeContext) {}

// EnterSetType is called when production setType is entered.
func (s *BaseASN_3gppListener) EnterSetType(ctx *SetTypeContext) {}

// ExitSetType is called when production setType is exited.
func (s *BaseASN_3gppListener) ExitSetType(ctx *SetTypeContext) {}

// EnterSetOfType is called when production setOfType is entered.
func (s *BaseASN_3gppListener) EnterSetOfType(ctx *SetOfTypeContext) {}

// ExitSetOfType is called when production setOfType is exited.
func (s *BaseASN_3gppListener) ExitSetOfType(ctx *SetOfTypeContext) {}

// EnterReferencedType is called when production referencedType is entered.
func (s *BaseASN_3gppListener) EnterReferencedType(ctx *ReferencedTypeContext) {}

// ExitReferencedType is called when production referencedType is exited.
func (s *BaseASN_3gppListener) ExitReferencedType(ctx *ReferencedTypeContext) {}

// EnterDefinedType is called when production definedType is entered.
func (s *BaseASN_3gppListener) EnterDefinedType(ctx *DefinedTypeContext) {}

// ExitDefinedType is called when production definedType is exited.
func (s *BaseASN_3gppListener) ExitDefinedType(ctx *DefinedTypeContext) {}

// EnterConstraint is called when production constraint is entered.
func (s *BaseASN_3gppListener) EnterConstraint(ctx *ConstraintContext) {}

// ExitConstraint is called when production constraint is exited.
func (s *BaseASN_3gppListener) ExitConstraint(ctx *ConstraintContext) {}

// EnterConstraintSpec is called when production constraintSpec is entered.
func (s *BaseASN_3gppListener) EnterConstraintSpec(ctx *ConstraintSpecContext) {}

// ExitConstraintSpec is called when production constraintSpec is exited.
func (s *BaseASN_3gppListener) ExitConstraintSpec(ctx *ConstraintSpecContext) {}

// EnterUserDefinedConstraint is called when production userDefinedConstraint is entered.
func (s *BaseASN_3gppListener) EnterUserDefinedConstraint(ctx *UserDefinedConstraintContext) {}

// ExitUserDefinedConstraint is called when production userDefinedConstraint is exited.
func (s *BaseASN_3gppListener) ExitUserDefinedConstraint(ctx *UserDefinedConstraintContext) {}

// EnterGeneralConstraint is called when production generalConstraint is entered.
func (s *BaseASN_3gppListener) EnterGeneralConstraint(ctx *GeneralConstraintContext) {}

// ExitGeneralConstraint is called when production generalConstraint is exited.
func (s *BaseASN_3gppListener) ExitGeneralConstraint(ctx *GeneralConstraintContext) {}

// EnterUserDefinedConstraintParameter is called when production userDefinedConstraintParameter is entered.
func (s *BaseASN_3gppListener) EnterUserDefinedConstraintParameter(ctx *UserDefinedConstraintParameterContext) {
}

// ExitUserDefinedConstraintParameter is called when production userDefinedConstraintParameter is exited.
func (s *BaseASN_3gppListener) ExitUserDefinedConstraintParameter(ctx *UserDefinedConstraintParameterContext) {
}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseASN_3gppListener) EnterTableConstraint(ctx *TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseASN_3gppListener) ExitTableConstraint(ctx *TableConstraintContext) {}

// EnterSimpleTableConstraint is called when production simpleTableConstraint is entered.
func (s *BaseASN_3gppListener) EnterSimpleTableConstraint(ctx *SimpleTableConstraintContext) {}

// ExitSimpleTableConstraint is called when production simpleTableConstraint is exited.
func (s *BaseASN_3gppListener) ExitSimpleTableConstraint(ctx *SimpleTableConstraintContext) {}

// EnterContentsConstraint is called when production contentsConstraint is entered.
func (s *BaseASN_3gppListener) EnterContentsConstraint(ctx *ContentsConstraintContext) {}

// ExitContentsConstraint is called when production contentsConstraint is exited.
func (s *BaseASN_3gppListener) ExitContentsConstraint(ctx *ContentsConstraintContext) {}

// EnterComponentPresenceLists is called when production componentPresenceLists is entered.
func (s *BaseASN_3gppListener) EnterComponentPresenceLists(ctx *ComponentPresenceListsContext) {}

// ExitComponentPresenceLists is called when production componentPresenceLists is exited.
func (s *BaseASN_3gppListener) ExitComponentPresenceLists(ctx *ComponentPresenceListsContext) {}

// EnterComponentPresenceList is called when production componentPresenceList is entered.
func (s *BaseASN_3gppListener) EnterComponentPresenceList(ctx *ComponentPresenceListContext) {}

// ExitComponentPresenceList is called when production componentPresenceList is exited.
func (s *BaseASN_3gppListener) ExitComponentPresenceList(ctx *ComponentPresenceListContext) {}

// EnterComponentPresence is called when production componentPresence is entered.
func (s *BaseASN_3gppListener) EnterComponentPresence(ctx *ComponentPresenceContext) {}

// ExitComponentPresence is called when production componentPresence is exited.
func (s *BaseASN_3gppListener) ExitComponentPresence(ctx *ComponentPresenceContext) {}

// EnterSubtypeConstraint is called when production subtypeConstraint is entered.
func (s *BaseASN_3gppListener) EnterSubtypeConstraint(ctx *SubtypeConstraintContext) {}

// ExitSubtypeConstraint is called when production subtypeConstraint is exited.
func (s *BaseASN_3gppListener) ExitSubtypeConstraint(ctx *SubtypeConstraintContext) {}

// EnterValue is called when production value is entered.
func (s *BaseASN_3gppListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseASN_3gppListener) ExitValue(ctx *ValueContext) {}

// EnterBuiltinValue is called when production builtinValue is entered.
func (s *BaseASN_3gppListener) EnterBuiltinValue(ctx *BuiltinValueContext) {}

// ExitBuiltinValue is called when production builtinValue is exited.
func (s *BaseASN_3gppListener) ExitBuiltinValue(ctx *BuiltinValueContext) {}

// EnterObjectIdentifierValue is called when production objectIdentifierValue is entered.
func (s *BaseASN_3gppListener) EnterObjectIdentifierValue(ctx *ObjectIdentifierValueContext) {}

// ExitObjectIdentifierValue is called when production objectIdentifierValue is exited.
func (s *BaseASN_3gppListener) ExitObjectIdentifierValue(ctx *ObjectIdentifierValueContext) {}

// EnterObjIdComponentsList is called when production objIdComponentsList is entered.
func (s *BaseASN_3gppListener) EnterObjIdComponentsList(ctx *ObjIdComponentsListContext) {}

// ExitObjIdComponentsList is called when production objIdComponentsList is exited.
func (s *BaseASN_3gppListener) ExitObjIdComponentsList(ctx *ObjIdComponentsListContext) {}

// EnterObjIdComponents is called when production objIdComponents is entered.
func (s *BaseASN_3gppListener) EnterObjIdComponents(ctx *ObjIdComponentsContext) {}

// ExitObjIdComponents is called when production objIdComponents is exited.
func (s *BaseASN_3gppListener) ExitObjIdComponents(ctx *ObjIdComponentsContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseASN_3gppListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseASN_3gppListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterChoiceValue is called when production choiceValue is entered.
func (s *BaseASN_3gppListener) EnterChoiceValue(ctx *ChoiceValueContext) {}

// ExitChoiceValue is called when production choiceValue is exited.
func (s *BaseASN_3gppListener) ExitChoiceValue(ctx *ChoiceValueContext) {}

// EnterEnumeratedValue is called when production enumeratedValue is entered.
func (s *BaseASN_3gppListener) EnterEnumeratedValue(ctx *EnumeratedValueContext) {}

// ExitEnumeratedValue is called when production enumeratedValue is exited.
func (s *BaseASN_3gppListener) ExitEnumeratedValue(ctx *EnumeratedValueContext) {}

// EnterSignedNumber is called when production signedNumber is entered.
func (s *BaseASN_3gppListener) EnterSignedNumber(ctx *SignedNumberContext) {}

// ExitSignedNumber is called when production signedNumber is exited.
func (s *BaseASN_3gppListener) ExitSignedNumber(ctx *SignedNumberContext) {}

// EnterChoiceType is called when production choiceType is entered.
func (s *BaseASN_3gppListener) EnterChoiceType(ctx *ChoiceTypeContext) {}

// ExitChoiceType is called when production choiceType is exited.
func (s *BaseASN_3gppListener) ExitChoiceType(ctx *ChoiceTypeContext) {}

// EnterAlternativeTypeLists is called when production alternativeTypeLists is entered.
func (s *BaseASN_3gppListener) EnterAlternativeTypeLists(ctx *AlternativeTypeListsContext) {}

// ExitAlternativeTypeLists is called when production alternativeTypeLists is exited.
func (s *BaseASN_3gppListener) ExitAlternativeTypeLists(ctx *AlternativeTypeListsContext) {}

// EnterExtensionAdditionAlternatives is called when production extensionAdditionAlternatives is entered.
func (s *BaseASN_3gppListener) EnterExtensionAdditionAlternatives(ctx *ExtensionAdditionAlternativesContext) {
}

// ExitExtensionAdditionAlternatives is called when production extensionAdditionAlternatives is exited.
func (s *BaseASN_3gppListener) ExitExtensionAdditionAlternatives(ctx *ExtensionAdditionAlternativesContext) {
}

// EnterExtensionAdditionAlternativesList is called when production extensionAdditionAlternativesList is entered.
func (s *BaseASN_3gppListener) EnterExtensionAdditionAlternativesList(ctx *ExtensionAdditionAlternativesListContext) {
}

// ExitExtensionAdditionAlternativesList is called when production extensionAdditionAlternativesList is exited.
func (s *BaseASN_3gppListener) ExitExtensionAdditionAlternativesList(ctx *ExtensionAdditionAlternativesListContext) {
}

// EnterExtensionAdditionAlternative is called when production extensionAdditionAlternative is entered.
func (s *BaseASN_3gppListener) EnterExtensionAdditionAlternative(ctx *ExtensionAdditionAlternativeContext) {
}

// ExitExtensionAdditionAlternative is called when production extensionAdditionAlternative is exited.
func (s *BaseASN_3gppListener) ExitExtensionAdditionAlternative(ctx *ExtensionAdditionAlternativeContext) {
}

// EnterExtensionAdditionAlternativesGroup is called when production extensionAdditionAlternativesGroup is entered.
func (s *BaseASN_3gppListener) EnterExtensionAdditionAlternativesGroup(ctx *ExtensionAdditionAlternativesGroupContext) {
}

// ExitExtensionAdditionAlternativesGroup is called when production extensionAdditionAlternativesGroup is exited.
func (s *BaseASN_3gppListener) ExitExtensionAdditionAlternativesGroup(ctx *ExtensionAdditionAlternativesGroupContext) {
}

// EnterRootAlternativeTypeList is called when production rootAlternativeTypeList is entered.
func (s *BaseASN_3gppListener) EnterRootAlternativeTypeList(ctx *RootAlternativeTypeListContext) {}

// ExitRootAlternativeTypeList is called when production rootAlternativeTypeList is exited.
func (s *BaseASN_3gppListener) ExitRootAlternativeTypeList(ctx *RootAlternativeTypeListContext) {}

// EnterAlternativeTypeList is called when production alternativeTypeList is entered.
func (s *BaseASN_3gppListener) EnterAlternativeTypeList(ctx *AlternativeTypeListContext) {}

// ExitAlternativeTypeList is called when production alternativeTypeList is exited.
func (s *BaseASN_3gppListener) ExitAlternativeTypeList(ctx *AlternativeTypeListContext) {}

// EnterNamedType is called when production namedType is entered.
func (s *BaseASN_3gppListener) EnterNamedType(ctx *NamedTypeContext) {}

// ExitNamedType is called when production namedType is exited.
func (s *BaseASN_3gppListener) ExitNamedType(ctx *NamedTypeContext) {}

// EnterEnumeratedType is called when production enumeratedType is entered.
func (s *BaseASN_3gppListener) EnterEnumeratedType(ctx *EnumeratedTypeContext) {}

// ExitEnumeratedType is called when production enumeratedType is exited.
func (s *BaseASN_3gppListener) ExitEnumeratedType(ctx *EnumeratedTypeContext) {}

// EnterEnumerations is called when production enumerations is entered.
func (s *BaseASN_3gppListener) EnterEnumerations(ctx *EnumerationsContext) {}

// ExitEnumerations is called when production enumerations is exited.
func (s *BaseASN_3gppListener) ExitEnumerations(ctx *EnumerationsContext) {}

// EnterRootEnumeration is called when production rootEnumeration is entered.
func (s *BaseASN_3gppListener) EnterRootEnumeration(ctx *RootEnumerationContext) {}

// ExitRootEnumeration is called when production rootEnumeration is exited.
func (s *BaseASN_3gppListener) ExitRootEnumeration(ctx *RootEnumerationContext) {}

// EnterEnumeration is called when production enumeration is entered.
func (s *BaseASN_3gppListener) EnterEnumeration(ctx *EnumerationContext) {}

// ExitEnumeration is called when production enumeration is exited.
func (s *BaseASN_3gppListener) ExitEnumeration(ctx *EnumerationContext) {}

// EnterEnumerationItem is called when production enumerationItem is entered.
func (s *BaseASN_3gppListener) EnterEnumerationItem(ctx *EnumerationItemContext) {}

// ExitEnumerationItem is called when production enumerationItem is exited.
func (s *BaseASN_3gppListener) ExitEnumerationItem(ctx *EnumerationItemContext) {}

// EnterNamedNumber is called when production namedNumber is entered.
func (s *BaseASN_3gppListener) EnterNamedNumber(ctx *NamedNumberContext) {}

// ExitNamedNumber is called when production namedNumber is exited.
func (s *BaseASN_3gppListener) ExitNamedNumber(ctx *NamedNumberContext) {}

// EnterDefinedValue is called when production definedValue is entered.
func (s *BaseASN_3gppListener) EnterDefinedValue(ctx *DefinedValueContext) {}

// ExitDefinedValue is called when production definedValue is exited.
func (s *BaseASN_3gppListener) ExitDefinedValue(ctx *DefinedValueContext) {}

// EnterParameterizedValue is called when production parameterizedValue is entered.
func (s *BaseASN_3gppListener) EnterParameterizedValue(ctx *ParameterizedValueContext) {}

// ExitParameterizedValue is called when production parameterizedValue is exited.
func (s *BaseASN_3gppListener) ExitParameterizedValue(ctx *ParameterizedValueContext) {}

// EnterSimpleDefinedValue is called when production simpleDefinedValue is entered.
func (s *BaseASN_3gppListener) EnterSimpleDefinedValue(ctx *SimpleDefinedValueContext) {}

// ExitSimpleDefinedValue is called when production simpleDefinedValue is exited.
func (s *BaseASN_3gppListener) ExitSimpleDefinedValue(ctx *SimpleDefinedValueContext) {}

// EnterActualParameterList is called when production actualParameterList is entered.
func (s *BaseASN_3gppListener) EnterActualParameterList(ctx *ActualParameterListContext) {}

// ExitActualParameterList is called when production actualParameterList is exited.
func (s *BaseASN_3gppListener) ExitActualParameterList(ctx *ActualParameterListContext) {}

// EnterActualParameter is called when production actualParameter is entered.
func (s *BaseASN_3gppListener) EnterActualParameter(ctx *ActualParameterContext) {}

// ExitActualParameter is called when production actualParameter is exited.
func (s *BaseASN_3gppListener) ExitActualParameter(ctx *ActualParameterContext) {}

// EnterExceptionSpec is called when production exceptionSpec is entered.
func (s *BaseASN_3gppListener) EnterExceptionSpec(ctx *ExceptionSpecContext) {}

// ExitExceptionSpec is called when production exceptionSpec is exited.
func (s *BaseASN_3gppListener) ExitExceptionSpec(ctx *ExceptionSpecContext) {}

// EnterExceptionIdentification is called when production exceptionIdentification is entered.
func (s *BaseASN_3gppListener) EnterExceptionIdentification(ctx *ExceptionIdentificationContext) {}

// ExitExceptionIdentification is called when production exceptionIdentification is exited.
func (s *BaseASN_3gppListener) ExitExceptionIdentification(ctx *ExceptionIdentificationContext) {}

// EnterAdditionalEnumeration is called when production additionalEnumeration is entered.
func (s *BaseASN_3gppListener) EnterAdditionalEnumeration(ctx *AdditionalEnumerationContext) {}

// ExitAdditionalEnumeration is called when production additionalEnumeration is exited.
func (s *BaseASN_3gppListener) ExitAdditionalEnumeration(ctx *AdditionalEnumerationContext) {}

// EnterIntegerType is called when production integerType is entered.
func (s *BaseASN_3gppListener) EnterIntegerType(ctx *IntegerTypeContext) {}

// ExitIntegerType is called when production integerType is exited.
func (s *BaseASN_3gppListener) ExitIntegerType(ctx *IntegerTypeContext) {}

// EnterNamedNumberList is called when production namedNumberList is entered.
func (s *BaseASN_3gppListener) EnterNamedNumberList(ctx *NamedNumberListContext) {}

// ExitNamedNumberList is called when production namedNumberList is exited.
func (s *BaseASN_3gppListener) ExitNamedNumberList(ctx *NamedNumberListContext) {}

// EnterObjectidentifiertype is called when production objectidentifiertype is entered.
func (s *BaseASN_3gppListener) EnterObjectidentifiertype(ctx *ObjectidentifiertypeContext) {}

// ExitObjectidentifiertype is called when production objectidentifiertype is exited.
func (s *BaseASN_3gppListener) ExitObjectidentifiertype(ctx *ObjectidentifiertypeContext) {}

// EnterComponentRelationConstraint is called when production componentRelationConstraint is entered.
func (s *BaseASN_3gppListener) EnterComponentRelationConstraint(ctx *ComponentRelationConstraintContext) {
}

// ExitComponentRelationConstraint is called when production componentRelationConstraint is exited.
func (s *BaseASN_3gppListener) ExitComponentRelationConstraint(ctx *ComponentRelationConstraintContext) {
}

// EnterAtNotation is called when production atNotation is entered.
func (s *BaseASN_3gppListener) EnterAtNotation(ctx *AtNotationContext) {}

// ExitAtNotation is called when production atNotation is exited.
func (s *BaseASN_3gppListener) ExitAtNotation(ctx *AtNotationContext) {}

// EnterLevel is called when production level is entered.
func (s *BaseASN_3gppListener) EnterLevel(ctx *LevelContext) {}

// ExitLevel is called when production level is exited.
func (s *BaseASN_3gppListener) ExitLevel(ctx *LevelContext) {}

// EnterComponentIdList is called when production componentIdList is entered.
func (s *BaseASN_3gppListener) EnterComponentIdList(ctx *ComponentIdListContext) {}

// ExitComponentIdList is called when production componentIdList is exited.
func (s *BaseASN_3gppListener) ExitComponentIdList(ctx *ComponentIdListContext) {}

// EnterOctetStringType is called when production octetStringType is entered.
func (s *BaseASN_3gppListener) EnterOctetStringType(ctx *OctetStringTypeContext) {}

// ExitOctetStringType is called when production octetStringType is exited.
func (s *BaseASN_3gppListener) ExitOctetStringType(ctx *OctetStringTypeContext) {}

// EnterBitStringType is called when production bitStringType is entered.
func (s *BaseASN_3gppListener) EnterBitStringType(ctx *BitStringTypeContext) {}

// ExitBitStringType is called when production bitStringType is exited.
func (s *BaseASN_3gppListener) ExitBitStringType(ctx *BitStringTypeContext) {}

// EnterNamedBitList is called when production namedBitList is entered.
func (s *BaseASN_3gppListener) EnterNamedBitList(ctx *NamedBitListContext) {}

// ExitNamedBitList is called when production namedBitList is exited.
func (s *BaseASN_3gppListener) ExitNamedBitList(ctx *NamedBitListContext) {}

// EnterNamedBit is called when production namedBit is entered.
func (s *BaseASN_3gppListener) EnterNamedBit(ctx *NamedBitContext) {}

// ExitNamedBit is called when production namedBit is exited.
func (s *BaseASN_3gppListener) ExitNamedBit(ctx *NamedBitContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseASN_3gppListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseASN_3gppListener) ExitBooleanValue(ctx *BooleanValueContext) {}
