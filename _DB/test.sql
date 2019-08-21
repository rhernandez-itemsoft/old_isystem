/*
 Navicat Premium Data Transfer

 Source Server         : Mobngo - localhost
 Source Server Type    : MongoDB
 Source Server Version : 40200
 Source Host           : localhost:27017
 Source Schema         : test

 Target Server Type    : MongoDB
 Target Server Version : 40200
 File Encoding         : 65001

 Date: 25/07/2019 23:52:47
*/


// ----------------------------
// Collection structure for log_access
// ----------------------------
db.getCollection("log_access").drop();
db.createCollection("log_access");

// ----------------------------
// Collection structure for log_httprequest
// ----------------------------
db.getCollection("log_httprequest").drop();
db.createCollection("log_httprequest");

// ----------------------------
// Collection structure for users
// ----------------------------
db.getCollection("users").drop();
db.createCollection("users");
db.getCollection("users").createIndex({
    username: NumberInt("1")
}, {
    name: "llaveunica",
    unique: true
});
db.getCollection("users").createIndex({
    token: NumberInt("1")
}, {
    name: "token",
    unique: true
});
db.getCollection("users").createIndex({
    email: NumberInt("1"),
    password: NumberInt("1")
}, {
    name: "signin",
    unique: true
});
