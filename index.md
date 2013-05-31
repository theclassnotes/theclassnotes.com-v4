---
layout: default
title: Home
---

{% for post in site.posts limit:5 %}
  {% include post_listing.html %}
{% endfor %}

That's all the posts on this page. Check out the <a href="/archive/">Archive</a for the rest!
