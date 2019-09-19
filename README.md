# Quoteboard helper

### Why?
often times as i'm pulling out letters for a new quote on the letterboard i struggle to know what letters i'll need. so i built this to help me save time.

#### usage:
```bash
go run quote_board.go  -h
```

#### interactive mode:
```bash
go run quote_board.go
```

ex:
```bash
rho:quote_board rho$ go run quote_board.go
Enter quote here: the most important thing is to keep the most important thing the most important thing
a : 3
e : 5
g : 3
h : 6
i : 7
k : 1
m : 6
n : 6
o : 7
p : 4
r : 3
s : 4
t : 16
```
this is particularly helpful because it tells me that i need 16 T's which I don't have.
#### file as input mode:
```bash
go run quote_board.go  -f sample_files/important.txt
```