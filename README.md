## Mission

I want to write a program in go that does render any link in a given text input unclickable.
This is important if you publish malware links to others, in order to prevent people from clicking those links and get infected by accident.

## Roadmap

- Version 0.1 Proof of Concept
- Version 0.2 take input from command line, like: 'dlink http://www.host.tld'
- Version 0.3 take piped input like in: 'cat file | dlink > new_file'
- Version 0.4 take a file as argument and replace it, like in: 'dlink file'
- Version 0.5 take input from the network via http

## Webservice

dlink is also available as a webservice.

- Version 0.1 - paste some text in a text input field and get back the modified version.
- Version 0.2 - upload a file, and get back the modified file as download.
