# Quick Serve

This is a very small little app to statically serve directories over 
HTTPS without any setup.  It will generate self-signed https cert and 
key on run, by default it will serve the current directory.

## Why?

Because serving directories over HTTP(s) is often the least painful way 
to quickly share something on a local network.  I use it mostly as a 
quick and dirty way to stream stuff to my tablets from my main desktop.

I got sick of writing little "serve.go" files that just had

    ...
    func main() {
            ...
            panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
    }

inside them that I would run with "go run serve.go"

## Binaries

If you don't want to install Go in order to build this yourself, feel free to grab
your platform specific binary from the binaries folder, or just "save as" one of the below links

- Windows: https://github.com/robertmeta/quickserve/blob/master/binaries/windows/quickserve.exe?raw=true
- Linux: https://github.com/robertmeta/quickserve/blob/master/binaries/linux/quickserve?raw=true 
- OS-X: https://github.com/robertmeta/quickserve/blob/master/binaries/osx/quickserve?raw=true

## Examples

- quickserve :: will serve the current directory
- quickserve -d /some/directory/to/serve -d other/local/directory :: will serve those two directories

## Usage of quickserve

- -a="localhost": The address to serve https on
- -c="cert.pem": The name of the cert to use or generate
- -d=[]: List of directories to serve (use multiple -d flags)
- -k="key.pem": The name of the key to use or generate
- -n=false: Force generation of new certs
- -po=443: The port to serve https on

## TODO

- Remove binaries, fresh directory, use github releases
- Tests
- Add BasicAuth
