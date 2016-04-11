package db

service Service

type Sys
 Version uint32

type Model
 ModelID string
 CloudName string
 CloudAddress string
 Data []byte

type Service
 Caption string
 Description string
 Source string
 Target string
 IsBuilt bool



