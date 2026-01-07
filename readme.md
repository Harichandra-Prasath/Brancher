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

# If needed to do Auth operations like pull, PV_KEY_FILE is expected which contains the private ssh key file under $HOME/.ssh/, then
PV_KEY_FILE=id_rsa Brancher
```
This will bring up the TUI of the Brancher  

#### Key-Bindings

`c` - checks out the selected branch  
`d` - deletes the selected branch  
`n` - creates a new branch  
`r` - renames the selected branch  
`p` - pulls the selected branch

General status messages will be displayed in the lower bar for better experience.  
 

