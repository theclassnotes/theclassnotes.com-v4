---
layout: default
title: Members
permalink: /members/
banner_photo: 2016-notes-at-washington-square-park.jpg
---

# Ladies

<div class="current-notes">

{% for note in site.data.current_notes.ladies %}
<div class="current-note">
  <img src="/images/notes/{{ note.picture }}" alt="{{ note.name }}" width="150" height="150">
  <span class="name">{{ note.name }}</span>
  <span class="part">{{ note.part }}</span>
  <span class="position">{{ note.position }}</span>
</div>
{% endfor %}

</div>

# Gentlemen

<div class="current-notes">

{% for note in site.data.current_notes.gentlemen %}
<div class="current-note">
  <img src="/images/notes/{{ note.picture }}" alt="{{ note.name }}" width="150" height="150">
  <span class="name">{{ note.name }}</span>
  <span class="part">{{ note.part }}</span>
  <span class="position">{{ note.position }}</span>
</div>
{% endfor %}

</div>