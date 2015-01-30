<!DOCTYPE html>
<html>
    <head>
        {{template "head"}}
    </head>
    <body>
        <h1>register</h1>
        <form action="/register" method="post">
            用户名:<input type="text" name="username"><br>
            密码:<input type="password" name="pwd"><br>
            <input type="submit" value="注册">
        </form>
    </body>
</html>
