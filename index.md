---
layout: default
title: Home
---

## The Latest

{% for post in site.posts limit:5 %}
  {% include post_listing.html %}
{% endfor %}

That's all the posts on this page. Check out the [Archive](/archive/) for the rest!
