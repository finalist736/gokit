# Summary
## GoKitMan
GoKit Manager призван создавать пустой проект-болванку
для дальнейшего написания функционала.
# Files and directories structure
```
%GOPATH%/src/yourproject/
    web/
        router/
            router.go
        handlers/
            api/
                items/
                    get.go
        templates/
            user/
                login.html
                register.html
                forgot.html
            blog/
                post.html
                postslist.html
        ctx/
            context.go
            cache.go
    model/
        migrations/
            1.sql
            2.sql
        mysql/
            connection.go
            usermodel/
                data.go
            blogmodel/
                data.go
    service/
        main.go

    config.ini
    db_users.go
    db_posts.go
```
## yourproject
root folder with db structure files
## web
#### router
router create
#### handlers
http handlers
## templates
folder contains html templates
## model
connection to database you use
mysql,pg,sqlite, etc...
#### mysql
models with mysql realizations
## service
main service loop