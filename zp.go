/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"dm/util"
	"os"
	"strconv"
	"strings"
)

const (
	Dm_build_596 = "7.6.0.0"

	Dm_build_597 = "7.0.0.9"

	Dm_build_598 = "8.0.0.73"

	Dm_build_599 = "7.1.2.128"

	Dm_build_600 = "7.1.5.144"

	Dm_build_601 = "7.1.6.123"

	Dm_build_602 = 32768 - 128

	Dm_build_603 int16 = 1

	Dm_build_604 int16 = 2

	Dm_build_605 int16 = 3

	Dm_build_606 int16 = 4

	Dm_build_607 int16 = 5

	Dm_build_608 int16 = 6

	Dm_build_609 int16 = 7

	Dm_build_610 int16 = 8

	Dm_build_611 int16 = 9

	Dm_build_612 int16 = 13

	Dm_build_613 int16 = 14

	Dm_build_614 int16 = 15

	Dm_build_615 int16 = 17

	Dm_build_616 int16 = 21

	Dm_build_617 int16 = 24

	Dm_build_618 int16 = 27

	Dm_build_619 int16 = 29

	Dm_build_620 int16 = 30

	Dm_build_621 int16 = 31

	Dm_build_622 int16 = 32

	Dm_build_623 int16 = 44

	Dm_build_624 int16 = 52

	Dm_build_625 int16 = 60

	Dm_build_626 int16 = 71

	Dm_build_627 int16 = 90

	Dm_build_628 int16 = 91

	Dm_build_629 int16 = 200

	Dm_build_630 = 64

	Dm_build_631 = 20

	Dm_build_632 = 0

	Dm_build_633 = 4

	Dm_build_634 = 6

	Dm_build_635 = 10

	Dm_build_636 = 14

	Dm_build_637 = 18

	Dm_build_638 = 19

	Dm_build_639 = 128

	Dm_build_640 = 256

	Dm_build_641 = 0xffff

	Dm_build_642 int32 = 2

	Dm_build_643 = -1

	Dm_build_644 uint16 = 0xFFFE

	Dm_build_645 uint16 = uint16(Dm_build_644 - 3)

	Dm_build_646 uint16 = Dm_build_644

	Dm_build_647 int32 = 0xFFFF

	Dm_build_648 int32 = 0x80

	Dm_build_649 byte = 0x60

	Dm_build_650 uint16 = uint16(Dm_build_646)

	Dm_build_651 uint16 = uint16(Dm_build_647)

	Dm_build_652 int16 = 0x00

	Dm_build_653 int16 = 0x03

	Dm_build_654 int32 = 0x80

	Dm_build_655 byte = 0

	Dm_build_656 byte = 1

	Dm_build_657 byte = 2

	Dm_build_658 byte = 3

	Dm_build_659 byte = 4

	Dm_build_660 byte = Dm_build_655

	Dm_build_661 int = 10

	Dm_build_662 int32 = 32

	Dm_build_663 int32 = 65536

	Dm_build_664 byte = 0

	Dm_build_665 byte = 1

	Dm_build_666 int32 = 0x00000000

	Dm_build_667 int32 = 0x00000020

	Dm_build_668 int32 = 0x00000040

	Dm_build_669 int32 = 0x00000FFF

	Dm_build_670 int32 = 0

	Dm_build_671 int32 = 1

	Dm_build_672 int32 = 2

	Dm_build_673 int32 = 3

	Dm_build_674 = 8192

	Dm_build_675 = 1

	Dm_build_676 = 2

	Dm_build_677 = 0

	Dm_build_678 = 0

	Dm_build_679 = 1

	Dm_build_680 = -1

	Dm_build_681 int16 = 0

	Dm_build_682 int16 = 1

	Dm_build_683 int16 = 2

	Dm_build_684 int16 = 3

	Dm_build_685 int16 = 4

	Dm_build_686 int16 = 127

	Dm_build_687 int16 = Dm_build_686 + 20

	Dm_build_688 int16 = Dm_build_686 + 21

	Dm_build_689 int16 = Dm_build_686 + 22

	Dm_build_690 int16 = Dm_build_686 + 24

	Dm_build_691 int16 = Dm_build_686 + 25

	Dm_build_692 int16 = Dm_build_686 + 26

	Dm_build_693 int16 = Dm_build_686 + 30

	Dm_build_694 int16 = Dm_build_686 + 31

	Dm_build_695 int16 = Dm_build_686 + 32

	Dm_build_696 int16 = Dm_build_686 + 33

	Dm_build_697 int16 = Dm_build_686 + 35

	Dm_build_698 int16 = Dm_build_686 + 38

	Dm_build_699 int16 = Dm_build_686 + 39

	Dm_build_700 int16 = Dm_build_686 + 51

	Dm_build_701 int16 = Dm_build_686 + 71

	Dm_build_702 int16 = Dm_build_686 + 124

	Dm_build_703 int16 = Dm_build_686 + 125

	Dm_build_704 int16 = Dm_build_686 + 126

	Dm_build_705 int16 = Dm_build_686 + 127

	Dm_build_706 int16 = Dm_build_686 + 128

	Dm_build_707 int16 = Dm_build_686 + 129

	Dm_build_708 byte = 0

	Dm_build_709 byte = 2

	Dm_build_710 = 2048

	Dm_build_711 = -1

	Dm_build_712 = 0

	Dm_build_713 = 16000

	Dm_build_714 = 32000

	Dm_build_715 = 0x00000000

	Dm_build_716 = 0x00000020

	Dm_build_717 = 0x00000040

	Dm_build_718 = 0x00000FFF
)

type dm_build_719 interface {
	dm_build_720()
	dm_build_721() error
	dm_build_722()
	dm_build_723(imsg dm_build_719) error
	dm_build_724() error
	dm_build_725() (interface{}, error)
	dm_build_726()
	dm_build_727(imsg dm_build_719) (interface{}, error)
	dm_build_728()
	dm_build_729() error
	dm_build_730() byte
	dm_build_731() int32
	dm_build_732(length int32)
	dm_build_733() int16
}

type dm_build_734 struct {
	dm_build_735 *dm_build_332

	dm_build_736 int16

	dm_build_737 int32

	dm_build_738 *DmStatement
}

func (dm_build_740 *dm_build_734) dm_build_739(dm_build_741 *dm_build_332, dm_build_742 int16) *dm_build_734 {
	dm_build_740.dm_build_735 = dm_build_741
	dm_build_740.dm_build_736 = dm_build_742
	return dm_build_740
}

func (dm_build_744 *dm_build_734) dm_build_743(dm_build_745 *dm_build_332, dm_build_746 int16, dm_build_747 *DmStatement) *dm_build_734 {
	dm_build_744.dm_build_739(dm_build_745, dm_build_746).dm_build_738 = dm_build_747
	return dm_build_744
}

func dm_build_748(dm_build_749 *dm_build_332, dm_build_750 int16) *dm_build_734 {
	return new(dm_build_734).dm_build_739(dm_build_749, dm_build_750)
}

func dm_build_751(dm_build_752 *dm_build_332, dm_build_753 int16, dm_build_754 *DmStatement) *dm_build_734 {
	return new(dm_build_734).dm_build_743(dm_build_752, dm_build_753, dm_build_754)
}

func (dm_build_756 *dm_build_734) dm_build_720() {
	dm_build_756.dm_build_735.dm_build_335.Dm_build_14(0)
	dm_build_756.dm_build_735.dm_build_335.Dm_build_25(Dm_build_630, true, true)
}

func (dm_build_758 *dm_build_734) dm_build_721() error {
	return nil
}

func (dm_build_760 *dm_build_734) dm_build_722() {
	if dm_build_760.dm_build_738 == nil {
		dm_build_760.dm_build_735.dm_build_335.Dm_build_191(Dm_build_632, 0)
	} else {
		dm_build_760.dm_build_735.dm_build_335.Dm_build_191(Dm_build_632, dm_build_760.dm_build_738.id)
	}

	dm_build_760.dm_build_735.dm_build_335.Dm_build_187(Dm_build_633, dm_build_760.dm_build_736)
	dm_build_760.dm_build_735.dm_build_335.Dm_build_191(Dm_build_634, int32(dm_build_760.dm_build_735.dm_build_335.Dm_build_12()-Dm_build_630))
}

func (dm_build_762 *dm_build_734) dm_build_724() error {
	dm_build_762.dm_build_735.dm_build_335.Dm_build_17(0)
	dm_build_762.dm_build_735.dm_build_335.Dm_build_25(Dm_build_630, false, true)
	return dm_build_762.dm_build_767()
}

func (dm_build_764 *dm_build_734) dm_build_725() (interface{}, error) {
	return nil, nil
}

func (dm_build_766 *dm_build_734) dm_build_726() {
}

func (dm_build_768 *dm_build_734) dm_build_767() error {
	dm_build_768.dm_build_737 = dm_build_768.dm_build_735.dm_build_335.Dm_build_269(Dm_build_635)
	if dm_build_768.dm_build_737 < 0 && dm_build_768.dm_build_737 != EC_RN_EXCEED_ROWSET_SIZE.ErrCode {
		return (&DmError{dm_build_768.dm_build_737, dm_build_768.dm_build_769(), nil, ""}).throw()
	} else if dm_build_768.dm_build_737 > 0 {

	} else if dm_build_768.dm_build_736 == Dm_build_629 || dm_build_768.dm_build_736 == Dm_build_603 {
		dm_build_768.dm_build_769()
	}

	return nil
}

func (dm_build_770 *dm_build_734) dm_build_769() string {

	dm_build_771 := dm_build_770.dm_build_735.dm_build_336.getServerEncoding()

	if dm_build_771 != "" && dm_build_771 == ENCODING_EUCKR && Locale != LANGUAGE_EN {
		dm_build_771 = ENCODING_GB18030
	}

	dm_build_770.dm_build_735.dm_build_335.Dm_build_25(int(dm_build_770.dm_build_735.dm_build_335.Dm_build_125()), false, true)

	dm_build_770.dm_build_735.dm_build_335.Dm_build_25(int(dm_build_770.dm_build_735.dm_build_335.Dm_build_125()), false, true)

	dm_build_770.dm_build_735.dm_build_335.Dm_build_25(int(dm_build_770.dm_build_735.dm_build_335.Dm_build_125()), false, true)

	return dm_build_770.dm_build_735.dm_build_335.Dm_build_167(dm_build_771, dm_build_770.dm_build_735.dm_build_336)
}

func (dm_build_773 *dm_build_734) dm_build_723(dm_build_774 dm_build_719) (dm_build_775 error) {
	dm_build_774.dm_build_720()
	if dm_build_775 = dm_build_774.dm_build_721(); dm_build_775 != nil {
		return dm_build_775
	}
	dm_build_774.dm_build_722()
	return nil
}

func (dm_build_777 *dm_build_734) dm_build_727(dm_build_778 dm_build_719) (dm_build_779 interface{}, dm_build_780 error) {
	dm_build_780 = dm_build_778.dm_build_724()
	if dm_build_780 != nil {
		return nil, dm_build_780
	}
	dm_build_779, dm_build_780 = dm_build_778.dm_build_725()
	if dm_build_780 != nil {
		return nil, dm_build_780
	}
	dm_build_778.dm_build_726()
	return dm_build_779, nil
}

func (dm_build_782 *dm_build_734) dm_build_728() {
	dm_build_782.dm_build_735.dm_build_335.Dm_build_183(Dm_build_638, dm_build_782.dm_build_730())
}

func (dm_build_784 *dm_build_734) dm_build_729() error {
	dm_build_785 := dm_build_784.dm_build_735.dm_build_335.Dm_build_263(Dm_build_638)
	dm_build_786 := dm_build_784.dm_build_730()
	if dm_build_785 != dm_build_786 {
		return ECGO_MSG_CHECK_ERROR.throw()
	}
	return nil
}

func (dm_build_788 *dm_build_734) dm_build_730() byte {
	dm_build_789 := dm_build_788.dm_build_735.dm_build_335.Dm_build_263(0)

	for i := 1; i < Dm_build_638; i++ {
		dm_build_789 ^= dm_build_788.dm_build_735.dm_build_335.Dm_build_263(i)
	}

	return dm_build_789
}

