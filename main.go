package main

import (
	"io"
	"log"
	"net/http"

	"github.com/jessebarton/get/install"
)

// var tpl *template.Template

// func init() {
// 	tpl = template.Must(template.ParseGlob("templates/*"))
// }

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,
		` <html lang="en">
		<head>
				<meta charset="utf-8">
				<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
				<!-- Bootstrap CSS -->
				<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/css/bootstrap.min.css" integrity="sha384-PsH8R72JQ3SOdhVi3uxftmaW6Vc51MKb0q5P2rRUpPvrszuE4W1povHYgTpBfshb" crossorigin="anonymous">
			<title>IT - Images</title>
				<nav class="navbar navbar-light bg-light">
					<span class="navbar-brand mb-0 h1">Life.Church IT</span>
				</nav>
		</head>
		
		<body>
		<div>
			<div>
				<div class="jumbotron">
					<h1 class="display-3">Welcome!</h1>
					<p class="lead">All Documentation for configuring computers can be found at this link <a class="link" href="">Configuration Documentation</a></p>
					<hr class="my-4">
					<p>Select the Image that you would like to download for the machine that you are configuring. </p>
					<p class="lead">
					<form method="POST">
						<option>Type Lighting, Checkin, Protools, Propresenter, Usermac, or Smaart</option>
						<input type="text" name="i">
						<input type="submit" class="btn btn-dark btn-lg">
					</form>
					<!-- <form method="POST">
						<fieldset>
							<div class="form-group">
								<label>Images</label>
									<select class="form-control" name="Images">
										<option>Checkin</option>
										<option>Protools</option>
										<option>Propresenter</option>
										<option>Usermac</option>
										<option>Smaart</option>
										<option>Lighting</option>
									</select>
							</div>
						</fieldset>
						<button type="submit" class="btn btn-dark btn-lg">Submit</button>
					</form> -->
					</p>
				</div>
			</div>
		</div>
		
			<!-- Optional JavaScript -->
			<!-- jQuery first, then Popper.js, then Bootstrap JS -->
			<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.3/umd/popper.min.js" integrity="sha384-vFJXuSJphROIrBnz7yo7oB41mKfc8JzQZiCq4NCceLEaO4IHwicKwpJf9c9IpFgh" crossorigin="anonymous"></script>
			<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/js/bootstrap.min.js" integrity="sha384-alpBpkh1PFOepccYVYDB4do5UnbKysX5WZXm3XxPqe5iKTfUKjNkCk9SaVuEZflJ" crossorigin="anonymous"></script>
		</body>
		
		</html>`)

	i := req.FormValue("i")
	install.Install(i)

}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
