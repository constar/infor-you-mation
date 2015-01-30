<!DOCTYPE html>

<html>
    <head>
        {{template "head"}}
    </head>
  	<body>
        <!--{{template "header"}} -->
        <h1>login</h1>
        <form action="/" method="post">
            用户名:<input type="text" name="username"><br>
            密码:<input type="password" name="pwd"><br>
            <input type="submit" value="登录"><br>
            <a href="/register">注册</a>
        </form>
    </body>
	</body>
</html>
