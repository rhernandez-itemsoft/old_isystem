/*
 Navicat Premium Data Transfer

 Source Server         : Mobngo - localhost
 Source Server Type    : MongoDB
 Source Server Version : 40200
 Source Host           : localhost:27017
 Source Schema         : license

 Target Server Type    : MongoDB
 Target Server Version : 40200
 File Encoding         : 65001

 Date: 25/07/2019 23:52:03
*/


// ----------------------------
// Collection structure for licenses
// ----------------------------
db.getCollection("licenses").drop();
db.createCollection("licenses");
db.getCollection("licenses").createIndex({
    company: NumberInt("1"),
    store: NumberInt("1")
}, {
    name: "unica",
    unique: true
});
db.getCollection("licenses").createIndex({
    uid: NumberInt("1")
}, {
    name: "kuid"
});
db.getCollection("licenses").createIndex({
    serial: NumberInt("1")
}, {
    name: "kserial"
});

// ----------------------------
// Documents of licenses
// ----------------------------
db.getCollection("licenses").insert([ {
    _id: ObjectId("5d27832a2b44e510b9b250d8"),
    company: "ItemSoftMX",
    store: "Tienda 1",
    uid: "WDC_WD5000AAKS-00TMA0_WD-WCAPW4822028",
    token: "9136706b091a8f321d34b0249c313d62830a6e821ca132c66be86f65948dcf7f719babb7207f3e95cb8e7b29ccf1543f6046d44a91b9af3c63d4bc5e3238c198",
    "created_at": NumberLong("1562870570237309200"),
    "updated_at": NumberLong("1563375719168113700"),
    "uiupdated_atd": NumberLong("1562988319524917600")
} ]);
