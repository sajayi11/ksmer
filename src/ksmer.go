package main 

import (
	"fmt"
//	"math"
	"io/ioutil"

)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getBaseId(b byte) int {
	switch(b) {
		case 'A':
			return 0;
		case 'C':
			return 1;
		case 'T':
			return 2;
		case 'G':
			return 3;
		case 'N':
			return -1;
	}
	return -1;
}

func main() {
	k := 16;
	//s := 0;
	
	str, err := ioutil.ReadFile("input\\sample.contig");
	check(err);
	
	for i:=0; i<(len(str)-k); i++ {
		fmt.Printf("%c %d\n", str[i], getBaseId(str[i]));
	}
	
	//fmt.Println(str);

}

