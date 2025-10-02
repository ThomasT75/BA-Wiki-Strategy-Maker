package main

import (
	"fmt"
	"os"
	"strconv"
)

// token list type
type TokenList struct {
	list []string
	idx int
	lineNum int
	fieldNum int
}

func NewTokenList() TokenList {
	return TokenList{
		lineNum: 1,
	}
}

// add more tokens
func (tl *TokenList) Append(s ...string) {
	tl.list = append(tl.list, s...)
}

// get current token text
func (tl *TokenList) GetCurrent() string {
	if tl.idx >= len(tl.list) {
		return "EOF"
	}
	return tl.list[tl.idx]
}

// check current token
func (tl *TokenList) IsCurrent(t string) bool {
	if tl.GetCurrent() == t {
		return true
	}
	return false
}

func (tl *TokenList) IsCurrentAny(ts ...string) bool {
	for _, t := range ts {
		if tl.IsCurrent(t) {
			return true
		}
	}
	return false
}

// next token 
func (tl *TokenList) Next() {
	if tl.IsCurrent("\n") {
		tl.lineNum += 1
		tl.fieldNum = 1
	} else {
		tl.fieldNum += 1
	}
	tl.idx += 1
}

// required token
func (tl *TokenList) Expect(t string) bool {
	if tl.Accept(t) {
		return true
	}
	tl.ErrTxt("Expected token: %s got: %s\n", t, tl.GetCurrent())
	os.Exit(2)
	return false
}

func (tl *TokenList) ExpectAny(ts ...string) bool {
	if tl.AcceptAny(ts...) {
		return true
	}
	tl.ErrTxt("Expected any of: %#v got: %s\n", ts, tl.GetCurrent())
	os.Exit(2)
	return false
}

// optional token
func (tl *TokenList) Accept(t string) bool {
	if tl.IsCurrent(t) {
		tl.Next()
		return true
	}
	return false
}

func (tl *TokenList) AcceptAny(ts ...string) bool {
	if tl.IsCurrentAny(ts...) {
		tl.Next()
		return true
	}
	return false
}

func (tl *TokenList) Err() {
	tl.ErrTxt("Invalid token: %s\n", tl.GetCurrent())
}

func (tlo *TokenList) errGetCurrentLine() string {
	var str string
	var tl = *tlo
	var lineNum = tl.lineNum
	tl.idx = 0 
	for range lineNum {
		str = ""
		for !tl.IsCurrent("\n") {
			str += tl.GetCurrent()
			tl.Next()
			if !tl.IsCurrent("\n") {
				str += " "
			}
		}
		tl.Next()
	}
	return str
}

func (tl *TokenList) ErrTxt(format string, a ...any) {
	for i, v := range a {
		switch value := v.(type) {
		case string:
			a[i] = strconv.Quote(value)
		default:
			a[i] = value
		}
	}
	fstr := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "L%d:%s:F%d:%s\n  Error: %s", tl.lineNum, strconv.Quote(tl.errGetCurrentLine()), tl.fieldNum, strconv.Quote(tl.GetCurrent()), fstr)
	os.Exit(2)
}
