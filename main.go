package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	for scanner.Scan(){
		ScanMail(scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		log.Fatal("Error while scanning input: %v \n", err)
	}
}

func ScanMail(domain string){

	var hasMx, hasSPF, hasDMARC bool
	var  spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil{
		fmt.Printf("error: %v", err)
	}

	if len(mxRecord) > 0 {
		hasMx = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	for _, record := range txtRecord{
		if strings.HasPrefix(domain, "v=spf1"){
			spfRecord = record
			hasSPF = true
			break
		}
	}

	RecordDmarc, err := net.LookupTXT("_dmarc." + domain)
	if err != nil{
		fmt.Printf("Error: %v", err)
	}

	for _, record := range RecordDmarc{
		if strings.HasPrefix(record, "v=DMARC1"){
			dmarcRecord = record
			hasDMARC = true
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}