---
layout: page
title: Alumni
permalink: /alumni/
---

## Before us came...

To our beloved alumni, help us stay in touch by updating your photo and contact info with
our [GM](mailto:gm@theclassnotes.com). Also, be sure to [join the alumni
listserve](mailto:gm@theclassnotes.com?subject=Alumni%20Listserve) to receive news and
updates from all generations of Class Notes.

{% for year in (1983..2012) reversed %}
  {% include alumni.html year=year %}
{% endfor %}
