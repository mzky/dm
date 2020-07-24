/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"crypto/tls"
	"dm/security"
	"net"
	"strconv"
	"time"
)

const (
	Dm_build_330 = 8192
	Dm_build_331 = 2 * time.Second
)

type dm_build_332 struct {
	dm_build_333 *net.TCPConn
	dm_build_334 *tls.Conn
	dm_build_335 *Dm_build_0
	dm_build_336 *DmConnection
	dm_build_337 security.Cipher
	dm_build_338 bool
	dm_build_339 bool
	dm_build_340 *security.DhKey
	dm_build_341 string
	dm_build_342 bool
}

func dm_build_343(dm_build_344 *DmConnection) (*dm_build_332, error) {
	dm_build_345, dm_build_346 := dm_build_348(dm_build_344.dmConnector.host+":"+strconv.Itoa(dm_build_344.dmConnector.port), time.Duration(dm_build_344.dmConnector.socketTimeout)*time.Second)
	if dm_build_346 != nil {
		return nil, dm_build_346
	}

	dm_build_347 := dm_build_332{}
	dm_build_347.dm_build_333 = dm_build_345
	dm_build_347.dm_build_335 = Dm_build_3(Dm_build_602)
	dm_build_347.dm_build_336 = dm_build_344
	dm_build_347.dm_build_338 = false
	dm_build_347.dm_build_339 = false
	dm_build_347.dm_build_341 = ""
	dm_build_347.dm_build_342 = false
	dm_build_344.Access = &dm_build_347

	return &dm_build_347, nil
}

func dm_build_348(dm_build_349 string, dm_build_350 time.Duration) (*net.TCPConn, error) {
	dm_build_351, dm_build_352 := net.DialTimeout("tcp", dm_build_349, dm_build_350)
	if dm_build_352 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_349).throw()
	}

	if tcpConn, ok := dm_build_351.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_331)
		tcpConn.SetNoDelay(true)

		tcpConn.SetReadBuffer(Dm_build_330)
		tcpConn.SetWriteBuffer(Dm_build_330)
		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_354 *dm_build_332) dm_build_353(dm_build_355 dm_build_719) bool {
	var dm_build_356 = dm_build_354.dm_build_336.dmConnector.compress
	if dm_build_355.dm_build_733() == Dm_build_629 || dm_build_356 == Dm_build_677 {
		return false
	}

	if dm_build_356 == Dm_build_675 {
		return true
	} else if dm_build_356 == Dm_build_676 {
		return !dm_build_354.dm_build_336.Local && dm_build_355.dm_build_731() > Dm_build_674
	}

	return false
}

func (dm_build_358 *dm_build_332) dm_build_357(dm_build_359 dm_build_719) bool {
	var dm_build_360 = dm_build_358.dm_build_336.dmConnector.compress
	if dm_build_359.dm_build_733() == Dm_build_629 || dm_build_360 == Dm_build_677 {
		return false
	}

	if dm_build_360 == Dm_build_675 {
		return true
	} else if dm_build_360 == Dm_build_676 {
		return dm_build_358.dm_build_335.Dm_build_263(Dm_build_637) == 1
	}

	return false
}

