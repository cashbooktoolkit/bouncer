## Bouncer

![Image](bouncer.png)

An OAuth2 based authentication server/relay.  I say relay because it delegates the actual autentication calls to a separate webservice. It's written in #golang and uses the Osin libary for its OAuth2 implementation.  

You will need to write a simple web service that does the lookup and password verification.  
I've written an example you can use as a reference: ( https://github.com/sourdoughlabs/bouncer-example-backend )

The delegated webservice can be implented in whatever programming language best suits the implementation of the athentication lookup and verify operations.

https://github.com/sourdoughlabs/omniauth-bouncer is an omniauth compatible gem to simplify the development of Rails based applications.

## Minimal dependencies

It's intended that Bouncer stands alone, with no runtime dependencies.  Specifically:

* It handles SSL (Coming soon) and Static assets (so no need for a front end server like Nginx or Apache)
* There is no database (rather there is an internal InMemory datastore - Even for tens of thoundands  of users, the memory use is quite modest)

These help reduce operational complexity and minimize the attack surface.

## Getting started

The Login page is fully customisable. CSS and javascript files may be placed into the assets folder.  The login.html file is a golang html/template and things like the error reporting should be left alone, but feel free to style it.

## Adding Clients/Apps

Similar to the way you would register apps for Facebook or Twitter, client apps need to be known to Bouncer.  They are added to clients.json.

## Contributing

1. Fork it ( https://github.com/sourdoughlabs/drivethru/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## LICENSE: MIT


