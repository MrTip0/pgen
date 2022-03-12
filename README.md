# pgen
A simple password generator  

## Installation
### Requierements
- go >= 1.17
- go bin folder in your $PATH variable

### Installation
```sh
go install github.com/MrTip0/pgen@latest
```

## How to use  

simply  
```sh
pgen
```

The default lenght is 20 characters but you can change it with  
```sh
pgen -l 20
```

If you want to chose an alphabet  
```sh
pegn -alfa "AaBbCcDd123#!="
```
> **Note**  
> The alphabet is case sensitive,  
> so 'A' is different from 'a'
