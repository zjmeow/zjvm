package instructions

import (
	"fmt"
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/instructions/comparisons"
	"github.com/zjmeow/zjvm/instructions/constants"
	"github.com/zjmeow/zjvm/instructions/control"
	"github.com/zjmeow/zjvm/instructions/conversions"
	"github.com/zjmeow/zjvm/instructions/loads"
	"github.com/zjmeow/zjvm/instructions/math"
	"github.com/zjmeow/zjvm/instructions/references"
	"github.com/zjmeow/zjvm/instructions/stack"
	"github.com/zjmeow/zjvm/instructions/stores"
)

// NoOperandsInstruction singletons
var (
	nop         = &constants.Nop{}
	aconst_null = &constants.ACONST_NULL{}
	iconst_m1   = constants.NewConstInt(-1)
	iconst_0    = constants.NewConstInt(0)
	iconst_1    = constants.NewConstInt(1)
	iconst_2    = constants.NewConstInt(2)
	iconst_3    = constants.NewConstInt(3)
	iconst_4    = constants.NewConstInt(4)
	iconst_5    = constants.NewConstInt(5)
	lconst_0    = constants.NewConstLong(0)
	lconst_1    = constants.NewConstLong(1)
	fconst_0    = constants.NewConstFloat(0)
	fconst_1    = constants.NewConstFloat(1.0)
	fconst_2    = constants.NewConstFloat(2.0)
	dconst_0    = constants.NewConstDouble(0.0)
	dconst_1    = constants.NewConstDouble(1.0)
	iload_0     = loads.NewLoadInt(0)
	iload_1     = loads.NewLoadInt(1)
	iload_2     = loads.NewLoadInt(2)
	iload_3     = loads.NewLoadInt(3)
	lload_0     = loads.NewLongLoad(0)
	lload_1     = loads.NewLongLoad(1)
	lload_2     = loads.NewLongLoad(2)
	lload_3     = loads.NewLongLoad(3)
	fload_0     = loads.NewFloatLoad(0)
	fload_1     = loads.NewFloatLoad(1)
	fload_2     = loads.NewFloatLoad(2)
	fload_3     = loads.NewFloatLoad(3)
	dload_0     = loads.NewDoubleLoad(0)
	dload_1     = loads.NewDoubleLoad(1)
	dload_2     = loads.NewDoubleLoad(2)
	dload_3     = loads.NewDoubleLoad(3)
	aload_0     = loads.NewALoad(0)
	aload_1     = loads.NewALoad(1)
	aload_2     = loads.NewALoad(2)
	aload_3     = loads.NewALoad(3)
	iaload      = loads.NewIALoad()
	laload      = loads.NewLALoad()
	faload      = loads.NewFALoad()
	daload      = loads.NewDALoad()
	aaload      = loads.NewAALoad()
	baload      = loads.NewBALoad()
	caload      = loads.NewCALoad()
	saload      = loads.NewSALoad()
	istore_0    = stores.NewStoreInt(0)
	istore_1    = stores.NewStoreInt(1)
	istore_2    = stores.NewStoreInt(2)
	istore_3    = stores.NewStoreInt(3)
	lstore_0    = stores.NewStoreLong(0)
	lstore_1    = stores.NewStoreLong(1)
	lstore_2    = stores.NewStoreLong(2)
	lstore_3    = stores.NewStoreLong(3)
	fstore_0    = stores.NewStoreFloat(0)
	fstore_1    = stores.NewStoreFloat(1)
	fstore_2    = stores.NewStoreFloat(2)
	fstore_3    = stores.NewStoreFloat(3)
	dstore_0    = stores.NewStoreDouble(0)
	dstore_1    = stores.NewStoreDouble(1)
	dstore_2    = stores.NewStoreDouble(2)
	dstore_3    = stores.NewStoreDouble(3)
	astore_0    = stores.NewStoreA(0)
	astore_1    = stores.NewStoreA(1)
	astore_2    = stores.NewStoreA(2)
	astore_3    = stores.NewStoreA(3)
	iastore     = stores.NewIAStore()
	lastore     = stores.NewLAStore()
	fastore     = stores.NewFAStore()
	dastore     = stores.NewDAStore()
	aastore     = stores.NewAAStore()
	bastore     = stores.NewBAStore()
	castore     = stores.NewCAStore()
	sastore     = stores.NewSAStore()
	pop         = &stack.Pop{}
	pop2        = &stack.Pop2{}
	dup         = &stack.Dup{}
	dup_x1      = &stack.DupX1{}
	dup_x2      = &stack.DupX2{}
	dup2        = &stack.Dup2{}
	dup2_x1     = &stack.Dup2X1{}
	dup2_x2     = &stack.Dup2X2{}
	swap        = &stack.Swap{}
	iadd        = math.NewIAdd()
	ladd        = math.NewLAdd()
	fadd        = math.NewFAdd()
	dadd        = math.NewDAdd()
	isub        = math.NewISub()
	lsub        = math.NewLSub()
	fsub        = math.NewFSub()
	dsub        = math.NewDSub()
	imul        = math.NewIMul()
	lmul        = math.NewLMul()
	fmul        = math.NewFMul()
	dmul        = math.NewDMul()
	idiv        = math.NewIDiv()
	ldiv        = math.NewLDiv()
	fdiv        = math.NewFDiv()
	ddiv        = math.NewDDiv()
	irem        = math.NewIRem()
	lrem        = math.NewLRem()
	frem        = math.NewFRem()
	drem        = math.NewDRem()
	ineg        = math.NewINeg()
	lneg        = math.NewLNeg()
	fneg        = math.NewFNeg()
	dneg        = math.NewDNeg()
	ishl        = math.NewIShl()
	lshl        = math.NewLShl()
	ishr        = math.NewIShr()
	lshr        = math.NewLShr()
	iushr       = math.NewIUShr()
	lushr       = math.NewLUShr()
	iand        = math.NewIAnd()
	land        = math.NewLAnd()
	ior         = math.NewIOr()
	lor         = math.NewLOr()
	ixor        = math.NewIXor()
	lxor        = math.NewLXor()
	i2l         = conversions.NewI2L()
	i2f         = conversions.NewI2F()
	i2d         = conversions.NewI2D()
	l2i         = conversions.NewL2I()
	l2f         = conversions.NewL2F()
	l2d         = conversions.NewL2D()
	f2i         = conversions.NewF2I()
	f2l         = conversions.NewF2L()
	f2d         = conversions.NewF2D()
	d2i         = conversions.NewD2I()
	d2l         = conversions.NewD2L()
	d2f         = conversions.NewD2F()
	i2b         = conversions.NewI2B()
	i2c         = conversions.NewI2C()
	i2s         = conversions.NewI2S()
	lcmp        = comparisons.NewLCMP()
	fcmpl       = comparisons.NewFCMPL()
	fcmpg       = comparisons.NewFCMPG()
	dcmpl       = comparisons.NewDCMPL()
	dcmpg       = comparisons.NewDCMPG()
	ireturn     = control.NewReturn("i")
	lreturn     = control.NewReturn("l")
	freturn     = control.NewReturn("f")
	dreturn     = control.NewReturn("d")
	areturn     = control.NewReturn("a")
	_return     = control.NewReturn("v")
	arraylength = &references.ArrayLength{}
	//athrow        = &AThrow{}
	//monitorenter  = &MonitorEnter{}
	//monitorexit   = &MonitorExit{}
	//invoke_native = &InvokeNative{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case OpNop:
		return nop
	case OpAConstNull:
		return aconst_null
	case OpIConstM1:
		return iconst_m1
	case OpIConst0:
		return iconst_0
	case OpIConst1:
		return iconst_1
	case OpIConst2:
		return iconst_2
	case OpIConst3:
		return iconst_3
	case OpIConst4:
		return iconst_4
	case OpIConst5:
		return iconst_5
	case OpLConst0:
		return lconst_0
	case OpLConst1:
		return lconst_1
	case OpFConst0:
		return fconst_0
	case OpFConst1:
		return fconst_1
	case OpFConst2:
		return fconst_2
	case OpDConst0:
		return dconst_0
	case OpDConst1:
		return dconst_1
	case OpBIPush:
		return &constants.BIPush{}
	case OpSIPush:
		return &constants.SIPush{}
	case OpLDC:
		return &constants.Ldc{}
	case OpLDCw:
		return &constants.LdcW{}
	case OpLDC2w:
		return &constants.Ldc2W{}
	case OpILoad:
		return &loads.ILOAD{}
	case OpLLoad:
		return &loads.LLOAD{}
	case OpFLoad:
		return &loads.FLOAD{}
	case OpDLoad:
		return &loads.DLOAD{}
	case OpALoad:
		return &loads.ALOAD{}
	case OpILoad0:
		return iload_0
	case OpILoad1:
		return iload_1
	case OpILoad2:
		return iload_2
	case OpILoad3:
		return iload_3
	case OpLLoad0:
		return lload_0
	case OpLLoad1:
		return lload_1
	case OpLLoad2:
		return lload_2
	case OpLLoad3:
		return lload_3
	case OpFLoad0:
		return fload_0
	case OpFLoad1:
		return fload_1
	case OpFLoad2:
		return fload_2
	case OpFLoad3:
		return fload_3
	case OpDLoad0:
		return dload_0
	case OpDLoad1:
		return dload_1
	case OpDLoad2:
		return dload_2
	case OpDLoad3:
		return dload_3
	case OpALoad0:
		return aload_0
	case OpALoad1:
		return aload_1
	case OpALoad2:
		return aload_2
	case OpALoad3:
		return aload_3
	case OpIALoad:
		return iaload
	case OpLALoad:
		return laload
	case OpFALoad:
		return faload
	case OpDALoad:
		return daload
	case OpAALoad:
		return aaload
	case OpBALoad:
		return baload
	case OpCALoad:
		return caload
	case OpSALoad:
		return saload
	case OpIStore:
		return &stores.ISTORE{}
	case OpLStore:
		return &stores.LSTORE{}
	case OpFStore:
		return &stores.FSTORE{}
	case OpDStore:
		return &stores.DSTORE{}
	case OpAStore:
		return &stores.ASTORE{}
	case OpIStore0:
		return istore_0
	case OpIStore1:
		return istore_1
	case OpIStore2:
		return istore_2
	case OpIStore3:
		return istore_3
	case OpLStore0:
		return lstore_0
	case OpLStore1:
		return lstore_1
	case OpLStore2:
		return lstore_2
	case OpLStore3:
		return lstore_3
	case OpFStore0:
		return fstore_0
	case OpFStore1:
		return fstore_1
	case OpFStore2:
		return fstore_2
	case OpFStore3:
		return fstore_3
	case OpDStore0:
		return dstore_0
	case OpDStore1:
		return dstore_1
	case OpDStore2:
		return dstore_2
	case OpDStore3:
		return dstore_3
	case OpAStore0:
		return astore_0
	case OpAStore1:
		return astore_1
	case OpAStore2:
		return astore_2
	case OpAStore3:
		return astore_3
	case OpIAStore:
		return iastore
	case OpLAStore:
		return lastore
	case OpFAStore:
		return fastore
	case OpDAStore:
		return dastore
	case OpAAStore:
		return aastore
	case OpBAStore:
		return bastore
	case OpCAStore:
		return castore
	case OpSAStore:
		return sastore
	case OpPop:
		return pop
	case OpPop2:
		return pop2
	case OpDup:
		return dup
	case OpDupX1:
		return dup_x1
	case OpDupX2:
		return dup_x2
	case OpDup2:
		return dup2
	case OpDup2X1:
		return dup2_x1
	case OpDup2X2:
		return dup2_x2
	case OpSwap:
		return swap
	case OpIAdd:
		return iadd
	case OpLAdd:
		return ladd
	case OpFAdd:
		return fadd
	case OpDAdd:
		return dadd
	case OpISub:
		return isub
	case OpLSub:
		return lsub
	case OpFSub:
		return fsub
	case OpDSub:
		return dsub
	case OpIMul:
		return imul
	case OpLMul:
		return lmul
	case OpFMul:
		return fmul
	case OpDMul:
		return dmul
	case OpIDiv:
		return idiv
	case OpLDiv:
		return ldiv
	case OpFDiv:
		return fdiv
	case OpDDiv:
		return ddiv
	case OpIRem:
		return irem
	case OpLRem:
		return lrem
	case OpFRem:
		return frem
	case OpDRem:
		return drem
	case OpINeg:
		return ineg
	case OpLNeg:
		return lneg
	case OpFNeg:
		return fneg
	case OpDNeg:
		return dneg
	case OpIShl:
		return ishl
	case OpLShl:
		return lshl
	case OpIShr:
		return ishr
	case OpLShr:
		return lshr
	case OpIUshr:
		return iushr
	case OpLUshr:
		return lushr
	case OpIAnd:
		return iand
	case OpLAnd:
		return land
	case OpIOr:
		return ior
	case OpLOr:
		return lor
	case OpIXor:
		return ixor
	case OpLXor:
		return lxor
	case OpIInc:
		return math.NewIInc()
	case OpI2L:
		return i2l
	case OpI2F:
		return i2f
	case OpI2D:
		return i2d
	case OpL2I:
		return l2i
	case OpL2F:
		return l2f
	case OpL2D:
		return l2d
	case OpF2I:
		return f2i
	case OpF2L:
		return f2l
	case OpF2D:
		return f2d
	case OpD2I:
		return d2i
	case OpD2L:
		return d2l
	case OpD2F:
		return d2f
	case OpI2B:
		return i2b
	case OpI2C:
		return i2c
	case OpI2S:
		return i2s
	case OpLCmp:
		return lcmp
	case OpFCmpL:
		return fcmpl
	case OpFCmpG:
		return fcmpg
	case OpDCmpL:
		return dcmpl
	case OpDCmpG:
		return dcmpg
	case OpIfEQ:
		return comparisons.NewIfEQ()
	case OpIfNE:
		return comparisons.NewIfNE()
	case OpIfLT:
		return comparisons.NewIfLT()
	case OpIfGE:
		return comparisons.NewIfGE()
	case OpIfGT:
		return comparisons.NewIfGT()
	case OpIfLE:
		return comparisons.NewIfLE()
	case OpIfICmpEQ:
		return comparisons.NewIfICmpEQ()
	case OpIfICmpNE:
		return comparisons.NewIfICmpNE()
	case OpIfICmpLT:
		return comparisons.NewIfICmpLT()
	case OpIfICmpGE:
		return comparisons.NewIfICmpGE()
	case OpIfICmpGT:
		return comparisons.NewIfICmpGT()
	case OpIfICmpLE:
		return comparisons.NewIfICmpLE()
	case OpIfACmpEQ:
		return comparisons.NewIfACmpEQ()
	case OpIfACmpNE:
		return comparisons.NewIfACmpNE()
	case OpGoto:
		return &control.Goto{}
	//case OpJSR:
	//	return &JSR{}
	//case OpRET:
	//	return &RET{}
	case OpTableSwitch:
		return &control.TableSwitch{}
	case OpLookupSwitch:
		return &control.LookupSwitch{}
	case OpIReturn:
		return ireturn
	case OpLReturn:
		return lreturn
	case OpFReturn:
		return freturn
	case OpDReturn:
		return dreturn
	case OpAReturn:
		return areturn
	case OpReturn:
		return _return
	case OpGetStatic:
		return &references.GetStatic{}
	case OpPutStatic:
		return &references.PutStatic{}
	case OpGetField:
		return &references.GetField{}
	case OpPutField:
		return &references.PutField{}
	case OpInvokeVirtual:
		return &references.InvokeVirtual{}
	case OpInvokeSpecial:
		return &references.InvokeSpecial{}
	case OpInvokeStatic:
		return &references.InvokeStatic{}
	case OpInvokeInterface:
		return &references.InvokeInterface{}
	//case OpInvokeDynamic:
	//	return &InvokeDynamic{}
	case OpNew:
		return &references.New{}
	case OpNewArray:
		return &references.NewArray{}
	case OpANewArray:
		return &references.ANewArray{}
	case OpArrayLength:
		return arraylength
	//case OpAThrow:
	//	return athrow
	case OpCheckCast:
		return &references.CheckCast{}
	case OpInstanceOf:
		return &references.InstanceOf{}
	//case OpMonitorEnter:
	//	return monitorenter
	//case OpMonitorExit:
	//	return monitorexit
	//case OpWide:
	//	return &Wide{}
	//case OpMultiANewArray:
	//	return &MultiANewArray{}
	case OpIfNull:
		return comparisons.NewIfNull()
	case OpIfNonNull:
		return comparisons.NewIfNonNull()
	//case OpGotoW:
	//	return &GotoW{}
	//case OpJSRw:
	//	return &JSR_W{}
	////case 0xca: todo breakpoint
	//case OpBreakpoint:
	//	return invoke_native // impdep1
	//case OpInvokeNative:
	//	return &Bootstrap{} // impdep2
	default:
		panic(fmt.Errorf("invalid opcode: %v", opcode))
	}
}
