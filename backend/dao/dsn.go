package dao

// 此处的内容可以根据自己的数据库进行更改
var DBUser = "fzsgball"
var DBPassword = "NTSyyds0927"
var DBHost = "127.0.0.1"
var DBPort = "3306"
var DBName = "official_book_library"

var DSN = DBUser + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