func (dm_build_791 *dm_build_734) dm_build_731() int32 {
	return dm_build_791.dm_build_735.dm_build_335.Dm_build_269(Dm_build_634)
}

func (dm_build_793 *dm_build_734) dm_build_732(dm_build_794 int32) {
	dm_build_793.dm_build_735.dm_build_335.Dm_build_191(Dm_build_634, dm_build_794)
}

func (dm_build_796 *dm_build_734) dm_build_733() int16 {
	return dm_build_796.dm_build_736
}

type dm_build_797 struct {
	dm_build_734
}

func dm_build_798(dm_build_799 *dm_build_332) *dm_build_797 {
	dm_build_800 := new(dm_build_797)
	dm_build_800.dm_build_739(dm_build_799, Dm_build_610)
	return dm_build_800
}

type dm_build_801 struct {
	dm_build_734
	dm_build_802 string
}

func dm_build_803(dm_build_804 *dm_build_332, dm_build_805 *DmStatement, dm_build_806 string) *dm_build_801 {
	dm_build_807 := new(dm_build_801)
	dm_build_807.dm_build_743(dm_build_804, Dm_build_618, dm_build_805)
	dm_build_807.dm_build_802 = dm_build_806
	dm_build_807.dm_build_738.cursorName = dm_build_806
	return dm_build_807
}

func (dm_build_809 *dm_build_801) dm_build_721() error {
	dm_build_809.dm_build_735.dm_build_335.Dm_build_113(dm_build_809.dm_build_802, dm_build_809.dm_build_735.dm_build_336.getServerEncoding(), dm_build_809.dm_build_735.dm_build_336)
	dm_build_809.dm_build_735.dm_build_335.Dm_build_51(1)
	return nil
}

type Dm_build_810 struct {
	dm_build_826
	dm_build_811 []OptParameter
}

func dm_build_812(dm_build_813 *dm_build_332, dm_build_814 *DmStatement, dm_build_815 []OptParameter) *Dm_build_810 {
	dm_build_816 := new(Dm_build_810)
	dm_build_816.dm_build_743(dm_build_813, Dm_build_628, dm_build_814)
	dm_build_816.dm_build_811 = dm_build_815
	return dm_build_816
}

func (dm_build_818 *Dm_build_810) dm_build_721() error {
	dm_build_819 := len(dm_build_818.dm_build_811)

	dm_build_818.dm_build_840(int16(dm_build_819), 1)

	dm_build_818.dm_build_735.dm_build_335.Dm_build_113(dm_build_818.dm_build_738.nativeSql, dm_build_818.dm_build_738.dmConn.getServerEncoding(), dm_build_818.dm_build_738.dmConn)

	for _, param := range dm_build_818.dm_build_811 {
		dm_build_818.dm_build_735.dm_build_335.Dm_build_43(param.ioType)
		dm_build_818.dm_build_735.dm_build_335.Dm_build_51(int32(param.tp))
		dm_build_818.dm_build_735.dm_build_335.Dm_build_51(int32(param.prec))
		dm_build_818.dm_build_735.dm_build_335.Dm_build_51(int32(param.scale))
	}

	for _, param := range dm_build_818.dm_build_811 {
		if param.bytes == nil {
			dm_build_818.dm_build_735.dm_build_335.Dm_build_59(Dm_build_646)
		} else {
			dm_build_818.dm_build_735.dm_build_335.Dm_build_89(param.bytes[:len(param.bytes)])
		}
	}
	return nil
}

func (dm_build_821 *Dm_build_810) dm_build_725() (interface{}, error) {
	return dm_build_821.dm_build_826.dm_build_725()
}

const (
	Dm_build_822 int = 0x01

	Dm_build_823 int = 0x02

	Dm_build_824 int = 0x04

	Dm_build_825 int = 0x08
)

type dm_build_826 struct {
	dm_build_734
	dm_build_827 [][]interface{}
	dm_build_828 []parameter
	dm_build_829 bool
}

func dm_build_830(dm_build_831 *dm_build_332, dm_build_832 int16, dm_build_833 *DmStatement) *dm_build_826 {
	dm_build_834 := new(dm_build_826)
	dm_build_834.dm_build_743(dm_build_831, dm_build_832, dm_build_833)
	dm_build_834.dm_build_829 = true
	return dm_build_834
}

func dm_build_835(dm_build_836 *dm_build_332, dm_build_837 *DmStatement, dm_build_838 [][]interface{}) *dm_build_826 {
	dm_build_839 := new(dm_build_826)

	if dm_build_836.dm_build_336.Execute2 {
		dm_build_839.dm_build_743(dm_build_836, Dm_build_612, dm_build_837)
	} else {
		dm_build_839.dm_build_743(dm_build_836, Dm_build_608, dm_build_837)
	}

	dm_build_839.dm_build_828 = dm_build_837.params
	dm_build_839.dm_build_827 = dm_build_838
	dm_build_839.dm_build_829 = true
	return dm_build_839
}

func (dm_build_841 *dm_build_826) dm_build_840(dm_build_842 int16, dm_build_843 int64) {

	dm_build_844 := Dm_build_631

	if dm_build_841.dm_build_735.dm_build_336.autoCommit {
		dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_183(dm_build_844, 1)
	} else {
		dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_183(dm_build_844, 0)
	}

	dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_187(dm_build_844, dm_build_842)

	dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_183(dm_build_844, 1)

	dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_195(dm_build_844, dm_build_843)

	dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_195(dm_build_844, dm_build_841.dm_build_738.cursorUpdateRow)

	if dm_build_841.dm_build_738.maxRows <= 0 || dm_build_841.dm_build_738.dmConn.dmConnector.enRsCache {
		dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_195(dm_build_844, INT64_MAX)
	} else {
		dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_195(dm_build_844, dm_build_841.dm_build_738.maxRows)
	}

	dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_183(dm_build_844, 1)

	if dm_build_841.dm_build_735.dm_build_336.dmConnector.continueBatchOnError {
		dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_183(dm_build_844, 1)
	} else {
		dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_183(dm_build_844, 0)
	}

	dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_183(dm_build_844, 0)

	dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_183(dm_build_844, 0)

	if dm_build_841.dm_build_738.queryTimeout == 0 {
		dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_191(dm_build_844, -1)
	} else {
		dm_build_844 += dm_build_841.dm_build_735.dm_build_335.Dm_build_191(dm_build_844, dm_build_841.dm_build_738.queryTimeout)
	}
}

