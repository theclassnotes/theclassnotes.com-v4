---
layout: default
title: Members
permalink: /members/
---

# Ladies

<table class="current_notes">
    <tbody>
        <tr>
            {% for note in site.data.current_notes.ladies limit:4 %}
            <td class="picture">
                <img src="/images/notes/{{ note.picture }}" alt="{{ note.name }}" width="150" height="150">
            </td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.ladies limit:4 %}
            <td class="name">{{ note.name }}</td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.ladies limit:4 %}
            <td class="part">{{ note.part }}</td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.ladies limit:4 %}
            <td class="position">{{ note.position }}</td>
            {% endfor %}
        </tr>
    </tbody>
</table>

<table class="current_notes">
    <tbody>
        <tr>
            {% for note in site.data.current_notes.ladies limit:4 offset:4 %}
            <td class="picture">
                <img src="/images/notes/{{ note.picture }}" alt="{{ note.name }}" width="150" height="150">
            </td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.ladies limit:4 offset:4 %}
            <td class="name">{{ note.name }}</td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.ladies limit:4 offset:4 %}
            <td class="part">{{ note.part }}</td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.ladies limit:4 offset:4 %}
            <td class="position">{{ note.position }}</td>
            {% endfor %}
        </tr>
    </tbody>
</table>

# Gentlemen

<table class="current_notes">
    <tbody>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 %}
            <td class="picture">
                <img src="/images/notes/{{ note.picture }}" alt="{{ note.name }}" width="150" height="150">
            </td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 %}
            <td class="name">{{ note.name }}</td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 %}
            <td class="part">{{ note.part }}</td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 %}
            <td class="position">{{ note.position }}</td>
            {% endfor %}
        </tr>
    </tbody>
</table>

<table class="current_notes">
    <tbody>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 offset:4 %}
            <td class="picture">
                <img src="/images/notes/{{ note.picture }}" alt="{{ note.name }}" width="150" height="150">
            </td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 offset:4 %}
            <td class="name">{{ note.name }}</td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 offset:4 %}
            <td class="part">{{ note.part }}</td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 offset:4 %}
            <td class="position">{{ note.position }}</td>
            {% endfor %}
        </tr>
    </tbody>
</table>

<table class="current_notes">
    <tbody>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 offset:8 %}
            <td class="picture">
                <img src="/images/notes/{{ note.picture }}" alt="{{ note.name }}" width="150" height="150">
            </td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 offset:8 %}
            <td class="name">{{ note.name }}</td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 offset:8 %}
            <td class="part">{{ note.part }}</td>
            {% endfor %}
        </tr>
        <tr>
            {% for note in site.data.current_notes.gentlemen limit:4 offset:8 %}
            <td class="position">{{ note.position }}</td>
            {% endfor %}
        </tr>
    </tbody>
</table>

