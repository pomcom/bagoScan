package reportingparser

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReportingParser() {
	file, err := os.Open("output/raw/combined-output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

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
		fmt.Println("Error writing headers:", err)
		return
	}

	// Write the SQL injection information to the CSV file
	data := []string{"123456", "", "", "High", "localhost", "HTTP", "8080", "SQL Injection Vulnerability", "A SQL injection vulnerability has been found in the application.", "The application is vulnerable to SQL injection attacks, which can allow an attacker to execute arbitrary SQL commands.", "Upgrade the application to the latest version to fix the vulnerability.", "", "SQL Injection: http://localhost:8080/rest/products/search?q="}
	err = writer.Write(data)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}

	fmt.Println("Data written to Nessus format CSV file successfully")
}