func (dm_build_846 *dm_build_826) dm_build_721() error {
	var dm_build_847 int16
	var dm_build_848 int64

	if dm_build_846.dm_build_828 != nil {
		dm_build_847 = int16(len(dm_build_846.dm_build_828))
	} else {
		dm_build_847 = 0
	}

	if dm_build_846.dm_build_827 != nil {
		dm_build_848 = int64(len(dm_build_846.dm_build_827))
	} else {
		dm_build_848 = 0
	}

	dm_build_846.dm_build_840(dm_build_847, dm_build_848)

	if dm_build_847 > 0 {
		err := dm_build_846.dm_build_849(dm_build_846.dm_build_828)
		if err != nil {
			return err
		}
		if dm_build_846.dm_build_827 != nil && len(dm_build_846.dm_build_827) > 0 {
			for _, paramObject := range dm_build_846.dm_build_827 {
				if err := dm_build_846.dm_build_852(paramObject); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (dm_build_850 *dm_build_826) dm_build_849(dm_build_851 []parameter) error {
	for _, param := range dm_build_851 {

		if param.colType == CURSOR && param.ioType == IO_TYPE_OUT {
			dm_build_850.dm_build_735.dm_build_335.Dm_build_43(IO_TYPE_INOUT)
		} else {
			dm_build_850.dm_build_735.dm_build_335.Dm_build_43(param.ioType)
		}

		dm_build_850.dm_build_735.dm_build_335.Dm_build_51(param.colType)

		lprec := param.prec
		lscale := param.scale
		typeDesc := param.typeDescriptor
		switch param.colType {
		case ARRAY, SARRAY:
			tmp, err := getPackArraySize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case PLTYPE_RECORD:
			tmp, err := getPackRecordSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case CLASS:
			tmp, err := getPackClassSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case BLOB:
			if isComplexType(int(param.colType), int(param.scale)) {
				lprec = int32(typeDesc.getObjId())
				if lprec == 4 {
					lprec = int32(typeDesc.getOuterId())
				}
			}
		}

		dm_build_850.dm_build_735.dm_build_335.Dm_build_51(lprec)

		dm_build_850.dm_build_735.dm_build_335.Dm_build_51(lscale)

		switch param.colType {
		case ARRAY, SARRAY:
			err := packArray(typeDesc, dm_build_850.dm_build_735.dm_build_335)
			if err != nil {
				return err
			}

		case PLTYPE_RECORD:
			err := packRecord(typeDesc, dm_build_850.dm_build_735.dm_build_335)
			if err != nil {
				return err
			}

		case CLASS:
			err := packClass(typeDesc, dm_build_850.dm_build_735.dm_build_335)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (dm_build_853 *dm_build_826) dm_build_852(dm_build_854 []interface{}) error {
	for i := 0; i < len(dm_build_853.dm_build_828); i++ {

		if dm_build_853.dm_build_828[i].colType == CURSOR {
			dm_build_853.dm_build_735.dm_build_335.Dm_build_47(ULINT_SIZE)
			dm_build_853.dm_build_735.dm_build_335.Dm_build_51(dm_build_853.dm_build_828[i].cursorStmt.id)
			continue
		}

		if dm_build_853.dm_build_828[i].ioType == IO_TYPE_OUT {
			continue
		}

		switch dm_build_854[i].(type) {
		case []byte:
			if dataBytes, ok := dm_build_854[i].([]byte); ok {
				if len(dataBytes) > Dm_build_641 {
					return ECGO_DATA_TOO_LONG.throw()
				}
				dm_build_853.dm_build_735.dm_build_335.Dm_build_89(dataBytes)
			}
		case int:
			if dm_build_854[i] == ParamDataEnum_Null {
				dm_build_853.dm_build_735.dm_build_335.Dm_build_59(Dm_build_646)
			} else if dm_build_854[i] == ParamDataEnum_OFF_ROW {
				dm_build_853.dm_build_735.dm_build_335.Dm_build_47(0)
			}
		case lobCtl:
			dm_build_853.dm_build_735.dm_build_335.Dm_build_59(uint16(Dm_build_645))
			dm_build_853.dm_build_735.dm_build_335.Dm_build_79(dm_build_854[i].(lobCtl).value)
		default:
			panic("Bind param data failed by invalid param data type: ")
		}
	}

	return nil
}

func (dm_build_856 *dm_build_826) dm_build_725() (interface{}, error) {
	dm_build_857 := execInfo{}
	dm_build_858 := dm_build_856.dm_build_738.dmConn

	dm_build_859 := Dm_build_631

	dm_build_857.retSqlType = dm_build_856.dm_build_735.dm_build_335.Dm_build_266(dm_build_859)
	dm_build_859 += USINT_SIZE

	dm_build_860 := dm_build_856.dm_build_735.dm_build_335.Dm_build_266(dm_build_859)
	dm_build_859 += USINT_SIZE

	dm_build_857.updateCount = dm_build_856.dm_build_735.dm_build_335.Dm_build_272(dm_build_859)
	dm_build_859 += DDWORD_SIZE

	dm_build_861 := dm_build_856.dm_build_735.dm_build_335.Dm_build_266(dm_build_859)
	dm_build_859 += USINT_SIZE

	dm_build_857.rsUpdatable = dm_build_856.dm_build_735.dm_build_335.Dm_build_263(dm_build_859) != 0
	dm_build_859 += BYTE_SIZE

	dm_build_862 := dm_build_856.dm_build_735.dm_build_335.Dm_build_266(dm_build_859)
	dm_build_859 += ULINT_SIZE

	dm_build_857.printLen = dm_build_856.dm_build_735.dm_build_335.Dm_build_269(dm_build_859)
	dm_build_859 += ULINT_SIZE

	var dm_build_863 int16 = -1
	if dm_build_857.retSqlType == Dm_build_696 || dm_build_857.retSqlType == Dm_build_697 {
		dm_build_857.rowid = 0

		dm_build_857.rsBdta = dm_build_856.dm_build_735.dm_build_335.Dm_build_263(dm_build_859) == Dm_build_709
		dm_build_859 += BYTE_SIZE

		dm_build_863 = dm_build_856.dm_build_735.dm_build_335.Dm_build_266(dm_build_859)
		dm_build_859 += USINT_SIZE
		dm_build_859 += 5
	} else {
		dm_build_857.rowid = dm_build_856.dm_build_735.dm_build_335.Dm_build_272(dm_build_859)
		dm_build_859 += DDWORD_SIZE
	}

	dm_build_857.execId = dm_build_856.dm_build_735.dm_build_335.Dm_build_269(dm_build_859)
	dm_build_859 += ULINT_SIZE

	dm_build_857.rsCacheOffset = dm_build_856.dm_build_735.dm_build_335.Dm_build_269(dm_build_859)
	dm_build_859 += ULINT_SIZE

	dm_build_864 := dm_build_856.dm_build_735.dm_build_335.Dm_build_263(dm_build_859)
	dm_build_859 += BYTE_SIZE
	dm_build_865 := (dm_build_864 & 0x01) == 0x01
	dm_build_866 := (dm_build_864 & 0x02) == 0x02

	dm_build_858.TrxStatus = dm_build_856.dm_build_735.dm_build_335.Dm_build_269(dm_build_859)
	dm_build_858.setTrxFinish(dm_build_858.TrxStatus)
	dm_build_859 += ULINT_SIZE

	if dm_build_857.printLen > 0 {
		bytes := dm_build_856.dm_build_735.dm_build_335.Dm_build_146(int(dm_build_857.printLen))
		dm_build_857.printMsg = Dm_build_1195.Dm_build_1349(bytes, 0, len(bytes), dm_build_858.getServerEncoding(), dm_build_858)
	}

	if dm_build_861 > 0 {
		dm_build_857.outParamDatas = dm_build_856.dm_build_867(int(dm_build_861))
	}

	switch dm_build_857.retSqlType {
	case Dm_build_698:
		dm_build_858.dmConnector.localTimezone = dm_build_856.dm_build_735.dm_build_335.Dm_build_122()
	case Dm_build_696:
		dm_build_857.hasResultSet = true
		if dm_build_860 > 0 {
			dm_build_856.dm_build_738.columns = dm_build_856.dm_build_876(int(dm_build_860), dm_build_857.rsBdta)
		}
		dm_build_856.dm_build_886(&dm_build_857, len(dm_build_856.dm_build_738.columns), int(dm_build_862), int(dm_build_863))
	case Dm_build_697:
		if dm_build_860 > 0 || dm_build_862 > 0 {
			dm_build_857.hasResultSet = true
		}
		if dm_build_860 > 0 {
			dm_build_856.dm_build_738.columns = dm_build_856.dm_build_876(int(dm_build_860), dm_build_857.rsBdta)
		}
		dm_build_856.dm_build_886(&dm_build_857, len(dm_build_856.dm_build_738.columns), int(dm_build_862), int(dm_build_863))
	case Dm_build_699:
		dm_build_858.IsoLevel = int32(dm_build_856.dm_build_735.dm_build_335.Dm_build_122())
		dm_build_858.ReadOnly = dm_build_856.dm_build_735.dm_build_335.Dm_build_119() == 1
	case Dm_build_692:
		dm_build_858.Schema = dm_build_856.dm_build_735.dm_build_335.Dm_build_167(dm_build_858.getServerEncoding(), dm_build_858)
	case Dm_build_689:
		dm_build_857.explain = dm_build_856.dm_build_735.dm_build_335.Dm_build_167(dm_build_858.getServerEncoding(), dm_build_858)

	case Dm_build_693, Dm_build_695, Dm_build_694:
		if dm_build_865 {

			counts := dm_build_856.dm_build_735.dm_build_335.Dm_build_125()
			rowCounts := make([]int64, counts)
			for i := 0; i < int(counts); i++ {
				rowCounts[i] = dm_build_856.dm_build_735.dm_build_335.Dm_build_128()
			}
			dm_build_857.updateCounts = rowCounts
		}

		if dm_build_866 {
			rows := dm_build_856.dm_build_735.dm_build_335.Dm_build_125()

			var lastInsertId int64
			for i := 0; i < int(rows); i++ {
				lastInsertId = dm_build_856.dm_build_735.dm_build_335.Dm_build_128()
			}
			dm_build_857.lastInsertId = lastInsertId

		} else if dm_build_857.updateCount == 1 {
			dm_build_857.lastInsertId = dm_build_857.rowid
		}

		if dm_build_856.dm_build_737 == EC_BP_WITH_ERROR.ErrCode {
			dm_build_856.dm_build_892(dm_build_857.updateCounts)
		}
	case Dm_build_702:
		len := dm_build_856.dm_build_735.dm_build_335.Dm_build_137()
		dm_build_858.OracleDateFormat = dm_build_856.dm_build_735.dm_build_335.Dm_build_162(int(len), dm_build_858.getServerEncoding(), dm_build_858)
	case Dm_build_704:

		len := dm_build_856.dm_build_735.dm_build_335.Dm_build_137()
		dm_build_858.OracleTimestampFormat = dm_build_856.dm_build_735.dm_build_335.Dm_build_162(int(len), dm_build_858.getServerEncoding(), dm_build_858)
	case Dm_build_705:

		len := dm_build_856.dm_build_735.dm_build_335.Dm_build_137()
		dm_build_858.OracleTimestampTZFormat = dm_build_856.dm_build_735.dm_build_335.Dm_build_162(int(len), dm_build_858.getServerEncoding(), dm_build_858)
	case Dm_build_703:
		len := dm_build_856.dm_build_735.dm_build_335.Dm_build_137()
		dm_build_858.OracleTimeFormat = dm_build_856.dm_build_735.dm_build_335.Dm_build_162(int(len), dm_build_858.getServerEncoding(), dm_build_858)
	case Dm_build_706:
		len := dm_build_856.dm_build_735.dm_build_335.Dm_build_137()
		dm_build_858.OracleTimeTZFormat = dm_build_856.dm_build_735.dm_build_335.Dm_build_162(int(len), dm_build_858.getServerEncoding(), dm_build_858)
	case Dm_build_707:
		dm_build_858.OracleDateLanguage = dm_build_856.dm_build_735.dm_build_335.Dm_build_137()
	}

	return &dm_build_857, nil
}

func (dm_build_868 *dm_build_826) dm_build_867(dm_build_869 int) [][]byte {
	dm_build_870 := make([]int, dm_build_869)

	dm_build_871 := 0
	for i := 0; i < len(dm_build_868.dm_build_828); i++ {
		if dm_build_868.dm_build_828[i].ioType == IO_TYPE_INOUT || dm_build_868.dm_build_828[i].ioType == IO_TYPE_OUT {
			dm_build_870[dm_build_871] = i
			dm_build_871++
		}
	}

	dm_build_872 := make([][]byte, len(dm_build_868.dm_build_828))
	var dm_build_873 int32
	var dm_build_874 bool
	var dm_build_875 []byte = nil
	for i := 0; i < dm_build_869; i++ {
		dm_build_874 = false
		dm_build_873 = int32(dm_build_868.dm_build_735.dm_build_335.Dm_build_140())

		if dm_build_873 == int32(Dm_build_646) {
			dm_build_873 = 0
			dm_build_874 = true
		} else if dm_build_873 == int32(Dm_build_647) {
			dm_build_873 = dm_build_868.dm_build_735.dm_build_335.Dm_build_125()
		}

		if dm_build_874 {
			dm_build_872[dm_build_870[i]] = nil
		} else {
			dm_build_875 = dm_build_868.dm_build_735.dm_build_335.Dm_build_146(int(dm_build_873))
			dm_build_872[dm_build_870[i]] = dm_build_875
		}
	}

	return dm_build_872
}

func (dm_build_877 *dm_build_826) dm_build_876(dm_build_878 int, dm_build_879 bool) []column {
	dm_build_880 := dm_build_877.dm_build_735.dm_build_336.getServerEncoding()
	var dm_build_881, dm_build_882, dm_build_883, dm_build_884 int16
	dm_build_885 := make([]column, dm_build_878)
	for i := 0; i < dm_build_878; i++ {
		dm_build_885[i].InitColumn()

		dm_build_885[i].colType = dm_build_877.dm_build_735.dm_build_335.Dm_build_125()

		dm_build_885[i].prec = dm_build_877.dm_build_735.dm_build_335.Dm_build_125()

		dm_build_885[i].scale = dm_build_877.dm_build_735.dm_build_335.Dm_build_125()

		dm_build_885[i].nullable = dm_build_877.dm_build_735.dm_build_335.Dm_build_125() != 0

		itemFlag := dm_build_877.dm_build_735.dm_build_335.Dm_build_122()
		dm_build_885[i].lob = int(itemFlag)&Dm_build_823 != 0
		dm_build_885[i].identity = int(itemFlag)&Dm_build_822 != 0
		dm_build_885[i].readonly = int(itemFlag)&Dm_build_824 != 0

		dm_build_877.dm_build_735.dm_build_335.Dm_build_25(4, false, true)

		dm_build_877.dm_build_735.dm_build_335.Dm_build_25(2, false, true)

		dm_build_881 = dm_build_877.dm_build_735.dm_build_335.Dm_build_122()

		dm_build_882 = dm_build_877.dm_build_735.dm_build_335.Dm_build_122()

		dm_build_883 = dm_build_877.dm_build_735.dm_build_335.Dm_build_122()

		dm_build_884 = dm_build_877.dm_build_735.dm_build_335.Dm_build_122()
		dm_build_885[i].name = dm_build_877.dm_build_735.dm_build_335.Dm_build_162(int(dm_build_881), dm_build_880, dm_build_877.dm_build_735.dm_build_336)
		dm_build_885[i].typeName = dm_build_877.dm_build_735.dm_build_335.Dm_build_162(int(dm_build_882), dm_build_880, dm_build_877.dm_build_735.dm_build_336)
		dm_build_885[i].tableName = dm_build_877.dm_build_735.dm_build_335.Dm_build_162(int(dm_build_883), dm_build_880, dm_build_877.dm_build_735.dm_build_336)
		dm_build_885[i].schemaName = dm_build_877.dm_build_735.dm_build_335.Dm_build_162(int(dm_build_884), dm_build_880, dm_build_877.dm_build_735.dm_build_336)

		if dm_build_877.dm_build_738.readBaseColName {
			dm_build_885[i].baseName = dm_build_877.dm_build_735.dm_build_335.Dm_build_175(dm_build_880, dm_build_877.dm_build_735.dm_build_336)
		}

		if dm_build_885[i].lob {
			dm_build_885[i].lobTabId = dm_build_877.dm_build_735.dm_build_335.Dm_build_125()
			dm_build_885[i].lobColId = dm_build_877.dm_build_735.dm_build_335.Dm_build_122()
		}

	}

	for i := 0; i < dm_build_878; i++ {

		if isComplexType(int(dm_build_885[i].colType), int(dm_build_885[i].scale)) {
			strDesc := newTypeDescriptor(dm_build_877.dm_build_735.dm_build_336)
			strDesc.unpack(dm_build_877.dm_build_735.dm_build_335)
			dm_build_885[i].typeDescriptor = strDesc
		}
	}

	return dm_build_885
}

func (dm_build_887 *dm_build_826) dm_build_886(dm_build_888 *execInfo, dm_build_889 int, dm_build_890 int, dm_build_891 int) {
	if dm_build_890 > 0 {
		startOffset := dm_build_887.dm_build_735.dm_build_335.Dm_build_20()
		if dm_build_888.rsBdta {
			dm_build_888.rsDatas = dm_build_887.dm_build_905(dm_build_887.dm_build_738.columns, dm_build_891)
		} else {
			datas := make([][][]byte, dm_build_890)

			for i := 0; i < dm_build_890; i++ {

				datas[i] = make([][]byte, dm_build_889+1)

				dm_build_887.dm_build_735.dm_build_335.Dm_build_25(2, false, true)

				datas[i][0] = dm_build_887.dm_build_735.dm_build_335.Dm_build_146(LINT64_SIZE)

				dm_build_887.dm_build_735.dm_build_335.Dm_build_25(2*dm_build_889, false, true)

				for j := 1; j < dm_build_889+1; j++ {

					colLen := dm_build_887.dm_build_735.dm_build_335.Dm_build_140()
					if colLen == Dm_build_650 {
						datas[i][j] = nil
					} else if colLen != Dm_build_651 {
						datas[i][j] = dm_build_887.dm_build_735.dm_build_335.Dm_build_146(int(colLen))
					} else {
						datas[i][j] = dm_build_887.dm_build_735.dm_build_335.Dm_build_150()
					}
				}
			}

			dm_build_888.rsDatas = datas
		}
		dm_build_888.rsSizeof = dm_build_887.dm_build_735.dm_build_335.Dm_build_20() - startOffset
	}

	if dm_build_888.rsCacheOffset > 0 {
		tbCount := dm_build_887.dm_build_735.dm_build_335.Dm_build_122()

		ids := make([]int32, tbCount)
		tss := make([]int64, tbCount)

		for i := 0; i < int(tbCount); i++ {
			ids[i] = dm_build_887.dm_build_735.dm_build_335.Dm_build_125()
			tss[i] = dm_build_887.dm_build_735.dm_build_335.Dm_build_128()
		}

		dm_build_888.tbIds = ids[:]
		dm_build_888.tbTss = tss[:]
	}
}

func (dm_build_893 *dm_build_826) dm_build_892(dm_build_894 []int64) error {

	dm_build_893.dm_build_735.dm_build_335.Dm_build_25(4, false, true)

	dm_build_895 := dm_build_893.dm_build_735.dm_build_335.Dm_build_125()

	dm_build_896 := make([]string, 0, 8)
	for i := 0; i < int(dm_build_895); i++ {
		irow := dm_build_893.dm_build_735.dm_build_335.Dm_build_125()

		dm_build_894[irow] = -3

		code := dm_build_893.dm_build_735.dm_build_335.Dm_build_125()

		errStr := dm_build_893.dm_build_735.dm_build_335.Dm_build_175(dm_build_893.dm_build_735.dm_build_336.getServerEncoding(), dm_build_893.dm_build_735.dm_build_336)

		dm_build_896 = append(dm_build_896, "row["+strconv.Itoa(int(irow))+"]:"+strconv.Itoa(int(code))+", "+errStr)
	}

	if len(dm_build_896) > 0 {
		builder := &strings.Builder{}
		for _, str := range dm_build_896 {
			builder.WriteString(util.LINE_SEPARATOR)
			builder.WriteString(str)
		}
		EC_BP_WITH_ERROR.ErrText += builder.String()
		return EC_BP_WITH_ERROR.throw()
	}
	return nil
}

const (
	Dm_build_897 = 0

	Dm_build_898 = Dm_build_897 + ULINT_SIZE

	Dm_build_899 = Dm_build_898 + USINT_SIZE

	Dm_build_900 = Dm_build_899 + ULINT_SIZE

	Dm_build_901 = Dm_build_900 + ULINT_SIZE

	Dm_build_902 = Dm_build_901 + BYTE_SIZE

	Dm_build_903 = -2

	Dm_build_904 = -3
)

func (dm_build_906 *dm_build_826) dm_build_905(dm_build_907 []column, dm_build_908 int) [][][]byte {

	dm_build_909 := dm_build_906.dm_build_735.dm_build_335.Dm_build_143()

	dm_build_910 := dm_build_906.dm_build_735.dm_build_335.Dm_build_140()

	var dm_build_911 bool

	if dm_build_908 >= 0 && int(dm_build_910) == len(dm_build_907)+1 {
		dm_build_911 = true
	} else {
		dm_build_911 = false
	}

	dm_build_906.dm_build_735.dm_build_335.Dm_build_25(ULINT_SIZE, false, true)

	dm_build_906.dm_build_735.dm_build_335.Dm_build_25(ULINT_SIZE, false, true)

	dm_build_906.dm_build_735.dm_build_335.Dm_build_25(BYTE_SIZE, false, true)

	dm_build_912 := make([]uint16, dm_build_910)
	for icol := 0; icol < int(dm_build_910); icol++ {
		dm_build_912[icol] = dm_build_906.dm_build_735.dm_build_335.Dm_build_140()
	}

	dm_build_913 := make([]uint32, dm_build_910)
	dm_build_914 := make([][][]byte, dm_build_909)

	for i := uint32(0); i < dm_build_909; i++ {
		dm_build_914[i] = make([][]byte, len(dm_build_907)+1)
	}

	for icol := 0; icol < int(dm_build_910); icol++ {
		dm_build_913[icol] = dm_build_906.dm_build_735.dm_build_335.Dm_build_143()
	}

	for icol := 0; icol < int(dm_build_910); icol++ {

		dataCol := icol + 1
		if dm_build_911 && icol == dm_build_908 {
			dataCol = 0
		} else if dm_build_911 && icol > dm_build_908 {
			dataCol = icol
		}

		allNotNull := dm_build_906.dm_build_735.dm_build_335.Dm_build_125() == 1
		var isNull []bool = nil
		if !allNotNull {
			isNull = make([]bool, dm_build_909)
			for irow := uint32(0); irow < dm_build_909; irow++ {
				isNull[irow] = dm_build_906.dm_build_735.dm_build_335.Dm_build_119() == 0
			}
		}

		for irow := uint32(0); irow < dm_build_909; irow++ {
			if allNotNull || !isNull[irow] {
				dm_build_914[irow][dataCol] = dm_build_906.dm_build_915(int(dm_build_912[icol]))
			}
		}
	}

	if !dm_build_911 && dm_build_908 >= 0 {
		for irow := uint32(0); irow < dm_build_909; irow++ {
			dm_build_914[irow][0] = dm_build_914[irow][dm_build_908+1]
		}
	}

	return dm_build_914
}

func (dm_build_916 *dm_build_826) dm_build_915(dm_build_917 int) []byte {

	dm_build_918 := dm_build_916.dm_build_921(dm_build_917)

	dm_build_919 := int32(0)
	if dm_build_918 == Dm_build_903 {
		dm_build_919 = dm_build_916.dm_build_735.dm_build_335.Dm_build_125()
		dm_build_918 = int(dm_build_916.dm_build_735.dm_build_335.Dm_build_125())
	} else if dm_build_918 == Dm_build_904 {
		dm_build_918 = int(dm_build_916.dm_build_735.dm_build_335.Dm_build_125())
	}

	dm_build_920 := dm_build_916.dm_build_735.dm_build_335.Dm_build_146(dm_build_918 + int(dm_build_919))
	if dm_build_919 == 0 {
		return dm_build_920
	}

	for i := dm_build_918; i < len(dm_build_920); i++ {
		dm_build_920[i] = ' '
	}
	return dm_build_920
}

func (dm_build_922 *dm_build_826) dm_build_921(dm_build_923 int) int {

	dm_build_924 := 0
	switch dm_build_923 {
	case INT:
	case BIT:
	case TINYINT:
	case SMALLINT:
	case BOOLEAN:
	case NULL:
		dm_build_924 = 4

	case BIGINT:

		dm_build_924 = 8

	case CHAR:
	case VARCHAR2:
	case VARCHAR:
	case BINARY:
	case VARBINARY:
	case BLOB:
	case CLOB:
		dm_build_924 = Dm_build_903

	case DECIMAL:
		dm_build_924 = Dm_build_904

	case REAL:
		dm_build_924 = 4

	case DOUBLE:
		dm_build_924 = 8

	case DATE:
	case TIME:
	case DATETIME:
	case TIME_TZ:
	case DATETIME_TZ:
		dm_build_924 = 12

	case INTERVAL_YM:
		dm_build_924 = 12

	case INTERVAL_DT:
		dm_build_924 = 24

	default:
		panic(ECGO_INVALID_COLUMN_TYPE)
	}
	return dm_build_924
}

const (
	Dm_build_925 = Dm_build_631

	Dm_build_926 = Dm_build_925 + DDWORD_SIZE

	Dm_build_927 = Dm_build_926 + LINT64_SIZE

	Dm_build_928 = Dm_build_927 + USINT_SIZE

	Dm_build_929 = Dm_build_631

	Dm_build_930 = Dm_build_929 + DDWORD_SIZE
)

type dm_build_931 struct {
	dm_build_826
	dm_build_932 *innerRows
	dm_build_933 int64
	dm_build_934 int64
}

func dm_build_935(dm_build_936 *dm_build_332, dm_build_937 *innerRows, dm_build_938 int64, dm_build_939 int64) *dm_build_931 {
	dm_build_940 := new(dm_build_931)
	dm_build_940.dm_build_743(dm_build_936, Dm_build_609, dm_build_937.dmStmt)
	dm_build_940.dm_build_932 = dm_build_937
	dm_build_940.dm_build_933 = dm_build_938
	dm_build_940.dm_build_934 = dm_build_939
	return dm_build_940
}

func (dm_build_942 *dm_build_931) dm_build_721() error {

	dm_build_942.dm_build_735.dm_build_335.Dm_build_195(Dm_build_925, dm_build_942.dm_build_933)

	dm_build_942.dm_build_735.dm_build_335.Dm_build_195(Dm_build_926, dm_build_942.dm_build_934)

	dm_build_942.dm_build_735.dm_build_335.Dm_build_187(Dm_build_927, dm_build_942.dm_build_932.id)

	dm_build_943 := dm_build_942.dm_build_932.dmStmt.dmConn.dmConnector.bufPrefetch
	var dm_build_944 int32
	if dm_build_942.dm_build_932.sizeOfRow != 0 && dm_build_942.dm_build_932.fetchSize != 0 {
		if dm_build_942.dm_build_932.sizeOfRow*dm_build_942.dm_build_932.fetchSize > int(INT32_MAX) {
			dm_build_944 = INT32_MAX
		} else {
			dm_build_944 = int32(dm_build_942.dm_build_932.sizeOfRow * dm_build_942.dm_build_932.fetchSize)
		}

		if dm_build_944 < Dm_build_662 {
			dm_build_943 = int(Dm_build_662)
		} else if dm_build_944 > Dm_build_663 {
			dm_build_943 = int(Dm_build_663)
		} else {
			dm_build_943 = int(dm_build_944)
		}

		dm_build_942.dm_build_735.dm_build_335.Dm_build_191(Dm_build_928, int32(dm_build_943))
	}

	return nil
}

func (dm_build_946 *dm_build_931) dm_build_725() (interface{}, error) {
	dm_build_947 := execInfo{}
	dm_build_947.rsBdta = dm_build_946.dm_build_932.isBdta

	dm_build_947.updateCount = dm_build_946.dm_build_735.dm_build_335.Dm_build_272(Dm_build_929)
	dm_build_948 := dm_build_946.dm_build_735.dm_build_335.Dm_build_269(Dm_build_930)

	dm_build_946.dm_build_886(&dm_build_947, len(dm_build_946.dm_build_932.columns), int(dm_build_948), -1)

	return &dm_build_947, nil
}

type dm_build_949 struct {
	dm_build_734
	dm_build_950 *lob
	dm_build_951 int
	dm_build_952 int
}

func dm_build_953(dm_build_954 *dm_build_332, dm_build_955 *lob, dm_build_956 int, dm_build_957 int) *dm_build_949 {
	dm_build_958 := new(dm_build_949)
	dm_build_958.dm_build_739(dm_build_954, Dm_build_622)
	dm_build_958.dm_build_950 = dm_build_955
	dm_build_958.dm_build_951 = dm_build_956
	dm_build_958.dm_build_952 = dm_build_957
	return dm_build_958
}

func (dm_build_960 *dm_build_949) dm_build_721() error {

	dm_build_960.dm_build_735.dm_build_335.Dm_build_43(byte(dm_build_960.dm_build_950.lobFlag))

	dm_build_960.dm_build_735.dm_build_335.Dm_build_51(dm_build_960.dm_build_950.tabId)

	dm_build_960.dm_build_735.dm_build_335.Dm_build_47(dm_build_960.dm_build_950.colId)

	dm_build_960.dm_build_735.dm_build_335.Dm_build_67(uint64(dm_build_960.dm_build_950.blobId))

	dm_build_960.dm_build_735.dm_build_335.Dm_build_47(dm_build_960.dm_build_950.groupId)

	dm_build_960.dm_build_735.dm_build_335.Dm_build_47(dm_build_960.dm_build_950.fileId)

	dm_build_960.dm_build_735.dm_build_335.Dm_build_51(dm_build_960.dm_build_950.pageNo)

	dm_build_960.dm_build_735.dm_build_335.Dm_build_47(dm_build_960.dm_build_950.curFileId)

	dm_build_960.dm_build_735.dm_build_335.Dm_build_51(dm_build_960.dm_build_950.curPageNo)

	dm_build_960.dm_build_735.dm_build_335.Dm_build_51(dm_build_960.dm_build_950.totalOffset)

	dm_build_960.dm_build_735.dm_build_335.Dm_build_51(int32(dm_build_960.dm_build_951))

	dm_build_960.dm_build_735.dm_build_335.Dm_build_51(int32(dm_build_960.dm_build_952))

	if dm_build_960.dm_build_735.dm_build_336.NewLobFlag {
		dm_build_960.dm_build_735.dm_build_335.Dm_build_67(uint64(dm_build_960.dm_build_950.rowId))
		dm_build_960.dm_build_735.dm_build_335.Dm_build_47(dm_build_960.dm_build_950.exGroupId)
		dm_build_960.dm_build_735.dm_build_335.Dm_build_47(dm_build_960.dm_build_950.exFileId)
		dm_build_960.dm_build_735.dm_build_335.Dm_build_51(dm_build_960.dm_build_950.exPageNo)
	}

	return nil
}

func (dm_build_962 *dm_build_949) dm_build_725() (interface{}, error) {

	dm_build_962.dm_build_950.readOver = dm_build_962.dm_build_735.dm_build_335.Dm_build_119() == 1
	var dm_build_963 = dm_build_962.dm_build_735.dm_build_335.Dm_build_143()
	if dm_build_963 <= 0 {
		return []byte(nil), nil
	}
	dm_build_962.dm_build_950.curFileId = dm_build_962.dm_build_735.dm_build_335.Dm_build_122()
	dm_build_962.dm_build_950.curPageNo = dm_build_962.dm_build_735.dm_build_335.Dm_build_125()
	dm_build_962.dm_build_950.totalOffset = dm_build_962.dm_build_735.dm_build_335.Dm_build_125()

	return dm_build_962.dm_build_735.dm_build_335.Dm_build_156(int(dm_build_963)), nil
}

type dm_build_964 struct {
	dm_build_734
	dm_build_965 *lob
}

func dm_build_966(dm_build_967 *dm_build_332, dm_build_968 *lob) *dm_build_964 {
	dm_build_969 := new(dm_build_964)
	dm_build_969.dm_build_739(dm_build_967, Dm_build_619)
	dm_build_969.dm_build_965 = dm_build_968
	return dm_build_969
}

func (dm_build_971 *dm_build_964) dm_build_721() error {

	dm_build_971.dm_build_735.dm_build_335.Dm_build_43(byte(dm_build_971.dm_build_965.lobFlag))

	dm_build_971.dm_build_735.dm_build_335.Dm_build_67(uint64(dm_build_971.dm_build_965.blobId))

	dm_build_971.dm_build_735.dm_build_335.Dm_build_47(dm_build_971.dm_build_965.groupId)

	dm_build_971.dm_build_735.dm_build_335.Dm_build_47(dm_build_971.dm_build_965.fileId)

	dm_build_971.dm_build_735.dm_build_335.Dm_build_51(dm_build_971.dm_build_965.pageNo)

	if dm_build_971.dm_build_735.dm_build_336.NewLobFlag {
		dm_build_971.dm_build_735.dm_build_335.Dm_build_51(dm_build_971.dm_build_965.tabId)
		dm_build_971.dm_build_735.dm_build_335.Dm_build_47(dm_build_971.dm_build_965.colId)
		dm_build_971.dm_build_735.dm_build_335.Dm_build_67(uint64(dm_build_971.dm_build_965.rowId))

		dm_build_971.dm_build_735.dm_build_335.Dm_build_47(dm_build_971.dm_build_965.exGroupId)
		dm_build_971.dm_build_735.dm_build_335.Dm_build_47(dm_build_971.dm_build_965.exFileId)
		dm_build_971.dm_build_735.dm_build_335.Dm_build_51(dm_build_971.dm_build_965.exPageNo)
	}

	return nil
}

func (dm_build_973 *dm_build_964) dm_build_725() (interface{}, error) {

	return int64(dm_build_973.dm_build_735.dm_build_335.Dm_build_125()), nil
}

type dm_build_974 struct {
	dm_build_734
	dm_build_975 *lob
	dm_build_976 int
}

func dm_build_977(dm_build_978 *dm_build_332, dm_build_979 *lob, dm_build_980 int) *dm_build_974 {
	dm_build_981 := new(dm_build_974)
	dm_build_981.dm_build_739(dm_build_978, Dm_build_621)
	dm_build_981.dm_build_975 = dm_build_979
	dm_build_981.dm_build_976 = dm_build_980
	return dm_build_981
}

func (dm_build_983 *dm_build_974) dm_build_721() error {

	dm_build_983.dm_build_735.dm_build_335.Dm_build_43(byte(dm_build_983.dm_build_975.lobFlag))

	dm_build_983.dm_build_735.dm_build_335.Dm_build_67(uint64(dm_build_983.dm_build_975.blobId))

	dm_build_983.dm_build_735.dm_build_335.Dm_build_47(dm_build_983.dm_build_975.groupId)

	dm_build_983.dm_build_735.dm_build_335.Dm_build_47(dm_build_983.dm_build_975.fileId)

	dm_build_983.dm_build_735.dm_build_335.Dm_build_51(dm_build_983.dm_build_975.pageNo)

	dm_build_983.dm_build_735.dm_build_335.Dm_build_51(dm_build_983.dm_build_975.tabId)
	dm_build_983.dm_build_735.dm_build_335.Dm_build_47(dm_build_983.dm_build_975.colId)
	dm_build_983.dm_build_735.dm_build_335.Dm_build_67(uint64(dm_build_983.dm_build_975.rowId))
	dm_build_983.dm_build_735.dm_build_335.Dm_build_79(Dm_build_1195.Dm_build_1394(uint32(dm_build_983.dm_build_976)))

	if dm_build_983.dm_build_735.dm_build_336.NewLobFlag {
		dm_build_983.dm_build_735.dm_build_335.Dm_build_47(dm_build_983.dm_build_975.exGroupId)
		dm_build_983.dm_build_735.dm_build_335.Dm_build_47(dm_build_983.dm_build_975.exFileId)
		dm_build_983.dm_build_735.dm_build_335.Dm_build_51(dm_build_983.dm_build_975.exPageNo)
	}
	return nil
}

func (dm_build_985 *dm_build_974) dm_build_725() (interface{}, error) {

	dm_build_986 := dm_build_985.dm_build_735.dm_build_335.Dm_build_143()
	dm_build_985.dm_build_975.blobId = dm_build_985.dm_build_735.dm_build_335.Dm_build_128()
	dm_build_985.dm_build_975.resetCurrentInfo()
	return int64(dm_build_986), nil
}

const (
	Dm_build_987 = Dm_build_631

	Dm_build_988 = Dm_build_987 + ULINT_SIZE

	Dm_build_989 = Dm_build_988 + ULINT_SIZE

	Dm_build_990 = Dm_build_989 + ULINT_SIZE

	Dm_build_991 = Dm_build_990 + BYTE_SIZE

	Dm_build_992 = Dm_build_991 + USINT_SIZE

	Dm_build_993 = Dm_build_992 + ULINT_SIZE

	Dm_build_994 = Dm_build_993 + BYTE_SIZE

	Dm_build_995 = Dm_build_994 + BYTE_SIZE

	Dm_build_996 = Dm_build_995 + BYTE_SIZE

	Dm_build_997 = Dm_build_631

	Dm_build_998 = Dm_build_997 + ULINT_SIZE

	Dm_build_999 = Dm_build_998 + ULINT_SIZE

	Dm_build_1000 = Dm_build_999 + BYTE_SIZE

	Dm_build_1001 = Dm_build_1000 + ULINT_SIZE

	Dm_build_1002 = Dm_build_1001 + BYTE_SIZE

	Dm_build_1003 = Dm_build_1002 + BYTE_SIZE

	Dm_build_1004 = Dm_build_1003 + USINT_SIZE

	Dm_build_1005 = Dm_build_1004 + USINT_SIZE

	Dm_build_1006 = Dm_build_1005 + BYTE_SIZE

	Dm_build_1007 = Dm_build_1006 + USINT_SIZE

	Dm_build_1008 = Dm_build_1007 + BYTE_SIZE

	Dm_build_1009 = Dm_build_1008 + BYTE_SIZE

	Dm_build_1010 = Dm_build_1009 + ULINT_SIZE
)

type dm_build_1011 struct {
	dm_build_734

	dm_build_1012 *DmConnection

	dm_build_1013 bool
}

func dm_build_1014(dm_build_1015 *dm_build_332) *dm_build_1011 {
	dm_build_1016 := new(dm_build_1011)
	dm_build_1016.dm_build_739(dm_build_1015, Dm_build_603)
	dm_build_1016.dm_build_1012 = dm_build_1015.dm_build_336
	return dm_build_1016
}

func (dm_build_1018 *dm_build_1011) dm_build_721() error {

	dm_build_1018.dm_build_735.dm_build_335.Dm_build_191(Dm_build_987, Dm_build_642)

	dm_build_1018.dm_build_735.dm_build_335.Dm_build_191(Dm_build_988, g2dbIsoLevel(dm_build_1018.dm_build_1012.IsoLevel))
	dm_build_1018.dm_build_735.dm_build_335.Dm_build_191(Dm_build_989, int32(Locale))
	dm_build_1018.dm_build_735.dm_build_335.Dm_build_187(Dm_build_991, dm_build_1018.dm_build_1012.dmConnector.localTimezone)

	if dm_build_1018.dm_build_1012.ReadOnly {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_183(Dm_build_990, Dm_build_665)
	} else {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_183(Dm_build_990, Dm_build_664)
	}

	dm_build_1018.dm_build_735.dm_build_335.Dm_build_191(Dm_build_992, int32(dm_build_1018.dm_build_1012.dmConnector.sessionTimeout))

	if dm_build_1018.dm_build_1012.dmConnector.mppLocal {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_183(Dm_build_993, 1)
	} else {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_183(Dm_build_993, 0)
	}

	if dm_build_1018.dm_build_1012.dmConnector.rwSeparate {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_183(Dm_build_994, 1)
	} else {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_183(Dm_build_994, 0)
	}

	if dm_build_1018.dm_build_1012.NewLobFlag {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_183(Dm_build_995, 1)
	} else {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_183(Dm_build_995, 0)
	}

	dm_build_1018.dm_build_735.dm_build_335.Dm_build_183(Dm_build_996, dm_build_1018.dm_build_1012.dmConnector.osAuthType)

	dm_build_1019 := dm_build_1018.dm_build_1012.getServerEncoding()

	if dm_build_1018.dm_build_735.dm_build_341 != "" {

	}

	dm_build_1020 := Dm_build_1195.Dm_build_1405(dm_build_1018.dm_build_1012.dmConnector.user, dm_build_1019, dm_build_1018.dm_build_735.dm_build_336)
	dm_build_1021 := Dm_build_1195.Dm_build_1405(dm_build_1018.dm_build_1012.dmConnector.password, dm_build_1019, dm_build_1018.dm_build_735.dm_build_336)
	if len(dm_build_1020) > Dm_build_639 {
		return ECGO_USERNAME_TOO_LONG.throw()
	}
	if len(dm_build_1021) > Dm_build_639 {
		return ECGO_PASSWORD_TOO_LONG.throw()
	}

	if dm_build_1018.dm_build_735.dm_build_338 && dm_build_1018.dm_build_1012.dmConnector.loginCertificate != "" {

	} else if dm_build_1018.dm_build_735.dm_build_338 {
		dm_build_1020 = dm_build_1018.dm_build_735.dm_build_337.Encrypt(dm_build_1020, false)
		dm_build_1021 = dm_build_1018.dm_build_735.dm_build_337.Encrypt(dm_build_1021, false)
	}

	dm_build_1018.dm_build_735.dm_build_335.Dm_build_83(dm_build_1020)
	dm_build_1018.dm_build_735.dm_build_335.Dm_build_83(dm_build_1021)

	dm_build_1018.dm_build_735.dm_build_335.Dm_build_95(dm_build_1018.dm_build_1012.dmConnector.appName, dm_build_1019, dm_build_1018.dm_build_735.dm_build_336)
	dm_build_1018.dm_build_735.dm_build_335.Dm_build_95(dm_build_1018.dm_build_1012.dmConnector.osName, dm_build_1019, dm_build_1018.dm_build_735.dm_build_336)

	if hostName, err := os.Hostname(); err != nil {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_95(hostName, dm_build_1019, dm_build_1018.dm_build_735.dm_build_336)
	} else {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_95("", dm_build_1019, dm_build_1018.dm_build_735.dm_build_336)
	}

	if dm_build_1018.dm_build_1012.dmConnector.rwStandby {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_43(1)
	} else {
		dm_build_1018.dm_build_735.dm_build_335.Dm_build_43(0)
	}

	return nil
}

func (dm_build_1023 *dm_build_1011) dm_build_725() (interface{}, error) {

	dm_build_1023.dm_build_1012.MaxRowSize = dm_build_1023.dm_build_735.dm_build_335.Dm_build_269(Dm_build_997)
	dm_build_1023.dm_build_1012.DDLAutoCommit = dm_build_1023.dm_build_735.dm_build_335.Dm_build_263(Dm_build_999) == 1
	dm_build_1023.dm_build_1012.IsoLevel = dm_build_1023.dm_build_735.dm_build_335.Dm_build_269(Dm_build_1000)
	dm_build_1023.dm_build_1012.dmConnector.caseSensitive = dm_build_1023.dm_build_735.dm_build_335.Dm_build_263(Dm_build_1001) == 1
	dm_build_1023.dm_build_1012.BackslashEscape = dm_build_1023.dm_build_735.dm_build_335.Dm_build_263(Dm_build_1002) == 1
	dm_build_1023.dm_build_1012.SvrStat = dm_build_1023.dm_build_735.dm_build_335.Dm_build_266(Dm_build_1004)
	dm_build_1023.dm_build_1012.SvrMode = dm_build_1023.dm_build_735.dm_build_335.Dm_build_266(Dm_build_1003)
	dm_build_1023.dm_build_1012.ConstParaOpt = dm_build_1023.dm_build_735.dm_build_335.Dm_build_263(Dm_build_1005) == 1
	dm_build_1023.dm_build_1012.DbTimezone = dm_build_1023.dm_build_735.dm_build_335.Dm_build_266(Dm_build_1006)
	dm_build_1023.dm_build_1012.NewLobFlag = dm_build_1023.dm_build_735.dm_build_335.Dm_build_263(Dm_build_1008) == 1

	if dm_build_1023.dm_build_1012.dmConnector.bufPrefetch == 0 {
		dm_build_1023.dm_build_1012.dmConnector.bufPrefetch = int(dm_build_1023.dm_build_735.dm_build_335.Dm_build_269(Dm_build_1009))
	}

	dm_build_1023.dm_build_1012.LifeTimeRemainder = dm_build_1023.dm_build_735.dm_build_335.Dm_build_266(Dm_build_1010)

	dm_build_1024 := dm_build_1023.dm_build_1012.getServerEncoding()

	dm_build_1023.dm_build_1012.InstanceName = dm_build_1023.dm_build_735.dm_build_335.Dm_build_167(dm_build_1024, dm_build_1023.dm_build_735.dm_build_336)
	dm_build_1023.dm_build_1012.Schema = dm_build_1023.dm_build_735.dm_build_335.Dm_build_167(dm_build_1024, dm_build_1023.dm_build_735.dm_build_336)
	dm_build_1023.dm_build_1012.LastLoginIP = dm_build_1023.dm_build_735.dm_build_335.Dm_build_167(dm_build_1024, dm_build_1023.dm_build_735.dm_build_336)
	dm_build_1023.dm_build_1012.LastLoginTime = dm_build_1023.dm_build_735.dm_build_335.Dm_build_167(dm_build_1024, dm_build_1023.dm_build_735.dm_build_336)
	dm_build_1023.dm_build_1012.FailedAttempts = dm_build_1023.dm_build_735.dm_build_335.Dm_build_125()
	dm_build_1023.dm_build_1012.LoginWarningID = dm_build_1023.dm_build_735.dm_build_335.Dm_build_125()
	dm_build_1023.dm_build_1012.GraceTimeRemainder = dm_build_1023.dm_build_735.dm_build_335.Dm_build_125()
	dm_build_1023.dm_build_1012.Guid = dm_build_1023.dm_build_735.dm_build_335.Dm_build_167(dm_build_1024, dm_build_1023.dm_build_735.dm_build_336)
	dm_build_1023.dm_build_1012.DbName = dm_build_1023.dm_build_735.dm_build_335.Dm_build_167(dm_build_1024, dm_build_1023.dm_build_735.dm_build_336)

	if dm_build_1023.dm_build_735.dm_build_335.Dm_build_263(Dm_build_1007) == 1 {
		dm_build_1023.dm_build_1012.StandbyHost = dm_build_1023.dm_build_735.dm_build_335.Dm_build_167(dm_build_1024, dm_build_1023.dm_build_735.dm_build_336)
		dm_build_1023.dm_build_1012.StandbyPort = dm_build_1023.dm_build_735.dm_build_335.Dm_build_125()
		dm_build_1023.dm_build_1012.StandbyCount = dm_build_1023.dm_build_735.dm_build_335.Dm_build_140()
	}

	if dm_build_1023.dm_build_735.dm_build_335.Dm_build_22(false) > 0 {
		dm_build_1023.dm_build_1012.SessionID = dm_build_1023.dm_build_735.dm_build_335.Dm_build_128()
	}

	if dm_build_1023.dm_build_735.dm_build_335.Dm_build_22(false) > 0 {
		if dm_build_1023.dm_build_735.dm_build_335.Dm_build_119() == 1 {

			dm_build_1023.dm_build_1012.OracleDateFormat = "DD-MON-YY"

			dm_build_1023.dm_build_1012.OracleTimestampFormat = "DD-MON-YY HH12.MI.SS.FF6 AM"

			dm_build_1023.dm_build_1012.OracleTimestampTZFormat = "DD-MON-YY HH12.MI.SS.FF6 AM +TZH:TZM"

			dm_build_1023.dm_build_1012.OracleTimeFormat = "HH12.MI.SS.FF6 AM"

			dm_build_1023.dm_build_1012.OracleTimeTZFormat = "HH12.MI.SS.FF6 AM +TZH:TZM"
		}
	}

	return nil, nil
}

const (
	Dm_build_1025 = Dm_build_631
)

type dm_build_1026 struct {
	dm_build_826
	dm_build_1027 int16
}

func dm_build_1028(dm_build_1029 *dm_build_332, dm_build_1030 *DmStatement, dm_build_1031 int16) *dm_build_1026 {
	dm_build_1032 := new(dm_build_1026)
	dm_build_1032.dm_build_743(dm_build_1029, Dm_build_623, dm_build_1030)
	dm_build_1032.dm_build_1027 = dm_build_1031
	return dm_build_1032
}

func (dm_build_1034 *dm_build_1026) dm_build_721() error {
	dm_build_1034.dm_build_735.dm_build_335.Dm_build_187(Dm_build_1025, dm_build_1034.dm_build_1027)
	return nil
}

func (dm_build_1036 *dm_build_1026) dm_build_725() (interface{}, error) {
	return dm_build_1036.dm_build_826.dm_build_725()
}

const (
	Dm_build_1037 = Dm_build_631
)

type dm_build_1038 struct {
	dm_build_826
	dm_build_1039 []parameter
}

func dm_build_1040(dm_build_1041 *dm_build_332, dm_build_1042 *DmStatement, dm_build_1043 []parameter) *dm_build_1038 {
	dm_build_1044 := new(dm_build_1038)
	dm_build_1044.dm_build_743(dm_build_1041, Dm_build_627, dm_build_1042)
	dm_build_1044.dm_build_1039 = dm_build_1043
	return dm_build_1044
}

func (dm_build_1046 *dm_build_1038) dm_build_721() error {

	if dm_build_1046.dm_build_1039 == nil {
		dm_build_1046.dm_build_735.dm_build_335.Dm_build_187(Dm_build_1037, 0)
	} else {
		dm_build_1046.dm_build_735.dm_build_335.Dm_build_187(Dm_build_1037, int16(len(dm_build_1046.dm_build_1039)))
	}

	return dm_build_1046.dm_build_849(dm_build_1046.dm_build_1039)
}

type dm_build_1047 struct {
	dm_build_826
	dm_build_1048 bool
	dm_build_1049 int16
}

func dm_build_1050(dm_build_1051 *dm_build_332, dm_build_1052 *DmStatement, dm_build_1053 bool, dm_build_1054 int16) *dm_build_1047 {
	dm_build_1055 := new(dm_build_1047)
	dm_build_1055.dm_build_743(dm_build_1051, Dm_build_607, dm_build_1052)
	dm_build_1055.dm_build_1048 = dm_build_1053
	dm_build_1055.dm_build_1049 = dm_build_1054
	return dm_build_1055
}

func (dm_build_1057 *dm_build_1047) dm_build_721() error {

	dm_build_1058 := Dm_build_631

	if dm_build_1057.dm_build_735.dm_build_336.autoCommit {
		dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 1)
	} else {
		dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 0)
	}

	if dm_build_1057.dm_build_1048 {
		dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 1)
	} else {
		dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 0)
	}

	dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 0)

	dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 1)

	if dm_build_1057.dm_build_735.dm_build_336.CompatibleOracle() {
		dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 0)
	} else {
		dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 1)
	}

	dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_187(dm_build_1058, dm_build_1057.dm_build_1049)

	if dm_build_1057.dm_build_738.maxRows <= 0 || dm_build_1057.dm_build_735.dm_build_336.dmConnector.enRsCache {
		dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_195(dm_build_1058, INT64_MAX)
	} else {
		dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_195(dm_build_1058, dm_build_1057.dm_build_738.maxRows)
	}

	if dm_build_1057.dm_build_735.dm_build_336.dmConnector.isBdtaRS {
		dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, Dm_build_709)
	} else {
		dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, Dm_build_708)
	}

	dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_187(dm_build_1058, 0)

	dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 1)

	dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 0)

	dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_183(dm_build_1058, 0)

	dm_build_1058 += dm_build_1057.dm_build_735.dm_build_335.Dm_build_191(dm_build_1058, dm_build_1057.dm_build_738.queryTimeout)

	dm_build_1057.dm_build_735.dm_build_335.Dm_build_113(dm_build_1057.dm_build_738.nativeSql, dm_build_1057.dm_build_735.dm_build_336.getServerEncoding(), dm_build_1057.dm_build_735.dm_build_336)

	return nil
}

