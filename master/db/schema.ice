package db

service Service

type Sys
 Version uint32

type Cloud
 EngineName string
 Size int
 ApplicationID string
 Address string
 Memory string
 Username string
 State string
 
type Model
 CloudName string
 Dataset string
 TargetName string
 MaxRuntime int
 JavaModelPath string
 GenModelPath string

type ScoringService
 ModelName string
 Address string
 Port int
 State string

type Engine
 Name string


