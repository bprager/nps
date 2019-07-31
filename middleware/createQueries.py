#!/usr/bin/env python3

"""Query creation for github.com/99designs/gqlgen

Complementary script to generate separate query files/functions.
"""

import getopt
import logging
import os
import os.path
import re
import sys

SOURCE = '../resolver.go'
DEST = '../resolver.bak'
PATTERN = r'^func\ \(r \*(\w+)Resolver\) (\w+)\((.*?)\)\ \((.*?)\)\ {'
GQL_PRIMITIVES = ['error', 'bool', 'string', 'int']

# set up logger
log = logging.getLogger('query')
log.setLevel(logging.INFO)
ch = logging.StreamHandler()
formatter = logging.Formatter(
    '%(asctime)s - %(name)s - %(levelname)s - %(message)s')
ch.setFormatter(formatter)
log.addHandler(ch)


def usageExit(msg=None):
    print(__doc__)
    print("""Usage: %s [options]\n\nOptions:
      -d, --debug
      \tLog DEBUG messages
      -q, --quiet
      \tSupress INFO log messages (show only WARN and above)
      -h, -H, --help
      \tShow this help message and exit'""" % sys.argv[0])
    sys.exit(2)


def main():
    try:
        options, _ = getopt.getopt(
            sys.argv[1:], 'hHdq', ['help', 'debug', 'quiet'])
    except getopt.error as msg:
        usageExit(msg)
    for opt, _ in options:
        if opt in ('-h', '-H', '--help'):
            usageExit()
        if opt in ('-d', '--debug'):
            log.setLevel(logging.DEBUG)
        if opt in ('-q', '--quiet'):
            log.setLevel(logging.WARN)

    # determine package names
    with open('../go.mod') as p:
        parent_module = p.readline().split()[1]
    log.info(f'found parent module {parent_module}')
    with open('go.mod') as p:
        module = p.readline().split()[-1]
        package = module.split('/')[-1]
    log.info(f'query package name is {package}')
    # backup 'resolver.go'
    log.info(f'backing up {SOURCE} to {DEST} ...')
    os.rename(SOURCE, DEST)
    # find all methods and create separate files
    log.info(f'processing resolver ...')
    last_line = ''
    with open(DEST, 'r') as f, open(SOURCE, 'w+') as d:
        p = re.compile(PATTERN)
        package_line = f'\tq "{module}"'
        for line in f:
            if line.startswith(') // THIS CODE'):
                log.debug(f'last line: {last_line.strip()}')
                if last_line.strip() == package_line.strip():
                    log.error(
                        'resolver.go has already been processed, please re-generate!')
                    sys.exit(-2)
                d.write(package_line + "\n")
            d.write(line)
            m = p.match(line)
            if m:
                name, params, return_types = create(m, package, parent_module)
                parameters = ', '.join(params.split()[::2])
                # add new created funtion
                d.write(
                    f'\treturn q.{name}({parameters}, DB) ({return_types})\n')
                # skip panic code
                next(f, None)
            last_line = line


def create(m, package, parent_module):
    qtype = m.group(1)
    name = m.group(2)
    filename = name[0].lower() + name[1:] + '.go'
    parameters = m.group(3)
    return_types = m.group(4)
    new_returns = return_types
    # if file exist leave it alone
    if os.path.isfile(filename):
        log.warning(f'{filename} already exist, being left alone')
        return name, parameters, return_types
    log.info(f'creating {filename} ...')
    # ignore derivations
    custom_return_types = return_types.translate({ord(i): None for i in '[]*'})
    # find all custom types
    custom_return_types = custom_return_types.split(', ')
    custom_return_types = set(custom_return_types) - set(GQL_PRIMITIVES)
    log.debug(f'custom return types: {custom_return_types}')
    # replace with parent type
    log.debug(f'old return types: {return_types}')
    for each in custom_return_types:
        new_returns = new_returns.replace(each, f'p.{each}')
    log.debug(f'new return types: {new_returns}')
    with open(filename, 'w+') as f:
        f.write(f"""
package queries

import (
    "context"
    "github.com/jmoiron/sqlx"
""")
        if len(custom_return_types) > 0:
            f.write(f'  p "{parent_module}"')
        f.write(f"""
)

// {name} ...
func {name}({parameters}, DB *sqlx.DB) ({new_returns}) {{
    panic("not implemented yet")
}}
""")
    return name, parameters, return_types


if __name__ == '__main__':
    main()