func (dm_build_1060 *dm_build_1047) dm_build_725() (interface{}, error) {

	if dm_build_1060.dm_build_1048 {
		return dm_build_1060.dm_build_826.dm_build_725()
	}

	dm_build_1061 := NewExceInfo()
	dm_build_1062 := Dm_build_631

	dm_build_1061.retSqlType = dm_build_1060.dm_build_735.dm_build_335.Dm_build_266(dm_build_1062)
	dm_build_1062 += USINT_SIZE

	dm_build_1063 := dm_build_1060.dm_build_735.dm_build_335.Dm_build_266(dm_build_1062)
	dm_build_1062 += USINT_SIZE

	dm_build_1064 := dm_build_1060.dm_build_735.dm_build_335.Dm_build_266(dm_build_1062)
	dm_build_1062 += USINT_SIZE

	dm_build_1060.dm_build_735.dm_build_335.Dm_build_272(dm_build_1062)
	dm_build_1062 += DDWORD_SIZE

	dm_build_1060.dm_build_735.dm_build_336.TrxStatus = dm_build_1060.dm_build_735.dm_build_335.Dm_build_269(dm_build_1062)
	dm_build_1062 += ULINT_SIZE

	if dm_build_1063 > 0 {
		dm_build_1060.dm_build_738.params = dm_build_1060.dm_build_1065(int(dm_build_1063))
		dm_build_1060.dm_build_738.paramCount = dm_build_1063
	} else {
		dm_build_1060.dm_build_738.params = make([]parameter, 0)
		dm_build_1060.dm_build_738.paramCount = 0
	}

	if dm_build_1064 > 0 {
		dm_build_1060.dm_build_738.columns = dm_build_1060.dm_build_876(int(dm_build_1064), dm_build_1061.rsBdta)
	} else {
		dm_build_1060.dm_build_738.columns = make([]column, 0)
	}

	return dm_build_1061, nil
}

