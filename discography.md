---
layout: page
title: Discography
permalink: /discography/
banner_photo: emma-many-the-miles.jpg
---

## Discography

In its short 30-year existence, The Class Notes has produced six studio albums
and two EPs.

{% for album in site.data.albums %}

<div class="album-listing">
  <div class="album-cover">
    <img src="/images/albums/{{ album.picture }}" alt="'{{ album.name }}' Album Cover">
  </div>

  <h2 class="album-title">{{ album.name }}</h2>
  <p>{{ album.blurb }}</p>

  <table class="album-tracks {{ album.name }}">
  {% tablerow track in album.tracks cols:2 %}
    <strong>{{ tablerowloop.index }}. {{ track.title }} - {{ track.soloist }}</strong>
    <br/>
    Orig: {{ track.opb }}
  {% endtablerow %}
  </table>
</div>

{% endfor %}
