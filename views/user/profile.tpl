<!DOCTYPE html>
<html lang="en">
    <head>
        <title>User profile page</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link href="/static/css/styles.css" rel="stylesheet">
    </head>
    <body>
       <div class="app">
           <header class="lks-header">
               <h2>Welcome {{.u}}</h2>
               <a href="/logout">Logout</a>
               <span>Session: {{.g}}</span>
           </header>
           <div class="lks-main">

           </div>
       </div>
    </body>
</html>