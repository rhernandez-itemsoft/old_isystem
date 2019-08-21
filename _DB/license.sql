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

 Date: 25/07/2019 23:52:16
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
