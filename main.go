package main

import "fmt"

const sample = `A0896/22 NOTAMN
Q) ZSHA/QRDCA/IV/BO/W/000/999/2908N12240E016
A) ZSHA B) 2204290320 C) 2204290437 
E) A TEMPORARY DANGER AREA ESTABLISHED,THE UNBURNED DEBRIS IS 
EXPECTED TO FALL IN THE N2909E12241,THE POSSIBLE FALLING AREA WILL 
NOT EXCEED THE RANGE WITHIN THE 
N2858E12227-N2924E12236-N2919E12254-N2854E12245,FOUR POINT 
CONNECTION RANGE. VERTICAL ALTITUDE:SFC-UNL.
F) SFC G) UNL`

func main() {
	icaoParser := NewParser()
	if header, seriesMap, err := icaoParser.tokenize(sample); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
	} else {
		fmt.Println(fmt.Sprintf("Header: %s", *header))
		for _, code := range itemCodeInOrder {
			if line, lineOK := seriesMap[code]; lineOK {
				fmt.Println(fmt.Sprintf("Series %s: %s", code, line))
			}
		}
	}

}
