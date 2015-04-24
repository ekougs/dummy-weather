# Dummy Weather

A Go language simple demo application simulating a weather station for some cities.
Compatible with Go 1.4 as some features have been introduced in this version.
It allows you to discover step by step new Go features as application grows.
Each step is a tag "Step \# : description of step with features introduced"

To build the application, you should run `make` command.
It will install in the bin dir of your GOPATH, the installation directory with the application at its root.

dummy-weather --help to find out how to use dummy-weather

On server mode, the application is deployed on port 1987
GET /cities returns cities list
GET /cities/{city_code}/temps?date=2015-04-25T13:00:00&duration=5D returns cities temperatures. date and duration params are optional. date default value is now.
