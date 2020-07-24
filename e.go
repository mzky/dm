/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_1194 struct{}

var Dm_build_1195 = &dm_build_1194{}

func (Dm_build_1197 *dm_build_1194) Dm_build_1196(dm_build_1198 []byte, dm_build_1199 int, dm_build_1200 byte) int {
	dm_build_1198[dm_build_1199] = dm_build_1200
	return 1
}

func (Dm_build_1202 *dm_build_1194) Dm_build_1201(dm_build_1203 []byte, dm_build_1204 int, dm_build_1205 int8) int {
	dm_build_1203[dm_build_1204] = byte(dm_build_1205)
	return 1
}

func (Dm_build_1207 *dm_build_1194) Dm_build_1206(dm_build_1208 []byte, dm_build_1209 int, dm_build_1210 int16) int {
	dm_build_1208[dm_build_1209] = byte(dm_build_1210)
	dm_build_1209++
	dm_build_1208[dm_build_1209] = byte(dm_build_1210 >> 8)
	return 2
}

func (Dm_build_1212 *dm_build_1194) Dm_build_1211(dm_build_1213 []byte, dm_build_1214 int, dm_build_1215 int32) int {
	dm_build_1213[dm_build_1214] = byte(dm_build_1215)
	dm_build_1214++
	dm_build_1213[dm_build_1214] = byte(dm_build_1215 >> 8)
	dm_build_1214++
	dm_build_1213[dm_build_1214] = byte(dm_build_1215 >> 16)
	dm_build_1214++
	dm_build_1213[dm_build_1214] = byte(dm_build_1215 >> 24)
	dm_build_1214++
	return 4
}

func (Dm_build_1217 *dm_build_1194) Dm_build_1216(dm_build_1218 []byte, dm_build_1219 int, dm_build_1220 int64) int {
	dm_build_1218[dm_build_1219] = byte(dm_build_1220)
	dm_build_1219++
	dm_build_1218[dm_build_1219] = byte(dm_build_1220 >> 8)
	dm_build_1219++
	dm_build_1218[dm_build_1219] = byte(dm_build_1220 >> 16)
	dm_build_1219++
	dm_build_1218[dm_build_1219] = byte(dm_build_1220 >> 24)
	dm_build_1219++
	dm_build_1218[dm_build_1219] = byte(dm_build_1220 >> 32)
	dm_build_1219++
	dm_build_1218[dm_build_1219] = byte(dm_build_1220 >> 40)
	dm_build_1219++
	dm_build_1218[dm_build_1219] = byte(dm_build_1220 >> 48)
	dm_build_1219++
	dm_build_1218[dm_build_1219] = byte(dm_build_1220 >> 56)
	return 8
}

func (Dm_build_1222 *dm_build_1194) Dm_build_1221(dm_build_1223 []byte, dm_build_1224 int, dm_build_1225 float32) int {
	return Dm_build_1222.Dm_build_1241(dm_build_1223, dm_build_1224, math.Float32bits(dm_build_1225))
}

func (Dm_build_1227 *dm_build_1194) Dm_build_1226(dm_build_1228 []byte, dm_build_1229 int, dm_build_1230 float64) int {
	return Dm_build_1227.Dm_build_1246(dm_build_1228, dm_build_1229, math.Float64bits(dm_build_1230))
}

func (Dm_build_1232 *dm_build_1194) Dm_build_1231(dm_build_1233 []byte, dm_build_1234 int, dm_build_1235 uint8) int {
	dm_build_1233[dm_build_1234] = byte(dm_build_1235)
	return 1
}

func (Dm_build_1237 *dm_build_1194) Dm_build_1236(dm_build_1238 []byte, dm_build_1239 int, dm_build_1240 uint16) int {
	dm_build_1238[dm_build_1239] = byte(dm_build_1240)
	dm_build_1239++
	dm_build_1238[dm_build_1239] = byte(dm_build_1240 >> 8)
	return 2
}

