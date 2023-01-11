# Check for WordPress Plugin Updates

![banner](images/banner.jpg)

## About
WordPress Plugin Update Checker (wpp-check), runs the standard `wp plugin list --update=available` command as well as some custom searches to grab those hard to find updates.

```console
Below is the current list of plugins requiring updates for test.blog.gov.bc.ca.

wpackagist-plugin/gutenberg:14.8.2
wpackagist-plugin/stackable-ultimate-gutenberg-blocks:3.6.3
wpackagist-plugin/styles-and-layouts-for-gravity-forms:4.3.10
wpackagist-plugin/tablepress:2.0.1
bcgov-plugin/events-calendar-pro:6.0.5.1
bcgov-plugin/event-tickets-plus:5.6.4
bcgov-plugin/gravityforms:2.6.8.2
```

## Prerequisite

Googles' [Go language](https://go.dev) installed locally to enable building executables from source code.

## Build

From the folder containing *main.go*, use the command that matches your environment:

Mac:

```bash
go build -o <build_location>/wpp-check main.go
```

Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o <build_location>/wpp-check main.go
```

## Run

```bash
./wpp-check <server>.dmz <WordPress_path> <site>
```

Example:

```bash
./wpp-check coeurl.dmz /data/www-app/test_blog_gov_bc_ca/current/web/wp test.blog.gov.bc.ca
```

## License
Code is under Version 3.0 of the [GNU General Public License](https://github.com/nausicaan/checker/blob/main/LICENSE.md).