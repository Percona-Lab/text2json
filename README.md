# text2json
POC for a text (pt-summary / pt-mysql-summary) reports to json converter

The idea is to be able to convert the text output of any program.  
Currently this program implements these filters:  
- All lines having the form `# title here ####...` are considered titles and the title is being used as the key of a json field.
- All lines having the form `left string | right string` are considered a hash entry where the `left string` is the hash key and 
`right string` is the hash value
- If a line cannot be parsed as a title nor a hash, it will be added as a row in a strings array.
  
### Example
   
```
# MySQL Executable ###########################################
       Path to executable | /home/karl/mysql/my-5.6.36/bin/mysqld
              Has symbols | Yes

# Status Counters (Wait 10 Seconds) ##########################
Variable                                Per day  Per second     11 secs
Aborted_connects                              2
Bytes_received                            30000                     350
Bytes_sent                               450000           4        2000
Com_select                                  150                       1
Com_show_binlogs                              5
Com_show_databases                            5

```
  
Tis block:  

```
# MySQL Executable ###########################################
       Path to executable | /home/karl/mysql/my-5.6.36/bin/mysqld
              Has symbols | Yes
```
will be parsed as a hash:
```
"MySQL Executable": {
       "Path to executable:": "/home/karl/mysql/my-5.6.36/bin/mysqld",
       "Has symbols": "Yes",
}
```
  
and this other block:  
```
# Status Counters (Wait 10 Seconds) ##########################
Variable                                Per day  Per second     11 secs
Aborted_connects                              2
Bytes_received                            30000                     350
Bytes_sent                               450000           4        2000
Com_select                                  150                       1
Com_show_binlogs                              5
Com_show_databases                            5
```
will be parsed as an array:
```
"Status Counters (Wait 10 Seconds)": [
"Variable                                Per day  Per second     11 secs",
"Aborted_connects                              2",
"Bytes_received                            30000                     350",
"Bytes_sent                               450000           4        2000",
"Com_select                                  150                       1",
"Com_show_binlogs                              5",
"Com_show_databases                            5"
]
```
  
The resulting json is:

```
{
   "MySQL Executable": {
          "Path to executable:": "/home/karl/mysql/my-5.6.36/bin/mysqld",
          "Has symbols": "Yes",
   },
   "Status Counters (Wait 10 Seconds)": [
   "Variable                                Per day  Per second     11 secs",
   "Aborted_connects                              2",
   "Bytes_received                            30000                     350",
   "Bytes_sent                               450000           4        2000",
   "Com_select                                  150                       1",
   "Com_show_binlogs                              5",
   "Com_show_databases                            5"
   ]
}
```
  
## Usage

`text2json [filename]`  
  
If no file name was specified, it reads from STADIN so you can `|` it to any command.  
  
### Examples
```
pt-summary | text2json
```  
```
pt-summary | text2json > pt-summary.json
```
```
pt-summary > pt-summary.txt
text2json pt-summary.txt
```

