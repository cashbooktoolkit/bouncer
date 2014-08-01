## Bouncer

An OAuth2 based authentication server/relay.  I say relay because it delegates the actual
autentication calls to a separate webservice. It's written in #golang and uses the Osin libary
for its OAuth2 implementation.

The delegated webservice can be implented in whatever programming language best suits the
implementation of the athentication lookup and verify operations.

An omniauth compatible strategy is (will be) available as a Ruby Gem to simplify the development of
Rails based applications.

## Minimal dependencies

It's intended that Bouncer stands alone, with no runtime dependencies.  Specifically:

* It handles SSL and Static assets (so no need for a front end server like Nginx or Apache)
* There is no database (rather there is an internal InMemory datastore - Even for tens of thoundands  of users, the memory use is quite modest)

These help reduce operational complexity and minimize the attack surface.

## Getting started

The Login page is fully customisable.  Begin by copying login-sample.html to login.html and editing it as you see fit.
CSS and javascript files may be placed into the assets folder.  The login.html file is a golang html/template and things 
like the error reporting should be left alone, but feel free to style it.
