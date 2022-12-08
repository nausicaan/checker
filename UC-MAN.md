UC(1)
#### NAME
>uc - find and email notifications of WordPress updates

#### SYNOPSIS
>**uc** [*OPTION*])

>**uc [-p | -t | -c]** server path url

#### DESCRIPTION
>**uc** searches for updates based on supplied arguments. These may be plugins, themes, or core updates for the WordPress platform.

>The options are as follows:

>-p
>>Search for WordPress Plugin updates

>-t 
>>Search for WordPress Theme updates

>-c 
>>Search for WordPress Core updates

#### OPTIONS
>**-h**, **--help**
>: Displays a friendly help message

#### EXAMPLE
>Search for Wordpress plugin updates on GDX Test server:
>>**$** -p coeurl.dmz /data/www-app/test_blog_gov_bc_ca/current/web/wp test.blog.gov.bc.ca

#### EXIT STATUS
>**O**
>: Success

>**1**
>: Invalid option

#### BUGS
>Don't be ridiculous.