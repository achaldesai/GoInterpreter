package object

import (
	"GoInterpreter/ast"
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Hashable interface {
	HashKey() HashKey
}

type BuiltinFunction func(args ...Object) Object

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
)

// INTEGER struct type
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// BOOLEAN object type
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// NULL object Type
type Null struct{}

func (nu *Null) Type() ObjectType { return NULL_OBJ }
func (nu *Null) Inspect() string  { return "null" }

// RETURN object Type
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// ERROR object type
type Error struct {
	Message string
}

func (err *Error) Type() ObjectType { return ERROR_OBJ }
func (err *Error) Inspect() string  { return "ERROR: " + err.Message }

// FUNCTION object type
type Function struct {
	Body       *ast.BlockStatement
	Env        *Environment
	Parameters []*ast.Identifier
}

func (fn *Function) Type() ObjectType { return FUNCTION_OBJ }
func (fn *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fn.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") {\n")
	out.WriteString(fn.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// STRING object type
type String struct {
	Value string
}

func (str *String) Type() ObjectType { return STRING_OBJ }
func (str *String) Inspect() string  { return str.Value }

// BUILTINFUNCTIONS type definition
type Builtin struct {
	Fn BuiltinFunction
}

func (bf *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (bf *Builtin) Inspect() string  { return "builtin function" }

// ARRAY object type definition
type Array struct {
	Elements []Object
}

func (arr *Array) Type() ObjectType { return ARRAY_OBJ }
func (arr *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range arr.Elements {
		elements = append(elements, el.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashKey type definition
type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HashTree object type definition
type HashPair struct {
	Key   Object
	Value Object
}

// Hash object type definition
type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
