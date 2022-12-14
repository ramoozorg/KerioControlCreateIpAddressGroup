package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	os.Remove("ipsKerio.xml")
	kerioXmlFile, err := os.OpenFile("ipsKerio.xml",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer kerioXmlFile.Close()
	iranIps := readLines("ips.txt")
	for i := 0; i < len(iranIps); i++ {
		s := fmt.Sprintf("<listitem>\n    <variable name=\"Id\">%d</variable>\n    <variable name=\"Enabled\">1</variable>\n    <variable name=\"Desc\">Iran</variable>\n    <variable name=\"Name\">IranIps</variable>\n    <variable name=\"Value\">prefix:%s</variable>\n    <variable name=\"SharedId\">0</variable>\n</listitem>\n", i+100, iranIps[i])

		if _, err := kerioXmlFile.WriteString(s); err != nil {
			log.Println(err)
		}
	}
}

func readLines(fileName string) []string {
	bytesRead, _ := ioutil.ReadFile(fileName)
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	return lines
}
