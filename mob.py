import click
from cli import setup as _setup
from cli import generate as _generate
from cli import pullrequest as _pullrequest

@click.group()
def cli():
    pass


@cli.command()
@click.option('--dev', is_flag=True, help="Install dev dependencies")
def setup(dev):
    _setup.run(dev)


@cli.command()
def generate():
    _generate.run()


@cli.command()
def pullrequest():
    _pullrequest.run()