func (Dm_build_1242 *dm_build_1194) Dm_build_1241(dm_build_1243 []byte, dm_build_1244 int, dm_build_1245 uint32) int {
	dm_build_1243[dm_build_1244] = byte(dm_build_1245)
	dm_build_1244++
	dm_build_1243[dm_build_1244] = byte(dm_build_1245 >> 8)
	dm_build_1244++
	dm_build_1243[dm_build_1244] = byte(dm_build_1245 >> 16)
	dm_build_1244++
	dm_build_1243[dm_build_1244] = byte(dm_build_1245 >> 24)
	return 3
}

func (Dm_build_1247 *dm_build_1194) Dm_build_1246(dm_build_1248 []byte, dm_build_1249 int, dm_build_1250 uint64) int {
	dm_build_1248[dm_build_1249] = byte(dm_build_1250)
	dm_build_1249++
	dm_build_1248[dm_build_1249] = byte(dm_build_1250 >> 8)
	dm_build_1249++
	dm_build_1248[dm_build_1249] = byte(dm_build_1250 >> 16)
	dm_build_1249++
	dm_build_1248[dm_build_1249] = byte(dm_build_1250 >> 24)
	dm_build_1249++
	dm_build_1248[dm_build_1249] = byte(dm_build_1250 >> 32)
	dm_build_1249++
	dm_build_1248[dm_build_1249] = byte(dm_build_1250 >> 40)
	dm_build_1249++
	dm_build_1248[dm_build_1249] = byte(dm_build_1250 >> 48)
	dm_build_1249++
	dm_build_1248[dm_build_1249] = byte(dm_build_1250 >> 56)
	return 3
}

func (Dm_build_1252 *dm_build_1194) Dm_build_1251(dm_build_1253 []byte, dm_build_1254 int, dm_build_1255 []byte, dm_build_1256 int, dm_build_1257 int) int {
	copy(dm_build_1253[dm_build_1254:dm_build_1254+dm_build_1257], dm_build_1255[dm_build_1256:dm_build_1256+dm_build_1257])
	return dm_build_1257
}

func (Dm_build_1259 *dm_build_1194) Dm_build_1258(dm_build_1260 []byte, dm_build_1261 int, dm_build_1262 []byte, dm_build_1263 int, dm_build_1264 int) int {
	dm_build_1261 += Dm_build_1259.Dm_build_1241(dm_build_1260, dm_build_1261, uint32(dm_build_1264))
	return 4 + Dm_build_1259.Dm_build_1251(dm_build_1260, dm_build_1261, dm_build_1262, dm_build_1263, dm_build_1264)
}

func (Dm_build_1266 *dm_build_1194) Dm_build_1265(dm_build_1267 []byte, dm_build_1268 int, dm_build_1269 []byte, dm_build_1270 int, dm_build_1271 int) int {
	dm_build_1268 += Dm_build_1266.Dm_build_1236(dm_build_1267, dm_build_1268, uint16(dm_build_1271))
	return 2 + Dm_build_1266.Dm_build_1251(dm_build_1267, dm_build_1268, dm_build_1269, dm_build_1270, dm_build_1271)
}

func (Dm_build_1273 *dm_build_1194) Dm_build_1272(dm_build_1274 []byte, dm_build_1275 int, dm_build_1276 string, dm_build_1277 string, dm_build_1278 *DmConnection) int {
	dm_build_1279 := Dm_build_1273.Dm_build_1405(dm_build_1276, dm_build_1277, dm_build_1278)
	dm_build_1275 += Dm_build_1273.Dm_build_1241(dm_build_1274, dm_build_1275, uint32(len(dm_build_1279)))
	return 4 + Dm_build_1273.Dm_build_1251(dm_build_1274, dm_build_1275, dm_build_1279, 0, len(dm_build_1279))
}

