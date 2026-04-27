// Code generated from antlr/Bridge.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Bridge

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by BridgeParser.
type BridgeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BridgeParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by BridgeParser#identifiers.
	VisitIdentifiers(ctx *IdentifiersContext) interface{}

	// Visit a parse tree produced by BridgeParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by BridgeParser#externalDecl.
	VisitExternalDecl(ctx *ExternalDeclContext) interface{}

	// Visit a parse tree produced by BridgeParser#op.
	VisitOp(ctx *OpContext) interface{}

	// Visit a parse tree produced by BridgeParser#arrayMembers.
	VisitArrayMembers(ctx *ArrayMembersContext) interface{}

	// Visit a parse tree produced by BridgeParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by BridgeParser#arrayMemberRef.
	VisitArrayMemberRef(ctx *ArrayMemberRefContext) interface{}

	// Visit a parse tree produced by BridgeParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by BridgeParser#typeofOperation.
	VisitTypeofOperation(ctx *TypeofOperationContext) interface{}

	// Visit a parse tree produced by BridgeParser#negateOperation.
	VisitNegateOperation(ctx *NegateOperationContext) interface{}

	// Visit a parse tree produced by BridgeParser#incDecOP.
	VisitIncDecOP(ctx *IncDecOPContext) interface{}

	// Visit a parse tree produced by BridgeParser#prefixOperation.
	VisitPrefixOperation(ctx *PrefixOperationContext) interface{}

	// Visit a parse tree produced by BridgeParser#suffixOperation.
	VisitSuffixOperation(ctx *SuffixOperationContext) interface{}

	// Visit a parse tree produced by BridgeParser#operationWithReturn.
	VisitOperationWithReturn(ctx *OperationWithReturnContext) interface{}

	// Visit a parse tree produced by BridgeParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by BridgeParser#assignmentOp.
	VisitAssignmentOp(ctx *AssignmentOpContext) interface{}

	// Visit a parse tree produced by BridgeParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by BridgeParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by BridgeParser#callArgs.
	VisitCallArgs(ctx *CallArgsContext) interface{}

	// Visit a parse tree produced by BridgeParser#funcCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by BridgeParser#comparator.
	VisitComparator(ctx *ComparatorContext) interface{}

	// Visit a parse tree produced by BridgeParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by BridgeParser#elseBlock.
	VisitElseBlock(ctx *ElseBlockContext) interface{}

	// Visit a parse tree produced by BridgeParser#elseifBlock.
	VisitElseifBlock(ctx *ElseifBlockContext) interface{}

	// Visit a parse tree produced by BridgeParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by BridgeParser#namespaceAccess.
	VisitNamespaceAccess(ctx *NamespaceAccessContext) interface{}

	// Visit a parse tree produced by BridgeParser#memberAccess.
	VisitMemberAccess(ctx *MemberAccessContext) interface{}

	// Visit a parse tree produced by BridgeParser#timesDecl.
	VisitTimesDecl(ctx *TimesDeclContext) interface{}

	// Visit a parse tree produced by BridgeParser#eachDecl.
	VisitEachDecl(ctx *EachDeclContext) interface{}

	// Visit a parse tree produced by BridgeParser#whileDecl.
	VisitWhileDecl(ctx *WhileDeclContext) interface{}

	// Visit a parse tree produced by BridgeParser#loopDecl.
	VisitLoopDecl(ctx *LoopDeclContext) interface{}

	// Visit a parse tree produced by BridgeParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by BridgeParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by BridgeParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by BridgeParser#cases.
	VisitCases(ctx *CasesContext) interface{}

	// Visit a parse tree produced by BridgeParser#switchBody.
	VisitSwitchBody(ctx *SwitchBodyContext) interface{}

	// Visit a parse tree produced by BridgeParser#deleteOp.
	VisitDeleteOp(ctx *DeleteOpContext) interface{}

	// Visit a parse tree produced by BridgeParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by BridgeParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by BridgeParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by BridgeParser#builtin.
	VisitBuiltin(ctx *BuiltinContext) interface{}

	// Visit a parse tree produced by BridgeParser#arrayDecl.
	VisitArrayDecl(ctx *ArrayDeclContext) interface{}

	// Visit a parse tree produced by BridgeParser#arrayAccess.
	VisitArrayAccess(ctx *ArrayAccessContext) interface{}

	// Visit a parse tree produced by BridgeParser#typeKey.
	VisitTypeKey(ctx *TypeKeyContext) interface{}

	// Visit a parse tree produced by BridgeParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by BridgeParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by BridgeParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by BridgeParser#declarator.
	VisitDeclarator(ctx *DeclaratorContext) interface{}

	// Visit a parse tree produced by BridgeParser#decl.
	VisitDecl(ctx *DeclContext) interface{}

	// Visit a parse tree produced by BridgeParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by BridgeParser#classDecl.
	VisitClassDecl(ctx *ClassDeclContext) interface{}
}
