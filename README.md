# matthew

Matthew is client app for [Sandman](https://github.com/tarunKoyalwar/Sandman.git) 


This App is similar to linux `tee` and can be easily integrated to existing bash scripts .


# Installation Instructions

 - Install Go
 - Make Sure $GOPATH/bin is added to PATH env Variable
 
 ```sh
 go install github.com/tarunKoyalwar/matthew@latest
 ```
 
 
 Before Using this Make Sure that Sandman App is Running . And WebServer is Started
 
 You Can Start WebServer In Settings Page of Sandman App. By Default WebServer is always off


 # Usage

  - Gets all Inscope Data

```sh
matthew pages -p in-scope
```
 	 
  - Adds all unique subs to all-subs page

```sh
bash subdomains.sh | matthew pages -p all-subs --post
```
  - Adds all unique urls to all-urls page

```sh
echo abc.com | gau | matthew pages -p all-urls --post
```
  - Prints Your CheckList

```sh
matthew web -L 
```

  -  Fetches sublist3r results of abc.com subdomain

```sh
matthew org -s abc.com -t sublist3r
```

  - Posts assetfinder results to checklist entry
	that contains keyword assetfinder

```sh
asserfinder abc.com | matthew org -t assetfinder --post
```

  - Prints Your Web CheckList

```sh
matthew web -L
```

  - Fetches nmap results of abc.com subdomain

```sh
matthew web -s abc.com -t nmap
```
  - Posts assetfinder results 

```sh
asserfinder abc.com | matthew web -s abc.com -t assetfinder --post
```
  - Fetches Subdomain that is currently Used 

```sh
matthew web --working
```

## Other Options

```console
./matthew -h    
An Awesome Client for Sandman that can easily be integrated to recon

Usage:
  matthew [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  org         Organization Checklist Page
  pages       Interact with other pages
  web         Web Checklists Page

Flags:
  -a, --append         Append to Existing Data
  -c, --chunk int      Chunk Size While Writing to Network (default 10)
  -h, --help           help for matthew
      --post           Post Data to CheckList Item
  -x, --proxy string   Proxy URL
  -u, --url string     Rest API URL (default "http://127.0.0.1:8088")

Use "matthew [command] --help" for more information about a command.
```

