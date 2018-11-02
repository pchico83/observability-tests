from django.contrib.auth.models import AbstractUser


class Entry(AbstractUser):
    value = 5
    class Meta:
        app_label = '{{ project_name }}'


@app.task
def add(x):
    sleep(10)
    entry = Entry.objects.First()
    entry.value += x
    entry.save()

@app.task
def mult(x):
    sleep(10)
    entry = Entry.objects.First()
    entry.value *= x
    entry.save()

def op()
    add(2).tasks()
    mult(3).tasks()