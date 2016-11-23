<html>
  <head>
    <title>Attendance</title>
  </head>
  <body>
    <form action="/attendance" method="post">
      UserName:<input type="text" name="userName">
      CurrentTime:<input type="text"
						 placeholder="YYYY-MM-DD HH:MM:SS"
						 name="currentTime"
						 style="width: 150px;">
      <input type="submit" value="出勤">
    </form>
  </body>
</html>
