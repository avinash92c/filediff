# filediff
if you want to generate a simple diff report between 2 files you use this.

# How to Build
For Windows:
```cmd
go build -o diff.exe
```

For Linux:
```bash
GOOS=linux GOARCH=amd64 go build -o diff
```
# How to run
```
.\diff -format=text "\path\to\file1.yml" "\path\to\file2.yml"
```

# Report Formats
I'm Just printing the same diff in text format in pdf/html/text formats, no special decorations
set format to text or pdf or html to specify format
default is text

# Note
```
.\diff -format=text "\path\to\file1.yml" "\path\to\file2.yml" "\path\to\file3.yml"
```

# Output Format
```
Line 1 differs between file G:\git\filediff\testfiles\yml\file1.yml and file G:\git\filediff\testfiles\yml\file2.yml:
  G:\git\filediff\testfiles\yml\file1.yml: # Development Environment Configuration
  G:\git\filediff\testfiles\yml\file2.yml: # Staging Environment Configuration

Line 2 differs between file G:\git\filediff\testfiles\yml\file1.yml and file G:\git\filediff\testfiles\yml\file2.yml:
  G:\git\filediff\testfiles\yml\file1.yml: dev:
  G:\git\filediff\testfiles\yml\file2.yml: staging:\
```

# Footnote
I didn't write this program. i just generated it and tested the results.
