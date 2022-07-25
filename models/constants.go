package models

type Operations string
type Handler func(operand float64)

type FunctionMap map[Operations]Handler
