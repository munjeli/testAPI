# testAPI
scaffolding, templates and testing for APIs in Go
### Auth
I've spent my entire career as a DevOps doing immutable infrastructure and as a result, I've never authed anything except with a secret service. First there was a custom service at one of my jobs, then some platform level stuff, and then Hashicorp Vault. Running Vault on the instance occurred to me but seemed overblown, so I just used a password file with a module for middleware.  
### Application 
I messed about with rune slices to do this before going to `bufio.ScanWords`. In either case, there were a few outliers on the messy file, which was a text version of Robert's Rules of Order. Because the text is a manual, it has a lot of formatting, outlines, symbols, etc. I think `bufio` has advantages as well if I needed to scale here. 

I didn't strive for perfection here; in my experience, there's always a last effort in writing a parser where you're just handling something specifically ugly about *your* data. 

### API
There's a whole lotta stuff that can happen with headers here, expecially around cross domain, but there's no specific use case so I refrained. 

One strange thing was the `Content-Length` property. When I first built a request for the integration test using `Open(datafile)` it returned a reader that worked in the application but Go couldn't determine the content length, which is auto-generated with `http.NewRequest`. Bug or feature? I reject those kinds of posts bc I couldn't validate. 

### Make
I'm automating my own iteration with bash and make. It's the lowest common denominator in systems I think, most everyone has it and knows how to use it. The easiest iteration is `make` && `make clean` 

### Config
This directory has setup scripts and systemd resources. The secret file isn't here, I just couldn't face it, and put it up separately with `scp`. The setup script handles making the folders for the cert and secret, also the permissioning. 

### Package
The `make pack` command will build the executable and zip it up with the configs so I can `scp` it to the server. Maybe a real package later.
 
