---
layout: default
title: Home
---

<div class="home">
    {% for post in site.posts limit:5 %}
      {% include post_listing.html %}
    {% endfor %}
</div>
<div class="to-archive">
  That's all the posts on this page. Check out the [Archive](/archive/) for the rest!
</div>
