---
layout: default
title: Home
banner_photo: joseph-okitsalrightwithme-cuttin-class-xxx.jpg
---

## The Latest

{% for post in site.posts limit:5 %}
  {% include post_listing.html %}
{% endfor %}

That's all the posts on this page. Check out the [Archive](/archive/) for the rest!
