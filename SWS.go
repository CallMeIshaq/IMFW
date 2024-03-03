package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", ironManHandler)
	http.ListenAndServe(":8080", nil)
}

func ironManHandler(w http.ResponseWriter, r *http.Request) {
	// HTML content to display Iron Man fan details
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	    <title>Iron Man </title>
	    <style>
	        body {
	            background-color: #1e1e1e;
	            color: #ffffff;
	            font-family: Arial, sans-serif;
	            margin: 0;
	            padding: 0;
	        }

	        .container {
	            max-width: 1200px;
	            margin: 0 auto;
	            padding: 20px;
	        }

	        .header {
	            text-align: center;
	            padding: 20px 0;
	        }

	        .content {
	            padding: 20px 0;
	        }

	        .footer {
	            text-align: center;
	            padding: 20px 0;
	            color: #888888;
	        }
	        
	        img {
	            display: block;
	            margin: 0 auto;
	            max-width: 100%;
	        }
	    </style>
	</head>
	<body>
	    <div class="container">
	        <header class="header">
	            <h1>Welcome to the Iron Man Fan Website</h1>
	        </header>
	        <section class="content">
	            <img src="https://purepng.com/public/uploads/large/purepng.com-ironmanironmansuperheromarvel-comicscharactermarvel-studiosrobert-downey-jrtony-stark-1701528612042dkvmr.png" alt="Iron Man">
	            <p>Iron Man, also known as Tony Stark, is a fictional superhero appearing in American comic books published by Marvel Comics. He is a billionaire playboy, philanthropist, and ingenious engineer who creates a powered suit of armor to save the world. Throughout his storied history, Iron Man has faced an array of adversaries, ranging from corporate rivals like Obadiah Stane to cosmic entities like Thanos. His adventures have taken him to the far reaches of space and deep into the heart of conflict on Earth, showcasing his adaptability and resilience in the face of adversity. Whether teaming up with fellow Avengers or going solo, Iron Man's ingenuity and courage have saved the day countless times, earning him the admiration of fans and foes alike. Beyond his exploits in comic books, Iron Man has leaped off the page into various media, including blockbuster films, animated series, and video games. Robert Downey Jr.'s portrayal of Tony Stark in the Marvel Cinematic Universe brought the character to new heights of popularity, captivating audiences with his charm and charisma. Iron Man's influence extends beyond entertainment, inspiring real-world engineers, technologists, and futurists to push the boundaries of innovation and make the world a better place. As Iron Man continues to evolve and inspire new generations of fans, his legacy remains a testament to the enduring power of heroism, invention, and the indomitable human spirit. Whether soaring through the skies in his iconic suit of armor or grappling with the complexities of his own humanity, Iron Man's journey is a reminder that even the greatest heroes are not without their flaws, but it is their courage and determination that define them in the end. Iron Man's adventures are far from over, and as long as there are threats to humanity and injustice in the world, Tony Stark will be there, armored and ready to face whatever challenges come his way.</p>
	            <!-- Repeat the paragraph above to reach a total of 250 lines -->
	        </section>
	        <footer class="footer">
	            <p>Website by Ishaq</p>
	        </footer>
	    </div>
	</body>
	</html>
	`

	// Write the HTML content to the response writer
	fmt.Fprintf(w, html)
}
