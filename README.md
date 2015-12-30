# redirection_service

A simple Go program that's used to redirect clients via a 302 HTTP status code.

Each client is served in a separate goroutine.

## Usage

    ➜ $?=0 /Users/arastogi/Playground/repos/redirection_service [ 4:34PM] % sudo ./redirection_service -location=https://blog.abhijeetr.com -port=80                        0s elasped ⌛
    2015/12/30 16:34:57 Starting web server... Location: https://blog.abhijeetr.com, Port: 80
