# API Template

This is a generic API template that will help setup a complex GORM API. The template is setup to allow for multiple databases to connect on a single API if needed.

## Setup

**1) Clone the repository**
```
git clone https://github.com/wbabcock/api-template.git
```

**2) Rename the folders and files appropriately**
```
models\table_name.go
routes\table_name.go
```

**3) Setup the Go mod file**
```
go mod init ...
go mod vendor
```

**4) Configure the database settings**
```
db\controller.go
```


