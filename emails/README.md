# Emails

This module handles templating and sending emails. The email assets are built using [mjml](https://mjml.io) and then embedded using `go-bindata`. Postmark is used to send emails, but you can swap in your own email provider.

## Usage

```
make emails   # compiles all build/*.mjml files into build/*.html
make bindata  # embeds build/*.html into a bindata.go file
make clean    # deletes build folder and bindata file
make          # equivalent to `make clean emails bindata`
```

