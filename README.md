# sssComment

sssComment (super super simple comment) is a self-host comment system, the alternative of disqus.

sssComment is based on golang (backend) + sqlite/litestream (db) + htmx/alpinejs/daisyui (frontend). You can Run it in your host and embed it into your blog, docs, or other web pages.

## Usage
### run by bin program
TBD

### run in docker
use this [docker compose file](./deploy/docker-compose.yml) or embed to your docker compose file
the example is using traefik as http proxy, you can use your preferred proxy like nginx, caddy, etc as your http proxy

## Embed into page
Embed below "div" and "script" into your page, then it will insert a comment div into your page
```html
<div id="sssc-iframe-container"
     data-host="https://your-site-domain"
     data-slug="page-id"
     data-url="page-url"
     data-title="page-title">
</div>
<script async src="https://your-site-domain/dist/js/sssc.js"></script>
```

## Enable litestream replicate
Since sssComment use sqlite, we can use litestream to replicate the database to S3 storage. The program already integrated litestream.
To enable litestream, just setup env variable
```
LITESTREAM_URL="s3://s3serivice/bucket/filename"
LITESTREAM_ACCESS_KEY_ID=admin
LITESTREAM_SECRET_ACCESS_KEY=admin
```

## Enable telegram notifications and anonymous approve
Also can set a telegram bot to notify you when a new comment is posted. And you can use noti link to approve or reject the comment by anonymous link. To enable bot notification, need to set "tg_bot_url" to `https://api.telegram.org/bot<token>/sendMessage?chat_id=<chat-id>`