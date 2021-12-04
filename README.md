# AdventOfCode2021

golang this year

_Using golang plugin feature running under linux_

### Account
Log into adventofcode.com and create account

### Session cookie
_Retrieve session cookie from your browser using development tools_

- Ctrl+Shift+I 
- -> Application 
- -> Cookies 
- -> session

```bash
export AOC_SESSION_COOKIE=0123456789abcdef

# builds plugins
make
go build -buildmode=plugin -o bin/1.so days/1/one.go
go build -buildmode=plugin -o bin/2.so days/2/two.go

# run all days
make run
go run main.go -day=all
AdventOfCode2021 in golang this year :D
2021/12/03 00:44:21 day-one-1: x
2021/12/03 00:44:21 day-one-2: xx
2021/12/03 00:44:21 day-two-1: xxx
2021/12/03 00:44:21 day-two-2: xxxx

# running one day
DAY=2 make run
go run main.go -day=2
AdventOfCode2021 in golang this year :D
2021/12/03 00:53:31 day-two-1: xxx
2021/12/03 00:53:31 day-two-2: xxxx
```