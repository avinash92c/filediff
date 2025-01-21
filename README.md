# filediff
just because
if you want to generate a simple diff report between 2 files you use this.
why did i make this. for reasons

# How to Build
```cmd
go build -o diff.exe
```

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
this can handle files from 2 files upwards to n files. Go nuts
just because you can abuse this program doesnt mean you should pass n files. do those kind of stuff somewhere else
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
why
just because
not because i can't write
i didn't want to spend too much time on this trivial thing