package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"flag"
	"github.com/jung-kurt/gofpdf"
	"strings"
)

// compareFiles compares multiple files line by line and generates the result in the specified format.
func compareFiles(files []string, outputFormat string) error {
	// Open all files
	var fileHandlers []*os.File
	var scanners []*bufio.Scanner
	var differences []string

	// Open each file and prepare a scanner
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("could not open file %s: %v", file, err)
		}
		fileHandlers = append(fileHandlers, f)
		scanners = append(scanners, bufio.NewScanner(f))
		defer f.Close() // Ensure files are closed when the function exits
	}

	lineNumber := 1
	for {
		// Read one line from each file
		var lines []string
		var hasLine bool

		for _, scanner := range scanners {
			if scanner.Scan() {
				line := scanner.Text()
				if line != "" { // Skip blank lines
					lines = append(lines, line)
				} else {
					lines = append(lines, "")
				}
				hasLine = true
			} else {
				// If any file has no more lines, we mark it as having no line
				lines = append(lines, "")
			}
		}

		if !hasLine {
			break // End of all files
		}

		// Skip comparing blank lines entirely by ensuring we only compare non-empty lines
		var nonBlankLines []string
		for _, line := range lines {
			if line != "" {
				nonBlankLines = append(nonBlankLines, line)
			}
		}

		// If there are non-blank lines left, compare them
		if len(nonBlankLines) > 0 {
			// Check if all lines are the same or different
			for i, _ := range nonBlankLines {
				for j := i + 1; j < len(nonBlankLines); j++ {
					if nonBlankLines[i] != nonBlankLines[j] {
						diff := fmt.Sprintf("Line %d differs between file %s and file %s:\n  %s: %s\n  %s: %s",
							lineNumber, files[i], files[j], files[i], nonBlankLines[i], files[j], nonBlankLines[j])
						differences = append(differences, diff)
					}
				}
			}
		} else {
			// If the lines are empty (i.e., they only contained blank lines), skip them entirely
			lineNumber++
			continue
		}

		// Check for any lines that are missing in some files (ignore blank lines)
		for i, line := range lines {
			if line == "" {
				differences = append(differences, fmt.Sprintf("Line %d only in %s", lineNumber, files[i]))
			}
		}

		lineNumber++
	}

	// Check for any scanner errors
	for i, scanner := range scanners {
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading file %s: %v", files[i], err)
		}
	}

	// Generate the output in the selected format
	switch outputFormat {
	case "pdf":
		err := generatePDFReport(differences)
		if err != nil {
			return fmt.Errorf("failed to generate PDF report: %v", err)
		}
	case "html":
		err := generateHTMLReport(differences)
		if err != nil {
			return fmt.Errorf("failed to generate HTML report: %v", err)
		}
	case "text":
		err := generateTextReport(differences)
		if err != nil {
			return fmt.Errorf("failed to generate text report: %v", err)
		}
	default:
		return fmt.Errorf("invalid output format: %s", outputFormat)
	}

	return nil
}

// generatePDFReport generates a PDF report with the differences
func generatePDFReport(differences []string) error {
	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a page to the PDF
	pdf.AddPage()

	// Set font for the PDF
	pdf.SetFont("Arial", "", 12)

	// Add title to the PDF
	pdf.Cell(0, 10, "File Comparison Report")
	pdf.Ln(10)

	// Add the differences to the PDF
	for _, diff := range differences {
		pdf.MultiCell(0, 10, diff, "", "", false)
		pdf.Ln(5) // Line break between differences
	}

	// Output the PDF to a file
	err := pdf.OutputFileAndClose("report.pdf")
	if err != nil {
		return fmt.Errorf("failed to save PDF: %v", err)
	}

	log.Println("PDF report generated: report.pdf")
	return nil
}

// generateHTMLReport generates an HTML report with the differences
func generateHTMLReport(differences []string) error {
	// Create the HTML content
	htmlContent := "<html><head><title>File Comparison Report</title></head><body>"
	htmlContent += "<h1>File Comparison Report</h1>"

	// Add the differences to the HTML
	for _, diff := range differences {
		htmlContent += "<p>" + strings.ReplaceAll(diff, "\n", "<br>") + "</p>"
	}

	htmlContent += "</body></html>"

	// Save the HTML to a file
	err := os.WriteFile("report.html", []byte(htmlContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to save HTML report: %v", err)
	}

	log.Println("HTML report generated: report.html")
	return nil
}

// generateTextReport generates a plain text report with the differences
func generateTextReport(differences []string) error {
	// Open a text file to write the report
	file, err := os.Create("report.txt")
	if err != nil {
		return fmt.Errorf("failed to create text file: %v", err)
	}
	defer file.Close()

	// Write the differences to the text file
	for _, diff := range differences {
		_, err := file.WriteString(diff + "\n\n")
		if err != nil {
			return fmt.Errorf("failed to write to text file: %v", err)
		}
	}

	log.Println("Text report generated: report.txt")
	return nil
}

func main() {
	// Define command-line flags for file paths and output format
	var outputFormat string
	var files []string

	// Define the format flag with a default value of "text"
	flag.StringVar(&outputFormat, "format", "text", "Output format: text, pdf, or html")

	// Parse the flags
	flag.Parse()

	// Retrieve files from the command-line arguments (remaining args after flags)
	for _, arg := range flag.Args() {
		files = append(files, arg)
	}

	// Ensure at least two files are provided
	if len(files) < 2 {
		log.Fatal("You must provide at least two file paths to compare.")
	}

	// Call the compareFiles function with the file paths and output format
	if err := compareFiles(files, outputFormat); err != nil {
		log.Fatal(err)
	}
}
