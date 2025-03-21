# colorlines
Draw random colored lines in your terminal


## Installation one time
```bash
go mod init colorlines
go mod tidy
```

## Build

    ./build.sh

Tag the version

    git tag -a v0.1.0 -m "First release"
    git push origin v0.1.0

## Usage

    Usage of colorlines:
      -colors int
            Number of colors per line (default 5)
      -lines int
            Number of lines to display (default 1)
      -piece_size int
            Number of characters per piece (default 5)  -save_config
      -save_config
            Save the configuration to the file
      -shape string
            Use 'blocks', 'chars' or 'smileys' characters to display (default "blocks")
      -version
            Display the version of the program

The program will display random generated colored characters on the terminal using the parameters specified.

If a config file `$HOME/.config/colorlines/config.json` exists, it will be used to read the values. 
The commandline parameters will always overrule the values from the configuration file.

If the `-save_config` flag is used, the configuration will be saved to the file. That way you can use colorlines without parameters the next time and get the same configuration. Note that the `-save_config` will always be saved as `false` to avoid constantly saving the configuration.

## Examples

    colorlines 

    colorlines -version

    colorlines -help

    colorlines -lines 2 -colors 10 -piece_size 4 -shape blocks
    
    colorlines --lines 2 --colors 10 --piece_size 4 --shape chars

    colorlines -lines 1 -save_config