func (dm_build_1066 *dm_build_1047) dm_build_1065(dm_build_1067 int) []parameter {

	var dm_build_1068, dm_build_1069, dm_build_1070, dm_build_1071 int16

	dm_build_1072 := make([]parameter, dm_build_1067)
	for i := 0; i < dm_build_1067; i++ {

		dm_build_1072[i].InitParameter()

		dm_build_1072[i].colType = dm_build_1066.dm_build_735.dm_build_335.Dm_build_125()

		dm_build_1072[i].prec = dm_build_1066.dm_build_735.dm_build_335.Dm_build_125()

		dm_build_1072[i].scale = dm_build_1066.dm_build_735.dm_build_335.Dm_build_125()

		dm_build_1072[i].nullable = dm_build_1066.dm_build_735.dm_build_335.Dm_build_125() != 0

		itemFlag := dm_build_1066.dm_build_735.dm_build_335.Dm_build_122()

		if int(itemFlag)&Dm_build_825 != 0 {
			dm_build_1072[i].typeFlag = TYPE_FLAG_RECOMMEND
		} else {
			dm_build_1072[i].typeFlag = TYPE_FLAG_EXACT
		}

		dm_build_1072[i].lob = int(itemFlag)&Dm_build_823 != 0

		dm_build_1066.dm_build_735.dm_build_335.Dm_build_125()

		dm_build_1072[i].ioType = byte(dm_build_1066.dm_build_735.dm_build_335.Dm_build_122())

		dm_build_1068 = dm_build_1066.dm_build_735.dm_build_335.Dm_build_122()

		dm_build_1069 = dm_build_1066.dm_build_735.dm_build_335.Dm_build_122()

		dm_build_1070 = dm_build_1066.dm_build_735.dm_build_335.Dm_build_122()

		dm_build_1071 = dm_build_1066.dm_build_735.dm_build_335.Dm_build_122()
		dm_build_1072[i].name = dm_build_1066.dm_build_735.dm_build_335.Dm_build_162(int(dm_build_1068), dm_build_1066.dm_build_735.dm_build_336.getServerEncoding(), dm_build_1066.dm_build_735.dm_build_336)
		dm_build_1072[i].typeName = dm_build_1066.dm_build_735.dm_build_335.Dm_build_162(int(dm_build_1069), dm_build_1066.dm_build_735.dm_build_336.getServerEncoding(), dm_build_1066.dm_build_735.dm_build_336)
		dm_build_1072[i].tableName = dm_build_1066.dm_build_735.dm_build_335.Dm_build_162(int(dm_build_1070), dm_build_1066.dm_build_735.dm_build_336.getServerEncoding(), dm_build_1066.dm_build_735.dm_build_336)
		dm_build_1072[i].schemaName = dm_build_1066.dm_build_735.dm_build_335.Dm_build_162(int(dm_build_1071), dm_build_1066.dm_build_735.dm_build_336.getServerEncoding(), dm_build_1066.dm_build_735.dm_build_336)

		if dm_build_1072[i].lob {
			dm_build_1072[i].lobTabId = dm_build_1066.dm_build_735.dm_build_335.Dm_build_125()
			dm_build_1072[i].lobColId = dm_build_1066.dm_build_735.dm_build_335.Dm_build_122()
		}
	}

	for i := 0; i < dm_build_1067; i++ {

		if isComplexType(int(dm_build_1072[i].colType), int(dm_build_1072[i].scale)) {

			strDesc := newTypeDescriptor(dm_build_1066.dm_build_735.dm_build_336)
			strDesc.unpack(dm_build_1066.dm_build_735.dm_build_335)
			dm_build_1072[i].typeDescriptor = strDesc
		}
	}

	return dm_build_1072
}

