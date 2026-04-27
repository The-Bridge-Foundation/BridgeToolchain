// Code generated from antlr/Bridge.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Bridge

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BridgeListener is a complete listener for a parse tree produced by BridgeParser.
type BridgeListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterIdentifiers is called when entering the identifiers production.
	EnterIdentifiers(c *IdentifiersContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterExternalDecl is called when entering the externalDecl production.
	EnterExternalDecl(c *ExternalDeclContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// EnterArrayMembers is called when entering the arrayMembers production.
	EnterArrayMembers(c *ArrayMembersContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterArrayMemberRef is called when entering the arrayMemberRef production.
	EnterArrayMemberRef(c *ArrayMemberRefContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterTypeofOperation is called when entering the typeofOperation production.
	EnterTypeofOperation(c *TypeofOperationContext)

	// EnterNegateOperation is called when entering the negateOperation production.
	EnterNegateOperation(c *NegateOperationContext)

	// EnterIncDecOP is called when entering the incDecOP production.
	EnterIncDecOP(c *IncDecOPContext)

	// EnterPrefixOperation is called when entering the prefixOperation production.
	EnterPrefixOperation(c *PrefixOperationContext)

	// EnterSuffixOperation is called when entering the suffixOperation production.
	EnterSuffixOperation(c *SuffixOperationContext)

	// EnterOperationWithReturn is called when entering the operationWithReturn production.
	EnterOperationWithReturn(c *OperationWithReturnContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterAssignmentOp is called when entering the assignmentOp production.
	EnterAssignmentOp(c *AssignmentOpContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterCallArgs is called when entering the callArgs production.
	EnterCallArgs(c *CallArgsContext)

	// EnterFuncCall is called when entering the funcCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterElseifBlock is called when entering the elseifBlock production.
	EnterElseifBlock(c *ElseifBlockContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterNamespaceAccess is called when entering the namespaceAccess production.
	EnterNamespaceAccess(c *NamespaceAccessContext)

	// EnterMemberAccess is called when entering the memberAccess production.
	EnterMemberAccess(c *MemberAccessContext)

	// EnterTimesDecl is called when entering the timesDecl production.
	EnterTimesDecl(c *TimesDeclContext)

	// EnterEachDecl is called when entering the eachDecl production.
	EnterEachDecl(c *EachDeclContext)

	// EnterWhileDecl is called when entering the whileDecl production.
	EnterWhileDecl(c *WhileDeclContext)

	// EnterLoopDecl is called when entering the loopDecl production.
	EnterLoopDecl(c *LoopDeclContext)

	// EnterLoopStatement is called when entering the loopStatement production.
	EnterLoopStatement(c *LoopStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterCases is called when entering the cases production.
	EnterCases(c *CasesContext)

	// EnterSwitchBody is called when entering the switchBody production.
	EnterSwitchBody(c *SwitchBodyContext)

	// EnterDeleteOp is called when entering the deleteOp production.
	EnterDeleteOp(c *DeleteOpContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterBuiltin is called when entering the builtin production.
	EnterBuiltin(c *BuiltinContext)

	// EnterArrayDecl is called when entering the arrayDecl production.
	EnterArrayDecl(c *ArrayDeclContext)

	// EnterArrayAccess is called when entering the arrayAccess production.
	EnterArrayAccess(c *ArrayAccessContext)

	// EnterTypeKey is called when entering the typeKey production.
	EnterTypeKey(c *TypeKeyContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterFuncDecl is called when entering the funcDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterDeclarator is called when entering the declarator production.
	EnterDeclarator(c *DeclaratorContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterClassDecl is called when entering the classDecl production.
	EnterClassDecl(c *ClassDeclContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitIdentifiers is called when exiting the identifiers production.
	ExitIdentifiers(c *IdentifiersContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitExternalDecl is called when exiting the externalDecl production.
	ExitExternalDecl(c *ExternalDeclContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)

	// ExitArrayMembers is called when exiting the arrayMembers production.
	ExitArrayMembers(c *ArrayMembersContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitArrayMemberRef is called when exiting the arrayMemberRef production.
	ExitArrayMemberRef(c *ArrayMemberRefContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitTypeofOperation is called when exiting the typeofOperation production.
	ExitTypeofOperation(c *TypeofOperationContext)

	// ExitNegateOperation is called when exiting the negateOperation production.
	ExitNegateOperation(c *NegateOperationContext)

	// ExitIncDecOP is called when exiting the incDecOP production.
	ExitIncDecOP(c *IncDecOPContext)

	// ExitPrefixOperation is called when exiting the prefixOperation production.
	ExitPrefixOperation(c *PrefixOperationContext)

	// ExitSuffixOperation is called when exiting the suffixOperation production.
	ExitSuffixOperation(c *SuffixOperationContext)

	// ExitOperationWithReturn is called when exiting the operationWithReturn production.
	ExitOperationWithReturn(c *OperationWithReturnContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitAssignmentOp is called when exiting the assignmentOp production.
	ExitAssignmentOp(c *AssignmentOpContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitCallArgs is called when exiting the callArgs production.
	ExitCallArgs(c *CallArgsContext)

	// ExitFuncCall is called when exiting the funcCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitElseifBlock is called when exiting the elseifBlock production.
	ExitElseifBlock(c *ElseifBlockContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitNamespaceAccess is called when exiting the namespaceAccess production.
	ExitNamespaceAccess(c *NamespaceAccessContext)

	// ExitMemberAccess is called when exiting the memberAccess production.
	ExitMemberAccess(c *MemberAccessContext)

	// ExitTimesDecl is called when exiting the timesDecl production.
	ExitTimesDecl(c *TimesDeclContext)

	// ExitEachDecl is called when exiting the eachDecl production.
	ExitEachDecl(c *EachDeclContext)

	// ExitWhileDecl is called when exiting the whileDecl production.
	ExitWhileDecl(c *WhileDeclContext)

	// ExitLoopDecl is called when exiting the loopDecl production.
	ExitLoopDecl(c *LoopDeclContext)

	// ExitLoopStatement is called when exiting the loopStatement production.
	ExitLoopStatement(c *LoopStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitCases is called when exiting the cases production.
	ExitCases(c *CasesContext)

	// ExitSwitchBody is called when exiting the switchBody production.
	ExitSwitchBody(c *SwitchBodyContext)

	// ExitDeleteOp is called when exiting the deleteOp production.
	ExitDeleteOp(c *DeleteOpContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitBuiltin is called when exiting the builtin production.
	ExitBuiltin(c *BuiltinContext)

	// ExitArrayDecl is called when exiting the arrayDecl production.
	ExitArrayDecl(c *ArrayDeclContext)

	// ExitArrayAccess is called when exiting the arrayAccess production.
	ExitArrayAccess(c *ArrayAccessContext)

	// ExitTypeKey is called when exiting the typeKey production.
	ExitTypeKey(c *TypeKeyContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitFuncDecl is called when exiting the funcDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitDeclarator is called when exiting the declarator production.
	ExitDeclarator(c *DeclaratorContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitClassDecl is called when exiting the classDecl production.
	ExitClassDecl(c *ClassDeclContext)
}
