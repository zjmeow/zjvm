package instructions

import (
	"fmt"
	"github.com/zjmeow/zjvm/instructions/base"
	"github.com/zjmeow/zjvm/instructions/constants"
	"github.com/zjmeow/zjvm/instructions/loads"
)

// NoOperandsInstruction singletons
var (
	nop           = &Nop
	aconst_null   = NewConstNull()
	iconst_m1     = constants.NewConstInt(-1)
	iconst_0      = constants.NewConstInt(0)
	iconst_1      = constants.NewConstInt(1)
	iconst_2      = constants.NewConstInt(2)
	iconst_3      = constants.NewConstInt(3)
	iconst_4      = constants.NewConstInt(4)
	iconst_5      = constants.NewConstInt(5)
	lconst_0      = constants.NewConstLong(0)
	lconst_1      = constants.NewConstLong(1)
	fconst_0      = constants.NewConstFloat(0)
	fconst_1      = constants.NewConstFloat(1.0)
	fconst_2      = constants.NewConstFloat(2.0)
	dconst_0      = constants.NewConstDouble(0.0)
	dconst_1      = constants.NewConstDouble(1.0)
	iload_0       = loads.NewLoadInt(0)
	iload_1       = loads.NewLoadInt(1)
	iload_2       = loads.NewLoadInt(2)
	iload_3       = loads.NewLoadInt(3)
	lload_0       = loads.NewLongLoad(0)
	lload_1       = loads.NewLongLoad(1)
	lload_2       = loads.NewLongLoad(2)
	lload_3       = loads.NewLongLoad(3)
	fload_0       = loads.NewFloatLoad(0)
	fload_1       = loads.NewFloatLoad(1)
	fload_2       = loads.NewFloatLoad(2)
	fload_3       = loads.NewFloatLoad(3)
	dload_0       = loads.NewDoubleLoad(0)
	dload_1       = loads.NewDoubleLoad(1)
	dload_2       = loads.NewDoubleLoad(2)
	dload_3       = loads.NewDoubleLoad(3)
	aload_0       = loads.NewALoad(0)
	aload_1       = loads.NewALoad(1)
	aload_2       = loads.NewALoad(2)
	aload_3       = loads.NewALoad(3)
	iaload        = NewIALoad()
	laload        = NewLALoad()
	faload        = NewFALoad()
	daload        = NewDALoad()
	aaload        = NewAALoad()
	baload        = NewBALoad()
	caload        = NewCALoad()
	saload        = NewSALoad()
	istore_0      = NewStoreN(0, false)
	istore_1      = NewStoreN(1, false)
	istore_2      = NewStoreN(2, false)
	istore_3      = NewStoreN(3, false)
	lstore_0      = NewStoreN(0, true)
	lstore_1      = NewStoreN(1, true)
	lstore_2      = NewStoreN(2, true)
	lstore_3      = NewStoreN(3, true)
	fstore_0      = NewStoreN(0, false)
	fstore_1      = NewStoreN(1, false)
	fstore_2      = NewStoreN(2, false)
	fstore_3      = NewStoreN(3, false)
	dstore_0      = NewStoreN(0, true)
	dstore_1      = NewStoreN(1, true)
	dstore_2      = NewStoreN(2, true)
	dstore_3      = NewStoreN(3, true)
	astore_0      = NewStoreN(0, false)
	astore_1      = NewStoreN(1, false)
	astore_2      = NewStoreN(2, false)
	astore_3      = NewStoreN(3, false)
	iastore       = NewIAStore()
	lastore       = NewLAStore()
	fastore       = NewFAStore()
	dastore       = NewDAStore()
	aastore       = NewAAStore()
	bastore       = NewBAStore()
	castore       = NewCAStore()
	sastore       = NewSAStore()
	pop           = &Pop{}
	pop2          = &Pop2{}
	dup           = &Dup{}
	dup_x1        = &DupX1{}
	dup_x2        = &DupX2{}
	dup2          = &Dup2{}
	dup2_x1       = &Dup2X1{}
	dup2_x2       = &Dup2X2{}
	swap          = &Swap{}
	iadd          = NewIAdd()
	ladd          = NewLAdd()
	fadd          = NewFAdd()
	dadd          = NewDAdd()
	isub          = NewISub()
	lsub          = NewLSub()
	fsub          = NewFSub()
	dsub          = NewDSub()
	imul          = NewIMul()
	lmul          = NewLMul()
	fmul          = NewFMul()
	dmul          = NewDMul()
	idiv          = NewIDiv()
	ldiv          = NewLDiv()
	fdiv          = NewFDiv()
	ddiv          = NewDDiv()
	irem          = NewIRem()
	lrem          = NewLRem()
	frem          = NewFRem()
	drem          = NewDRem()
	ineg          = NewINeg()
	lneg          = NewLNeg()
	fneg          = NewFNeg()
	dneg          = NewDNeg()
	ishl          = NewIShl()
	lshl          = NewLShl()
	ishr          = NewIShr()
	lshr          = NewLShr()
	iushr         = NewIUShr()
	lushr         = NewLUShr()
	iand          = NewIAnd()
	land          = NewLAnd()
	ior           = NewIOr()
	lor           = NewLOr()
	ixor          = NewIXor()
	lxor          = NewLXor()
	i2l           = NewI2L()
	i2f           = NewI2F()
	i2d           = NewI2D()
	l2i           = NewL2I()
	l2f           = NewL2F()
	l2d           = NewL2D()
	f2i           = NewF2I()
	f2l           = NewF2L()
	f2d           = NewF2D()
	d2i           = NewD2I()
	d2l           = NewD2L()
	d2f           = NewD2F()
	i2b           = NewI2B()
	i2c           = NewI2C()
	i2s           = NewI2S()
	lcmp          = NewLCMP()
	fcmpl         = NewFCMPL()
	fcmpg         = NewFCMPG()
	dcmpl         = NewDCMPL()
	dcmpg         = NewDCMPG()
	ireturn       = NewXReturn(false)
	lreturn       = NewXReturn(true)
	freturn       = NewXReturn(false)
	dreturn       = NewXReturn(true)
	areturn       = NewXReturn(false)
	_return       = &Return{}
	arraylength   = &ArrayLength{}
	athrow        = &AThrow{}
	monitorenter  = &MonitorEnter{}
	monitorexit   = &MonitorExit{}
	invoke_native = &InvokeNative{}
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
		return &BIPush{}
	case OpSIPush:
		return &SIPush{}
	case OpLDC:
		return &LDC{}
	case OpLDCw:
		return &LDC_W{}
	case OpLDC2w:
		return &LDC2_W{}
	case OpILoad:
		return NewLoad(false)
	case OpLLoad:
		return NewLoad(true)
	case OpFLoad:
		return NewLoad(false)
	case OpDLoad:
		return NewLoad(true)
	case OpALoad:
		return NewLoad(false)
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
		return NewStore(false)
	case OpLStore:
		return NewStore(true)
	case OpFStore:
		return NewStore(false)
	case OpDStore:
		return NewStore(true)
	case OpAStore:
		return NewStore(false)
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
		return &IInc{}
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
		return NewIfEQ()
	case OpIfNE:
		return NewIfNE()
	case OpIfLT:
		return NewIfLT()
	case OpIfGE:
		return NewIfGE()
	case OpIfGT:
		return NewIfGT()
	case OpIfLE:
		return NewIfLE()
	case OpIfICmpEQ:
		return NewIfICmpEQ()
	case OpIfICmpNE:
		return NewIfICmpNE()
	case OpIfICmpLT:
		return NewIfICmpLT()
	case OpIfICmpGE:
		return NewIfICmpGE()
	case OpIfICmpGT:
		return NewIfICmpGT()
	case OpIfICmpLE:
		return NewIfICmpLE()
	case OpIfACmpEQ:
		return NewIfACmpEQ()
	case OpIfACmpNE:
		return NewIfACmpNE()
	case OpGoto:
		return &Goto{}
	case OpJSR:
		return &JSR{}
	case OpRET:
		return &RET{}
	case OpTableSwitch:
		return &TableSwitch{}
	case OpLookupSwitch:
		return &LookupSwitch{}
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
		return &GetStatic{}
	case OpPupStatic:
		return &PupStatic{}
	case OpGetField:
		return &GetField{}
	case OpPutField:
		return &PutField{}
	case OpInvokeVirtual:
		return &InvokeVirtual{}
	case OpInvokeSpecial:
		return &InvokeSpecial{}
	case OpInvokeStatic:
		return &InvokeStatic{}
	case OpInvokeInterface:
		return &InvokeInterface{}
	case OpInvokeDynamic:
		return &InvokeDynamic{}
	case OpNew:
		return &New{}
	case OpNewArray:
		return &NewArray{}
	case OpANewArray:
		return &ANewArray{}
	case OpArrayLength:
		return arraylength
	case OpAThrow:
		return athrow
	case OpCheckCast:
		return &CheckCast{}
	case OpInstanceOf:
		return &InstanceOf{}
	case OpMonitorEnter:
		return monitorenter
	case OpMonitorExit:
		return monitorexit
	case OpWide:
		return &Wide{}
	case OpMultiANewArray:
		return &MultiANewArray{}
	case OpIfNull:
		return NewIfNull()
	case OpIfNonNull:
		return NewIfNonNull()
	case OpGotoW:
		return &GotoW{}
	case OpJSRw:
		return &JSR_W{}
	//case 0xca: todo breakpoint
	case OpBreakpoint:
		return invoke_native // impdep1
	case OpInvokeNative:
		return &Bootstrap{} // impdep2
	default:
		panic(fmt.Errorf("invalid opcode: %v", opcode))
	}
}