const (
	Dm_build_1073 = Dm_build_631
)

type dm_build_1074 struct {
	dm_build_734
	dm_build_1075 int16
	dm_build_1076 *Dm_build_1471
	dm_build_1077 int32
}

func dm_build_1078(dm_build_1079 *dm_build_332, dm_build_1080 *DmStatement, dm_build_1081 int16, dm_build_1082 *Dm_build_1471, dm_build_1083 int32) *dm_build_1074 {
	dm_build_1084 := new(dm_build_1074)
	dm_build_1084.dm_build_743(dm_build_1079, Dm_build_613, dm_build_1080)
	dm_build_1084.dm_build_1075 = dm_build_1081
	dm_build_1084.dm_build_1076 = dm_build_1082
	dm_build_1084.dm_build_1077 = dm_build_1083
	return dm_build_1084
}

func (dm_build_1086 *dm_build_1074) dm_build_721() error {
	dm_build_1086.dm_build_735.dm_build_335.Dm_build_187(Dm_build_1073, dm_build_1086.dm_build_1075)

	dm_build_1086.dm_build_735.dm_build_335.Dm_build_51(dm_build_1086.dm_build_1077)

	if dm_build_1086.dm_build_735.dm_build_336.NewLobFlag {
		dm_build_1086.dm_build_735.dm_build_335.Dm_build_51(-1)
	}
	dm_build_1086.dm_build_1076.Dm_build_1478(dm_build_1086.dm_build_735.dm_build_335, int(dm_build_1086.dm_build_1077))
	return nil
}

