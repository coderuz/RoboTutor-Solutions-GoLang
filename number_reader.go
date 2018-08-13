package main

import (
	"fmt"
		)

func main() {
	fmt.Print("Sonni kiriting: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(reader(input))
}

func reader(son string) string {
	uch_xonaliklar := []string{"", "ming ","million ", "milliard ", "trillion ", "kvatrillion ", "kvintillion ", "sekstillion ","septillion ", "oktillion ","nonillion ", "detsillion ","undetsillion ", "dodetsillion ","tredetsillion ","kvattuordetsillion "}
	sonUzunligi := int(len(son))
	for sonUzunligi%3!=0 {
		son = "0"+son
		sonUzunligi++
	}
	//fmt.Println(son)
	for i:=0;i<sonUzunligi/3;i++{
		beingRead := son[i*3:(i+1)*3]
		littleReader(beingRead)
		if beingRead !="000" {
			fmt.Print(uch_xonaliklar[int(sonUzunligi/3)-i-1])
		}
	}
	return ""
}

func littleReader(n string) string {
	son_uzunligi := len(n)
	for i,raqam := range n {
		switch son_uzunligi-i {
		case 1:	birlik(string(raqam))
		case 2: onlik(string(raqam))
		case 3: birlik(string(raqam)); if string(raqam)!="0" {fmt.Print("yuz ")}
		}
	}
	return ""
}

func birlik(raqam string) string  {
	switch string(raqam) {
		case "1": fmt.Print("bir", " ")
		case "2": fmt.Print("ikki", " ")
		case "3": fmt.Print("uch", " ")
		case "4": fmt.Print("to'rt", " ")
		case "5": fmt.Print("besh", " ")
		case "6": fmt.Print("olti", " ")
		case "7": fmt.Print("yetti", " ")
		case "8": fmt.Print("sakkiz", " ")
		case "9": fmt.Print("to'qqiz", " ")
	}
	return ""
}

func onlik(raqam string) string {
	switch string(raqam) {
	case "1": fmt.Print("o'n", " ")
	case "2": fmt.Print("yigirma", " ")
	case "3": fmt.Print("o'ttiz", " ")
	case "4": fmt.Print("qirq", " ")
	case "5": fmt.Print("ellik", " ")
	case "6": fmt.Print("oltmish", " ")
	case "7": fmt.Print("yetmish", " ")
	case "8": fmt.Print("sakson", " ")
	case "9": fmt.Print("to'qson", " ")
	}
	return ""
}