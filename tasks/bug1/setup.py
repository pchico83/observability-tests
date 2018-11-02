from setuptools import setup
from setuptools.command.test import test as TestCommand
import sys


setup_requires = []

dev_requires = [
    'flake8',
    'ipdb',
]

tests_require = [
    'pytest-cov',
    'pytest-django',
    'factory_boy',
]

install_requires = [
    'Django==1.6.2',
    'Pillow==2.3.0',
    'psycopg2==2.5.2',
    'South==0.8.4',
    'celery[redis]==3.1.8',
    'raven==4.0.4',
    'gunicorn==18.0',
    'python3-memcached==1.51',
]
