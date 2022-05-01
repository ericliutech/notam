package main

import (
	"errors"
	"strings"
)

const (
	itemCodeQ string = "Q)"
	itemCodeA string = "A)"
	itemCodeB string = "B)"
	itemCodeC string = "C)"
	itemCodeD string = "D)"
	itemCodeE string = "E)"
	itemCodeF string = "F)"
	itemCodeG string = "G)"
)

var itemCodeInOrder = []string{
	itemCodeQ,
	itemCodeA,
	itemCodeB,
	itemCodeC,
	itemCodeD,
	itemCodeE,
	itemCodeF,
	itemCodeG,
}

type itemName uint8

const (
	itemNameQualifier itemName = iota
	itemNameLocation
	itemNameFrom
	itemNameTo
	itemNameSchedule
	itemNameCondition
	itemNameLower
	itemNameUpper
)

var (
	itemNameToCodeMap = map[itemName]string{
		itemNameQualifier: itemCodeQ,
		itemNameLocation:  itemCodeA,
		itemNameFrom:      itemCodeB,
		itemNameTo:        itemCodeC,
		itemNameSchedule:  itemCodeD,
		itemNameCondition: itemCodeE,
		itemNameLower:     itemCodeF,
		itemNameUpper:     itemCodeG,
	}
	itemCodeToNameMap = map[string]itemName{
		itemCodeQ: itemNameQualifier,
		itemCodeA: itemNameLocation,
		itemCodeB: itemNameFrom,
		itemCodeC: itemNameTo,
		itemCodeD: itemNameSchedule,
		itemCodeE: itemNameCondition,
		itemCodeF: itemNameLower,
		itemCodeG: itemNameUpper,
	}
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) findQualifierIndex(text string) int {
	return p.findItemStartingIndex(text, itemCodeQ)
}

func (p *Parser) findItemStartingIndex(text, itemCode string) int {
	idx := strings.Index(text, itemCode)
	return idx
}

func (p *Parser) tokenize(text string) (*string, map[string]string, error) {
	presentCodeList := make([]string, 0)
	presentCodeStartingIndices := make([]int, 0)
	itemMap := map[string]string{}
	var qIdx int

	for _, code := range itemCodeInOrder {
		idx := p.findItemStartingIndex(text, code)
		if idx >= 0 {
			presentCodeList = append(presentCodeList, code)
			presentCodeStartingIndices = append(presentCodeStartingIndices, idx)
			if code == itemCodeQ {
				qIdx = idx
			}
		}
	}

	itemsLen := len(presentCodeList)

	for pos, code := range presentCodeList {
		if pos == itemsLen-1 {
			itemMap[code] = (strings.TrimSpace(text[presentCodeStartingIndices[pos]+2:]))
			break
		}
		itemMap[code] = strings.TrimSpace(text[presentCodeStartingIndices[pos]+2 : presentCodeStartingIndices[pos+1]])
	}

	// first line
	firstLine := strings.TrimSpace(text[:qIdx])

	if len(firstLine) == 0 {
		return nil, nil, errors.New("invalid format: empty header row")
	}

	if len(itemMap) == 0 {
		return nil, nil, errors.New("invalid format: contains no items")
	}

	return &firstLine, itemMap, nil
}
