<!DOCTYPE html>
<html lang="en">
    <head>
        <title>Lakexyde's GoSocial</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link href="/static/css/home.css" rel="stylesheet">
    </head>
    <body>
        <div class="app">
            <header class="lks-home-header">
                <div class="container">
                    <h1 class="h-logo">Go Social</h1>
                    <form class="user-login" method="POST" action="/login">
                        <div clas="form-div">
                            <label for="Username">Username</label>
                            <input type="text" name="Username" required>
                        </div>
                        <div clas="form-div">
                            <label for="Password">Password</label>
                            <input type="password" name="Password" required>
                            <a href="#">Forgotten account?</a>
                        </div>
                        <div clas="form-div">
                            <button type="submit" class="login-btn">Log In</button>
                        </div>
                        
                    </form>
                </div>                   
            </header>
            <div class="lks-main flex">
                <div class="lks-left-content">
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit. Quia possimus ratione, officiis accusamus, delectus iste nulla odio. Deleniti corporis nesciunt consequuntur, facere aliquam facilis maxime odit animi, blanditiis quaerat commodi.
                </div>
                <div class="lks-right-content">
                    <h3 class="form-head">Create an account</h3>
                    <form action="/register" method="POST" class="user-reg">
                        <input type="text" id="f-name" placeholder="First name" name="FirstName" required>
                        <input type="text" id="l-name" placeholder="Surname" name="LastName" required>
                        <input type="email" id="email" placeholder="Email" name="Email" required>
                        <input type="password" class="password" placeholder="Password" required>
                        <input type="password" class="password" placeholder="Confirm Password" name="Password">
                        <div class="f-radio">
                            <input type="radio" name="Gender" value="Female" required> Female   <input type="radio" name="Gender" value="Male" required> Male  
                        </div>
                        <button type="submit" class="reg-btn">Create an account</button>
                    </form>
                </div>
            </div>
            <footer class="lks-footer">
                <div class="container">
                    <div class="footer-top-section">
                        <a href="#">Custom</a>  <a href="#">Another</a>  <a href="#">Lorem</a>
                    </div>
                    <div class="footer-bottom-section"></div>
                </div>
            </footer>
        </div>
    </body>
</html>