type dm_build_1087 struct {
	dm_build_734
}

func dm_build_1088(dm_build_1089 *dm_build_332) *dm_build_1087 {
	dm_build_1090 := new(dm_build_1087)
	dm_build_1090.dm_build_739(dm_build_1089, Dm_build_611)
	return dm_build_1090
}

type dm_build_1091 struct {
	dm_build_734
	dm_build_1092 int32
}

func dm_build_1093(dm_build_1094 *dm_build_332, dm_build_1095 int32) *dm_build_1091 {
	dm_build_1096 := new(dm_build_1091)
	dm_build_1096.dm_build_739(dm_build_1094, Dm_build_624)
	dm_build_1096.dm_build_1092 = dm_build_1095
	return dm_build_1096
}

func (dm_build_1098 *dm_build_1091) dm_build_721() error {

	dm_build_1099 := Dm_build_631
	dm_build_1099 += dm_build_1098.dm_build_735.dm_build_335.Dm_build_191(dm_build_1099, g2dbIsoLevel(dm_build_1098.dm_build_1092))
	return nil
}

type dm_build_1100 struct {
	dm_build_734
	dm_build_1101 *lob
	dm_build_1102 byte
	dm_build_1103 int
	dm_build_1104 []byte
	dm_build_1105 int
	dm_build_1106 int
}

func dm_build_1107(dm_build_1108 *dm_build_332, dm_build_1109 *lob, dm_build_1110 byte, dm_build_1111 int, dm_build_1112 []byte,
	dm_build_1113 int, dm_build_1114 int) *dm_build_1100 {
	dm_build_1115 := new(dm_build_1100)
	dm_build_1115.dm_build_739(dm_build_1108, Dm_build_620)
	dm_build_1115.dm_build_1101 = dm_build_1109
	dm_build_1115.dm_build_1102 = dm_build_1110
	dm_build_1115.dm_build_1103 = dm_build_1111
	dm_build_1115.dm_build_1104 = dm_build_1112
	dm_build_1115.dm_build_1105 = dm_build_1113
	dm_build_1115.dm_build_1106 = dm_build_1114
	return dm_build_1115
}

func (dm_build_1117 *dm_build_1100) dm_build_721() error {

	dm_build_1117.dm_build_735.dm_build_335.Dm_build_43(byte(dm_build_1117.dm_build_1101.lobFlag))
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_43(dm_build_1117.dm_build_1102)
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_67(uint64(dm_build_1117.dm_build_1101.blobId))
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_47(dm_build_1117.dm_build_1101.groupId)
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_47(dm_build_1117.dm_build_1101.fileId)
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_51(dm_build_1117.dm_build_1101.pageNo)
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_47(dm_build_1117.dm_build_1101.curFileId)
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_51(dm_build_1117.dm_build_1101.curPageNo)
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_51(dm_build_1117.dm_build_1101.totalOffset)
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_51(dm_build_1117.dm_build_1101.tabId)
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_47(dm_build_1117.dm_build_1101.colId)
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_67(uint64(dm_build_1117.dm_build_1101.rowId))

	dm_build_1117.dm_build_735.dm_build_335.Dm_build_51(int32(dm_build_1117.dm_build_1103))
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_51(int32(dm_build_1117.dm_build_1106))
	dm_build_1117.dm_build_735.dm_build_335.Dm_build_79(dm_build_1117.dm_build_1104)

	if dm_build_1117.dm_build_735.dm_build_336.NewLobFlag {
		dm_build_1117.dm_build_735.dm_build_335.Dm_build_47(dm_build_1117.dm_build_1101.exGroupId)
		dm_build_1117.dm_build_735.dm_build_335.Dm_build_47(dm_build_1117.dm_build_1101.exFileId)
		dm_build_1117.dm_build_735.dm_build_335.Dm_build_51(dm_build_1117.dm_build_1101.exPageNo)
	}
	return nil
}

func (dm_build_1119 *dm_build_1100) dm_build_725() (interface{}, error) {

	var dm_build_1120 = dm_build_1119.dm_build_735.dm_build_335.Dm_build_125()
	dm_build_1119.dm_build_1101.blobId = dm_build_1119.dm_build_735.dm_build_335.Dm_build_128()
	dm_build_1119.dm_build_1101.fileId = dm_build_1119.dm_build_735.dm_build_335.Dm_build_122()
	dm_build_1119.dm_build_1101.pageNo = dm_build_1119.dm_build_735.dm_build_335.Dm_build_125()
	dm_build_1119.dm_build_1101.curFileId = dm_build_1119.dm_build_735.dm_build_335.Dm_build_122()
	dm_build_1119.dm_build_1101.curPageNo = dm_build_1119.dm_build_735.dm_build_335.Dm_build_125()
	dm_build_1119.dm_build_1101.totalOffset = dm_build_1119.dm_build_735.dm_build_335.Dm_build_125()
	return dm_build_1120, nil
}

