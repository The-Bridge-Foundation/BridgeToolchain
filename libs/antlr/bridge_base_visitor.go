// Code generated from antlr/Bridge.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Bridge

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseBridgeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBridgeVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitIdentifiers(ctx *IdentifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitExternalDecl(ctx *ExternalDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitOp(ctx *OpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitArrayMembers(ctx *ArrayMembersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitArrayMemberRef(ctx *ArrayMemberRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitTypeofOperation(ctx *TypeofOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitNegateOperation(ctx *NegateOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitIncDecOP(ctx *IncDecOPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitPrefixOperation(ctx *PrefixOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitSuffixOperation(ctx *SuffixOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitOperationWithReturn(ctx *OperationWithReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitAssignmentOp(ctx *AssignmentOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitCallArgs(ctx *CallArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitComparator(ctx *ComparatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitElseBlock(ctx *ElseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitElseifBlock(ctx *ElseifBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitNamespaceAccess(ctx *NamespaceAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitMemberAccess(ctx *MemberAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitTimesDecl(ctx *TimesDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitEachDecl(ctx *EachDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitWhileDecl(ctx *WhileDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitLoopDecl(ctx *LoopDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitCases(ctx *CasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitSwitchBody(ctx *SwitchBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitDeleteOp(ctx *DeleteOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitBuiltin(ctx *BuiltinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitArrayDecl(ctx *ArrayDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitArrayAccess(ctx *ArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitTypeKey(ctx *TypeKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitFuncDecl(ctx *FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitDeclarator(ctx *DeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitDecl(ctx *DeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitClassBody(ctx *ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBridgeVisitor) VisitClassDecl(ctx *ClassDeclContext) interface{} {
	return v.VisitChildren(ctx)
}
