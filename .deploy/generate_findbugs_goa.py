#!/bin/python

def main():

    with open('bugs.txt') as f:
        with open('findbugs.design', 'w') as g:
            for bug in f:
                data = bug.split(':')
                abbreviation    = data[0].strip()
                description     = data[1].strip()

                s = """Member("{0}", Boolean, func() {{ \n Description( "{1}" ) \n }}) \n""".format(abbreviation, description)
                g.write(s)

if __name__ == '__main__':
    main()
