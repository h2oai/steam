package db

service Service

type Sys
 Version uint32

type Cloud
 ApplicationID string
 Size int
 
type Model
 CloudName string
 CloudAddress string
 Data []byte

type Service
 Caption string
 Description string
 Source string
 Target string
 IsBuilt bool

type Engine
 Name string