func (Dm_build_1281 *dm_build_1194) Dm_build_1280(dm_build_1282 []byte, dm_build_1283 int, dm_build_1284 string, dm_build_1285 string, dm_build_1286 *DmConnection) int {
	dm_build_1287 := Dm_build_1281.Dm_build_1405(dm_build_1284, dm_build_1285, dm_build_1286)

	dm_build_1283 += Dm_build_1281.Dm_build_1236(dm_build_1282, dm_build_1283, uint16(len(dm_build_1287)))
	return 2 + Dm_build_1281.Dm_build_1251(dm_build_1282, dm_build_1283, dm_build_1287, 0, len(dm_build_1287))
}

func (Dm_build_1289 *dm_build_1194) Dm_build_1288(dm_build_1290 []byte, dm_build_1291 int) byte {
	return dm_build_1290[dm_build_1291]
}

func (Dm_build_1293 *dm_build_1194) Dm_build_1292(dm_build_1294 []byte, dm_build_1295 int) int16 {
	var dm_build_1296 int16
	dm_build_1296 = int16(dm_build_1294[dm_build_1295] & 0xff)
	dm_build_1295++
	dm_build_1296 |= int16(dm_build_1294[dm_build_1295]&0xff) << 8
	return dm_build_1296
}

func (Dm_build_1298 *dm_build_1194) Dm_build_1297(dm_build_1299 []byte, dm_build_1300 int) int32 {
	var dm_build_1301 int32
	dm_build_1301 = int32(dm_build_1299[dm_build_1300] & 0xff)
	dm_build_1300++
	dm_build_1301 |= int32(dm_build_1299[dm_build_1300]&0xff) << 8
	dm_build_1300++
	dm_build_1301 |= int32(dm_build_1299[dm_build_1300]&0xff) << 16
	dm_build_1300++
	dm_build_1301 |= int32(dm_build_1299[dm_build_1300]&0xff) << 24
	return dm_build_1301
}

func (Dm_build_1303 *dm_build_1194) Dm_build_1302(dm_build_1304 []byte, dm_build_1305 int) int64 {
	var dm_build_1306 int64
	dm_build_1306 = int64(dm_build_1304[dm_build_1305] & 0xff)
	dm_build_1305++
	dm_build_1306 |= int64(dm_build_1304[dm_build_1305]&0xff) << 8
	dm_build_1305++
	dm_build_1306 |= int64(dm_build_1304[dm_build_1305]&0xff) << 16
	dm_build_1305++
	dm_build_1306 |= int64(dm_build_1304[dm_build_1305]&0xff) << 24
	dm_build_1305++
	dm_build_1306 |= int64(dm_build_1304[dm_build_1305]&0xff) << 32
	dm_build_1305++
	dm_build_1306 |= int64(dm_build_1304[dm_build_1305]&0xff) << 40
	dm_build_1305++
	dm_build_1306 |= int64(dm_build_1304[dm_build_1305]&0xff) << 48
	dm_build_1305++
	dm_build_1306 |= int64(dm_build_1304[dm_build_1305]&0xff) << 56
	return dm_build_1306
}

func (Dm_build_1308 *dm_build_1194) Dm_build_1307(dm_build_1309 []byte, dm_build_1310 int) float32 {
	return math.Float32frombits(Dm_build_1308.Dm_build_1324(dm_build_1309, dm_build_1310))
}

func (Dm_build_1312 *dm_build_1194) Dm_build_1311(dm_build_1313 []byte, dm_build_1314 int) float64 {
	return math.Float64frombits(Dm_build_1312.Dm_build_1329(dm_build_1313, dm_build_1314))
}

func (Dm_build_1316 *dm_build_1194) Dm_build_1315(dm_build_1317 []byte, dm_build_1318 int) uint8 {
	return uint8(dm_build_1317[dm_build_1318] & 0xff)
}

