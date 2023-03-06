#!/usr/bin/env python
"""
run golang-migrate - shortcut
"""
import shutil
from envex import Env
import subprocess

migrate = shutil.which("migrate")
env = Env(readenv=True, parents=True)
database_url = env("DATABASE_URL") + "?sslmode=disable"

def migrate_create(extension: str, directory: str, timestamps: bool, name: str, extra: str):
    args = [
        migrate,
        'create',
        '-ext', extension,
        '-dir', directory,
    ]
    if not timestamps:
        args.append('-seq')
    if extra:
        args.extend(extra)
    else:
        args.append(name)
    print(args)
    subprocess.call(args, 0, migrate)

def migrate_exec(up_down, directory, extra):
    args = [
        migrate,
       '-database', database_url,
        "-path", directory,
        up_down,
    ] + extra
    print(args)
    subprocess.call(args, 0, migrate)

if __name__ == '__main__':
    from argparse import ArgumentParser

    d = dict(
        path = migrate,
        ext = "sql",
        dir = "repository/migrations",
        name = "unnamed_migration",
    )

    parser = ArgumentParser(description=__doc__)
    parser.add_argument('-m', '--migrate', action='store', metavar='path', default=d["path"], help=f'golang-migrate path, default={d["path"]}')
    parser.add_argument('-e', '--ext', action='store', default=d["ext"], help=f'migrate file extension default={d["ext"]}')
    parser.add_argument('-d', '--dir', action='store', default=d["dir"], help=f'where to store migrations default={d["dir"]}')
    parser.add_argument('-t', '--timestamps', action='store_true', default=False, help=f'use timestamp instead of sequence')
    parser.add_argument('-n', '--name', action='store', default=d["name"], help=f'set migration file name')

    parser.add_argument('action', action='store', choices=('up', 'down', "create"), help='migration action')
    parser.add_argument('extra', action='store', nargs='*', default=[], help='additional arguments')


    args = parser.parse_args()

    match args.action:
        case 'create':
            migrate_create(args.ext, args.dir, args.timestamps, args.name, args.extra)
        case 'up':
            migrate_exec("up", args.dir, args.extra)
        case 'down':
            migrate_exec("down", args.dir, args.extra)