from django.contrib.auth.models import AbstractUser


class Entry(AbstractUser):
    value = 5
    class Meta:
        app_label = '{{ project_name }}'


@app.task
def add(x, l):
    sleep(10)
    entry = Entry.objects.First()
    entry.value += x
    entry.save()
    if len(l) > 0:
        next = l[0]
        rest = l[1:]
        next(rest)

@app.task
def mult(x, l):
    sleep(10)
    entry = Entry.objects.First()
    entry.value *= x
    entry.save()
    if len(l) > 0:
        next = l[0]
        rest = l[1:]
        next(rest)

def op()
    add(2, mult(3, [])).tasks()
