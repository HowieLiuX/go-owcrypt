package owcrypt

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func slide_cmp_equ(a []byte, b []byte, length uint16) bool {
	i := uint16(0)
	for i = 0; i < length; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_sm2_keyagreement(*testing.T) {
	////////////////////////////////////////////////////////////////////发起方 - initiator////////////////////////////////////////////////////////////////////
	//发起方标识符
	IDinitiator := [4]byte{0x11, 0x22, 0x33, 0x44}
	//发起方私钥
	prikeyInitiator := [32]byte{0xF5, 0xA6, 0xF5, 0x9C, 0x50, 0xCD, 0x3D, 0x57, 0xCF, 0xC7, 0xC8, 0xB2, 0xA9, 0x9C, 0x98, 0x3C, 0x3B, 0xC0, 0x9A, 0xB6, 0x6E, 0x86, 0xAB, 0x64, 0x35, 0xC5, 0x18, 0x5B, 0x70, 0x15, 0xEA, 0x37}
	//发起方共钥
	pubkeyInitiator := [64]byte{0xDE, 0x6E, 0xB0, 0xD4, 0x70, 0x42, 0xF9, 0x51, 0x61, 0xC6, 0xC7, 0x18, 0x75, 0xEE, 0x62, 0x7D, 0xE1, 0xDE, 0x49, 0xE4, 0x23, 0x61, 0xF5, 0x3B, 0x2A, 0x15, 0x13, 0x2D, 0xA9, 0x8A, 0xEB, 0x2E, 0xFE, 0xA1, 0xDF, 0xC8, 0x2C, 0xCE, 0xF3, 0x0A, 0xB1, 0xB0, 0xAD, 0x8F, 0xB3, 0x05, 0x90, 0x55, 0x25, 0xF8, 0x4A, 0x8C, 0xEC, 0x81, 0x85, 0xB8, 0x0A, 0x51, 0x63, 0x8B, 0x5A, 0x10, 0x05, 0xB7}
	//发起方临时私钥 tmpPrikeyInitiator := [32]byte{}
	//发起方临时公钥tmpPubkeyInitiator := [64]byte{}
	//发起方向响应方发送的验证值 SA := [32]byte{}
	//发起方协商的结果retA := [32]byte{}

	////////////////////////////////////////////////////////////////////响应方 - responder////////////////////////////////////////////////////////////////////
	//响应方标识符
	IDresponder := [4]byte{0x55, 0x66, 0x77, 0x88}
	//相应方私钥
	prikeyResponder := [32]byte{0xB3, 0x93, 0xF6, 0xDB, 0xAB, 0x4E, 0xB4, 0x7C, 0x89, 0x03, 0xB3, 0x3A, 0xAA, 0x6E, 0x36, 0x70, 0xE1, 0xAE, 0x1A, 0xFD, 0xE3, 0x7F, 0x44, 0x1B, 0x7C, 0x78, 0xF1, 0x9E, 0x68, 0xEA, 0x7A, 0x90}
	//响应方公钥
	pubkeyResponder := [64]byte{0x01, 0x99, 0xB9, 0x57, 0x9B, 0x44, 0x83, 0x95, 0x62, 0x91, 0x12, 0x36, 0xA1, 0x44, 0x8E, 0x1B, 0xF2, 0xFF, 0x7B, 0xC0, 0xAE, 0xD9, 0x77, 0xFD, 0x88, 0x67, 0x1B, 0x16, 0x21, 0x13, 0x59, 0x73, 0x4D, 0x3F, 0x9A, 0xC4, 0xC1, 0x11, 0x2B, 0x4B, 0xE8, 0x8B, 0x30, 0x93, 0x84, 0x9F, 0xB8, 0x3E, 0x8D, 0xAB, 0xD0, 0xCE, 0x6F, 0xA4, 0x5F, 0x90, 0x41, 0xC5, 0x38, 0x16, 0xB2, 0x6B, 0x14, 0xBB}
	//响应方临时公钥 tmpPubkeyResponder := [64]byte{}
	//响应方本地验证S 2 := [32]byte{}
	//响应方向发起方发送的验证值 SB := [32]byte{}
	//响应方协商的结果 retB := [32]byte{}

	///////////////////////////////////////////////////////////////////////////协商开始////////////////////////////////////////////////////////////////////
	//协商开始前，发起方掌握的信息有：  自身的私钥、公钥、响应方的公钥，以及提前约定好的曲线参数
	//          响应方掌握的信息有：  自身的私钥、公钥、发送方的公钥，以及提前约定好的曲线参数
	//第一步：
	//1.1 发起方在本地产生一组临时的公私钥对，然后发起协商
	fmt.Println("--------------------------发起方第一步--------------------------")
	tmpPrikeyInitiator, tmpPubkeyInitiator := KeyAgreement_initiator_step1(ECC_CURVE_SM2_STANDARD)

	fmt.Println("发起方产生临时公私钥对，产生结果为：")
	fmt.Println("发起方临时私钥：", hex.EncodeToString(tmpPrikeyInitiator[:]))
	fmt.Println("发起方临时公钥：", hex.EncodeToString(tmpPubkeyInitiator[:]))

	//1.2 发起方将临时私钥保存在本地，用于第二步操作的输入
	//1.3 发起方将临时公钥发送给响应方来发起协商，同时会指定协商的具体长度

	//第二步：
	//2.1 响应方获得发送方的临时公钥和协商长度，然后开始进行协商计算
	fmt.Println("--------------------------响应方第一步--------------------------")
	retB, tmpPubkeyResponder, S2, SB, ret := KeyAgreement_responder_step1(IDinitiator[:],
		4,
		IDresponder[:],
		4,
		prikeyResponder[:],
		pubkeyResponder[:],
		pubkeyInitiator[:],
		tmpPubkeyInitiator[:],
		32,
		ECC_CURVE_SM2_STANDARD)
	if ret != SUCCESS {
		fmt.Println("响应方协商第一步出错！")
		return

	} else {
		fmt.Println("响应方产生临时公钥 ：", hex.EncodeToString(tmpPubkeyResponder[:]))
		fmt.Println("响应方本地校验值： ", hex.EncodeToString(S2[:]))
		fmt.Println("响应方发送给发起方的校验值： ", hex.EncodeToString(SB[:]))
		fmt.Println("响应方获得的协商结果： ", hex.EncodeToString(retB[:]))
	}

	//2.2 响应方此时获得临时公钥、用于本地校验的S2、用于发送给发起方的校验值SB， 协商结果
	//2.3 响应方将S2和协商保存在本地，用于第二步的校验
	//2.4 响应方将临时公钥和校验值SB发送给发起方

	//第三步：
	//发起方获得响应方的临时公钥和校验值，开始进行协商计算
	fmt.Println("--------------------------发起方第二步--------------------------")
	retA, SA, err := KeyAgreement_initiator_step2(IDinitiator[:],
		4,
		IDresponder[:],
		4,
		prikeyInitiator[:],
		pubkeyInitiator[:],
		pubkeyResponder[:],
		tmpPrikeyInitiator[:],
		tmpPubkeyInitiator[:],
		tmpPubkeyResponder[:],
		SB[:],
		32,
		ECC_CURVE_SM2_STANDARD)
	if err != SUCCESS {
		fmt.Println("发起方协商第一步出错！")
		return
	} else {
		fmt.Println("发起方发送给响应方的校验值： ", hex.EncodeToString(SA[:]))
		fmt.Println("发起方获得的协商结果： ", hex.EncodeToString(retA[:]))
	}

	//此时，发起方已经获得协商结果，如果接口返回SUCCESS，则说明接口内部已经与响应方发来的校验值完成校验
	//即：发起方的协商流程已经完成
	//然后，发起方需要将输出的校验值SA发送给响应方进行校验

	//第四步：
	//响应方拿到发起方发来的最终校验值SA， 与之前本地保存的校验值S2进行比对，返回SUCCESS则响应方协商通过
	fmt.Println("--------------------------响应方第二步--------------------------")
	if SUCCESS != KeyAgreement_responder_step2(SA[:], S2[:], ECC_CURVE_SM2_STANDARD) {
		fmt.Println("响应方校验未通过")
		return
	} else {
		fmt.Println("响应方校验通过")
	}

	if slide_cmp_equ(retA[:], retB[:], 32) {
		fmt.Println("双方协商结果一致")
	} else {
		fmt.Println("双方协商结果不一致")
	}

}

func Test_sm2_keyagreement_ElGamal(*testing.T) {
	////////////////////////////////////////////////////////////////////发起方 - initiator////////////////////////////////////////////////////////////////////
	//发起方标识符
	IDinitiator := [4]byte{0x11, 0x22, 0x33, 0x44}
	//发起方私钥
	prikeyInitiator := [32]byte{0xF5, 0xA6, 0xF5, 0x9C, 0x50, 0xCD, 0x3D, 0x57, 0xCF, 0xC7, 0xC8, 0xB2, 0xA9, 0x9C, 0x98, 0x3C, 0x3B, 0xC0, 0x9A, 0xB6, 0x6E, 0x86, 0xAB, 0x64, 0x35, 0xC5, 0x18, 0x5B, 0x70, 0x15, 0xEA, 0x37}
	//发起方共钥
	pubkeyInitiator := [64]byte{0xDE, 0x6E, 0xB0, 0xD4, 0x70, 0x42, 0xF9, 0x51, 0x61, 0xC6, 0xC7, 0x18, 0x75, 0xEE, 0x62, 0x7D, 0xE1, 0xDE, 0x49, 0xE4, 0x23, 0x61, 0xF5, 0x3B, 0x2A, 0x15, 0x13, 0x2D, 0xA9, 0x8A, 0xEB, 0x2E, 0xFE, 0xA1, 0xDF, 0xC8, 0x2C, 0xCE, 0xF3, 0x0A, 0xB1, 0xB0, 0xAD, 0x8F, 0xB3, 0x05, 0x90, 0x55, 0x25, 0xF8, 0x4A, 0x8C, 0xEC, 0x81, 0x85, 0xB8, 0x0A, 0x51, 0x63, 0x8B, 0x5A, 0x10, 0x05, 0xB7}
	//发起方临时私钥 tmpPrikeyInitiator := [32]byte{}
	//发起方临时公钥tmpPubkeyInitiator := [64]byte{}
	//发起方向响应方发送的验证值 SA := [32]byte{}
	//发起方协商的结果retA := [32]byte{}

	////////////////////////////////////////////////////////////////////响应方 - responder////////////////////////////////////////////////////////////////////
	//响应方临时私钥固定
	fixedRandom := [32]byte{0x8e, 0x81, 0x24, 0x36, 0xa0, 0xe3, 0x32, 0x31, 0x66, 0xe1, 0xf0, 0xe8, 0xba, 0x79, 0xe1, 0x9e, 0x21, 0x7b, 0x2c, 0x4a, 0x53, 0xc9, 0x70, 0xd4, 0xcc, 0xa0, 0xcf, 0xb1, 0x07, 0x89, 0x79, 0xdf}
	//响应方标识符
	IDresponder := [4]byte{0x55, 0x66, 0x77, 0x88}
	//相应方私钥
	prikeyResponder := [32]byte{0xB3, 0x93, 0xF6, 0xDB, 0xAB, 0x4E, 0xB4, 0x7C, 0x89, 0x03, 0xB3, 0x3A, 0xAA, 0x6E, 0x36, 0x70, 0xE1, 0xAE, 0x1A, 0xFD, 0xE3, 0x7F, 0x44, 0x1B, 0x7C, 0x78, 0xF1, 0x9E, 0x68, 0xEA, 0x7A, 0x90}
	//响应方公钥
	pubkeyResponder := [64]byte{0x01, 0x99, 0xB9, 0x57, 0x9B, 0x44, 0x83, 0x95, 0x62, 0x91, 0x12, 0x36, 0xA1, 0x44, 0x8E, 0x1B, 0xF2, 0xFF, 0x7B, 0xC0, 0xAE, 0xD9, 0x77, 0xFD, 0x88, 0x67, 0x1B, 0x16, 0x21, 0x13, 0x59, 0x73, 0x4D, 0x3F, 0x9A, 0xC4, 0xC1, 0x11, 0x2B, 0x4B, 0xE8, 0x8B, 0x30, 0x93, 0x84, 0x9F, 0xB8, 0x3E, 0x8D, 0xAB, 0xD0, 0xCE, 0x6F, 0xA4, 0x5F, 0x90, 0x41, 0xC5, 0x38, 0x16, 0xB2, 0x6B, 0x14, 0xBB}
	//响应方临时公钥 tmpPubkeyResponder := [64]byte{}
	//响应方本地验证S 2 := [32]byte{}
	//响应方向发起方发送的验证值 SB := [32]byte{}
	//响应方协商的结果 retB := [32]byte{}

	///////////////////////////////////////////////////////////////////////////协商开始////////////////////////////////////////////////////////////////////
	//协商开始前，发起方掌握的信息有：  自身的私钥、公钥、响应方的公钥，以及提前约定好的曲线参数
	//          响应方掌握的信息有：  自身的私钥、公钥、发送方的公钥，以及提前约定好的曲线参数
	//第一步：
	//1.1 发起方在本地产生一组临时的公私钥对，然后发起协商
	fmt.Println("--------------------------发起方第一步--------------------------")
	tmpPrikeyInitiator, tmpPubkeyInitiator := KeyAgreement_initiator_step1(ECC_CURVE_SM2_STANDARD)

	fmt.Println("发起方产生临时公私钥对，产生结果为：")
	fmt.Println("发起方临时私钥：", hex.EncodeToString(tmpPrikeyInitiator[:]))
	fmt.Println("发起方临时公钥：", hex.EncodeToString(tmpPubkeyInitiator[:]))

	//1.2 发起方将临时私钥保存在本地，用于第二步操作的输入
	//1.3 发起方将临时公钥发送给响应方来发起协商，同时会指定协商的具体长度

	//第二步：
	//2.1 响应方获得发送方的临时公钥和协商长度，然后开始进行协商计算
	fmt.Println("--------------------------响应方第一步--------------------------")
	retB, tmpPubkeyResponder, S2, SB, ret := KeyAgreement_responder_ElGamal_step1(IDinitiator[:],
		4,
		IDresponder[:],
		4,
		prikeyResponder[:],
		pubkeyResponder[:],
		pubkeyInitiator[:],
		tmpPubkeyInitiator[:],
		32,
		fixedRandom[:],
		ECC_CURVE_SM2_STANDARD)
	if ret != SUCCESS {
		fmt.Println("响应方协商第一步出错！")
		return

	} else {
		fmt.Println("响应方产生临时公钥 ：", hex.EncodeToString(tmpPubkeyResponder[:]))
		fmt.Println("响应方本地校验值： ", hex.EncodeToString(S2[:]))
		fmt.Println("响应方发送给发起方的校验值： ", hex.EncodeToString(SB[:]))
		fmt.Println("响应方获得的协商结果： ", hex.EncodeToString(retB[:]))
	}

	//2.2 响应方此时获得临时公钥、用于本地校验的S2、用于发送给发起方的校验值SB， 协商结果
	//2.3 响应方将S2和协商保存在本地，用于第二步的校验
	//2.4 响应方将临时公钥和校验值SB发送给发起方

	//第三步：
	//发起方获得响应方的临时公钥和校验值，开始进行协商计算
	fmt.Println("--------------------------发起方第二步--------------------------")
	retA, SA, err := KeyAgreement_initiator_step2(IDinitiator[:],
		4,
		IDresponder[:],
		4,
		prikeyInitiator[:],
		pubkeyInitiator[:],
		pubkeyResponder[:],
		tmpPrikeyInitiator[:],
		tmpPubkeyInitiator[:],
		tmpPubkeyResponder[:],
		SB[:],
		32,
		ECC_CURVE_SM2_STANDARD)
	if err != SUCCESS {
		fmt.Println("发起方协商第一步出错！")
		return
	} else {
		fmt.Println("发起方发送给响应方的校验值： ", hex.EncodeToString(SA[:]))
		fmt.Println("发起方获得的协商结果： ", hex.EncodeToString(retA[:]))
	}

	//此时，发起方已经获得协商结果，如果接口返回SUCCESS，则说明接口内部已经与响应方发来的校验值完成校验
	//即：发起方的协商流程已经完成
	//然后，发起方需要将输出的校验值SA发送给响应方进行校验

	//第四步：
	//响应方拿到发起方发来的最终校验值SA， 与之前本地保存的校验值S2进行比对，返回SUCCESS则响应方协商通过
	fmt.Println("--------------------------响应方第二步--------------------------")
	if SUCCESS != KeyAgreement_responder_step2(SA[:], S2[:], ECC_CURVE_SM2_STANDARD) {
		fmt.Println("响应方校验未通过")
		return
	} else {
		fmt.Println("响应方校验通过")
	}

	if slide_cmp_equ(retA[:], retB[:], 32) {
		fmt.Println("双方协商结果一致")
	} else {
		fmt.Println("双方协商结果不一致")
	}

}

func Test_sm2_genpubkey(t *testing.T) {
	prikey_0 := [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	prikey_illegal := [32]byte{0xFF, 0xFF, 0xFF, 0xFf, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xff, 0xff, 0xff, 0xff, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFe}
	prikey := [32]byte{0xF6, 0xAA, 0xD7, 0x09, 0x86, 0x01, 0x19, 0x2E, 0x0E, 0x1C, 0xB3, 0x6B, 0x5F, 0x12, 0xF7, 0x3B, 0xBA, 0x39, 0xD7, 0x59, 0x6A, 0x10, 0x0F, 0x41, 0x84, 0xC0, 0xFE, 0x43, 0xFC, 0x3C, 0x11, 0x0A}
	//pubkey_ecdsasecp256r1 := [64]byte{0xB2,0xC0,0xC3,0x07,0x28,0xAB,0x3C,0x10,0xCC,0x28,0x9B,0x78,0x0D,0x79,0xFB,0x92,0xF0,0x62,0xB7,0x2E,0x7D,0x3A,0x69,0xEC,0x0C,0x88,0xF1,0xE1,0x49,0x1D,0xDF,0xC1,0x24,0x8B,0x13,0xB2,0x6C,0xCD,0x60,0x2A,0xDA,0x23,0x29,0x69,0x65,0x70,0x94,0x01,0x33,0x53,0xFE,0xB7,0x94,0xDA,0x04,0xDD,0xF3,0xF7,0xF4,0x5C,0x26,0x70,0xFD,0x7B}
	pubkey_sm2 := [64]byte{0x6C, 0x1E, 0xFB, 0x83, 0xBD, 0xEC, 0xBC, 0x47, 0x99, 0xCE, 0x03, 0xAB, 0xE2, 0x64, 0x3C, 0x18, 0x4B, 0x75, 0x0A, 0x28, 0xA9, 0x33, 0x7A, 0x22, 0x1C, 0x58, 0x1B, 0x7B, 0x11, 0x61, 0xDA, 0x01, 0xB0, 0x5E, 0x98, 0x96, 0x58, 0x61, 0xC8, 0x78, 0x16, 0xFB, 0x5D, 0x57, 0xFA, 0xD6, 0xB6, 0x30, 0x6F, 0x98, 0x2A, 0x36, 0x97, 0xA0, 0x11, 0x80, 0x7A, 0x5C, 0x4C, 0xBB, 0xD4, 0xEF, 0x8A, 0xC2}
	//pubkey := [64]byte{}

	//私钥全0的情况
	pubkey, err := GenPubkey(prikey_0[:], ECC_CURVE_SM2_STANDARD)
	if err != ECC_PRIKEY_ILLEGAL {
		fmt.Println("sm2产生公钥接口未对全零私钥正确检查")
		return
	} else {
		fmt.Println("sm2产生公钥接口全零私钥检查通过")
	}
	pubkey, err = GenPubkey(prikey_illegal[:], ECC_CURVE_SM2_STANDARD)
	if err != ECC_PRIKEY_ILLEGAL {
		fmt.Println("sm2产生公钥接口未对非法私钥正确检查")
		return
	} else {
		fmt.Println("sm2产生公钥接口非法私钥检查通过")
	}
	pubkey, err = GenPubkey(prikey[:], ECC_CURVE_SM2_STANDARD)
	if err != SUCCESS {
		fmt.Println("sm2产生公钥接口返回值错误")
		return
	} else {
		fmt.Println("sm2产生公钥接口返回值正确")
	}
	if slide_cmp_equ(pubkey_sm2[:], pubkey[:], 64) != true {
		fmt.Println("sm2产生了错误的公钥")
		return
	} else {
		fmt.Println("sm2 genPubkey pass")
	}
}

func Test_sm2_encdec(t *testing.T) {
	prikey := [32]byte{0xF6, 0xAA, 0xD7, 0x09, 0x86, 0x01, 0x19, 0x2E, 0x0E, 0x1C, 0xB3, 0x6B, 0x5F, 0x12, 0xF7, 0x3B, 0xBA, 0x39, 0xD7, 0x59, 0x6A, 0x10, 0x0F, 0x41, 0x84, 0xC0, 0xFE, 0x43, 0xFC, 0x3C, 0x11, 0x0A}
	pubkey := [64]byte{0x6C, 0x1E, 0xFB, 0x83, 0xBD, 0xEC, 0xBC, 0x47, 0x99, 0xCE, 0x03, 0xAB, 0xE2, 0x64, 0x3C, 0x18, 0x4B, 0x75, 0x0A, 0x28, 0xA9, 0x33, 0x7A, 0x22, 0x1C, 0x58, 0x1B, 0x7B, 0x11, 0x61, 0xDA, 0x01, 0xB0, 0x5E, 0x98, 0x96, 0x58, 0x61, 0xC8, 0x78, 0x16, 0xFB, 0x5D, 0x57, 0xFA, 0xD6, 0xB6, 0x30, 0x6F, 0x98, 0x2A, 0x36, 0x97, 0xA0, 0x11, 0x80, 0x7A, 0x5C, 0x4C, 0xBB, 0xD4, 0xEF, 0x8A, 0xC2}

	plain := [32]byte{0xF6, 0xAA, 0xD7, 0x09, 0x86, 0x01, 0x19, 0x2E, 0x0E, 0x1C, 0xB3, 0x6B, 0x5F, 0x12, 0xF7, 0x3B, 0xBA, 0x39, 0xD7, 0x59, 0x6A, 0x10, 0x0F, 0x41, 0x84, 0xC0, 0xFE, 0x43, 0xFC, 0x3C, 0x11, 0x0A}

	i := uint16(0)

	for i = 1; i < 32; i++ {
		cipher, ret := Encryption(pubkey[:], plain[:i], ECC_CURVE_SM2_STANDARD)

		if ret != SUCCESS || uint16(len(cipher)) != i+97 {
			fmt.Println("sm2加密错误")
			return
		} else {
			fmt.Println("sm2 enc pass")
			fmt.Println("sm2 cipher value : ", hex.EncodeToString(cipher[:i+97]))
		}

		check, ret := Decryption(prikey[:], cipher[:], ECC_CURVE_SM2_STANDARD)

		if ret != SUCCESS || uint16(len(check)) != i {
			fmt.Println("sm2解密错误")
			return
		}

		if slide_cmp_equ(plain[:], check[:], i) != true {
			fmt.Println("sm2解密失败")
			return
		} else {
			fmt.Println("sm2 dec pass")
			fmt.Println("sm2 source plain value    : ", hex.EncodeToString(plain[:i]))
			fmt.Println("sm2 decrypted plain value : ", hex.EncodeToString(check[:i]))
		}
	}
}

func Test_sm2_signverify(t *testing.T) {
	prikey := [32]byte{0xF6, 0xAA, 0xD7, 0x09, 0x86, 0x01, 0x19, 0x2E, 0x0E, 0x1C, 0xB3, 0x6B, 0x5F, 0x12, 0xF7, 0x3B, 0xBA, 0x39, 0xD7, 0x59, 0x6A, 0x10, 0x0F, 0x41, 0x84, 0xC0, 0xFE, 0x43, 0xFC, 0x3C, 0x11, 0x0A}
	pubkey := [64]byte{0x6C, 0x1E, 0xFB, 0x83, 0xBD, 0xEC, 0xBC, 0x47, 0x99, 0xCE, 0x03, 0xAB, 0xE2, 0x64, 0x3C, 0x18, 0x4B, 0x75, 0x0A, 0x28, 0xA9, 0x33, 0x7A, 0x22, 0x1C, 0x58, 0x1B, 0x7B, 0x11, 0x61, 0xDA, 0x01, 0xB0, 0x5E, 0x98, 0x96, 0x58, 0x61, 0xC8, 0x78, 0x16, 0xFB, 0x5D, 0x57, 0xFA, 0xD6, 0xB6, 0x30, 0x6F, 0x98, 0x2A, 0x36, 0x97, 0xA0, 0x11, 0x80, 0x7A, 0x5C, 0x4C, 0xBB, 0xD4, 0xEF, 0x8A, 0xC2}

	message := [32]byte{0xF6, 0xAA, 0xD7, 0x09, 0x86, 0x01, 0x19, 0x2E, 0x0E, 0x1C, 0xB3, 0x6B, 0x5F, 0x12, 0xF7, 0x3B, 0xBA, 0x39, 0xD7, 0x59, 0x6A, 0x10, 0x0F, 0x41, 0x84, 0xC0, 0xFE, 0x43, 0xFC, 0x3C, 0x11, 0x0A}

	ID := [4]byte{1, 2, 3, 4}

	i := uint16(0)

	for i = 1; i < 32; i++ {
		signature, err := Signature(prikey[:], ID[:], 4, message[:], i, ECC_CURVE_SM2_STANDARD)
		if err != SUCCESS {
			fmt.Println("sm2签名错误")
			return
		} else {
			fmt.Println("sm2 sign pass")
			fmt.Println("sm2 sign value : ", hex.EncodeToString(signature[:]))
		}
		if Verify(pubkey[:], ID[:], 4, message[:], i, signature[:], ECC_CURVE_SM2_STANDARD) != 0x0001 {
			fmt.Println("sm2验签错误")
			return
		} else {
			fmt.Println("sm2 verify pass")
		}
	}
}

func Test_getcurveorder(t *testing.T) {

	ret := GetCurveOrder(ECC_CURVE_SECP256K1)
	sret := hex.EncodeToString(ret[:])
	if sret == "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141" {
		fmt.Println("曲线secp256k1的阶是：")
		fmt.Println(sret)
	} else {
		fmt.Println("secp256k1获取失败！")
		return
	}

	ret = GetCurveOrder(ECC_CURVE_SECP256R1)
	sret = hex.EncodeToString(ret[:])
	if sret == "ffffffff00000000ffffffffffffffffbce6faada7179e84f3b9cac2fc632551" {
		fmt.Println("曲线secp256r1的阶是：")
		fmt.Println(sret)
	} else {
		fmt.Println("secp256r1获取失败！")
		return
	}

	ret = GetCurveOrder(ECC_CURVE_SM2_STANDARD)
	sret = hex.EncodeToString(ret[:])
	if sret == "fffffffeffffffffffffffffffffffff7203df6b21c6052b53bbf40939d54123" {
		fmt.Println("曲线sm2_std的阶是：")
		fmt.Println(sret)
	} else {
		fmt.Println("sm2_std获取失败！")
		return
	}

	ret = GetCurveOrder(ECC_CURVE_ED25519)
	sret = hex.EncodeToString(ret[:])
	if sret == "1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed" {
		fmt.Println("曲线ed25519的阶是：")
		fmt.Println(hex.EncodeToString(ret[:]))
	} else {
		fmt.Println("ed25519获取失败！")
		return
	}

	ret = GetCurveOrder(ECC_CURVE_ED25519_NORMAL)
	sret = hex.EncodeToString(ret[:])
	if sret == "1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed" {
		fmt.Println("曲线ed25519的阶是：")
		fmt.Println(hex.EncodeToString(ret[:]))
	} else {
		fmt.Println("ed25519获取失败！")
		return
	}
}

func Test_pointcompress(t *testing.T) {
	tmp1 := [65]byte{0x04, 0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73, 0x31, 0x85, 0x38, 0x77, 0x24, 0x08, 0x9D, 0x0E, 0x84, 0xC1, 0x7D, 0x54, 0x70, 0x5B, 0xD1, 0x23, 0xEB, 0x3A, 0x82, 0x54, 0xDA, 0x96, 0x43, 0x9E, 0xF6, 0x4B, 0x45, 0x07, 0x41, 0x2A, 0x94, 0x28}
	tmp2 := PointCompress(tmp1[:], ECC_CURVE_SECP256R1)
	fmt.Println(hex.EncodeToString(tmp2[:]))

	tmp3 := [64]byte{0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73, 0x31, 0x85, 0x38, 0x77, 0x24, 0x08, 0x9D, 0x0E, 0x84, 0xC1, 0x7D, 0x54, 0x70, 0x5B, 0xD1, 0x23, 0xEB, 0x3A, 0x82, 0x54, 0xDA, 0x96, 0x43, 0x9E, 0xF6, 0x4B, 0x45, 0x07, 0x41, 0x2A, 0x94, 0x28}
	tmp4 := PointCompress(tmp3[:], ECC_CURVE_SECP256R1)
	fmt.Println(hex.EncodeToString(tmp4[:]))

	// tmp5 := [65]byte{0x04, 0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73, 0x31, 0x85, 0x38, 0x77, 0x24, 0x08, 0x9D, 0x0E, 0x84, 0xC1, 0x7D, 0x54, 0x70, 0x5B, 0xD1, 0x23, 0xEB, 0x3A, 0x82, 0x54, 0xDA, 0x96, 0x43, 0x9E, 0xF6, 0x4B, 0x45, 0x07, 0x41, 0x2A, 0x94, 0x28}
	// tmp6 := PointCompress(tmp5[:], ECC_CURVE_SECP256R1)
	// fmt.Println(hex.EncodeToString(tmp6[:]))

	// tmp7 := [64]byte{0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73, 0x31, 0x85, 0x38, 0x77, 0x24, 0x08, 0x9D, 0x0E, 0x84, 0xC1, 0x7D, 0x54, 0x70, 0x5B, 0xD1, 0x23, 0xEB, 0x3A, 0x82, 0x54, 0xDA, 0x96, 0x43, 0x9E, 0xF6, 0x4B, 0x45, 0x07, 0x41, 0x2A, 0x94, 0x28}
	// tmp8 := PointCompress(tmp7[:], ECC_CURVE_SECP256R1)
	// fmt.Println(hex.EncodeToString(tmp8[:]))
}

func Test_pointdecompress(t *testing.T) {
	tmp1 := [33]byte{0x02, 0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73}
	tmp2 := PointDecompress(tmp1[:], ECC_CURVE_SECP256R1)
	fmt.Println(hex.EncodeToString(tmp2[:]))
}
func Test_hashset(t *testing.T) {
	//sha1 test......
	sha1_msg := []byte{0xBE, 0xA5, 0xA4, 0xDB, 0x64, 0x29, 0x0A, 0xC5, 0x62, 0xAA, 0x63, 0x3A, 0xB0, 0x04, 0x73, 0xEF, 0xC0, 0xA7, 0x3C, 0x56, 0x28, 0x2B, 0xC9, 0xFC, 0x41, 0x2A, 0x6E, 0xCC, 0xE7, 0xE3, 0x84, 0x22, 0x79, 0x04, 0xBF, 0x2F, 0x10, 0x22, 0xF2, 0xF4, 0xA0, 0x46, 0x20, 0xF8, 0xE6, 0xC3, 0x20, 0xD0, 0x74, 0xBD, 0xAE, 0x19, 0x1B, 0x76, 0x8A, 0x9B, 0x10, 0xFD, 0x5F, 0xCC, 0xEB, 0xFC, 0xF5, 0xB2, 0xD7}
	sha1_digest := Hash(sha1_msg[:], 0, HASH_ALG_SHA1)
	//standar result: BC22A6EEF346E688E06E44EA97AF765C0056C7A6
	fmt.Println("sha1 test result:", hex.EncodeToString(sha1_digest[:]))
	//sha256 test......
	sha256_msg := []byte{0xBE, 0xA5, 0xA4, 0xDB, 0x64, 0x29, 0x0A, 0xC5, 0x62, 0xAA, 0x63, 0x3A, 0xB0, 0x04, 0x73, 0xEF, 0xC0, 0xA7, 0x3C, 0x56, 0x28, 0x2B, 0xC9, 0xFC, 0x41, 0x2A, 0x6E, 0xCC, 0xE7, 0xE3, 0x84, 0x22, 0x79, 0x04, 0xBF, 0x2F, 0x10, 0x22, 0xF2, 0xF4, 0xA0, 0x46, 0x20, 0xF8, 0xE6, 0xC3, 0x20, 0xD0, 0x74, 0xBD, 0xAE, 0x19, 0x1B, 0x76, 0x8A, 0x9B, 0x10, 0xFD, 0x5F, 0xCC, 0xEB, 0xFC, 0xF5, 0xB2, 0xD7}
	sha256_digest := Hash(sha256_msg[:], 0, HASH_ALG_SHA256)
	//standar result: FCBEF90EDA56AE9E60AD8CEA860ACC7EBCA3C83EC418EFC7BE48037D95DEF18E
	fmt.Println("sha256 test result:", hex.EncodeToString(sha256_digest[:]))
	//sha512 test......
	sha512_msg := []byte{0xBE, 0xA5, 0xA4, 0xDB, 0x64, 0x29, 0x0A, 0xC5, 0x62, 0xAA, 0x63, 0x3A, 0xB0, 0x04, 0x73, 0xEF, 0xC0, 0xA7, 0x3C, 0x56, 0x28, 0x2B, 0xC9, 0xFC, 0x41, 0x2A, 0x6E, 0xCC, 0xE7, 0xE3, 0x84, 0x22, 0x79, 0x04, 0xBF, 0x2F, 0x10, 0x22, 0xF2, 0xF4, 0xA0, 0x46, 0x20, 0xF8, 0xE6, 0xC3, 0x20, 0xD0, 0x74, 0xBD, 0xAE, 0x19, 0x1B, 0x76, 0x8A, 0x9B, 0x10, 0xFD, 0x5F, 0xCC, 0xEB, 0xFC, 0xF5, 0xB2, 0xD7}
	sha512_digest := Hash(sha512_msg[:], 0, HASH_ALG_SHA512)
	//standar result: BE293563BEFD2070468CFA63CB86E4745E77B0287CA845D946C8069762BBBEC4A14C5052A5D0038B2CF99E3E88656DFEE56747EC1A7A3F21CFA7DD1637DBD607
	fmt.Println("sha512 test result:", hex.EncodeToString(sha512_digest[:]))
	//sm3 test......
	sm3_msg := []byte{0xBE, 0xA5, 0xA4, 0xDB, 0x64, 0x29, 0x0A, 0xC5, 0x62, 0xAA, 0x63, 0x3A, 0xB0, 0x04, 0x73, 0xEF, 0xC0, 0xA7, 0x3C, 0x56, 0x28, 0x2B, 0xC9, 0xFC, 0x41, 0x2A, 0x6E, 0xCC, 0xE7, 0xE3, 0x84, 0x22, 0x79, 0x04, 0xBF, 0x2F, 0x10, 0x22, 0xF2, 0xF4, 0xA0, 0x46, 0x20, 0xF8, 0xE6, 0xC3, 0x20, 0xD0, 0x74, 0xBD, 0xAE, 0x19, 0x1B, 0x76, 0x8A, 0x9B, 0x10, 0xFD, 0x5F, 0xCC, 0xEB, 0xFC, 0xF5, 0xB2, 0xD7}
	sm3_digest := Hash(sm3_msg[:], 0, HASH_ALG_SM3)
	//standar result: D187AD8F2D5C93FF46D1B13A9FAFC7F94D9CCC9D12C7CDF7864C8788FE180D14
	fmt.Println("sm3 test result:", hex.EncodeToString(sm3_digest[:]))
	//md4 test......
	md4_msg := []byte{0xBE, 0xA5, 0xA4, 0xDB, 0x64, 0x29, 0x0A, 0xC5, 0x62, 0xAA, 0x63, 0x3A, 0xB0, 0x04, 0x73, 0xEF, 0xC0, 0xA7, 0x3C, 0x56, 0x28, 0x2B, 0xC9, 0xFC, 0x41, 0x2A, 0x6E, 0xCC, 0xE7, 0xE3, 0x84, 0x22, 0x79, 0x04, 0xBF, 0x2F, 0x10, 0x22, 0xF2, 0xF4, 0xA0, 0x46, 0x20, 0xF8, 0xE6, 0xC3, 0x20, 0xD0, 0x74, 0xBD, 0xAE, 0x19, 0x1B, 0x76, 0x8A, 0x9B, 0x10, 0xFD, 0x5F, 0xCC, 0xEB, 0xFC, 0xF5, 0xB2, 0xD7}
	md4_digest := Hash(md4_msg[:], 0, HASH_ALG_MD4)
	//standar result: 77B1C8B6FE20866C41B94ED92886478B
	fmt.Println("md4 test result:", hex.EncodeToString(md4_digest[:]))
	//md5 test......
	md5_msg := []byte{0xBE, 0xA5, 0xA4, 0xDB, 0x64, 0x29, 0x0A, 0xC5, 0x62, 0xAA, 0x63, 0x3A, 0xB0, 0x04, 0x73, 0xEF, 0xC0, 0xA7, 0x3C, 0x56, 0x28, 0x2B, 0xC9, 0xFC, 0x41, 0x2A, 0x6E, 0xCC, 0xE7, 0xE3, 0x84, 0x22, 0x79, 0x04, 0xBF, 0x2F, 0x10, 0x22, 0xF2, 0xF4, 0xA0, 0x46, 0x20, 0xF8, 0xE6, 0xC3, 0x20, 0xD0, 0x74, 0xBD, 0xAE, 0x19, 0x1B, 0x76, 0x8A, 0x9B, 0x10, 0xFD, 0x5F, 0xCC, 0xEB, 0xFC, 0xF5, 0xB2, 0xD7}
	md5_digest := Hash(md5_msg[:], 0, HASH_ALG_MD5)
	//standar result: CF9CDB434DC7E3A275648EC554915349
	fmt.Println("md5 test result:", hex.EncodeToString(md5_digest[:]))
	//blake256 test......
	blake256_msg := []byte{0x84, 0x5D, 0xD3, 0x38, 0x30, 0xCF, 0x38, 0xAA, 0x74, 0x9E, 0x0D, 0x10, 0x5F, 0x1B, 0xDE, 0x0B, 0x6B, 0x3C, 0x74, 0x44, 0x44, 0x1F, 0xBB, 0x6D, 0x9A, 0x45, 0xFB, 0x06, 0x61, 0x5B, 0xF0, 0x4D}
	blake256_digest := Hash(blake256_msg[:], 0, HASH_ALG_BLAKE256)
	//standar result: 20c209644bedf5433858795d6671cae3d9af25fab89ae52f1cfe78797daf0547
	fmt.Println("blake256 test result:", hex.EncodeToString(blake256_digest[:]))

}

func Test_hmacset(t *testing.T) {
	msg := []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66}
	key := []byte{0x66, 0x65, 0x64, 0x63, 0x62, 0x61, 0x30, 0x39, 0x38, 0x37, 0x36, 0x35, 0x34, 0x33, 0x32, 0x31}
	result_sha256 := []byte{0x34, 0x70, 0x8f, 0xc5, 0xb5, 0x57, 0x4a, 0xb1, 0x6b, 0x89, 0xc6, 0xb3, 0x0b, 0x97, 0x09, 0xd0, 0x28, 0x17, 0x55, 0x23, 0xd6, 0x59, 0xd6, 0x40, 0x3c, 0x9d, 0x6b, 0x64, 0x4a, 0x57, 0xbf, 0x3c}
	result_sha512 := []byte{0x1f, 0x6e, 0x50, 0xcd, 0xbb, 0xe9, 0xdc, 0xc9, 0xd4, 0xa0, 0x1f, 0x2c, 0x2c, 0xee, 0x22, 0xdf, 0x3d, 0x0f, 0xb9, 0xb5, 0xd8, 0x85, 0x30, 0x67, 0x53, 0x6a, 0x5a, 0x50, 0x89, 0xce, 0x69, 0x45, 0x9f, 0xc9, 0x99, 0x9e, 0x7b, 0xbb, 0x70, 0x43, 0x41, 0x23, 0xe0, 0xde, 0xe1, 0x3d, 0x7a, 0xa3, 0x1a, 0x25, 0x13, 0x7d, 0x7d, 0xa4, 0x9f, 0x17, 0x39, 0xbd, 0x7e, 0xfe, 0x9c, 0x58, 0xd3, 0xd0}
	ret1 := Hmac(key[:], msg[:], HMAC_SHA256_ALG)
	if reflect.DeepEqual(ret1, result_sha256) {
		fmt.Println("hmac-sha256 test success")
	} else {
		fmt.Println("hmac-sha256 test fail")
	}
	ret2 := Hmac(key[:], msg[:], HMAC_SHA512_ALG)
	if reflect.DeepEqual(ret2, result_sha512) {
		fmt.Println("hmac-sha512 test success")
	} else {
		fmt.Println("hmac-sha512 test fail")
	}
}

func Test_pbkdf2(t *testing.T) {
	pwstr := "passDATAb00AB7YxDTTlRH2dqxDx19GDxDV1zFMz7E6QVqKIzwOtMnlxQLttpE57Un4u12D2YD7oOPpiEvCDYvntXEe4NNPLCnGGeJArbYDEu6xDoCfWH6kbuV6awi04U"
	saltstr := "saltKEYbcTcXHCBxtjD2PnBh44AIQ6XUOCESOhXpEp3HrcGMwbjzQKMSaf63IJemkURWoqHusIeVB8Il91NjiCGQacPUu9qTFaShLbKG0Yj4RCMV56WPj7E14EMpbxy6P"
	var iterations uint32
	var outlen uint32
	iterations = 100000
	outlen = 64
	pw := []byte(pwstr)
	salt := []byte(saltstr)
	ret := pbkdf2_hmac_sha512(pw, salt, iterations, outlen)
	fmt.Println(hex.EncodeToString(ret[:]))
}
func Test_EthSign(t *testing.T) {
	/**---test 1-----**/
	//standard result:9e3813f1beb98607ed0cfa9199a41000ee12ac57f551d46ed944705a2cfad52e713d2ba16e48e58f4c427df185bd73b7142afed19c37752978acf7417aa517af01
	//hash:=[]byte{0xce,0x99,0x90,0x1c,0x86,0xcb,0x50,0x0b,0xb1,0xd1,0x35,0xb6,0xab,0x6b,0xf1,0xd9,0x18,0xde,0x2a,0x7b,0xc5,0x54,0xc9,0xad,0xc2,0x5b,0x5a,0x09,0xd2,0x5c,0xf3,0x03}
	//prikey:=[]byte{0xc1,0xa1,0x19,0xf1,0x63,0xde,0xf0,0x72,0xf2,0x1a,0x1b,0x3b,0x96,0x05,0x2c,0x65,0x9d,0x71,0x77,0x3f,0x20,0x75,0xec,0x00,0x5d,0x4a,0xd8,0x49,0x24,0x71,0x7e,0x6a}
	//pubkey:=[]byte{0xD7,0xE4,0xC0,0x1B,0x63,0xAA,0xBE,0x17,0x75,0x89,0xFA,0x90,0x05,0xD6,0xC7,0xB1,0x18,0x22,0x83,0xC8,0x04,0xD8,0x43,0xF5,0xF4,0xC4,0xD6,0x16,0x17,0xBC,0x9F,0x23,0xFD,0x27,0xEB,0xDD,0x1E,0x67,0xFF,0x6C,0x93,0xA2,0x56,0x11,0xB5,0xC4,0xC3,0xA3,0xDD,0x7C,0x87,0xBD,0x6E,0x3E,0x63,0x62,0x71,0x7F,0x5E,0x67,0x0D,0xE3,0x66,0x32}
	/**----test 2-----**/
	//standard result:4d98f9b5ac76d314ba249a37d64de347a7f406132c8f7624b69c74b5badf9743668c89edfc1743cd8bd58fe935383dd8d4b72b25ac21112f66d45dbebcc7b1af00
	//hash:=[]byte{0xA4,0x4C,0x69,0x32,0x00,0xC3,0x7B,0x00,0x32,0x68,0x76,0x27,0x17,0x6E,0x41,0xDF,0xAC,0xC9,0x53,0xCC,0x77,0xEB,0x97,0x63,0x81,0xCD,0xB7,0xA6,0x6B,0x17,0x21,0x58}
	//prikey:=[]byte{0xBC,0xB9,0x71,0xDD,0x9A,0x73,0x1B,0x66,0xA4,0x25,0x51,0x7F,0x1F,0x02,0xC8,0xC3,0xAF,0x46,0xAF,0x74,0xFF,0x2F,0x62,0xF4,0xEF,0x21,0x14,0x70,0x41,0xC6,0xBB,0xA5}
	//pubkey:=[]byte{0x3D,0x91,0xE1,0xF9,0xC2,0x3E,0xAA,0x38,0x09,0x7C,0x87,0xAC,0xC0,0x6F,0x02,0xC9,0x57,0xDD,0x98,0x8F,0x0A,0x24,0x76,0x36,0xCD,0xDD,0x0F,0x91,0x43,0xA4,0xA9,0x5D,0x6A,0x08,0x19,0x58,0x6E,0xE3,0xF3,0xC7,0x31,0xA2,0x76,0xEF,0x74,0x2B,0xEF,0xB1,0xAE,0x61,0x5B,0xBF,0x48,0xCE,0x7D,0xD2,0xA6,0xE8,0x91,0x67,0x63,0x2F,0xE9,0x73}
	/**----test 3-----**/
	//standard result: a8064a1b1eab7f28bd0f26cdbdf2315e280b17eacab834bc27ab86e40307a9822e2b6bc2901fa439ce408dd13ff7ee930af51e47fc362bb8e44977e7009d1b5f00
	hash := []byte{0xA4, 0x4C, 0x69, 0x32, 0x00, 0xC3, 0x7B, 0x00, 0x32, 0x68, 0x76, 0x27, 0x17, 0x6E, 0x41, 0xDF, 0xAC, 0xC9, 0x53, 0xCC, 0x77, 0xEB, 0x97, 0x63, 0x81, 0xCD, 0xB7, 0xA6, 0x6B, 0x17, 0x21, 0x58}
	prikey := []byte{0xA8, 0xDE, 0xCB, 0xDF, 0x2A, 0x5C, 0x92, 0xF8, 0xD8, 0xFC, 0x4D, 0x53, 0x36, 0x7F, 0x3A, 0x21, 0x55, 0x84, 0xB0, 0xDD, 0xA9, 0x2E, 0xFC, 0x30, 0xBE, 0x89, 0x51, 0x44, 0xD3, 0xD5, 0x6F, 0x97}
	//pubkey :=[]byte{0x0B,0xF0,0xAE,0xD1,0x07,0x11,0xCC,0xE9,0xC0,0x7D,0x6F,0xFB,0xB4,0xCD,0x9D,0x93,0xA0,0x0B,0xF5,0x3A,0x97,0x22,0x08,0x1E,0x5A,0x1A,0x6C,0xB5,0x94,0xB0,0xF0,0x4E,0xAF,0x97,0x8B,0x8F,0x7B,0x7F,0xCA,0xFE,0xEF,0x85,0xA3,0x6F,0xBA,0xF6,0x6C,0x6F,0xA0,0xEA,0xC0,0x5D,0x46,0x8E,0x83,0x41,0x80,0xDE,0x34,0xCB,0x74,0xDD,0x45,0xCA}
	sig, err := ETHsignature(prikey, hash)
	if err != SUCCESS {
		t.Error("ETH sign fail")
	} else {
		fmt.Println("eth sign result:", hex.EncodeToString(sig))
	}
}

func Test_RecoverPubkey(t *testing.T) {
	/**----------test 1-----**/
	//standard result:0xD7,0xE4,0xC0,0x1B,0x63,0xAA,0xBE,0x17,0x75,0x89,0xFA,0x90,0x05,0xD6,0xC7,0xB1,0x18,0x22,0x83,0xC8,0x04,0xD8,0x43,0xF5,0xF4,0xC4,0xD6,0x16,0x17,0xBC,0x9F,0x23,0xFD,0x27,0xEB,0xDD,0x1E,0x67,0xFF,0x6C,0x93,0xA2,0x56,0x11,0xB5,0xC4,0xC3,0xA3,0xDD,0x7C,0x87,0xBD,0x6E,0x3E,0x63,0x62,0x71,0x7F,0x5E,0x67,0x0D,0xE3,0x66,0x32
	//sig := []byte{0x9e,0x38,0x13,0xf1,0xbe,0xb9,0x86,0x07,0xed,0x0c,0xfa,0x91,0x99,0xa4,0x10,0x00,0xee,0x12,0xac,0x57,0xf5,0x51,0xd4,0x6e,0xd9,0x44,0x70,0x5a,0x2c,0xfa,0xd5,0x2e,0x71,0x3d,0x2b,0xa1,0x6e,0x48,0xe5,0x8f,0x4c,0x42,0x7d,0xf1,0x85,0xbd,0x73,0xb7,0x14,0x2a,0xfe,0xd1,0x9c,0x37,0x75,0x29,0x78,0xac,0xf7,0x41,0x7a,0xa5,0x17,0xaf,0x01}
	//msg :=[]byte{0xce,0x99,0x90,0x1c,0x86,0xcb,0x50,0x0b,0xb1,0xd1,0x35,0xb6,0xab,0x6b,0xf1,0xd9,0x18,0xde,0x2a,0x7b,0xc5,0x54,0xc9,0xad,0xc2,0x5b,0x5a,0x09,0xd2,0x5c,0xf3,0x03}
	/**----------test 2-----**/
	//standard result: 0x3D,0x91,0xE1,0xF9,0xC2,0x3E,0xAA,0x38,0x09,0x7C,0x87,0xAC,0xC0,0x6F,0x02,0xC9,0x57,0xDD,0x98,0x8F,0x0A,0x24,0x76,0x36,0xCD,0xDD,0x0F,0x91,0x43,0xA4,0xA9,0x5D,0x6A,0x08,0x19,0x58,0x6E,0xE3,0xF3,0xC7,0x31,0xA2,0x76,0xEF,0x74,0x2B,0xEF,0xB1,0xAE,0x61,0x5B,0xBF,0x48,0xCE,0x7D,0xD2,0xA6,0xE8,0x91,0x67,0x63,0x2F,0xE9,0x73
	//sig :=[]byte{0x4d,0x98,0xf9,0xb5,0xac,0x76,0xd3,0x14,0xba,0x24,0x9a,0x37,0xd6,0x4d,0xe3,0x47,0xa7,0xf4,0x06,0x13,0x2c,0x8f,0x76,0x24,0xb6,0x9c,0x74,0xb5,0xba,0xdf,0x97,0x43,0x66,0x8c,0x89,0xed,0xfc,0x17,0x43,0xcd,0x8b,0xd5,0x8f,0xe9,0x35,0x38,0x3d,0xd8,0xd4,0xb7,0x2b,0x25,0xac,0x21,0x11,0x2f,0x66,0xd4,0x5d,0xbe,0xbc,0xc7,0xb1,0xaf,0x00}
	//msg :=[]byte{0xA4,0x4C,0x69,0x32,0x00,0xC3,0x7B,0x00,0x32,0x68,0x76,0x27,0x17,0x6E,0x41,0xDF,0xAC,0xC9,0x53,0xCC,0x77,0xEB,0x97,0x63,0x81,0xCD,0xB7,0xA6,0x6B,0x17,0x21,0x58}
	/**----------test 3-----**/
	//standard result: 0BF0AED10711CCE9C07D6FFBB4CD9D93A00BF53A9722081E5A1A6CB594B0F04EAF978B8F7B7FCAFEEF85A36FBAF66C6FA0EAC05D468E834180DE34CB74DD45CA
	//sig :=[]byte{0xa8,0x06,0x4a,0x1b,0x1e,0xab,0x7f,0x28,0xbd,0x0f,0x26,0xcd,0xbd,0xf2,0x31,0x5e,0x28,0x0b,0x17,0xea,0xca,0xb8,0x34,0xbc,0x27,0xab,0x86,0xe4,0x03,0x07,0xa9,0x82,0x2e,0x2b,0x6b,0xc2,0x90,0x1f,0xa4,0x39,0xce,0x40,0x8d,0xd1,0x3f,0xf7,0xee,0x93,0x0a,0xf5,0x1e,0x47,0xfc,0x36,0x2b,0xb8,0xe4,0x49,0x77,0xe7,0x00,0x9d,0x1b,0x5f,0x00}
	//msg :=[]byte{0xA4,0x4C,0x69,0x32,0x00,0xC3,0x7B,0x00,0x32,0x68,0x76,0x27,0x17,0x6E,0x41,0xDF,0xAC,0xC9,0x53,0xCC,0x77,0xEB,0x97,0x63,0x81,0xCD,0xB7,0xA6,0x6B,0x17,0x21,0x58}

	/**----------test 4-----**/
	//standard result:A5BB3B28466F578E6E93FBFD5F75CEE1AE86033AA4BBEA690E3312C087181EB366F9A1D1D6A437A9BF9FC65EC853B9FD60FA322BE3997C47144EB20DA658B3D1
	sig := []byte{0x38, 0xb7, 0xda, 0xc5, 0xee, 0x93, 0x2a, 0xc1, 0xbf, 0x2b, 0xc6, 0x2c, 0x05, 0xb7, 0x92, 0xcd, 0x93, 0xc3, 0xb4, 0xaf, 0x61, 0xdc, 0x02, 0xdb, 0xb4, 0xb9, 0x3d, 0xac, 0xb7, 0x58, 0x12, 0x3f, 0x08, 0xbf, 0x12, 0x3e, 0xab, 0xe7, 0x74, 0x80, 0x78, 0x7d, 0x66, 0x4c, 0xa2, 0x80, 0xdc, 0x1f, 0x20, 0xd9, 0x20, 0x57, 0x25, 0x32, 0x06, 0x58, 0xc3, 0x9c, 0x6c, 0x14, 0x3f, 0xd5, 0x64, 0x2d, 0x00}
	msg := []byte{0x15, 0x98, 0x17, 0xa0, 0x85, 0xf1, 0x13, 0xd0, 0x99, 0xd3, 0xd9, 0x3c, 0x05, 0x14, 0x10, 0xe9, 0xbf, 0xe0, 0x43, 0xcc, 0x5c, 0x20, 0xe4, 0x3a, 0xa9, 0xa0, 0x83, 0xbf, 0x73, 0x66, 0x01, 0x45}
	pubkey, err := RecoverPubkey(sig, msg, ECC_CURVE_SECP256K1)
	if err != SUCCESS {
		fmt.Println("recover pubkey fail!!!")
	} else {
		fmt.Println("pubkey result:", hex.EncodeToString(pubkey))
	}
}

func Test_tronSign(t *testing.T) {
	//standard result:38b7dac5ee932ac1bf2bc62c05b792cd93c3b4af61dc02dbb4b93dacb758123f 08bf123eabe77480787d664ca280dc1f20d9205725320658c39c6c143fd5642d 00
	prikey := []byte{0x8e, 0x81, 0x24, 0x36, 0xa0, 0xe3, 0x32, 0x31, 0x66, 0xe1, 0xf0, 0xe8, 0xba, 0x79, 0xe1, 0x9e, 0x21, 0x7b, 0x2c, 0x4a, 0x53, 0xc9, 0x70, 0xd4, 0xcc, 0xa0, 0xcf, 0xb1, 0x07, 0x89, 0x79, 0xdf}
	hash := []byte{0x15, 0x98, 0x17, 0xa0, 0x85, 0xf1, 0x13, 0xd0, 0x99, 0xd3, 0xd9, 0x3c, 0x05, 0x14, 0x10, 0xe9, 0xbf, 0xe0, 0x43, 0xcc, 0x5c, 0x20, 0xe4, 0x3a, 0xa9, 0xa0, 0x83, 0xbf, 0x73, 0x66, 0x01, 0x45}
	sig, err := Tron_signature(prikey, hash)
	if err != SUCCESS {
		t.Error("tron sign fail")
	} else {
		fmt.Println("tron sign result:", hex.EncodeToString(sig))
	}
}

func Test_NASsign(t *testing.T) {
	//standard result:e226e5ecb8d640cb2d2fad00bb83089ac5475ba091764001d7015f11f244b4bf4a5f633417cf36e5a45b37f4992cc8d2365988c6f981b2679590ef8c2636e7e401
	prikey := []byte{0x1d, 0x3f, 0xe0, 0x6a, 0x53, 0x91, 0x9e, 0x72, 0x83, 0x15, 0xe2, 0xcc, 0xca, 0x41, 0xd4, 0xaa, 0x5b, 0x19, 0x08, 0x45, 0xa7, 0x90, 0x07, 0x79, 0x75, 0x17, 0xe6, 0x2d, 0xbc, 0x0d, 0xf4, 0x54}
	hash := []byte{0xde, 0x9a, 0xb0, 0xb3, 0xa8, 0x0a, 0x68, 0x44, 0xd3, 0xc3, 0xff, 0xc8, 0x6b, 0xdb, 0xd7, 0x5c, 0x70, 0x60, 0xe2, 0x30, 0xc3, 0x1b, 0xb8, 0x4a, 0xc8, 0xd2, 0x0e, 0xf2, 0x1e, 0x22, 0x5e, 0x4b}
	sig, err := NAS_signature(prikey, hash)
	if err != SUCCESS {
		t.Error("NAS sign fail")
	} else {
		fmt.Println("NAS sign result:", hex.EncodeToString(sig))
	}
}

func Test_IcxSign(t *testing.T) {
	/**---test 1-----**/
	//standard result:9e3813f1beb98607ed0cfa9199a41000ee12ac57f551d46ed944705a2cfad52e713d2ba16e48e58f4c427df185bd73b7142afed19c37752978acf7417aa517af01
	//	hash:=[]byte{0xce,0x99,0x90,0x1c,0x86,0xcb,0x50,0x0b,0xb1,0xd1,0x35,0xb6,0xab,0x6b,0xf1,0xd9,0x18,0xde,0x2a,0x7b,0xc5,0x54,0xc9,0xad,0xc2,0x5b,0x5a,0x09,0xd2,0x5c,0xf3,0x03}
	//prikey:=[]byte{0xc1,0xa1,0x19,0xf1,0x63,0xde,0xf0,0x72,0xf2,0x1a,0x1b,0x3b,0x96,0x05,0x2c,0x65,0x9d,0x71,0x77,0x3f,0x20,0x75,0xec,0x00,0x5d,0x4a,0xd8,0x49,0x24,0x71,0x7e,0x6a}
	//pubkey:=[]byte{0xD7,0xE4,0xC0,0x1B,0x63,0xAA,0xBE,0x17,0x75,0x89,0xFA,0x90,0x05,0xD6,0xC7,0xB1,0x18,0x22,0x83,0xC8,0x04,0xD8,0x43,0xF5,0xF4,0xC4,0xD6,0x16,0x17,0xBC,0x9F,0x23,0xFD,0x27,0xEB,0xDD,0x1E,0x67,0xFF,0x6C,0x93,0xA2,0x56,0x11,0xB5,0xC4,0xC3,0xA3,0xDD,0x7C,0x87,0xBD,0x6E,0x3E,0x63,0x62,0x71,0x7F,0x5E,0x67,0x0D,0xE3,0x66,0x32}
	/**----test 2-----**/
	//standard result:4d98f9b5ac76d314ba249a37d64de347a7f406132c8f7624b69c74b5badf9743668c89edfc1743cd8bd58fe935383dd8d4b72b25ac21112f66d45dbebcc7b1af00
	//hash:=[]byte{0xA4,0x4C,0x69,0x32,0x00,0xC3,0x7B,0x00,0x32,0x68,0x76,0x27,0x17,0x6E,0x41,0xDF,0xAC,0xC9,0x53,0xCC,0x77,0xEB,0x97,0x63,0x81,0xCD,0xB7,0xA6,0x6B,0x17,0x21,0x58}
	//prikey:=[]byte{0xBC,0xB9,0x71,0xDD,0x9A,0x73,0x1B,0x66,0xA4,0x25,0x51,0x7F,0x1F,0x02,0xC8,0xC3,0xAF,0x46,0xAF,0x74,0xFF,0x2F,0x62,0xF4,0xEF,0x21,0x14,0x70,0x41,0xC6,0xBB,0xA5}
	//pubkey:=[]byte{0x3D,0x91,0xE1,0xF9,0xC2,0x3E,0xAA,0x38,0x09,0x7C,0x87,0xAC,0xC0,0x6F,0x02,0xC9,0x57,0xDD,0x98,0x8F,0x0A,0x24,0x76,0x36,0xCD,0xDD,0x0F,0x91,0x43,0xA4,0xA9,0x5D,0x6A,0x08,0x19,0x58,0x6E,0xE3,0xF3,0xC7,0x31,0xA2,0x76,0xEF,0x74,0x2B,0xEF,0xB1,0xAE,0x61,0x5B,0xBF,0x48,0xCE,0x7D,0xD2,0xA6,0xE8,0x91,0x67,0x63,0x2F,0xE9,0x73}
	/**----test 3-----**/
	//standard result: a8064a1b1eab7f28bd0f26cdbdf2315e280b17eacab834bc27ab86e40307a9822e2b6bc2901fa439ce408dd13ff7ee930af51e47fc362bb8e44977e7009d1b5f00
	hash := []byte{0xA4, 0x4C, 0x69, 0x32, 0x00, 0xC3, 0x7B, 0x00, 0x32, 0x68, 0x76, 0x27, 0x17, 0x6E, 0x41, 0xDF, 0xAC, 0xC9, 0x53, 0xCC, 0x77, 0xEB, 0x97, 0x63, 0x81, 0xCD, 0xB7, 0xA6, 0x6B, 0x17, 0x21, 0x58}
	prikey := []byte{0xA8, 0xDE, 0xCB, 0xDF, 0x2A, 0x5C, 0x92, 0xF8, 0xD8, 0xFC, 0x4D, 0x53, 0x36, 0x7F, 0x3A, 0x21, 0x55, 0x84, 0xB0, 0xDD, 0xA9, 0x2E, 0xFC, 0x30, 0xBE, 0x89, 0x51, 0x44, 0xD3, 0xD5, 0x6F, 0x97}
	//pubkey :=[]byte{0x0B,0xF0,0xAE,0xD1,0x07,0x11,0xCC,0xE9,0xC0,0x7D,0x6F,0xFB,0xB4,0xCD,0x9D,0x93,0xA0,0x0B,0xF5,0x3A,0x97,0x22,0x08,0x1E,0x5A,0x1A,0x6C,0xB5,0x94,0xB0,0xF0,0x4E,0xAF,0x97,0x8B,0x8F,0x7B,0x7F,0xCA,0xFE,0xEF,0x85,0xA3,0x6F,0xBA,0xF6,0x6C,0x6F,0xA0,0xEA,0xC0,0x5D,0x46,0x8E,0x83,0x41,0x80,0xDE,0x34,0xCB,0x74,0xDD,0x45,0xCA}
	/**-----test 4-----**/
	//hash:=[]byte{0x23,0x87,0x14,0x9b,0x9d,0x9e,0x58,0x75,0x2b,0x7b,0xf8,0xc5,0x69,0x52,0x21,0xd9,0xfe,0xb7,0x27,0x4f}
	//hash:=[]byte{0x80,0xa1,0x2d,0x64,0x64,0xfa,0x73,0x47,0xab,0x85,0x41,0x1a,0x4c,0x2b,0xe0,0x6b,0xfa,0xd4,0x3c,0x1a,0xb6,0x69,0x13,0xd5,0x0e,0x5a,0x12,0x0c,0x9d,0x2b,0xdb,0xde}
	//prikey := []byte{226, 188, 53, 199, 56, 81, 55, 120, 229, 100, 198, 237, 124, 61, 112, 10, 13, 182, 115, 9, 204, 194, 126, 244, 0, 18, 136, 109, 0, 113, 139, 157}
	sig, err := ICX_signature(prikey, hash)
	if err != SUCCESS {
		t.Error("ICX sign fail")
	} else {
		fmt.Println("ICX sign result:", hex.EncodeToString(sig))
	}
}

func Test_ed25519_ref10(t *testing.T) {
	pri := []byte{0x18, 0x6f, 0xdc, 0x45, 0xdb, 0x17, 0x67, 0x2d, 0x00, 0x56, 0x22, 0x03, 0x8f, 0x4c, 0x9e, 0x1c, 0x42, 0x4a, 0xce, 0xe6, 0x61, 0x10, 0x8f, 0xc7, 0x0a, 0xde, 0xe9, 0xfb, 0x78, 0x71, 0xa5, 0x56}
	pub, ret := GenPubkey(pri, ECC_CURVE_X25519)
	if ret != SUCCESS {
		t.Error("failed to gen pubkey!")
	} else {
		pubStr := hex.EncodeToString(pub)
		if pubStr != "94eb983c9b63d459b6d334ce8d69513e33114916f5da01233815836c1b26b574" {
			t.Error("gen pubkey error!")
		} else {
			fmt.Println("pubkey: ", pubStr)
		}
	}
	sig := []byte{0x2b, 0x05, 0xd0, 0x2d, 0xc9, 0x83, 0x35, 0xd3, 0xdc, 0x7e, 0x42, 0x10, 0x35, 0xea, 0x04, 0x3e, 0x9e, 0x50, 0x60, 0x06, 0x16, 0xec, 0x1e, 0x4a, 0x8f, 0xec, 0x45, 0x59, 0xee, 0x5a, 0x9d, 0xd8, 0x10, 0x62, 0xc1, 0xa8, 0xb6, 0x0e, 0xe9, 0x72, 0x7b, 0xb9, 0x6d, 0xad, 0xa5, 0xd2, 0xfa, 0xe9, 0x92, 0x25, 0x00, 0x83, 0xfa, 0xda, 0xc0, 0x67, 0xb4, 0x39, 0x58, 0xa8, 0x6d, 0xf5, 0x68, 0x05}
	data := []byte{0x02, 0x15, 0x7a, 0x9d, 0x02, 0xac, 0x57, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3b, 0x9a, 0xca, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x98, 0x96, 0x80, 0x00, 0x64, 0x05, 0x54, 0x9c, 0x6d, 0xf7, 0xb3, 0x76, 0x77, 0x1b, 0x19, 0xff, 0x3b, 0xdb, 0x58, 0xd0, 0x4b, 0x49, 0x99, 0x91, 0x66, 0x3c, 0x47, 0x44, 0x4e, 0x42, 0x5f, 0x00, 0x00}

	pass := Verify(pub, nil, 0, data[:54], 54, sig, ECC_CURVE_X25519)
	if pass != 0 {
		t.Error("verify failed!")
	} else {
		fmt.Println("success!")
	}

	pass = Verify(pub, nil, 0, data, 55, sig, ECC_CURVE_X25519)
	if pass != 1 {
		t.Error("verify failed!")
	} else {
		fmt.Println("success!")
	}

}

func Test_curve25519_point_convert(t *testing.T) {
	x25519 := []byte{0x94, 0xeb, 0x98, 0x3c, 0x9b, 0x63, 0xd4, 0x59, 0xb6, 0xd3, 0x34, 0xce, 0x8d, 0x69, 0x51, 0x3e, 0x33, 0x11, 0x49, 0x16, 0xf5, 0xda, 0x01, 0x23, 0x38, 0x15, 0x83, 0x6c, 0x1b, 0x26, 0xb5, 0x74}
	ed25519 := []byte{0x25, 0x00, 0x5C, 0x9D, 0x6F, 0x79, 0xE6, 0x5F, 0x09, 0x6B, 0x7B, 0x71, 0x7C, 0x69, 0x57, 0x94, 0x53, 0x04, 0xB5, 0xFA, 0x86, 0x11, 0x68, 0x83, 0xE2, 0x6C, 0x31, 0xB8, 0xD4, 0xC8, 0x6E, 0x23}

	edchk, err := CURVE25519_convert_X_to_Ed(x25519)

	if err != nil {
		t.Error("convert failed!")
	} else {
		if hex.EncodeToString(ed25519) == hex.EncodeToString(edchk) {
			fmt.Println("convert success!")
		} else {
			t.Error("convert passed but result wrong!")
		}
	}

	xchk, err := CURVE25519_convert_Ed_to_X(ed25519)
	if err != nil {
		t.Error("convert failed!")
	} else {
		if hex.EncodeToString(x25519) == hex.EncodeToString(xchk) {
			fmt.Println("convert success!")
		} else {
			t.Error("convert passed but result wrong!")
		}
	}
}