func (Dm_build_1320 *dm_build_1194) Dm_build_1319(dm_build_1321 []byte, dm_build_1322 int) uint16 {
	var dm_build_1323 uint16
	dm_build_1323 = uint16(dm_build_1321[dm_build_1322] & 0xff)
	dm_build_1322++
	dm_build_1323 |= uint16(dm_build_1321[dm_build_1322]&0xff) << 8
	return dm_build_1323
}

func (Dm_build_1325 *dm_build_1194) Dm_build_1324(dm_build_1326 []byte, dm_build_1327 int) uint32 {
	var dm_build_1328 uint32
	dm_build_1328 = uint32(dm_build_1326[dm_build_1327] & 0xff)
	dm_build_1327++
	dm_build_1328 |= uint32(dm_build_1326[dm_build_1327]&0xff) << 8
	dm_build_1327++
	dm_build_1328 |= uint32(dm_build_1326[dm_build_1327]&0xff) << 16
	dm_build_1327++
	dm_build_1328 |= uint32(dm_build_1326[dm_build_1327]&0xff) << 24
	return dm_build_1328
}

func (Dm_build_1330 *dm_build_1194) Dm_build_1329(dm_build_1331 []byte, dm_build_1332 int) uint64 {
	var dm_build_1333 uint64
	dm_build_1333 = uint64(dm_build_1331[dm_build_1332] & 0xff)
	dm_build_1332++
	dm_build_1333 |= uint64(dm_build_1331[dm_build_1332]&0xff) << 8
	dm_build_1332++
	dm_build_1333 |= uint64(dm_build_1331[dm_build_1332]&0xff) << 16
	dm_build_1332++
	dm_build_1333 |= uint64(dm_build_1331[dm_build_1332]&0xff) << 24
	dm_build_1332++
	dm_build_1333 |= uint64(dm_build_1331[dm_build_1332]&0xff) << 32
	dm_build_1332++
	dm_build_1333 |= uint64(dm_build_1331[dm_build_1332]&0xff) << 40
	dm_build_1332++
	dm_build_1333 |= uint64(dm_build_1331[dm_build_1332]&0xff) << 48
	dm_build_1332++
	dm_build_1333 |= uint64(dm_build_1331[dm_build_1332]&0xff) << 56
	return dm_build_1333
}

func (Dm_build_1335 *dm_build_1194) Dm_build_1334(dm_build_1336 []byte, dm_build_1337 int) []byte {
	dm_build_1338 := Dm_build_1335.Dm_build_1324(dm_build_1336, dm_build_1337)
	return dm_build_1336[dm_build_1337+4 : dm_build_1337+4+int(dm_build_1338)]

}

func (Dm_build_1340 *dm_build_1194) Dm_build_1339(dm_build_1341 []byte, dm_build_1342 int) []byte {
	dm_build_1343 := Dm_build_1340.Dm_build_1319(dm_build_1341, dm_build_1342)
	return dm_build_1341[dm_build_1342+2 : dm_build_1342+2+int(dm_build_1343)]

}

func (Dm_build_1345 *dm_build_1194) Dm_build_1344(dm_build_1346 []byte, dm_build_1347 int, dm_build_1348 int) []byte {
	return dm_build_1346[dm_build_1347 : dm_build_1347+dm_build_1348]

}

func (Dm_build_1350 *dm_build_1194) Dm_build_1349(dm_build_1351 []byte, dm_build_1352 int, dm_build_1353 int, dm_build_1354 string, dm_build_1355 *DmConnection) string {
	return Dm_build_1350.Dm_build_1442(dm_build_1351[dm_build_1352:dm_build_1352+dm_build_1353], dm_build_1354, dm_build_1355)
}

