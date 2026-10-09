# Redirector

A simple shortcode redirector for your website. It allows you to create short URLs that redirect to longer URLs, making it easier to share links.

## Usage

1. Build the project: `go build -o redirector main.go`
2. Create a plain text file with the following format:
   ```
   short-url-1 https://www.example.com/long-url-1
   short-url-2 https://www.example.com/long-url-2
   ```
3. Run the redirector with the path to your text file: `./redirector path/to/your/file.txt`
4. Access the short URLs in your browser: `http://localhost:8080/short-url-1` will redirect to `https://www.example.com/long-url-1`.
5. The HTTP listening address and port can be changed with the `-addr` and `-port` flags. For example: `./redirector -addr 0.0.0.0 -port 8080`