const (
	Dm_build_1121 = Dm_build_631

	Dm_build_1122 = Dm_build_1121 + ULINT_SIZE

	Dm_build_1123 = Dm_build_1122 + ULINT_SIZE

	Dm_build_1124 = Dm_build_1123 + BYTE_SIZE

	Dm_build_1125 = Dm_build_1124 + BYTE_SIZE

	Dm_build_1126 = Dm_build_1125 + BYTE_SIZE

	Dm_build_1127 = Dm_build_1126 + BYTE_SIZE

	Dm_build_1128 = Dm_build_631

	Dm_build_1129 = Dm_build_1128 + ULINT_SIZE

	Dm_build_1130 = Dm_build_1129 + ULINT_SIZE

	Dm_build_1131 = Dm_build_1130 + ULINT_SIZE

	Dm_build_1132 = Dm_build_1131 + ULINT_SIZE

	Dm_build_1133 = Dm_build_1132 + ULINT_SIZE

	Dm_build_1134 = Dm_build_1133 + BYTE_SIZE

	Dm_build_1135 = Dm_build_1134 + BYTE_SIZE

	Dm_build_1136 = Dm_build_1135 + BYTE_SIZE
)

type dm_build_1137 struct {
	dm_build_734
	dm_build_1138 *DmConnection
	dm_build_1139 int
	Dm_build_1140 int32
	Dm_build_1141 []byte
	dm_build_1142 byte
}

func dm_build_1143(dm_build_1144 *dm_build_332) *dm_build_1137 {
	dm_build_1145 := new(dm_build_1137)
	dm_build_1145.dm_build_739(dm_build_1144, Dm_build_629)
	dm_build_1145.dm_build_1138 = dm_build_1144.dm_build_336
	return dm_build_1145
}

func dm_build_1146(dm_build_1147 string, dm_build_1148 string) int {
	dm_build_1149 := strings.Split(dm_build_1147, ".")
	dm_build_1150 := strings.Split(dm_build_1148, ".")

	for i, serStr := range dm_build_1149 {
		ser, _ := strconv.ParseInt(serStr, 10, 32)
		global, _ := strconv.ParseInt(dm_build_1150[i], 10, 32)
		if ser < global {
			return -1
		} else if ser == global {
			continue
		} else {
			return 1
		}
	}

	return 0
}

func (dm_build_1152 *dm_build_1137) dm_build_721() error {

	dm_build_1152.dm_build_735.dm_build_335.Dm_build_191(Dm_build_1121, int32(0))
	dm_build_1152.dm_build_735.dm_build_335.Dm_build_191(Dm_build_1122, int32(dm_build_1152.dm_build_1138.dmConnector.compress))

	if dm_build_1152.dm_build_1138.dmConnector.loginEncrypt {
		dm_build_1152.dm_build_735.dm_build_335.Dm_build_183(Dm_build_1124, 2)
		dm_build_1152.dm_build_735.dm_build_335.Dm_build_183(Dm_build_1123, 1)
	} else {
		dm_build_1152.dm_build_735.dm_build_335.Dm_build_183(Dm_build_1124, 0)
		dm_build_1152.dm_build_735.dm_build_335.Dm_build_183(Dm_build_1123, 0)
	}

	if dm_build_1152.dm_build_1138.dmConnector.isBdtaRS {
		dm_build_1152.dm_build_735.dm_build_335.Dm_build_183(Dm_build_1125, Dm_build_709)
	} else {
		dm_build_1152.dm_build_735.dm_build_335.Dm_build_183(Dm_build_1125, Dm_build_708)
	}

	dm_build_1152.dm_build_735.dm_build_335.Dm_build_183(Dm_build_1126, byte(dm_build_1152.dm_build_1138.dmConnector.compressID))

	if dm_build_1152.dm_build_1138.dmConnector.loginCertificate != "" {
		dm_build_1152.dm_build_735.dm_build_335.Dm_build_183(Dm_build_1127, 1)
	} else {
		dm_build_1152.dm_build_735.dm_build_335.Dm_build_183(Dm_build_1127, 0)
	}

	dm_build_1153 := dm_build_1152.dm_build_1138.getServerEncoding()
	dm_build_1152.dm_build_735.dm_build_335.Dm_build_95(Dm_build_596, dm_build_1153, dm_build_1152.dm_build_735.dm_build_336)

	var dm_build_1154 byte
	if dm_build_1152.dm_build_1138.dmConnector.uKeyName != "" {
		dm_build_1154 = 1
	} else {
		dm_build_1154 = 0
	}

	dm_build_1152.dm_build_735.dm_build_335.Dm_build_43(0)

	if dm_build_1154 == 1 {

	}

	if dm_build_1152.dm_build_1138.dmConnector.loginEncrypt {
		clientPubKey, err := dm_build_1152.dm_build_735.dm_build_576()
		if err != nil {
			return err
		}
		dm_build_1152.dm_build_735.dm_build_335.Dm_build_83(clientPubKey)
	}

	return nil
}

func (dm_build_1156 *dm_build_1137) dm_build_724() error {
	dm_build_1156.dm_build_735.dm_build_335.Dm_build_17(0)
	dm_build_1156.dm_build_735.dm_build_335.Dm_build_25(Dm_build_630, false, true)
	return nil
}

func (dm_build_1158 *dm_build_1137) dm_build_725() (interface{}, error) {

	dm_build_1158.dm_build_1138.sslEncrypt = int(dm_build_1158.dm_build_735.dm_build_335.Dm_build_269(Dm_build_1128))
	dm_build_1158.dm_build_1138.GlobalServerSeries = int(dm_build_1158.dm_build_735.dm_build_335.Dm_build_269(Dm_build_1129))

	switch dm_build_1158.dm_build_735.dm_build_335.Dm_build_269(Dm_build_1130) {
	case 1:
		dm_build_1158.dm_build_1138.serverEncoding = ENCODING_UTF8
	case 2:
		dm_build_1158.dm_build_1138.serverEncoding = ENCODING_EUCKR
	default:
		dm_build_1158.dm_build_1138.serverEncoding = ENCODING_GB18030
	}

	dm_build_1158.dm_build_1138.dmConnector.compress = int(dm_build_1158.dm_build_735.dm_build_335.Dm_build_269(Dm_build_1131))
	dm_build_1159 := dm_build_1158.dm_build_735.dm_build_335.Dm_build_263(Dm_build_1133)
	dm_build_1160 := dm_build_1158.dm_build_735.dm_build_335.Dm_build_263(Dm_build_1134)
	dm_build_1158.dm_build_1138.dmConnector.isBdtaRS = dm_build_1158.dm_build_735.dm_build_335.Dm_build_263(Dm_build_1135) > 0
	dm_build_1158.dm_build_1138.dmConnector.compressID = int8(dm_build_1158.dm_build_735.dm_build_335.Dm_build_263(Dm_build_1136))

	dm_build_1161 := dm_build_1158.dm_build_767()
	if dm_build_1161 != nil {
		return nil, dm_build_1161
	}

	dm_build_1162 := dm_build_1158.dm_build_735.dm_build_335.Dm_build_167(dm_build_1158.dm_build_1138.getServerEncoding(), dm_build_1158.dm_build_735.dm_build_336)
	if dm_build_1146(dm_build_1162, Dm_build_597) < 0 {
		return nil, ECGO_ERROR_SERVER_VERSION.throw()
	}

	dm_build_1158.dm_build_1138.ServerVersion = dm_build_1162
	dm_build_1158.dm_build_1138.Malini2 = dm_build_1146(dm_build_1162, Dm_build_598) > 0
	dm_build_1158.dm_build_1138.Execute2 = dm_build_1146(dm_build_1162, Dm_build_599) > 0
	dm_build_1158.dm_build_1138.LobEmptyCompOrcl = dm_build_1146(dm_build_1162, Dm_build_600) > 0

	if dm_build_1158.dm_build_735.dm_build_336.dmConnector.uKeyName != "" {
		dm_build_1158.dm_build_1142 = 1
	} else {
		dm_build_1158.dm_build_1142 = 0
	}

	if dm_build_1158.dm_build_1142 == 1 {
		dm_build_1158.dm_build_735.dm_build_341 = dm_build_1158.dm_build_735.dm_build_335.Dm_build_162(16, dm_build_1158.dm_build_1138.getServerEncoding(), dm_build_1158.dm_build_735.dm_build_336)
	}

	dm_build_1158.dm_build_1139 = -1
	dm_build_1163 := false
	dm_build_1164 := false
	dm_build_1158.Dm_build_1140 = -1
	if dm_build_1160 > 0 {
		dm_build_1158.dm_build_1139 = int(dm_build_1158.dm_build_735.dm_build_335.Dm_build_125())
	}

	if dm_build_1159 > 0 {

		if dm_build_1158.dm_build_1139 == -1 {
			dm_build_1163 = true
		} else {
			dm_build_1164 = true
		}

		dm_build_1158.Dm_build_1141 = dm_build_1158.dm_build_735.dm_build_335.Dm_build_150()
	}

	if dm_build_1160 == 2 {
		dm_build_1158.Dm_build_1140 = dm_build_1158.dm_build_735.dm_build_335.Dm_build_125()
	}
	dm_build_1158.dm_build_735.dm_build_338 = dm_build_1163
	dm_build_1158.dm_build_735.dm_build_339 = dm_build_1164

	return nil, nil
}

type dm_build_1165 struct {
	dm_build_734
}

func dm_build_1166(dm_build_1167 *dm_build_332, dm_build_1168 *DmStatement) *dm_build_1165 {
	dm_build_1169 := new(dm_build_1165)
	dm_build_1169.dm_build_743(dm_build_1167, Dm_build_605, dm_build_1168)
	return dm_build_1169
}

func (dm_build_1171 *dm_build_1165) dm_build_721() error {

	dm_build_1171.dm_build_735.dm_build_335.Dm_build_183(Dm_build_631, 1)
	return nil
}

func (dm_build_1173 *dm_build_1165) dm_build_725() (interface{}, error) {

	dm_build_1173.dm_build_738.id = dm_build_1173.dm_build_735.dm_build_335.Dm_build_269(Dm_build_632)

	dm_build_1173.dm_build_738.readBaseColName = dm_build_1173.dm_build_735.dm_build_335.Dm_build_263(Dm_build_631) == 1
	return nil, nil
}

type dm_build_1174 struct {
	dm_build_734
	dm_build_1175 int32
}

func dm_build_1176(dm_build_1177 *dm_build_332, dm_build_1178 int32) *dm_build_1174 {
	dm_build_1179 := new(dm_build_1174)
	dm_build_1179.dm_build_739(dm_build_1177, Dm_build_606)
	dm_build_1179.dm_build_1175 = dm_build_1178
	return dm_build_1179
}

func (dm_build_1181 *dm_build_1174) dm_build_722() {
	dm_build_1181.dm_build_734.dm_build_722()
	dm_build_1181.dm_build_735.dm_build_335.Dm_build_191(Dm_build_632, dm_build_1181.dm_build_1175)
}

type dm_build_1182 struct {
	dm_build_734
	dm_build_1183 []uint32
}

func dm_build_1184(dm_build_1185 *dm_build_332, dm_build_1186 []uint32) *dm_build_1182 {
	dm_build_1187 := new(dm_build_1182)
	dm_build_1187.dm_build_739(dm_build_1185, Dm_build_626)
	dm_build_1187.dm_build_1183 = dm_build_1186
	return dm_build_1187
}

func (dm_build_1189 *dm_build_1182) dm_build_721() error {

	dm_build_1189.dm_build_735.dm_build_335.Dm_build_211(Dm_build_631, uint16(len(dm_build_1189.dm_build_1183)))

	for _, tableID := range dm_build_1189.dm_build_1183 {
		dm_build_1189.dm_build_735.dm_build_335.Dm_build_63(uint32(tableID))
	}

	return nil
}

func (dm_build_1191 *dm_build_1182) dm_build_725() (interface{}, error) {
	dm_build_1192 := dm_build_1191.dm_build_735.dm_build_335.Dm_build_284(Dm_build_631)
	if dm_build_1192 <= 0 {
		return nil, nil
	}

	dm_build_1193 := make([]int64, dm_build_1192)
	for i := 0; i < int(dm_build_1192); i++ {
		dm_build_1193[i] = dm_build_1191.dm_build_735.dm_build_335.Dm_build_128()
	}

	return dm_build_1193, nil
}