func (Dm_build_1357 *dm_build_1194) Dm_build_1356(dm_build_1358 []byte, dm_build_1359 int, dm_build_1360 string, dm_build_1361 *DmConnection) string {
	dm_build_1362 := Dm_build_1357.Dm_build_1324(dm_build_1358, dm_build_1359)
	dm_build_1359 += 4
	return Dm_build_1357.Dm_build_1349(dm_build_1358, dm_build_1359, int(dm_build_1362), dm_build_1360, dm_build_1361)
}

func (Dm_build_1364 *dm_build_1194) Dm_build_1363(dm_build_1365 []byte, dm_build_1366 int, dm_build_1367 string, dm_build_1368 *DmConnection) string {
	dm_build_1369 := Dm_build_1364.Dm_build_1319(dm_build_1365, dm_build_1366)
	dm_build_1366 += 2
	return Dm_build_1364.Dm_build_1349(dm_build_1365, dm_build_1366, int(dm_build_1369), dm_build_1367, dm_build_1368)
}

func (Dm_build_1371 *dm_build_1194) Dm_build_1370(dm_build_1372 byte) []byte {
	return []byte{dm_build_1372}
}

func (Dm_build_1374 *dm_build_1194) Dm_build_1373(dm_build_1375 int16) []byte {
	return []byte{byte(dm_build_1375), byte(dm_build_1375 >> 8)}
}

func (Dm_build_1377 *dm_build_1194) Dm_build_1376(dm_build_1378 int32) []byte {
	return []byte{byte(dm_build_1378), byte(dm_build_1378 >> 8), byte(dm_build_1378 >> 16), byte(dm_build_1378 >> 24)}
}

func (Dm_build_1380 *dm_build_1194) Dm_build_1379(dm_build_1381 int64) []byte {
	return []byte{byte(dm_build_1381), byte(dm_build_1381 >> 8), byte(dm_build_1381 >> 16), byte(dm_build_1381 >> 24), byte(dm_build_1381 >> 32),
		byte(dm_build_1381 >> 40), byte(dm_build_1381 >> 48), byte(dm_build_1381 >> 56)}
}

func (Dm_build_1383 *dm_build_1194) Dm_build_1382(dm_build_1384 float32) []byte {
	return Dm_build_1383.Dm_build_1394(math.Float32bits(dm_build_1384))
}

func (Dm_build_1386 *dm_build_1194) Dm_build_1385(dm_build_1387 float64) []byte {
	return Dm_build_1386.Dm_build_1397(math.Float64bits(dm_build_1387))
}

func (Dm_build_1389 *dm_build_1194) Dm_build_1388(dm_build_1390 uint8) []byte {
	return []byte{byte(dm_build_1390)}
}

func (Dm_build_1392 *dm_build_1194) Dm_build_1391(dm_build_1393 uint16) []byte {
	return []byte{byte(dm_build_1393), byte(dm_build_1393 >> 8)}
}

func (Dm_build_1395 *dm_build_1194) Dm_build_1394(dm_build_1396 uint32) []byte {
	return []byte{byte(dm_build_1396), byte(dm_build_1396 >> 8), byte(dm_build_1396 >> 16), byte(dm_build_1396 >> 24)}
}

func (Dm_build_1398 *dm_build_1194) Dm_build_1397(dm_build_1399 uint64) []byte {
	return []byte{byte(dm_build_1399), byte(dm_build_1399 >> 8), byte(dm_build_1399 >> 16), byte(dm_build_1399 >> 24), byte(dm_build_1399 >> 32), byte(dm_build_1399 >> 40), byte(dm_build_1399 >> 48), byte(dm_build_1399 >> 56)}
}

