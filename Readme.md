# first install go mod file for requirement package

    -- go mod init

# for env variables

    -- go install github.com/joho/dotenv

# for http requests

    -- go install github.com/gin-gonic/gin@latest

# for orm querys

    -- go get -u gorm.io/gorm

# for auto compliy logs

    -- go get github.com/githubnemo/CompileDaemon@latest

# add path for go

    -- export PATH=$PATH:$(go env GOPATH)/bin

# refersh source

    -- source ~/.zshrc

# then last run project

    -- CompileDaemon -command="./project name"
