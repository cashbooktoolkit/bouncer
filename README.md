## Bouncer

An OAuth2 based authentication server/relay.  I say relay because it delegates the actual autentication calls to a separate webservice. It's written in #golang and uses the Osin libary for its OAuth2 implementation.  

You will need to write a simple web service that does the lookup and password verification.  
I've written an example you can use as a reference: [bouncer example backend](https://github.com/sourdoughlabs/bouncer-example-backend)

The delegated webservice can be implented in whatever programming language best suits the implementation of the athentication lookup and verify operations.

An [omniauth compatible stratgey](https://github.com/sourdoughlabs/omniauth-bouncer) is (will be) available as a Ruby Gem to simplify the development of Rails based applications.

## Minimal dependencies

It's intended that Bouncer stands alone, with no runtime dependencies.  Specifically:

* It handles SSL and Static assets (so no need for a front end server like Nginx or Apache)
* There is no database (rather there is an internal InMemory datastore - Even for tens of thoundands  of users, the memory use is quite modest)

These help reduce operational complexity and minimize the attack surface.

## Getting started

The Login page is fully customisable. CSS and javascript files may be placed into the assets folder.  The login.html file is a golang html/template and things like the error reporting should be left alone, but feel free to style it.
