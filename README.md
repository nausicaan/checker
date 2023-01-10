# Easily check for WordPress plugin updates

![output](images/output.png)

## About
WordPress Plugin Update Checker (wpp-check), runs the standard `wp plugin list --update=available` command as well as some custom searches to grab those hard to find updates.

## How to Run
`wpp-check {server}.dmz {path} {site}`

Example:

`wpp-check coeurl.dmz /data/www-app/test_blog_gov_bc_ca/current/web/wp test.blog.gov.bc.ca`

## License
Code is under Version 3.0 of the [GNU General Public License](https://github.com/nausicaan/checker/blob/main/LICENSE.md).