func (dm_build_362 *dm_build_332) dm_build_361(dm_build_363 dm_build_719) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_365 := dm_build_363.dm_build_731()

	if dm_build_365 > 0 {

		if dm_build_362.dm_build_353(dm_build_363) {
			var retBytes, err = Compress(dm_build_362.dm_build_335, Dm_build_630, int(dm_build_365), int(dm_build_362.dm_build_336.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_362.dm_build_335.Dm_build_14(Dm_build_630)

			dm_build_362.dm_build_335.Dm_build_51(dm_build_365)

			dm_build_362.dm_build_335.Dm_build_79(retBytes)

			dm_build_363.dm_build_732(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_362.dm_build_335.Dm_build_183(Dm_build_637, 1)
		}

		if dm_build_362.dm_build_339 {
			dm_build_365 = dm_build_363.dm_build_731()
			var retBytes = dm_build_362.dm_build_337.Encrypt(dm_build_362.dm_build_335.Dm_build_290(Dm_build_630, int(dm_build_365)), true)

			dm_build_362.dm_build_335.Dm_build_14(Dm_build_630)

			dm_build_362.dm_build_335.Dm_build_79(retBytes)

			dm_build_363.dm_build_732(int32(len(retBytes)))
		}
	}

	dm_build_363.dm_build_728()
	if dm_build_362.dm_build_593(dm_build_363) {
		if dm_build_362.dm_build_334 != nil {
			dm_build_362.dm_build_335.Dm_build_17(0)
			dm_build_362.dm_build_335.Dm_build_36(dm_build_362.dm_build_334)
		}
	} else {
		dm_build_362.dm_build_335.Dm_build_17(0)
		dm_build_362.dm_build_335.Dm_build_36(dm_build_362.dm_build_333)
	}
	return nil
}

func (dm_build_367 *dm_build_332) dm_build_366(dm_build_368 dm_build_719) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_370 := int32(0)
	if dm_build_367.dm_build_593(dm_build_368) {
		if dm_build_367.dm_build_334 != nil {
			dm_build_367.dm_build_335.Dm_build_14(0)
			dm_build_367.dm_build_335.Dm_build_30(dm_build_367.dm_build_334, Dm_build_630)
			dm_build_370 = dm_build_368.dm_build_731()
			if dm_build_370 > 0 {
				dm_build_367.dm_build_335.Dm_build_30(dm_build_367.dm_build_334, int(dm_build_370))
			}
		}
	} else {

		dm_build_367.dm_build_335.Dm_build_14(0)
		dm_build_367.dm_build_335.Dm_build_30(dm_build_367.dm_build_333, Dm_build_630)
		dm_build_370 = dm_build_368.dm_build_731()

		if dm_build_370 > 0 {
			dm_build_367.dm_build_335.Dm_build_30(dm_build_367.dm_build_333, int(dm_build_370))
		}
	}

	dm_build_368.dm_build_729()

	if dm_build_370 <= 0 {
		return nil
	}

	if dm_build_367.dm_build_339 {
		dm_build_370 = dm_build_368.dm_build_731()
		ebytes := dm_build_367.dm_build_335.Dm_build_290(Dm_build_630, int(dm_build_370))
		bytes, err := dm_build_367.dm_build_337.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_367.dm_build_335.Dm_build_14(Dm_build_630)
		dm_build_367.dm_build_335.Dm_build_79(bytes)
		dm_build_368.dm_build_732(int32(len(bytes)))
	}

	if dm_build_367.dm_build_357(dm_build_368) {

		dm_build_370 = dm_build_368.dm_build_731()
		cbytes := dm_build_367.dm_build_335.Dm_build_290(Dm_build_630+ULINT_SIZE, int(dm_build_370-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_367.dm_build_336.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_367.dm_build_335.Dm_build_14(Dm_build_630)
		dm_build_367.dm_build_335.Dm_build_79(bytes)
		dm_build_368.dm_build_732(int32(len(bytes)))
	}
	return nil
}

func (dm_build_372 *dm_build_332) dm_build_371(dm_build_373 dm_build_719) (dm_build_374 interface{}, dm_build_375 error) {
	dm_build_375 = dm_build_373.dm_build_723(dm_build_373)
	if dm_build_375 != nil {
		return nil, dm_build_375
	}

	dm_build_375 = dm_build_372.dm_build_361(dm_build_373)
	if dm_build_375 != nil {
		return nil, dm_build_375
	}

	dm_build_375 = dm_build_372.dm_build_366(dm_build_373)
	if dm_build_375 != nil {
		return nil, dm_build_375
	}

	return dm_build_373.dm_build_727(dm_build_373)
}

func (dm_build_377 *dm_build_332) dm_build_376() (*dm_build_1137, error) {

	Dm_build_378 := dm_build_1143(dm_build_377)
	_, dm_build_379 := dm_build_377.dm_build_371(Dm_build_378)
	if dm_build_379 != nil {
		return nil, dm_build_379
	}

	return Dm_build_378, nil
}

func (dm_build_381 *dm_build_332) dm_build_380() error {

	dm_build_382 := dm_build_1014(dm_build_381)
	_, dm_build_383 := dm_build_381.dm_build_371(dm_build_382)
	if dm_build_383 != nil {
		return dm_build_383
	}

	return nil
}

func (dm_build_385 *dm_build_332) dm_build_384() error {

	var dm_build_386 *dm_build_1137
	var err error
	if dm_build_386, err = dm_build_385.dm_build_376(); err != nil {
		return err
	}

	if dm_build_385.dm_build_336.sslEncrypt == 2 {
		if err = dm_build_385.dm_build_589(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_385.dm_build_336.sslEncrypt == 1 {
		if err = dm_build_385.dm_build_589(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_385.dm_build_339 || dm_build_385.dm_build_338 {
		k, err := dm_build_385.dm_build_579()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_386.Dm_build_1141)
		encryptType := dm_build_386.dm_build_1139
		hashType := int(dm_build_386.Dm_build_1140)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_385.dm_build_582(encryptType, sessionKey, dm_build_385.dm_build_336.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_385.dm_build_380(); err != nil {
		return err
	}
	return nil
}

func (dm_build_389 *dm_build_332) Dm_build_388(dm_build_390 *DmStatement) error {
	dm_build_391 := dm_build_1166(dm_build_389, dm_build_390)
	_, dm_build_392 := dm_build_389.dm_build_371(dm_build_391)
	if dm_build_392 != nil {
		return dm_build_392
	}

	return nil
}

func (dm_build_394 *dm_build_332) Dm_build_393(dm_build_395 int32) error {
	dm_build_396 := dm_build_1176(dm_build_394, dm_build_395)
	_, dm_build_397 := dm_build_394.dm_build_371(dm_build_396)
	if dm_build_397 != nil {
		return dm_build_397
	}

	return nil
}

func (dm_build_399 *dm_build_332) Dm_build_398(dm_build_400 *DmStatement, dm_build_401 bool, dm_build_402 int16) (*execInfo, error) {
	dm_build_403 := dm_build_1050(dm_build_399, dm_build_400, dm_build_401, dm_build_402)
	dm_build_404, dm_build_405 := dm_build_399.dm_build_371(dm_build_403)
	if dm_build_405 != nil {
		return nil, dm_build_405
	}
	return dm_build_404.(*execInfo), nil
}

func (dm_build_407 *dm_build_332) Dm_build_406(dm_build_408 *DmStatement, dm_build_409 int16) (*execInfo, error) {
	return dm_build_407.Dm_build_398(dm_build_408, false, Dm_build_681)
}

func (dm_build_411 *dm_build_332) Dm_build_410(dm_build_412 *DmStatement, dm_build_413 []OptParameter) (*execInfo, error) {
	dm_build_414, dm_build_415 := dm_build_411.dm_build_371(dm_build_812(dm_build_411, dm_build_412, dm_build_413))
	if dm_build_415 != nil {
		return nil, dm_build_415
	}

	return dm_build_414.(*execInfo), nil
}

func (dm_build_417 *dm_build_332) Dm_build_416(dm_build_418 *DmStatement, dm_build_419 int16) (*execInfo, error) {
	return dm_build_417.Dm_build_398(dm_build_418, true, dm_build_419)
}

func (dm_build_421 *dm_build_332) Dm_build_420(dm_build_422 *DmStatement, dm_build_423 [][]interface{}) (*execInfo, error) {
	dm_build_424 := dm_build_835(dm_build_421, dm_build_422, dm_build_423)
	dm_build_425, dm_build_426 := dm_build_421.dm_build_371(dm_build_424)
	if dm_build_426 != nil {
		return nil, dm_build_426
	}
	return dm_build_425.(*execInfo), nil
}

func (dm_build_428 *dm_build_332) Dm_build_427(dm_build_429 *DmStatement, dm_build_430 [][]interface{}) (*execInfo, error) {
	var dm_build_431, dm_build_432 = 0, 0
	var dm_build_433 = len(dm_build_430)
	var dm_build_434 [][]interface{}
	var dm_build_435 = NewExceInfo()
	dm_build_435.updateCounts = make([]int64, dm_build_433)
	var dm_build_436 = false
	var dm_build_437 = false
	for dm_build_431 < dm_build_433 {
		for dm_build_432 = dm_build_431; dm_build_432 < dm_build_433; dm_build_432++ {
			paramData := dm_build_430[dm_build_432]
			bindData := make([]interface{}, dm_build_429.paramCount)
			dm_build_436 = false
			for icol := 0; icol < int(dm_build_429.paramCount); icol++ {
				if dm_build_429.params[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_428.dm_build_562(bindData, paramData, icol) {
					dm_build_436 = true
					break
				}
			}

			if dm_build_436 {
				break
			}
			dm_build_434 = append(dm_build_434, bindData)
		}

		if dm_build_432 != dm_build_431 {
			tmpExecInfo, err := dm_build_428.Dm_build_420(dm_build_429, dm_build_434)
			if err != nil {
				return nil, err
			}
			dm_build_434 = dm_build_434[0:0]
			if dm_build_432-dm_build_431 == 1 {
				dm_build_435.updateCounts[dm_build_431] = tmpExecInfo.updateCount
			} else if tmpExecInfo.updateCounts != nil {
				copy(dm_build_435.updateCounts[dm_build_431:dm_build_432], tmpExecInfo.updateCounts[0:dm_build_432-dm_build_431])
			}
		}

		if dm_build_432 < dm_build_433 {
			tmpExecInfo, err := dm_build_428.Dm_build_438(dm_build_429, dm_build_430[dm_build_432], dm_build_437)
			if err != nil {
				return nil, err
			}

			dm_build_437 = true
			dm_build_435.updateCounts[dm_build_432] = tmpExecInfo.updateCount
		}

		dm_build_431 = dm_build_432 + 1
	}
	for _, i := range dm_build_435.updateCounts {
		dm_build_435.updateCount += i
	}
	return dm_build_435, nil
}

func (dm_build_439 *dm_build_332) Dm_build_438(dm_build_440 *DmStatement, dm_build_441 []interface{}, dm_build_442 bool) (*execInfo, error) {

	var dm_build_443 = make([]interface{}, dm_build_440.paramCount)
	for icol := 0; icol < int(dm_build_440.paramCount); icol++ {
		if dm_build_440.params[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_439.dm_build_562(dm_build_443, dm_build_441, icol) {

			if !dm_build_442 {
				preExecute := dm_build_1040(dm_build_439, dm_build_440, dm_build_440.params)
				dm_build_439.dm_build_371(preExecute)
				dm_build_442 = true
			}

			dm_build_439.dm_build_568(dm_build_440, dm_build_440.params[icol], icol, dm_build_441[icol].(iOffRowBinder))
			dm_build_443[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_444 = make([][]interface{}, 1, 1)
	dm_build_444[0] = dm_build_443

	dm_build_445 := dm_build_835(dm_build_439, dm_build_440, dm_build_444)
	dm_build_446, dm_build_447 := dm_build_439.dm_build_371(dm_build_445)
	if dm_build_447 != nil {
		return nil, dm_build_447
	}
	return dm_build_446.(*execInfo), nil
}

func (dm_build_449 *dm_build_332) Dm_build_448(dm_build_450 *DmStatement, dm_build_451 int16) (*execInfo, error) {
	dm_build_452 := dm_build_1028(dm_build_449, dm_build_450, dm_build_451)

	dm_build_453, dm_build_454 := dm_build_449.dm_build_371(dm_build_452)
	if dm_build_454 != nil {
		return nil, dm_build_454
	}
	return dm_build_453.(*execInfo), nil
}

func (dm_build_456 *dm_build_332) Dm_build_455(dm_build_457 *innerRows, dm_build_458 int64) (*execInfo, error) {
	dm_build_459 := dm_build_935(dm_build_456, dm_build_457, dm_build_458, INT64_MAX)
	dm_build_460, dm_build_461 := dm_build_456.dm_build_371(dm_build_459)
	if dm_build_461 != nil {
		return nil, dm_build_461
	}
	return dm_build_460.(*execInfo), nil
}

func (dm_build_463 *dm_build_332) Commit() error {
	dm_build_464 := dm_build_798(dm_build_463)
	_, dm_build_465 := dm_build_463.dm_build_371(dm_build_464)
	if dm_build_465 != nil {
		return dm_build_465
	}

	return nil
}

func (dm_build_467 *dm_build_332) Rollback() error {
	dm_build_468 := dm_build_1088(dm_build_467)
	_, dm_build_469 := dm_build_467.dm_build_371(dm_build_468)
	if dm_build_469 != nil {
		return dm_build_469
	}

	return nil
}

func (dm_build_471 *dm_build_332) Dm_build_470(dm_build_472 *DmConnection) error {
	dm_build_473 := dm_build_1093(dm_build_471, dm_build_472.IsoLevel)
	_, dm_build_474 := dm_build_471.dm_build_371(dm_build_473)
	if dm_build_474 != nil {
		return dm_build_474
	}

	return nil
}

func (dm_build_476 *dm_build_332) Dm_build_475(dm_build_477 *DmStatement, dm_build_478 string) error {
	dm_build_479 := dm_build_803(dm_build_476, dm_build_477, dm_build_478)
	_, dm_build_480 := dm_build_476.dm_build_371(dm_build_479)
	if dm_build_480 != nil {
		return dm_build_480
	}

	return nil
}

func (dm_build_482 *dm_build_332) Dm_build_481(dm_build_483 []uint32) ([]int64, error) {
	dm_build_484 := dm_build_1184(dm_build_482, dm_build_483)
	dm_build_485, dm_build_486 := dm_build_482.dm_build_371(dm_build_484)
	if dm_build_486 != nil {
		return nil, dm_build_486
	}
	return dm_build_485.([]int64), nil
}

func (dm_build_488 *dm_build_332) Close() error {
	if dm_build_488.dm_build_342 {
		return nil
	}

	dm_build_489 := dm_build_488.dm_build_333.Close()
	if dm_build_489 != nil {
		return dm_build_489
	}

	dm_build_488.dm_build_336 = nil
	dm_build_488.dm_build_342 = true
	return nil
}

func (dm_build_491 *dm_build_332) dm_build_490(dm_build_492 *lob) (int64, error) {
	dm_build_493 := dm_build_966(dm_build_491, dm_build_492)
	dm_build_494, dm_build_495 := dm_build_491.dm_build_371(dm_build_493)
	if dm_build_495 != nil {
		return 0, dm_build_495
	}
	return dm_build_494.(int64), nil
}

func (dm_build_497 *dm_build_332) dm_build_496(dm_build_498 *lob, dm_build_499 int32, dm_build_500 int32) ([]byte, error) {
	dm_build_501 := dm_build_953(dm_build_497, dm_build_498, int(dm_build_499), int(dm_build_500))
	dm_build_502, dm_build_503 := dm_build_497.dm_build_371(dm_build_501)
	if dm_build_503 != nil {
		return nil, dm_build_503
	}
	return dm_build_502.([]byte), nil
}

func (dm_build_505 *dm_build_332) dm_build_504(dm_build_506 *DmBlob, dm_build_507 int32, dm_build_508 int32) ([]byte, error) {
	var dm_build_509 = make([]byte, dm_build_508)
	var dm_build_510 int32 = 0
	var dm_build_511 int32 = 0
	var dm_build_512 []byte
	var dm_build_513 error
	for dm_build_510 < dm_build_508 {
		dm_build_511 = dm_build_508 - dm_build_510
		if dm_build_511 > Dm_build_714 {
			dm_build_511 = Dm_build_714
		}
		dm_build_512, dm_build_513 = dm_build_505.dm_build_496(&dm_build_506.lob, dm_build_507, dm_build_508)
		if dm_build_513 != nil {
			return nil, dm_build_513
		}
		if dm_build_512 == nil || len(dm_build_512) == 0 {
			break
		}
		Dm_build_1195.Dm_build_1251(dm_build_509, int(dm_build_510), dm_build_512, 0, len(dm_build_512))
		dm_build_510 += int32(len(dm_build_512))
		dm_build_507 += int32(len(dm_build_512))
		if dm_build_506.readOver {
			break
		}
	}
	return dm_build_509, nil
}

func (dm_build_515 *dm_build_332) dm_build_514(dm_build_516 *DmClob, dm_build_517 int32, dm_build_518 int32) (string, error) {
	var dm_build_519 = ""
	var dm_build_520 int32 = 0
	var dm_build_521 int32 = 0
	var dm_build_522 []byte
	var dm_build_523 string
	var dm_build_524 error
	for dm_build_520 < dm_build_518 {
		dm_build_521 = dm_build_518 - dm_build_520
		if dm_build_521 > Dm_build_714/2 {
			dm_build_521 = Dm_build_714 / 2
		}
		dm_build_522, dm_build_524 = dm_build_515.dm_build_496(&dm_build_516.lob, dm_build_517, dm_build_518)
		if dm_build_524 != nil {
			return "", dm_build_524
		}
		if dm_build_522 == nil || len(dm_build_522) == 0 {
			break
		}
		dm_build_523 = Dm_build_1195.Dm_build_1349(dm_build_522, 0, len(dm_build_522), dm_build_516.serverEncoding, dm_build_515.dm_build_336)
		dm_build_519 += dm_build_523
		dm_build_520 += int32(len(dm_build_523))
		dm_build_517 += int32(len(dm_build_522))
		if dm_build_516.readOver {
			break
		}
	}
	return dm_build_519, nil
}

func (dm_build_526 *dm_build_332) dm_build_525(dm_build_527 *DmClob, dm_build_528 int, dm_build_529 string, dm_build_530 string) (int, error) {
	var dm_build_531 = Dm_build_1195.Dm_build_1405(dm_build_529, dm_build_530, dm_build_526.dm_build_336)
	var dm_build_532 = 0
	var dm_build_533 = len(dm_build_531)
	var dm_build_534 = 0
	var dm_build_535 = 0
	var dm_build_536 = 0
	var dm_build_537 = dm_build_533/Dm_build_713 + 1
	var dm_build_538 byte = 0
	var dm_build_539 byte = 0x01
	var dm_build_540 byte = 0x02
	for i := 0; i < dm_build_537; i++ {
		dm_build_538 = 0
		if i == 0 {
			dm_build_538 |= dm_build_539
		}
		if i == dm_build_537-1 {
			dm_build_538 |= dm_build_540
		}
		dm_build_536 = dm_build_533 - dm_build_535
		if dm_build_536 > Dm_build_713 {
			dm_build_536 = Dm_build_713
		}

		setLobData := dm_build_1107(dm_build_526, &dm_build_527.lob, dm_build_538, dm_build_528, dm_build_531, dm_build_532, dm_build_536)
		ret, err := dm_build_526.dm_build_371(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_534, nil
		} else {
			dm_build_528 += int(tmp)
			dm_build_534 += int(tmp)
			dm_build_535 += dm_build_536
			dm_build_532 += dm_build_536
		}
	}
	return dm_build_534, nil
}

func (dm_build_542 *dm_build_332) dm_build_541(dm_build_543 *DmBlob, dm_build_544 int, dm_build_545 []byte) (int, error) {
	var dm_build_546 = 0
	var dm_build_547 = len(dm_build_545)
	var dm_build_548 = 0
	var dm_build_549 = 0
	var dm_build_550 = 0
	var dm_build_551 = dm_build_547/Dm_build_713 + 1
	var dm_build_552 byte = 0
	var dm_build_553 byte = 0x01
	var dm_build_554 byte = 0x02
	for i := 0; i < dm_build_551; i++ {
		dm_build_552 = 0
		if i == 0 {
			dm_build_552 |= dm_build_553
		}
		if i == dm_build_551-1 {
			dm_build_552 |= dm_build_554
		}
		dm_build_550 = dm_build_547 - dm_build_549
		if dm_build_550 > Dm_build_713 {
			dm_build_550 = Dm_build_713
		}

		setLobData := dm_build_1107(dm_build_542, &dm_build_543.lob, dm_build_552, dm_build_544, dm_build_545, dm_build_546, dm_build_550)
		ret, err := dm_build_542.dm_build_371(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_548, nil
		} else {
			dm_build_544 += int(tmp)
			dm_build_548 += int(tmp)
			dm_build_549 += dm_build_550
			dm_build_546 += dm_build_550
		}
	}
	return dm_build_548, nil
}

func (dm_build_556 *dm_build_332) dm_build_555(dm_build_557 *lob, dm_build_558 int) (int64, error) {
	dm_build_559 := dm_build_977(dm_build_556, dm_build_557, dm_build_558)
	dm_build_560, dm_build_561 := dm_build_556.dm_build_371(dm_build_559)
	if dm_build_561 != nil {
		return dm_build_557.length, dm_build_561
	}
	return dm_build_560.(int64), nil
}

func (dm_build_563 *dm_build_332) dm_build_562(dm_build_564 []interface{}, dm_build_565 []interface{}, dm_build_566 int) bool {
	var dm_build_567 = false
	if dm_build_566 >= len(dm_build_565) || dm_build_565[dm_build_566] == nil {
		dm_build_564[dm_build_566] = ParamDataEnum_Null
	} else if binder, ok := dm_build_565[dm_build_566].(iOffRowBinder); ok {
		dm_build_567 = true
		dm_build_564[dm_build_566] = ParamDataEnum_OFF_ROW
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_563.dm_build_336) {
			dm_build_564[dm_build_566] = &lobCtl{lob.buildCtlData()}
			dm_build_567 = false
		}
	} else {
		dm_build_564[dm_build_566] = dm_build_565[dm_build_566]
	}
	return dm_build_567
}

func (dm_build_569 *dm_build_332) dm_build_568(dm_build_570 *DmStatement, dm_build_571 parameter, dm_build_572 int, dm_build_573 iOffRowBinder) error {
	var dm_build_574 = Dm_build_1475()
	dm_build_573.read(dm_build_574)
	var dm_build_575 = 0
	for !dm_build_573.isReadOver() || dm_build_574.Dm_build_1476() > 0 {
		if !dm_build_573.isReadOver() && dm_build_574.Dm_build_1476() < Dm_build_713 {
			dm_build_573.read(dm_build_574)
		}
		if dm_build_574.Dm_build_1476() > Dm_build_713 {
			dm_build_575 = Dm_build_713
		} else {
			dm_build_575 = dm_build_574.Dm_build_1476()
		}

		putData := dm_build_1078(dm_build_569, dm_build_570, int16(dm_build_572), dm_build_574, int32(dm_build_575))
		_, err := dm_build_569.dm_build_371(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_577 *dm_build_332) dm_build_576() ([]byte, error) {
	var dm_build_578 error
	if dm_build_577.dm_build_340 == nil {
		if dm_build_577.dm_build_340, dm_build_578 = security.NewClientKeyPair(); dm_build_578 != nil {
			return nil, dm_build_578
		}
	}
	return security.Bn2Bytes(dm_build_577.dm_build_340.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_580 *dm_build_332) dm_build_579() (*security.DhKey, error) {
	var dm_build_581 error
	if dm_build_580.dm_build_340 == nil {
		if dm_build_580.dm_build_340, dm_build_581 = security.NewClientKeyPair(); dm_build_581 != nil {
			return nil, dm_build_581
		}
	}
	return dm_build_580.dm_build_340, nil
}

func (dm_build_583 *dm_build_332) dm_build_582(dm_build_584 int, dm_build_585 []byte, dm_build_586 string, dm_build_587 int) (dm_build_588 error) {
	if dm_build_584 > 0 && dm_build_584 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_585 != nil {
		dm_build_583.dm_build_337, dm_build_588 = security.NewSymmCipher(dm_build_584, dm_build_585)
	} else if dm_build_584 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_583.dm_build_337, dm_build_588 = security.NewThirdPartCipher(dm_build_584, dm_build_585, dm_build_586, dm_build_587); dm_build_588 != nil {
			dm_build_588 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_588.Error()).throw()
		}
	}
	return
}

func (dm_build_590 *dm_build_332) dm_build_589(dm_build_591 bool) (dm_build_592 error) {
	if dm_build_590.dm_build_334, dm_build_592 = security.NewTLSFromTCP(dm_build_590.dm_build_333, dm_build_590.dm_build_336.dmConnector.sslCertPath, dm_build_590.dm_build_336.dmConnector.sslKeyPath, dm_build_590.dm_build_336.dmConnector.user); dm_build_592 != nil {
		return
	}
	if !dm_build_591 {
		dm_build_590.dm_build_334 = nil
	}
	return
}

func (dm_build_594 *dm_build_332) dm_build_593(dm_build_595 dm_build_719) bool {
	return dm_build_595.dm_build_733() != Dm_build_629 && dm_build_594.dm_build_336.sslEncrypt == 1
}
