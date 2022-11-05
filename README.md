![Golang](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1024px-Go_Logo_Blue.svg.png?20191207190041)

## Basic

## Package
#### <ins>Essential Command</ins>

##### initialize package
`go mod init <base-package-name>`

##### download all dependency
`go mod download`
##### ensure that all imports are satisfied
`go mod tidy`

#### <ins>additional notes</ins>

1. Blank identifier ( _ ) can be used to force unused import
2. Access Modifier depends on first letter of name at each functions or variables

- Capital letter means can be accessed from outside package (like public)
- Small letter means can only be accessed from same package (like private)

