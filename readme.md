# Brancher

Git Branch Manager for quick checkout, creation and deletion.

### Quick Start

Install it via `go`:  
```bash
go install github.com/Harichandra-Prasath/Brancher@latest
```  

Build:  
```bash
git clone https://github.com/Harichandra-Prasath/Brancher.git
cd Brancher
make build

# Install it 
go install .
```

### Usage

Run the binary (Assumed in `$PATH`)
```bash
Brancher
```
This will bring up the TUI of the Brancher  

#### Key-Bindings

`c` - checks out the selected branch  
`d` - deletes the selected branch  
`n` - creates a new branch  
`r` - renames the selected branch  

General status messages will be displayed in the lower bar for better experience.  
 

