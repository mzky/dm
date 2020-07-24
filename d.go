/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1471 struct {
	dm_build_1472 *list.List
	dm_build_1473 *dm_build_1525
	dm_build_1474 int
}

func Dm_build_1475() *Dm_build_1471 {
	return &Dm_build_1471{
		dm_build_1472: list.New(),
		dm_build_1474: 0,
	}
}

func (dm_build_1477 *Dm_build_1471) Dm_build_1476() int {
	return dm_build_1477.dm_build_1474
}

func (dm_build_1479 *Dm_build_1471) Dm_build_1478(dm_build_1480 *Dm_build_0, dm_build_1481 int) int {
	var dm_build_1482 = 0
	var dm_build_1483 = 0
	for dm_build_1482 < dm_build_1481 && dm_build_1479.dm_build_1473 != nil {
		dm_build_1483 = dm_build_1479.dm_build_1473.dm_build_1533(dm_build_1480, dm_build_1481-dm_build_1482)
		if dm_build_1479.dm_build_1473.dm_build_1528 == 0 {
			dm_build_1479.dm_build_1515()
		}
		dm_build_1482 += dm_build_1483
		dm_build_1479.dm_build_1474 -= dm_build_1483
	}
	return dm_build_1482
}

func (dm_build_1485 *Dm_build_1471) Dm_build_1484(dm_build_1486 []byte, dm_build_1487 int, dm_build_1488 int) int {
	var dm_build_1489 = 0
	var dm_build_1490 = 0
	for dm_build_1489 < dm_build_1488 && dm_build_1485.dm_build_1473 != nil {
		dm_build_1490 = dm_build_1485.dm_build_1473.dm_build_1537(dm_build_1486, dm_build_1487, dm_build_1488-dm_build_1489)
		if dm_build_1485.dm_build_1473.dm_build_1528 == 0 {
			dm_build_1485.dm_build_1515()
		}
		dm_build_1489 += dm_build_1490
		dm_build_1485.dm_build_1474 -= dm_build_1490
		dm_build_1487 += dm_build_1490
	}
	return dm_build_1489
}

func (dm_build_1492 *Dm_build_1471) Dm_build_1491(dm_build_1493 io.Writer, dm_build_1494 int) int {
	var dm_build_1495 = 0
	var dm_build_1496 = 0
	for dm_build_1495 < dm_build_1494 && dm_build_1492.dm_build_1473 != nil {
		dm_build_1496 = dm_build_1492.dm_build_1473.dm_build_1542(dm_build_1493, dm_build_1494-dm_build_1495)
		if dm_build_1492.dm_build_1473.dm_build_1528 == 0 {
			dm_build_1492.dm_build_1515()
		}
		dm_build_1495 += dm_build_1496
		dm_build_1492.dm_build_1474 -= dm_build_1496
	}
	return dm_build_1495
}

func (dm_build_1498 *Dm_build_1471) Dm_build_1497(dm_build_1499 []byte, dm_build_1500 int, dm_build_1501 int) {
	if dm_build_1501 == 0 {
		return
	}
	var dm_build_1502 = dm_build_1529(dm_build_1499, dm_build_1500, dm_build_1501)
	if dm_build_1498.dm_build_1473 == nil {
		dm_build_1498.dm_build_1473 = dm_build_1502
	} else {
		dm_build_1498.dm_build_1472.PushBack(dm_build_1502)
	}
	dm_build_1498.dm_build_1474 += dm_build_1501
}

func (dm_build_1504 *Dm_build_1471) dm_build_1503(dm_build_1505 int) byte {
	var dm_build_1506 = dm_build_1505
	var dm_build_1507 = dm_build_1504.dm_build_1473
	for dm_build_1506 > 0 && dm_build_1507 != nil {
		if dm_build_1507.dm_build_1528 == 0 {
			continue
		}
		if dm_build_1506 > dm_build_1507.dm_build_1528-1 {
			dm_build_1506 -= dm_build_1507.dm_build_1528
			dm_build_1507 = dm_build_1504.dm_build_1472.Front().Value.(*dm_build_1525)
		} else {
			break
		}
	}
	return dm_build_1507.dm_build_1546(dm_build_1506)
}
func (dm_build_1509 *Dm_build_1471) Dm_build_1508(dm_build_1510 *Dm_build_1471) {
	if dm_build_1510.dm_build_1474 == 0 {
		return
	}
	var dm_build_1511 = dm_build_1510.dm_build_1473
	for dm_build_1511 != nil {
		dm_build_1509.dm_build_1512(dm_build_1511)
		dm_build_1510.dm_build_1515()
		dm_build_1511 = dm_build_1510.dm_build_1473
	}
	dm_build_1510.dm_build_1474 = 0
}
func (dm_build_1513 *Dm_build_1471) dm_build_1512(dm_build_1514 *dm_build_1525) {
	if dm_build_1514.dm_build_1528 == 0 {
		return
	}
	if dm_build_1513.dm_build_1473 == nil {
		dm_build_1513.dm_build_1473 = dm_build_1514
	} else {
		dm_build_1513.dm_build_1472.PushBack(dm_build_1514)
	}
	dm_build_1513.dm_build_1474 += dm_build_1514.dm_build_1528
}

