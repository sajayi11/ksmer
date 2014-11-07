package main 

import (
	"fmt"
//	"math"
	"io/ioutil"
	"os"
	"encoding/gob"
)

var gkcMap map[uint32]map[uint16]int;

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getBaseId(b byte) uint32 {
	switch(b) {
		case 'A':
			return 0;
		case 'C':
			return 1;
		case 'G':
			return 2;
		case 'T':
			return 3;
		case 'N':
			return 4;
	}
	return 4;
}

func store(genomeId uint16, kmerId uint32) {
	gMap, ok := gkcMap[kmerId];
	if !ok {
		gMap = make(map[uint16]int);
		gkcMap[kmerId] = gMap;
	}
	gMap[genomeId]++;
}

func storeContigMappings(genomeId uint16, contigStr []byte, k uint32) {
	var i uint32;
	var kmerVal uint32 = 0;
	for i=0; i<(k-1); i++ {
		kmerVal = (kmerVal<<2) + getBaseId(contigStr[i]);
		//fmt.Printf("%d\n", kmerVal);
	}
	bitMask := uint32((1<<(k<<1))-1);
	length := uint32(len(contigStr));
	for i=k-1; i<length; i++ {
		kmerVal = (kmerVal<<2 & bitMask) + getBaseId(contigStr[i]);
		store(genomeId, kmerVal);
		//fmt.Printf("%d %c %d %d\n", i, contigStr[i], getBaseId(contigStr[i]), kmerVal);
	}
}

func storeContigMappingsSpaced(genomeId uint16, contigStr []byte, k1 uint32, s uint32, k2 uint32) {
	var i,j uint32;
	var kmerValFirst uint32 = 0;
	var kmerValSecond uint32 = 0;
	
	for i=0; i<(k1-1); i++ {
		kmerValFirst = (kmerValFirst<<2) + getBaseId(contigStr[i]);
	}
	for i=k1+s; i<(k1+s+k2-1); i++ {
		kmerValSecond = (kmerValSecond<<2) + getBaseId(contigStr[i]);
	}
	
	bitMask1 := uint32((1<<(k1<<1))-1);
	bitMask2 := uint32((1<<(k2<<1))-1);
	length := uint32(len(contigStr));
	var kmerVal uint32;
	i=k1-1;
	for j=k1+s+k2-1; j<length; j++ {
		kmerValFirst = (kmerValFirst<<2 & bitMask1) + getBaseId(contigStr[i]);
		kmerValSecond = (kmerValSecond<<2 & bitMask2) + getBaseId(contigStr[j]);
		kmerVal = kmerValFirst<<(k2<<1) + kmerValSecond;
		store(genomeId, kmerVal);
		fmt.Printf("%d %c %d %d\n", j, contigStr[j], getBaseId(contigStr[j]), kmerVal);
		i++;
	}
}


func saveIndexMapToFile(filePath string, m map[uint32]map[uint16]int) {
	fp, err := os.Create(filePath);
	if err != nil {
		panic("cant open file");
	}
	defer fp.Close();
	
	enc := gob.NewEncoder(fp);
	if err := enc.Encode(m); err != nil {
		panic("cant encode");
	}
}

func loadIndexMapToFile(filePath string) (m map[uint32]map[uint16]int) {
	fp, err := os.Open(filePath);
	if err != nil {
		panic("cant open file");
	}
	defer fp.Close();
	
	enc := gob.NewDecoder(fp);
	if err := enc.Decode(&m); err != nil {
		panic("cant decode");
	}
	return m;
}

func main() {
	
	gkcMap = make(map[uint32]map[uint16]int);
	
	contigStr, err := ioutil.ReadFile("input\\sample.contig");
	check(err);
	
	for i:=0; i<(len(contigStr)); i++ {
		//fmt.Printf("%d ", getBaseId(contigStr[i]));
	}
	
	
	genomeId := uint16(0);
	k := uint32(4);
	storeContigMappings(genomeId, contigStr, k);
	
	/*
	genomeId := uint16(0);
	k1 := uint32(4);
	s := uint32(3);
	k2 := uint32(4);
	storeContigMappingsSpaced(genomeId, contigStr, k1, s, k2);
	*/
	
	saveIndexMapToFile("input\\index.map", gkcMap);
	//gkcMap = loadIndexMapToFile("input\\index.map");
	
	for kmerId, gMap := range gkcMap {
		fmt.Print("kmerId:", kmerId);
		for genomeId, count := range gMap {
			fmt.Println(" genomeId:", genomeId, " count:", count);
			//fmt.Println("count:", gkcMap[kmerId][genomeId]);
		}
	}
	

}

