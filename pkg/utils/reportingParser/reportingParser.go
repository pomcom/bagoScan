package reportingparser

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
)

func ReportingParser() {
	outputFile, err := os.Open("output/raw/combined-output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer outputFile.Close()

	// Create a new CSV file
	csvFile, err := os.Create("output_nessus_format.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write the headers to the CSV file
	header := []string{
		"Plugin ID",
		"CVE",
		"CVSS v2.0 Base Score",
		"Risk",
		"Host",
		"Protocol",
		"Port",
		"Name",
		"Synopsis",
		"Description",
		"Solution",
		"See Also",
		"Plugin Output",
		"STIG Severity",
		"CVSS v3.0 Base Score",
		"CVSS v2.0 Temporal Score",
		"CVSS v3.0 Temporal Score",
		"Risk Factor",
		"BID",
		"XREF",
		"MSKB",
		"Plugin Publication Date",
		"Plugin Modification Date",
		"Metasploit",
		"Core Impact",
		"CANVAS",
	}

	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing header to CSV file:", err)
		return
	}

	//Regex patterns sqlmap - check for valid SQL injection

	// hostRegex := regexp.MustCompile(`GET (.*?)/rest`)
	sqlInjectionRegex := regexp.MustCompile(`do you want to exploit this SQL injection\? \[Y/n\] Y`)
	outputRegex := regexp.MustCompile(`\[INFO\] retrieved: (.*?)\n`)

	// injectionRegex := regexp.MustCompile(`(\d+)\s.*\n.*(GET)\n.*(\w+-based.*)\n.*(\w+.*)\n`)

	var host string
	var pluginOutput string
	var sqlInjectionFound bool

	scanner := bufio.NewScanner(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		if sqlInjectionRegex.MatchString(line) {
			sqlInjectionFound = true
			// host = hostRegex.FindStringSubmatch(line)[1]
		} else if sqlInjectionFound && outputRegex.MatchString(line) {
			pluginOutput = outputRegex.FindStringSubmatch(line)[1]
		}
	}

	// Write the SQL injection information to the CSV file
	data := []string{
		"123456",
		"",
		"",
		"High",
		host,
		"HTTP",
		"8080",
		"SQL Injection Vulnerability",
		"A SQL injection vulnerability has been discovered in this application.",
		"The application is vulnerable to SQL injection attacks, which can allow an attacker to execute arbitrary SQL commands on the underlying database. This can result in the theft of sensitive data, such as user credentials and other confidential information.",
		"Apply patches or upgrades to the application to eliminate the vulnerability.",
		"",
		pluginOutput,
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}

	if sqlInjectionFound {
		// Write the data to the CSV file
		err = writer.Write(data)
		if err != nil {
			fmt.Println("Error writing data to CSV file:", err)
			return
		}
	}

	defer csvFile.Close()
	defer writer.Flush()
	fmt.Println("Data written to Nessus format CSV file successfully")

}
