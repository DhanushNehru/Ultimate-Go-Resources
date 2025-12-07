# Ultimate Go Resources

<img width="1600" height="840" alt="Ultimate Golang Resources" src="https://github.com/user-attachments/assets/2364b0ce-f189-4311-8441-d5f0aa648eb1" />

<a href="https://www.producthunt.com/products/ultimate-golang-resources?embed=true&utm_source=badge-featured&utm_medium=badge&utm_source=badge-ultimate&#0045;golang&#0045;resources" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=992783&theme=light&t=1752676381336" alt="Ultimate&#0032;Golang&#0032;Resources - A&#0032;curated&#0032;list&#0032;of&#0032;collection&#0032;of&#0032;Go&#0032;projects&#0032;and&#0032;resources&#0032; | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

<div align="center">
  
[![Join Our Discord](https://img.shields.io/badge/Discord-Join%20Server-blue?logo=discord&style=for-the-badge)](https://discord.com/invite/Yn9g6KuWyA)
[![Subscribe on YouTube](https://img.shields.io/badge/YouTube-Subscribe-red?logo=youtube&style=for-the-badge)](https://www.youtube.com/@dhanushnehru?sub_confirmation=1)
[![Subscribe to Newsletter](https://img.shields.io/badge/Newsletter-Subscribe-orange?style=for-the-badge)](https://dhanushn.substack.com/)

</div>

# 📖 Documentation
- 🔗 Official GoLang Docs: [Go Dev](https://go.dev/)
- [Go Tour](https://go.dev/tour/welcome/1)
- [Go Packages](https://pkg.go.dev/)
- [Go Blog](https://go.dev/blog/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go FAQ](https://go.dev/doc/faq)

## 📚 Learning Resources
- [Go by Example](https://gobyexample.com/) - Learn Go with annotated example programs
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) - Learn Go by writing tests
- [Gophercises](https://gophercises.com/) - Free coding exercises for Go beginners
- [Go Web Examples](https://gowebexamples.com/) - Learn Go for web development with examples
- [Ultimate Go Programming](https://www.ardanlabs.com/ultimate-go/) - Professional Go training by Ardan Labs

## 📖 Books
- [The Go Programming Language](https://www.gopl.io/) - By Alan Donovan & Brian Kernighan
- [Go in Action](https://www.manning.com/books/go-in-action) - By William Kennedy, Brian Ketelsen & Erik St. Martin
- [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) - By Katherine Cox-Buday
- [Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/) - By Jon Bodner

## 🎥 Video Resources
- [JustForFunc](https://www.youtube.com/c/JustForFunc) - Programming in Go by Francesc Campoy
- [Gopher Academy](https://www.youtube.com/c/GopherAcademy) - Go conferences and talks
- [Todd McLeod](https://www.youtube.com/user/toddmcleod) - Go programming tutorials

## 📁 Available Projects
- [cards](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/cards): A simple Go program that defines a new custom type deck (a slice of strings) representing a deck of playing cards. It demonstrates creating, populating, and printing the deck using nested loops and receiver functions.
- [string_to_html](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/string_to_html): Converts a plain string into an HTML document using Go and writes the output to a file.
- [template_string_to_html](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_string_to_html): Defines an HTML template as a raw string in the Go file and renders it with inserted data.
- [template_pass_string_value_to_html](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_pass_string_value_to_html): Passes a string value into an external HTML template and renders it dynamically using Go templates.
- [template_struct_data_to_html](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_struct_data_to_html): Demonstrates how to pass structured data (using Go structs and slices) into an external HTML template and render dynamic content using Go’s text/template package.
- [template_with_custom_functions_to_html](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_with_custom_functions_to_html): Demonstrates how to register and use custom functions in Go templates to manipulate and display structured data dynamically.
- [template_with_pipelining](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_with_pipelining): Demonstrates using function pipelines in Go templates to transform and render data step-by-step with custom functions.
- [template_safe_html_rendering](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_safe_html_rendering): Demonstrates safe rendering of dynamic user input in Go templates using the `html/template` package, which automatically escapes potentially dangerous content to prevent XSS attacks.


- [tcp_hello_server](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/tcp_hello_server): A simple TCP server in Go that listens on port 8080 and sends a greeting message to any client that connects. Demonstrates basic use of the `net` package for creating TCP servers and handling connections.
- [tcp_echo_server_with_concurrency](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/tcp_echo_server_with_concurrency): A concurrent TCP server in Go that listens on port 8080 and reads text input from clients line-by-line. Demonstrates handling multiple client connections using goroutines and reading data from network streams with `bufio.Scanner`.
- [tcp_echo_response_server](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/tcp_echo_response_server): A concurrent TCP server in Go that listens on port 8080, reads input from clients line-by-line, and sends back a custom echoed response. Demonstrates bidirectional communication using `bufio.Scanner` and `fmt.Fprintf`, along with concurrent connection handling using goroutines.

- [tcp_rot13_server](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/tcp_rot13_server): A concurrent TCP server in Go that listens on port 8080, reads input from clients, converts it to lowercase, and sends back a ROT13-transformed response. Demonstrates text transformation, bidirectional communication using `bufio.Scanner` and `fmt.Fprintf`, and concurrent client handling using goroutines.

- [photo_blog](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/photo_blog): A simple web application in Go that allows users to upload image files via an HTML form, which are saved with a SHA1-hashed filename. Uploaded images are stored in the `public/pics` directory and displayed back to the user based on session-tracked cookies.

- [hmac_hash](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/hmac_hash): A simple Go program that demonstrates generating an HMAC using SHA-256. It takes input strings (e.g., email addresses) and returns a secure, consistent hash using a secret key. Illustrates the use of `crypto/hmac`, `crypto/sha256`, and `io.WriteString` for keyed hashing.
- [base64_encoder](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/base64_encoder): A simple Go program that demonstrates Base64 encoding and decoding of a plain text message using Go's `encoding/base64` package. Shows how to convert strings to Base64 format and decode them back, with basic error handling using `log.Fatalln`.

- [json_marshal_vs_encode](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/json_marshal_vs_encode): A simple Go web server that demonstrates the difference between `json.Marshal` and `json.Encoder`. It defines a struct and shows how to convert it to JSON using both methods. Includes endpoints to serve basic HTML, marshal data to JSON, and encode data directly into the HTTP response, with proper headers and error handling.
- [json_unmarshal](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/json_unmarshal): A simple Go program that demonstrates how to parse a JSON string into nested Go structs using `json.Unmarshal`. It shows how to access nested fields, iterate over slices, and explains the importance of passing a pointer (`&data`) so the unmarshal operation modifies the original variable. Ideal for beginners learning how to handle JSON in Go.
- [crud_mongodb](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/crud_mongodb): A basic CRUD API built with Go and MongoDB that demonstrates how to connect to a MongoDB instance using the official Go driver. The project follows clean structuring with models, controllers, and routing logic, making it a great starting point for developers looking to build RESTful services with MongoDB as the database layer.
- [crud_postgres](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/crud_postgres): A simple CRUD API built with Go and PostgreSQL that demonstrates how to connect to a PostgreSQL database using the `database/sql` package and `lib/pq` driver. The project uses a clean folder structure with separate `models` and `controllers`, and includes standard RESTful routes for creating, reading, updating, and deleting user records. Ideal for developers who want to learn how to build robust backend services with PostgreSQL as the relational data layer.

## 🛠️ Tools & Libraries
- [Gin](https://github.com/gin-gonic/gin) - Fast HTTP web framework written in Go
- [Echo](https://github.com/labstack/echo) - High performance, extensible web framework
- [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
- [Logrus](https://github.com/sirupsen/logrus) - Structured logger for Go
- [Testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks
- [Gorilla Mux](https://github.com/gorilla/mux) - A powerful HTTP router and URL matcher
- [Resty](https://github.com/go-resty/resty) - Simple HTTP and REST client library for Go

## 🌟 Open Source Projects
- [Kubernetes](https://github.com/kubernetes/kubernetes) - Production-Grade Container Scheduling and Management
- [Docker](https://github.com/moby/moby) - Moby Project - a collaborative project for the container ecosystem
- [Prometheus](https://github.com/prometheus/prometheus) - The Prometheus monitoring system and time series database
- [Hugo](https://github.com/gohugoio/hugo) - The world's fastest framework for building websites
- [Grafana](https://github.com/grafana/grafana) - The open observability platform
- [Terraform](https://github.com/hashicorp/terraform) - Terraform enables you to safely and predictably create and manage infrastructure
- [Consul](https://github.com/hashicorp/consul) - Consul is a distributed service mesh to connect, secure, and configure services
- [Vault](https://github.com/hashicorp/vault) - A tool for secrets management, encryption as a service, and privileged access management
- [Etcd](https://github.com/etcd-io/etcd) - Distributed reliable key-value store for the most critical data of a distributed system
- [CockroachDB](https://github.com/cockroachdb/cockroach) - CockroachDB - the open source, cloud-native distributed SQL database

## 📝 Blogs & Articles
- [The Go Blog](https://go.dev/blog/) - Official Go team blog with insights and updates
- [Dave Cheney's Blog](https://dave.cheney.net/) - Go best practices and performance insights
- [Ardan Labs Blog](https://www.ardanlabs.com/blog/) - Go training and development articles
- [Golang Weekly](https://golangweekly.com/) - Weekly newsletter about the Go programming language

## 🎪 Community
- [Go Forum](https://forum.golangbridge.org/) - The Go community forum
- [Gophers Slack](https://gophers.slack.com/) - Join via invite.slack.golangbridge.org
- [r/golang](https://www.reddit.com/r/golang/) - Reddit Go community
- [GopherCon](https://www.gophercon.com/) - Annual Go conference


### Connect With Me

Let's stay connected! Follow me on social media to stay updated:

[![X](https://img.shields.io/badge/Twitter-Follow-blue?style=flat-square&logo=X)](https://x.com/Dhanush_Nehru) 
[![Instagram](https://img.shields.io/badge/Instagram-Follow-blue?style=flat-square&logo=instagram)](https://www.instagram.com/dhanush_nehru/) 
[![YouTube](https://img.shields.io/badge/YouTube-Subscribe-red?style=flat-square&logo=youtube)](https://www.youtube.com/@dhanushnehru?sub_confirmation=1) 
[![GitHub](https://img.shields.io/badge/GitHub-Follow-blue?style=flat-square&logo=github)](https://github.com/DhanushNehru)


----

Feel free to update the README.md or raise issues if any to enhance the project