func (Dm_build_1401 *dm_build_1194) Dm_build_1400(dm_build_1402 []byte, dm_build_1403 string, dm_build_1404 *DmConnection) []byte {
	if dm_build_1403 == "UTF-8" {
		return dm_build_1402
	}

	if dm_build_1404 == nil {
		if e := dm_build_1447(dm_build_1403); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1402), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1404.encodeBuffer == nil {
		dm_build_1404.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1404.encode = dm_build_1447(dm_build_1404.getServerEncoding())
		dm_build_1404.transformReaderDst = make([]byte, 4096)
		dm_build_1404.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1404.encode; e != nil {

		dm_build_1404.encodeBuffer.Reset()

		n, err := dm_build_1404.encodeBuffer.ReadFrom(
			Dm_build_1461(bytes.NewReader(dm_build_1402), e.NewEncoder(), dm_build_1404.transformReaderDst, dm_build_1404.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1404.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1406 *dm_build_1194) Dm_build_1405(dm_build_1407 string, dm_build_1408 string, dm_build_1409 *DmConnection) []byte {
	return Dm_build_1406.Dm_build_1400([]byte(dm_build_1407), dm_build_1408, dm_build_1409)
}

func (Dm_build_1411 *dm_build_1194) Dm_build_1410(dm_build_1412 []byte) byte {
	return Dm_build_1411.Dm_build_1288(dm_build_1412, 0)
}

func (Dm_build_1414 *dm_build_1194) Dm_build_1413(dm_build_1415 []byte) int16 {
	return Dm_build_1414.Dm_build_1292(dm_build_1415, 0)
}

func (Dm_build_1417 *dm_build_1194) Dm_build_1416(dm_build_1418 []byte) int32 {
	return Dm_build_1417.Dm_build_1297(dm_build_1418, 0)
}

func (Dm_build_1420 *dm_build_1194) Dm_build_1419(dm_build_1421 []byte) int64 {
	return Dm_build_1420.Dm_build_1302(dm_build_1421, 0)
}

func (Dm_build_1423 *dm_build_1194) Dm_build_1422(dm_build_1424 []byte) float32 {
	return Dm_build_1423.Dm_build_1307(dm_build_1424, 0)
}

func (Dm_build_1426 *dm_build_1194) Dm_build_1425(dm_build_1427 []byte) float64 {
	return Dm_build_1426.Dm_build_1311(dm_build_1427, 0)
}

func (Dm_build_1429 *dm_build_1194) Dm_build_1428(dm_build_1430 []byte) uint8 {
	return Dm_build_1429.Dm_build_1315(dm_build_1430, 0)
}

func (Dm_build_1432 *dm_build_1194) Dm_build_1431(dm_build_1433 []byte) uint16 {
	return Dm_build_1432.Dm_build_1319(dm_build_1433, 0)
}

func (Dm_build_1435 *dm_build_1194) Dm_build_1434(dm_build_1436 []byte) uint32 {
	return Dm_build_1435.Dm_build_1324(dm_build_1436, 0)
}

func (Dm_build_1438 *dm_build_1194) Dm_build_1437(dm_build_1439 []byte, dm_build_1440 string, dm_build_1441 *DmConnection) []byte {
	if dm_build_1440 == "UTF-8" {
		return dm_build_1439
	}

	if dm_build_1441 == nil {
		if e := dm_build_1447(dm_build_1440); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1439), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1441.encodeBuffer == nil {
		dm_build_1441.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1441.encode = dm_build_1447(dm_build_1441.getServerEncoding())
		dm_build_1441.transformReaderDst = make([]byte, 4096)
		dm_build_1441.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1441.encode; e != nil {

		dm_build_1441.encodeBuffer.Reset()

		n, err := dm_build_1441.encodeBuffer.ReadFrom(
			Dm_build_1461(bytes.NewReader(dm_build_1439), e.NewDecoder(), dm_build_1441.transformReaderDst, dm_build_1441.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1441.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1443 *dm_build_1194) Dm_build_1442(dm_build_1444 []byte, dm_build_1445 string, dm_build_1446 *DmConnection) string {
	return string(Dm_build_1443.Dm_build_1437(dm_build_1444, dm_build_1445, dm_build_1446))
}

func dm_build_1447(dm_build_1448 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1448); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1449 struct {
	dm_build_1450 io.Reader
	dm_build_1451 transform.Transformer
	dm_build_1452 error

	dm_build_1453                []byte
	dm_build_1454, dm_build_1455 int

	dm_build_1456                []byte
	dm_build_1457, dm_build_1458 int

	dm_build_1459 bool
}

const dm_build_1460 = 4096

func Dm_build_1461(dm_build_1462 io.Reader, dm_build_1463 transform.Transformer, dm_build_1464 []byte, dm_build_1465 []byte) *Dm_build_1449 {
	dm_build_1463.Reset()
	return &Dm_build_1449{
		dm_build_1450: dm_build_1462,
		dm_build_1451: dm_build_1463,
		dm_build_1453: dm_build_1464,
		dm_build_1456: dm_build_1465,
	}
}

func (dm_build_1467 *Dm_build_1449) Read(dm_build_1468 []byte) (int, error) {
	dm_build_1469, dm_build_1470 := 0, error(nil)
	for {

		if dm_build_1467.dm_build_1454 != dm_build_1467.dm_build_1455 {
			dm_build_1469 = copy(dm_build_1468, dm_build_1467.dm_build_1453[dm_build_1467.dm_build_1454:dm_build_1467.dm_build_1455])
			dm_build_1467.dm_build_1454 += dm_build_1469
			if dm_build_1467.dm_build_1454 == dm_build_1467.dm_build_1455 && dm_build_1467.dm_build_1459 {
				return dm_build_1469, dm_build_1467.dm_build_1452
			}
			return dm_build_1469, nil
		} else if dm_build_1467.dm_build_1459 {
			return 0, dm_build_1467.dm_build_1452
		}

		if dm_build_1467.dm_build_1457 != dm_build_1467.dm_build_1458 || dm_build_1467.dm_build_1452 != nil {
			dm_build_1467.dm_build_1454 = 0
			dm_build_1467.dm_build_1455, dm_build_1469, dm_build_1470 = dm_build_1467.dm_build_1451.Transform(dm_build_1467.dm_build_1453, dm_build_1467.dm_build_1456[dm_build_1467.dm_build_1457:dm_build_1467.dm_build_1458], dm_build_1467.dm_build_1452 == io.EOF)
			dm_build_1467.dm_build_1457 += dm_build_1469

			switch {
			case dm_build_1470 == nil:
				if dm_build_1467.dm_build_1457 != dm_build_1467.dm_build_1458 {
					dm_build_1467.dm_build_1452 = nil
				}

				dm_build_1467.dm_build_1459 = dm_build_1467.dm_build_1452 != nil
				continue
			case dm_build_1470 == transform.ErrShortDst && (dm_build_1467.dm_build_1455 != 0 || dm_build_1469 != 0):

				continue
			case dm_build_1470 == transform.ErrShortSrc && dm_build_1467.dm_build_1458-dm_build_1467.dm_build_1457 != len(dm_build_1467.dm_build_1456) && dm_build_1467.dm_build_1452 == nil:

			default:
				dm_build_1467.dm_build_1459 = true

				if dm_build_1467.dm_build_1452 == nil || dm_build_1467.dm_build_1452 == io.EOF {
					dm_build_1467.dm_build_1452 = dm_build_1470
				}
				continue
			}
		}

		if dm_build_1467.dm_build_1457 != 0 {
			dm_build_1467.dm_build_1457, dm_build_1467.dm_build_1458 = 0, copy(dm_build_1467.dm_build_1456, dm_build_1467.dm_build_1456[dm_build_1467.dm_build_1457:dm_build_1467.dm_build_1458])
		}
		dm_build_1469, dm_build_1467.dm_build_1452 = dm_build_1467.dm_build_1450.Read(dm_build_1467.dm_build_1456[dm_build_1467.dm_build_1458:])
		dm_build_1467.dm_build_1458 += dm_build_1469
	}
}
