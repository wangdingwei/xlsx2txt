## usage 

    xlsx inputfile outputfile

## beyond compare setting

* open **File Formats** setting dialog
* new **Text Format**
* set external program to convert excel to txt
  `/path/xlsx2txt  %s %t`



![setting](docs/bc_setting.png)

### variables for external conversion program

* %s - source file and path
* %t - target file and path
* %n - source filename
* %x - extension of the source file

    
