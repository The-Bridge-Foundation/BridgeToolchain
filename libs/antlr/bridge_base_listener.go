// Code generated from antlr/Bridge.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Bridge

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBridgeListener is a complete listener for a parse tree produced by BridgeParser.
type BaseBridgeListener struct{}

var _ BridgeListener = &BaseBridgeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBridgeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBridgeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBridgeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBridgeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseBridgeListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseBridgeListener) ExitStart(ctx *StartContext) {}

// EnterIdentifiers is called when production identifiers is entered.
func (s *BaseBridgeListener) EnterIdentifiers(ctx *IdentifiersContext) {}

// ExitIdentifiers is called when production identifiers is exited.
func (s *BaseBridgeListener) ExitIdentifiers(ctx *IdentifiersContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseBridgeListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseBridgeListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterExternalDecl is called when production externalDecl is entered.
func (s *BaseBridgeListener) EnterExternalDecl(ctx *ExternalDeclContext) {}

// ExitExternalDecl is called when production externalDecl is exited.
func (s *BaseBridgeListener) ExitExternalDecl(ctx *ExternalDeclContext) {}

// EnterOp is called when production op is entered.
func (s *BaseBridgeListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseBridgeListener) ExitOp(ctx *OpContext) {}

// EnterArrayMembers is called when production arrayMembers is entered.
func (s *BaseBridgeListener) EnterArrayMembers(ctx *ArrayMembersContext) {}

// ExitArrayMembers is called when production arrayMembers is exited.
func (s *BaseBridgeListener) ExitArrayMembers(ctx *ArrayMembersContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseBridgeListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseBridgeListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterArrayMemberRef is called when production arrayMemberRef is entered.
func (s *BaseBridgeListener) EnterArrayMemberRef(ctx *ArrayMemberRefContext) {}

// ExitArrayMemberRef is called when production arrayMemberRef is exited.
func (s *BaseBridgeListener) ExitArrayMemberRef(ctx *ArrayMemberRefContext) {}

// EnterValue is called when production value is entered.
func (s *BaseBridgeListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseBridgeListener) ExitValue(ctx *ValueContext) {}

// EnterTypeofOperation is called when production typeofOperation is entered.
func (s *BaseBridgeListener) EnterTypeofOperation(ctx *TypeofOperationContext) {}

// ExitTypeofOperation is called when production typeofOperation is exited.
func (s *BaseBridgeListener) ExitTypeofOperation(ctx *TypeofOperationContext) {}

// EnterNegateOperation is called when production negateOperation is entered.
func (s *BaseBridgeListener) EnterNegateOperation(ctx *NegateOperationContext) {}

// ExitNegateOperation is called when production negateOperation is exited.
func (s *BaseBridgeListener) ExitNegateOperation(ctx *NegateOperationContext) {}

// EnterIncDecOP is called when production incDecOP is entered.
func (s *BaseBridgeListener) EnterIncDecOP(ctx *IncDecOPContext) {}

// ExitIncDecOP is called when production incDecOP is exited.
func (s *BaseBridgeListener) ExitIncDecOP(ctx *IncDecOPContext) {}

// EnterPrefixOperation is called when production prefixOperation is entered.
func (s *BaseBridgeListener) EnterPrefixOperation(ctx *PrefixOperationContext) {}

// ExitPrefixOperation is called when production prefixOperation is exited.
func (s *BaseBridgeListener) ExitPrefixOperation(ctx *PrefixOperationContext) {}

// EnterSuffixOperation is called when production suffixOperation is entered.
func (s *BaseBridgeListener) EnterSuffixOperation(ctx *SuffixOperationContext) {}

// ExitSuffixOperation is called when production suffixOperation is exited.
func (s *BaseBridgeListener) ExitSuffixOperation(ctx *SuffixOperationContext) {}

// EnterOperationWithReturn is called when production operationWithReturn is entered.
func (s *BaseBridgeListener) EnterOperationWithReturn(ctx *OperationWithReturnContext) {}

// ExitOperationWithReturn is called when production operationWithReturn is exited.
func (s *BaseBridgeListener) ExitOperationWithReturn(ctx *OperationWithReturnContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseBridgeListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseBridgeListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterAssignmentOp is called when production assignmentOp is entered.
func (s *BaseBridgeListener) EnterAssignmentOp(ctx *AssignmentOpContext) {}

// ExitAssignmentOp is called when production assignmentOp is exited.
func (s *BaseBridgeListener) ExitAssignmentOp(ctx *AssignmentOpContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseBridgeListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseBridgeListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseBridgeListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseBridgeListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterCallArgs is called when production callArgs is entered.
func (s *BaseBridgeListener) EnterCallArgs(ctx *CallArgsContext) {}

// ExitCallArgs is called when production callArgs is exited.
func (s *BaseBridgeListener) ExitCallArgs(ctx *CallArgsContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BaseBridgeListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BaseBridgeListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseBridgeListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseBridgeListener) ExitComparator(ctx *ComparatorContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseBridgeListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseBridgeListener) ExitCondition(ctx *ConditionContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseBridgeListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseBridgeListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterElseifBlock is called when production elseifBlock is entered.
func (s *BaseBridgeListener) EnterElseifBlock(ctx *ElseifBlockContext) {}

// ExitElseifBlock is called when production elseifBlock is exited.
func (s *BaseBridgeListener) ExitElseifBlock(ctx *ElseifBlockContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseBridgeListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseBridgeListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterNamespaceAccess is called when production namespaceAccess is entered.
func (s *BaseBridgeListener) EnterNamespaceAccess(ctx *NamespaceAccessContext) {}

// ExitNamespaceAccess is called when production namespaceAccess is exited.
func (s *BaseBridgeListener) ExitNamespaceAccess(ctx *NamespaceAccessContext) {}

// EnterMemberAccess is called when production memberAccess is entered.
func (s *BaseBridgeListener) EnterMemberAccess(ctx *MemberAccessContext) {}

// ExitMemberAccess is called when production memberAccess is exited.
func (s *BaseBridgeListener) ExitMemberAccess(ctx *MemberAccessContext) {}

// EnterTimesDecl is called when production timesDecl is entered.
func (s *BaseBridgeListener) EnterTimesDecl(ctx *TimesDeclContext) {}

// ExitTimesDecl is called when production timesDecl is exited.
func (s *BaseBridgeListener) ExitTimesDecl(ctx *TimesDeclContext) {}

// EnterEachDecl is called when production eachDecl is entered.
func (s *BaseBridgeListener) EnterEachDecl(ctx *EachDeclContext) {}

// ExitEachDecl is called when production eachDecl is exited.
func (s *BaseBridgeListener) ExitEachDecl(ctx *EachDeclContext) {}

// EnterWhileDecl is called when production whileDecl is entered.
func (s *BaseBridgeListener) EnterWhileDecl(ctx *WhileDeclContext) {}

// ExitWhileDecl is called when production whileDecl is exited.
func (s *BaseBridgeListener) ExitWhileDecl(ctx *WhileDeclContext) {}

// EnterLoopDecl is called when production loopDecl is entered.
func (s *BaseBridgeListener) EnterLoopDecl(ctx *LoopDeclContext) {}

// ExitLoopDecl is called when production loopDecl is exited.
func (s *BaseBridgeListener) ExitLoopDecl(ctx *LoopDeclContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *BaseBridgeListener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *BaseBridgeListener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseBridgeListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseBridgeListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseBridgeListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseBridgeListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCases is called when production cases is entered.
func (s *BaseBridgeListener) EnterCases(ctx *CasesContext) {}

// ExitCases is called when production cases is exited.
func (s *BaseBridgeListener) ExitCases(ctx *CasesContext) {}

// EnterSwitchBody is called when production switchBody is entered.
func (s *BaseBridgeListener) EnterSwitchBody(ctx *SwitchBodyContext) {}

// ExitSwitchBody is called when production switchBody is exited.
func (s *BaseBridgeListener) ExitSwitchBody(ctx *SwitchBodyContext) {}

// EnterDeleteOp is called when production deleteOp is entered.
func (s *BaseBridgeListener) EnterDeleteOp(ctx *DeleteOpContext) {}

// ExitDeleteOp is called when production deleteOp is exited.
func (s *BaseBridgeListener) ExitDeleteOp(ctx *DeleteOpContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseBridgeListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseBridgeListener) ExitStatement(ctx *StatementContext) {}

// EnterBody is called when production body is entered.
func (s *BaseBridgeListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseBridgeListener) ExitBody(ctx *BodyContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseBridgeListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseBridgeListener) ExitFunction(ctx *FunctionContext) {}

// EnterBuiltin is called when production builtin is entered.
func (s *BaseBridgeListener) EnterBuiltin(ctx *BuiltinContext) {}

// ExitBuiltin is called when production builtin is exited.
func (s *BaseBridgeListener) ExitBuiltin(ctx *BuiltinContext) {}

// EnterArrayDecl is called when production arrayDecl is entered.
func (s *BaseBridgeListener) EnterArrayDecl(ctx *ArrayDeclContext) {}

// ExitArrayDecl is called when production arrayDecl is exited.
func (s *BaseBridgeListener) ExitArrayDecl(ctx *ArrayDeclContext) {}

// EnterArrayAccess is called when production arrayAccess is entered.
func (s *BaseBridgeListener) EnterArrayAccess(ctx *ArrayAccessContext) {}

// ExitArrayAccess is called when production arrayAccess is exited.
func (s *BaseBridgeListener) ExitArrayAccess(ctx *ArrayAccessContext) {}

// EnterTypeKey is called when production typeKey is entered.
func (s *BaseBridgeListener) EnterTypeKey(ctx *TypeKeyContext) {}

// ExitTypeKey is called when production typeKey is exited.
func (s *BaseBridgeListener) ExitTypeKey(ctx *TypeKeyContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseBridgeListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseBridgeListener) ExitArgument(ctx *ArgumentContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseBridgeListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseBridgeListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterFuncDecl is called when production funcDecl is entered.
func (s *BaseBridgeListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production funcDecl is exited.
func (s *BaseBridgeListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseBridgeListener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseBridgeListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseBridgeListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseBridgeListener) ExitDecl(ctx *DeclContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseBridgeListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseBridgeListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassDecl is called when production classDecl is entered.
func (s *BaseBridgeListener) EnterClassDecl(ctx *ClassDeclContext) {}

// ExitClassDecl is called when production classDecl is exited.
func (s *BaseBridgeListener) ExitClassDecl(ctx *ClassDeclContext) {}
