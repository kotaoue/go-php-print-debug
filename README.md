# go-php-print-debug
This script search print debug from PHP code.  
Checking "print", "print_r", "var_dump", "var_export", "echo" as print debug.  
Exclude user-defined functions like echoOriginal().


## Usage
```ShellSession
$ go run tools/search_print_debug/main.go -directories=app,vendor
# root: app
## path: app/add.php
## path: app/mul.php
- [ ] find print debug@app/mul.php:5, var_export($a * $b);
## path: app/sub.php
- [ ] find print debug@app/sub.php:5, print_r($a);
- [ ] find print debug@app/sub.php:6, var_dump($b);
- [ ] find print debug@app/sub.php:7, echo $a - $b;
# root: vendor
## path: vendor/switcher.php
find print debug
exit status 1
```