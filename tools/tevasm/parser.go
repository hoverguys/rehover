package main

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type gxArgType uint8

const (
	gxaUint8  gxArgType = iota
	gxaUint32 gxArgType = iota
)

type gxArg struct {
	Type  gxArgType
	Value interface{}
}

type gxCall struct {
	FuncName  string
	Arguments []gxArg
}

func Parse(in io.Reader) ([]gxCall, error) {
	reader := bufio.NewReader(in)
	linen := 0
	last := false
	var calls []gxCall
	for {
		linen++

		// Read next line
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			last = true
		}

		// Trim line
		line = strings.TrimSpace(line)

		// Skip empty lines
		if len(line) < 1 {
			continue
		}

		// Skip single line comments
		if line[0:2] == "//" {
			continue
		}

		call := gxCall{}

		// Make sure it's a function call
		fnindex := strings.IndexRune(line, '(')
		if fnindex < 0 {
			// It's not?!!11
			return nil, fmt.Errorf("line %d: Missing '(', not a function call? [%s]", linen, line)
		}
		call.FuncName = line[:fnindex]

		// Check that the function call is supported
		_, ok := gxfunctions[call.FuncName]
		if !ok {
			return nil, fmt.Errorf("line %d: Function \"%s\" has incorrect name or is not supported", linen, line)
		}

		// Get list of arguments
		endindex := strings.IndexRune(line, ')')
		if endindex < 0 {
			return nil, fmt.Errorf("line %d: Missing ')', functions can't span over multiple line! [%s]", linen, line)
		}
		arglist := strings.Split(line[fnindex+1:endindex], ",")

		// Iterate over each argument
		for _, arg := range arglist {
			arg := strings.TrimSpace(arg)
			if numval, ok := constantsU8[arg]; ok {
				// Argument is a uint8 constant
				call.Arguments = append(call.Arguments, gxArg{
					Type:  gxaUint8,
					Value: numval,
				})
			} else if numval, ok := constantsU32[arg]; ok {
				// Argument is a uint32 constant
				call.Arguments = append(call.Arguments, gxArg{
					Type:  gxaUint32,
					Value: numval,
				})
			} else if unicode.IsDigit(rune(arg[0])) {
				// Argument might be a number
				num, err := strconv.Atoi(arg)
				if err != nil {
					return nil, fmt.Errorf("line %d: Argument \"%s\" is not a valid number", linen, arg)
				}
				call.Arguments = append(call.Arguments, gxArg{
					Type:  gxaUint32,
					Value: uint32(num),
				})
			} else {
				return nil, fmt.Errorf("line %d: Argument \"%s\" is not a valid value", linen, arg)
			}
		}

		// Add call to list
		calls = append(calls, call)

		if last {
			break
		}
	}

	return calls, nil
}

func (gx gxCall) String() (out string) {
	out += gx.FuncName + " [ "
	for _, arg := range gx.Arguments {
		out += fmt.Sprintf("%v ", arg)
	}
	out += "]"

	return
}

func (gx gxCall) Call() error {
	fn := reflect.ValueOf(gxfunctions[gx.FuncName])
	fntype := fn.Type()
	if len(gx.Arguments) != fntype.NumIn() {
		return fmt.Errorf("wrong number of arguments (required %d but provided %d)", fntype.NumIn(), len(gx.Arguments))
	}
	in := make([]reflect.Value, len(gx.Arguments))
	for k, arg := range gx.Arguments {
		Oargtype := fntype.In(k)
		Pargtype := reflect.TypeOf(arg)
		// Make sure that required arg and provided arg are the same type
		if Oargtype.Kind() != Pargtype.Kind() {
			// Get original argument in a neutral format (uint32)
			var arg32 uint32
			switch arg.Type {
			case gxaUint32:
				arg32 = arg.Value.(uint32)
			case gxaUint8:
				arg32 = uint32(arg.Value.(uint8))
			}

			// Convert arg32 to the required parameter
			switch Oargtype.Kind() {
			case reflect.Uint8:
				in[k] = reflect.ValueOf(uint8(arg32))
			case reflect.Uint32:
				in[k] = reflect.ValueOf(arg32)
			}
		} else {
			in[k] = reflect.ValueOf(arg.Value)
		}

	}
	fn.Call(in)
	return nil
}
