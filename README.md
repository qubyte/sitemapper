# sitemapper

A demo cli wrapper around [sitemap][1]. To use, run:

```bash
docker run qubyte/sitemapper --start-url <put a URL to start from here> --jobs <workers>
```

`--start-url` is mandatory, `--jobs` defaults to `1`. `--jobs` is the number of 
sites to allow the program to request and process simultanously. For large 
sites, consider using this argument.

[1]: https://github.com/qubyte/sitemap