func (dm_build_1516 *Dm_build_1471) dm_build_1515() {
	var dm_build_1517 = dm_build_1516.dm_build_1472.Front()
	if dm_build_1517 == nil {
		dm_build_1516.dm_build_1473 = nil
	} else {
		dm_build_1516.dm_build_1473 = dm_build_1517.Value.(*dm_build_1525)
		dm_build_1516.dm_build_1472.Remove(dm_build_1517)
	}
}

func (dm_build_1519 *Dm_build_1471) Dm_build_1518() []byte {
	var dm_build_1520 = make([]byte, dm_build_1519.dm_build_1474)
	var dm_build_1521 = dm_build_1519.dm_build_1473
	var dm_build_1522 = 0
	var dm_build_1523 = len(dm_build_1520)
	var dm_build_1524 = 0
	for dm_build_1521 != nil {
		if dm_build_1521.dm_build_1528 > 0 {
			if dm_build_1523 > dm_build_1521.dm_build_1528 {
				dm_build_1524 = dm_build_1521.dm_build_1528
			} else {
				dm_build_1524 = dm_build_1523
			}
			copy(dm_build_1520[dm_build_1522:dm_build_1522+dm_build_1524], dm_build_1521.dm_build_1526[dm_build_1521.dm_build_1527:dm_build_1521.dm_build_1527+dm_build_1524])
			dm_build_1522 += dm_build_1524
			dm_build_1523 -= dm_build_1524
		}
		if dm_build_1519.dm_build_1472.Front() == nil {
			dm_build_1521 = nil
		} else {
			dm_build_1521 = dm_build_1519.dm_build_1472.Front().Value.(*dm_build_1525)
		}
	}
	return dm_build_1520
}

type dm_build_1525 struct {
	dm_build_1526 []byte
	dm_build_1527 int
	dm_build_1528 int
}

func dm_build_1529(dm_build_1530 []byte, dm_build_1531 int, dm_build_1532 int) *dm_build_1525 {
	return &dm_build_1525{
		dm_build_1530,
		dm_build_1531,
		dm_build_1532,
	}
}

func (dm_build_1534 *dm_build_1525) dm_build_1533(dm_build_1535 *Dm_build_0, dm_build_1536 int) int {
	if dm_build_1534.dm_build_1528 <= dm_build_1536 {
		dm_build_1536 = dm_build_1534.dm_build_1528
	}
	dm_build_1535.Dm_build_79(dm_build_1534.dm_build_1526[dm_build_1534.dm_build_1527 : dm_build_1534.dm_build_1527+dm_build_1536])
	dm_build_1534.dm_build_1527 += dm_build_1536
	dm_build_1534.dm_build_1528 -= dm_build_1536
	return dm_build_1536
}

func (dm_build_1538 *dm_build_1525) dm_build_1537(dm_build_1539 []byte, dm_build_1540 int, dm_build_1541 int) int {
	if dm_build_1538.dm_build_1528 <= dm_build_1541 {
		dm_build_1541 = dm_build_1538.dm_build_1528
	}
	copy(dm_build_1539[dm_build_1540:dm_build_1540+dm_build_1541], dm_build_1538.dm_build_1526[dm_build_1538.dm_build_1527:dm_build_1538.dm_build_1527+dm_build_1541])
	dm_build_1538.dm_build_1527 += dm_build_1541
	dm_build_1538.dm_build_1528 -= dm_build_1541
	return dm_build_1541
}

func (dm_build_1543 *dm_build_1525) dm_build_1542(dm_build_1544 io.Writer, dm_build_1545 int) int {
	if dm_build_1543.dm_build_1528 <= dm_build_1545 {
		dm_build_1545 = dm_build_1543.dm_build_1528
	}
	dm_build_1544.Write(dm_build_1543.dm_build_1526[dm_build_1543.dm_build_1527 : dm_build_1543.dm_build_1527+dm_build_1545])
	dm_build_1543.dm_build_1527 += dm_build_1545
	dm_build_1543.dm_build_1528 -= dm_build_1545
	return dm_build_1545
}
func (dm_build_1547 *dm_build_1525) dm_build_1546(dm_build_1548 int) byte {
	return dm_build_1547.dm_build_1526[dm_build_1547.dm_build_1527+dm_build_1548]
}
